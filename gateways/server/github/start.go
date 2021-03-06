/*
Copyright 2018 KompiTech GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package github

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/argoproj/argo-events/gateways/server"
	"github.com/argoproj/argo-events/pkg/apis/events"

	"github.com/argoproj/argo-events/common"
	"github.com/argoproj/argo-events/gateways"
	"github.com/argoproj/argo-events/gateways/server/common/webhook"
	"github.com/argoproj/argo-events/pkg/apis/eventsource/v1alpha1"
	"github.com/argoproj/argo-events/store"
	"github.com/ghodss/yaml"
	gh "github.com/google/go-github/v31/github"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
)

// GitHub headers
const (
	githubEventHeader    = "X-GitHub-Event"
	githubDeliveryHeader = "X-GitHub-Delivery"
)

// controller controls the webhook operations
var (
	controller = webhook.NewController()
)

// set up the activation and inactivation channels to control the state of routes.
func init() {
	go webhook.ProcessRouteStatus(controller)
}

// getCredentials for retrieves credentials for GitHub connection
func (router *Router) getCredentials(keySelector *corev1.SecretKeySelector, namespace string) (*cred, error) {
	token, err := store.GetSecrets(router.k8sClient, namespace, keySelector.Name, keySelector.Key)
	if err != nil {
		return nil, err
	}
	return &cred{
		secret: token,
	}, nil
}

// Implement Router
// 1. GetRoute
// 2. HandleRoute
// 3. PostActivate
// 4. PostDeactivate

// GetRoute returns the route
func (router *Router) GetRoute() *webhook.Route {
	return router.route
}

// HandleRoute handles incoming requests on the route
func (router *Router) HandleRoute(writer http.ResponseWriter, request *http.Request) {
	route := router.route

	logger := route.Logger.WithFields(
		map[string]interface{}{
			common.LabelEventSource: route.EventSource.Name,
			common.LabelEndpoint:    route.Context.Endpoint,
			common.LabelPort:        route.Context.Port,
		})

	logger.Info("received a request, processing it...")

	if !route.Active {
		logger.Info("endpoint is not active, won't process the request")
		common.SendErrorResponse(writer, "endpoint is inactive")
		return
	}

	hook := router.hook
	secret := ""
	if s, ok := hook.Config["secret"]; ok {
		secret = s.(string)
	}

	body, err := parseValidateRequest(request, []byte(secret))
	if err != nil {
		logger.WithError(err).Error("request is not valid event notification, discarding it")
		common.SendErrorResponse(writer, err.Error())
		return
	}

	event := &events.GithubEventData{
		Headers: request.Header,
		Body:    (*json.RawMessage)(&body),
	}

	eventBody, err := json.Marshal(event)
	if err != nil {
		logger.Info("failed to marshal event")
		common.SendErrorResponse(writer, "invalid event")
		return
	}

	logger.Infoln("dispatching event on route's data channel")
	route.DataCh <- eventBody
	logger.Info("request successfully processed")

	common.SendSuccessResponse(writer, "success")
}

// PostActivate performs operations once the route is activated and ready to consume requests
func (router *Router) PostActivate() error {
	// In order to successfully setup a GitHub hook for the given repository,
	// 1. Get the API Token and Webhook secret from K8s secrets
	// 2. Configure the hook with url, content type, ssl etc.
	// 3. Set up a GitHub client
	// 4. Set the base and upload url for the client
	// 5. Create the hook if one doesn't exist already. If exists already, then use that one.

	route := router.route
	githubEventSource := router.githubEventSource

	logger := route.Logger.WithFields(map[string]interface{}{
		common.LabelEventSource: route.EventSource.Name,
		"repository":            githubEventSource.Repository,
	})

	logger.Infoln("retrieving api token credentials...")
	apiTokenCreds, err := router.getCredentials(githubEventSource.APIToken, githubEventSource.Namespace)
	if err != nil {
		return errors.Errorf("failed to retrieve api token credentials. err: %+v", err)
	}

	logger.Infoln("setting up auth with api token...")
	PATTransport := TokenAuthTransport{
		Token: apiTokenCreds.secret,
	}

	logger.Infoln("configuring GitHub hook...")
	formattedURL := common.FormattedURL(githubEventSource.Webhook.URL, githubEventSource.Webhook.Endpoint)
	hookConfig := map[string]interface{}{
		"url": &formattedURL,
	}

	if githubEventSource.ContentType != "" {
		hookConfig["content_type"] = githubEventSource.ContentType
	}

	if githubEventSource.Insecure {
		hookConfig["insecure_ssl"] = "1"
	} else {
		hookConfig["insecure_ssl"] = "0"
	}

	logger.Infoln("retrieving webhook secret credentials...")
	if githubEventSource.WebhookSecret != nil {
		webhookSecretCreds, err := router.getCredentials(githubEventSource.WebhookSecret, githubEventSource.Namespace)
		if err != nil {
			return errors.Errorf("failed to retrieve webhook secret. err: %+v", err)
		}
		hookConfig["secret"] = webhookSecretCreds.secret
	}

	router.hook = &gh.Hook{
		Events: githubEventSource.Events,
		Active: gh.Bool(githubEventSource.Active),
		Config: hookConfig,
	}

	logger.Infoln("setting up client for GitHub...")
	router.githubClient = gh.NewClient(PATTransport.Client())

	logger.Infoln("setting up base url for GitHub client...")
	if githubEventSource.GithubBaseURL != "" {
		baseURL, err := url.Parse(githubEventSource.GithubBaseURL)
		if err != nil {
			return errors.Errorf("failed to parse github base url. err: %s", err)
		}
		router.githubClient.BaseURL = baseURL
	}

	logger.Infoln("setting up the upload url for GitHub client...")
	if githubEventSource.GithubUploadURL != "" {
		uploadURL, err := url.Parse(githubEventSource.GithubUploadURL)
		if err != nil {
			return errors.Errorf("failed to parse github upload url. err: %s", err)
		}
		router.githubClient.UploadURL = uploadURL
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logger.Infoln("creating a GitHub hook for the repository...")
	hook, _, err := router.githubClient.Repositories.CreateHook(ctx, githubEventSource.Owner, githubEventSource.Repository, router.hook)
	if err != nil {
		// Continue if error is because hook already exists
		er, ok := err.(*gh.ErrorResponse)
		if !ok || er.Response.StatusCode != http.StatusUnprocessableEntity {
			return errors.Errorf("failed to create webhook. err: %+v", err)
		}
	}

	// if hook alreay exists then CreateHook returns hook value as nil
	if hook == nil {
		logger.Infoln("GitHub hook for the repository already exists, trying to use the existing hook...")
		ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		hooks, _, err := router.githubClient.Repositories.ListHooks(ctx, githubEventSource.Owner, githubEventSource.Repository, nil)
		if err != nil {
			return errors.Errorf("failed to list existing webhooks. err: %+v", err)
		}

		hook = getHook(hooks, formattedURL, githubEventSource.Events)
		if hook == nil {
			return errors.New("failed to find existing webhook")
		}
	}

	if githubEventSource.WebhookSecret != nil {
		// As secret in hook config is masked with asterisk (*), replace it with unmasked secret.
		hook.Config["secret"] = hookConfig["secret"]
	}

	router.hook = hook
	logger.Infoln("GitHub hook has been successfully set for the repository")

	return nil
}

// PostInactivate performs operations after the route is inactivated
func (router *Router) PostInactivate() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	githubEventSource := router.githubEventSource

	if githubEventSource.DeleteHookOnFinish {
		logger := router.route.Logger.WithFields(map[string]interface{}{
			common.LabelEventSource: router.route.EventSource.Name,
			"repository":            githubEventSource.Repository,
			"hook-id":               *router.hook.ID,
		})

		logger.Infoln("deleting GitHub hook...")
		if _, err := router.githubClient.Repositories.DeleteHook(ctx, githubEventSource.Owner, githubEventSource.Repository, *router.hook.ID); err != nil {
			return errors.Errorf("failed to delete hook. err: %+v", err)
		}
		logger.Infoln("GitHub hook deleted")
	}

	return nil
}

// StartEventSource starts an event source
func (listener *EventListener) StartEventSource(eventSource *gateways.EventSource, eventStream gateways.Eventing_StartEventSourceServer) error {
	defer server.Recover(eventSource.Name)

	listener.Logger.WithField(common.LabelEventSource, eventSource.Name).Infoln("started processing the event source...")

	var githubEventSource *v1alpha1.GithubEventSource
	if err := yaml.Unmarshal(eventSource.Value, &githubEventSource); err != nil {
		listener.Logger.WithError(err).WithField(common.LabelEventSource, eventSource.Name).Infoln("failed to parse the event source")
		return err
	}

	if githubEventSource.Namespace == "" {
		githubEventSource.Namespace = listener.Namespace
	}

	route := webhook.NewRoute(githubEventSource.Webhook, listener.Logger, eventSource)

	return webhook.ManageRoute(&Router{
		route:             route,
		k8sClient:         listener.K8sClient,
		githubEventSource: githubEventSource,
	}, controller, eventStream)
}

// parseValidateRequest parses a http request and checks if it is valid GitHub notification
func parseValidateRequest(r *http.Request, secret []byte) ([]byte, error) {
	body, err := gh.ValidatePayload(r, secret)
	if err != nil {
		return nil, err
	}

	payload := make(map[string]interface{})
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}
	for _, h := range []string{
		githubEventHeader,
		githubDeliveryHeader,
	} {
		payload[h] = r.Header.Get(h)
	}
	return json.Marshal(payload)
}
