// +build !ignore_autogenerated

/*
Copyright 2020 BlackRock, Inc.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	common "github.com/argoproj/argo-events/pkg/apis/common"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSLambdaTrigger) DeepCopyInto(out *AWSLambdaTrigger) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Payload != nil {
		in, out := &in.Payload, &out.Payload
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSLambdaTrigger.
func (in *AWSLambdaTrigger) DeepCopy() *AWSLambdaTrigger {
	if in == nil {
		return nil
	}
	out := new(AWSLambdaTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoWorkflowTrigger) DeepCopyInto(out *ArgoWorkflowTrigger) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(ArtifactLocation)
		(*in).DeepCopyInto(*out)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.GroupVersionResource = in.GroupVersionResource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoWorkflowTrigger.
func (in *ArgoWorkflowTrigger) DeepCopy() *ArgoWorkflowTrigger {
	if in == nil {
		return nil
	}
	out := new(ArgoWorkflowTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactLocation) DeepCopyInto(out *ArtifactLocation) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(common.S3Artifact)
		(*in).DeepCopyInto(*out)
	}
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = new(string)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(FileArtifact)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(URLArtifact)
		**out = **in
	}
	if in.Configmap != nil {
		in, out := &in.Configmap, &out.Configmap
		*out = new(ConfigmapArtifact)
		**out = **in
	}
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitArtifact)
		(*in).DeepCopyInto(*out)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(common.Resource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactLocation.
func (in *ArtifactLocation) DeepCopy() *ArtifactLocation {
	if in == nil {
		return nil
	}
	out := new(ArtifactLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	*out = *in
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuth.
func (in *BasicAuth) DeepCopy() *BasicAuth {
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigmapArtifact) DeepCopyInto(out *ConfigmapArtifact) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigmapArtifact.
func (in *ConfigmapArtifact) DeepCopy() *ConfigmapArtifact {
	if in == nil {
		return nil
	}
	out := new(ConfigmapArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomTrigger) DeepCopyInto(out *CustomTrigger) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Payload != nil {
		in, out := &in.Payload, &out.Payload
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomTrigger.
func (in *CustomTrigger) DeepCopy() *CustomTrigger {
	if in == nil {
		return nil
	}
	out := new(CustomTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataFilter) DeepCopyInto(out *DataFilter) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataFilter.
func (in *DataFilter) DeepCopy() *DataFilter {
	if in == nil {
		return nil
	}
	out := new(DataFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependencyGroup) DeepCopyInto(out *DependencyGroup) {
	*out = *in
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependencyGroup.
func (in *DependencyGroup) DeepCopy() *DependencyGroup {
	if in == nil {
		return nil
	}
	out := new(DependencyGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Event) DeepCopyInto(out *Event) {
	*out = *in
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(EventContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Event.
func (in *Event) DeepCopy() *Event {
	if in == nil {
		return nil
	}
	out := new(Event)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventContext) DeepCopyInto(out *EventContext) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventContext.
func (in *EventContext) DeepCopy() *EventContext {
	if in == nil {
		return nil
	}
	out := new(EventContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventDependency) DeepCopyInto(out *EventDependency) {
	*out = *in
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = new(EventDependencyFilter)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventDependency.
func (in *EventDependency) DeepCopy() *EventDependency {
	if in == nil {
		return nil
	}
	out := new(EventDependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventDependencyFilter) DeepCopyInto(out *EventDependencyFilter) {
	*out = *in
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(TimeFilter)
		**out = **in
	}
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(EventContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DataFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventDependencyFilter.
func (in *EventDependencyFilter) DeepCopy() *EventDependencyFilter {
	if in == nil {
		return nil
	}
	out := new(EventDependencyFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileArtifact) DeepCopyInto(out *FileArtifact) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileArtifact.
func (in *FileArtifact) DeepCopy() *FileArtifact {
	if in == nil {
		return nil
	}
	out := new(FileArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitArtifact) DeepCopyInto(out *GitArtifact) {
	*out = *in
	if in.Creds != nil {
		in, out := &in.Creds, &out.Creds
		*out = new(GitCreds)
		(*in).DeepCopyInto(*out)
	}
	if in.Remote != nil {
		in, out := &in.Remote, &out.Remote
		*out = new(GitRemoteConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitArtifact.
func (in *GitArtifact) DeepCopy() *GitArtifact {
	if in == nil {
		return nil
	}
	out := new(GitArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitCreds) DeepCopyInto(out *GitCreds) {
	*out = *in
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitCreds.
func (in *GitCreds) DeepCopy() *GitCreds {
	if in == nil {
		return nil
	}
	out := new(GitCreds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRemoteConfig) DeepCopyInto(out *GitRemoteConfig) {
	*out = *in
	if in.URLS != nil {
		in, out := &in.URLS, &out.URLS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRemoteConfig.
func (in *GitRemoteConfig) DeepCopy() *GitRemoteConfig {
	if in == nil {
		return nil
	}
	out := new(GitRemoteConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSubscription) DeepCopyInto(out *HTTPSubscription) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSubscription.
func (in *HTTPSubscription) DeepCopy() *HTTPSubscription {
	if in == nil {
		return nil
	}
	out := new(HTTPSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPTrigger) DeepCopyInto(out *HTTPTrigger) {
	*out = *in
	if in.Payload != nil {
		in, out := &in.Payload, &out.Payload
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(BasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPTrigger.
func (in *HTTPTrigger) DeepCopy() *HTTPTrigger {
	if in == nil {
		return nil
	}
	out := new(HTTPTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8SResourcePolicy) DeepCopyInto(out *K8SResourcePolicy) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Backoff.DeepCopyInto(&out.Backoff)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SResourcePolicy.
func (in *K8SResourcePolicy) DeepCopy() *K8SResourcePolicy {
	if in == nil {
		return nil
	}
	out := new(K8SResourcePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaTrigger) DeepCopyInto(out *KafkaTrigger) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	if in.Payload != nil {
		in, out := &in.Payload, &out.Payload
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaTrigger.
func (in *KafkaTrigger) DeepCopy() *KafkaTrigger {
	if in == nil {
		return nil
	}
	out := new(KafkaTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATSSubscription) DeepCopyInto(out *NATSSubscription) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATSSubscription.
func (in *NATSSubscription) DeepCopy() *NATSSubscription {
	if in == nil {
		return nil
	}
	out := new(NATSSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATSTrigger) DeepCopyInto(out *NATSTrigger) {
	*out = *in
	if in.Payload != nil {
		in, out := &in.Payload, &out.Payload
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATSTrigger.
func (in *NATSTrigger) DeepCopy() *NATSTrigger {
	if in == nil {
		return nil
	}
	out := new(NATSTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	in.CompletedAt.DeepCopyInto(&out.CompletedAt)
	if in.Event != nil {
		in, out := &in.Event, &out.Event
		*out = new(Event)
		(*in).DeepCopyInto(*out)
	}
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
	in.ResolvedAt.DeepCopyInto(&out.ResolvedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenWhiskTrigger) DeepCopyInto(out *OpenWhiskTrigger) {
	*out = *in
	if in.AuthToken != nil {
		in, out := &in.AuthToken, &out.AuthToken
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Payload != nil {
		in, out := &in.Payload, &out.Payload
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenWhiskTrigger.
func (in *OpenWhiskTrigger) DeepCopy() *OpenWhiskTrigger {
	if in == nil {
		return nil
	}
	out := new(OpenWhiskTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sensor) DeepCopyInto(out *Sensor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sensor.
func (in *Sensor) DeepCopy() *Sensor {
	if in == nil {
		return nil
	}
	out := new(Sensor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sensor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorList) DeepCopyInto(out *SensorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sensor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorList.
func (in *SensorList) DeepCopy() *SensorList {
	if in == nil {
		return nil
	}
	out := new(SensorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SensorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorResources) DeepCopyInto(out *SensorResources) {
	*out = *in
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(metav1.ObjectMeta)
		(*in).DeepCopyInto(*out)
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(metav1.ObjectMeta)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorResources.
func (in *SensorResources) DeepCopy() *SensorResources {
	if in == nil {
		return nil
	}
	out := new(SensorResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorSpec) DeepCopyInto(out *SensorSpec) {
	*out = *in
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]EventDependency, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]Trigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	if in.Subscription != nil {
		in, out := &in.Subscription, &out.Subscription
		*out = new(Subscription)
		(*in).DeepCopyInto(*out)
	}
	if in.DependencyGroups != nil {
		in, out := &in.DependencyGroups, &out.DependencyGroups
		*out = make([]DependencyGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceLabels != nil {
		in, out := &in.ServiceLabels, &out.ServiceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceAnnotations != nil {
		in, out := &in.ServiceAnnotations, &out.ServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorSpec.
func (in *SensorSpec) DeepCopy() *SensorSpec {
	if in == nil {
		return nil
	}
	out := new(SensorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorStatus) DeepCopyInto(out *SensorStatus) {
	*out = *in
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	in.CompletedAt.DeepCopyInto(&out.CompletedAt)
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]NodeStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.LastCycleTime.DeepCopyInto(&out.LastCycleTime)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(SensorResources)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorStatus.
func (in *SensorStatus) DeepCopy() *SensorStatus {
	if in == nil {
		return nil
	}
	out := new(SensorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackTrigger) DeepCopyInto(out *SlackTrigger) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SlackToken != nil {
		in, out := &in.SlackToken, &out.SlackToken
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackTrigger.
func (in *SlackTrigger) DeepCopy() *SlackTrigger {
	if in == nil {
		return nil
	}
	out := new(SlackTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandardK8STrigger) DeepCopyInto(out *StandardK8STrigger) {
	*out = *in
	out.GroupVersionResource = in.GroupVersionResource
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(ArtifactLocation)
		(*in).DeepCopyInto(*out)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandardK8STrigger.
func (in *StandardK8STrigger) DeepCopy() *StandardK8STrigger {
	if in == nil {
		return nil
	}
	out := new(StandardK8STrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusPolicy) DeepCopyInto(out *StatusPolicy) {
	*out = *in
	if in.Allow != nil {
		in, out := &in.Allow, &out.Allow
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusPolicy.
func (in *StatusPolicy) DeepCopy() *StatusPolicy {
	if in == nil {
		return nil
	}
	out := new(StatusPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscription) DeepCopyInto(out *Subscription) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(HTTPSubscription)
		**out = **in
	}
	if in.NATS != nil {
		in, out := &in.NATS, &out.NATS
		*out = new(NATSSubscription)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscription.
func (in *Subscription) DeepCopy() *Subscription {
	if in == nil {
		return nil
	}
	out := new(Subscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Template) DeepCopyInto(out *Template) {
	*out = *in
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(v1.Container)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(v1.PodSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Template.
func (in *Template) DeepCopy() *Template {
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeFilter) DeepCopyInto(out *TimeFilter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeFilter.
func (in *TimeFilter) DeepCopy() *TimeFilter {
	if in == nil {
		return nil
	}
	out := new(TimeFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(TriggerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]TriggerParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(TriggerPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerParameter) DeepCopyInto(out *TriggerParameter) {
	*out = *in
	if in.Src != nil {
		in, out := &in.Src, &out.Src
		*out = new(TriggerParameterSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerParameter.
func (in *TriggerParameter) DeepCopy() *TriggerParameter {
	if in == nil {
		return nil
	}
	out := new(TriggerParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerParameterSource) DeepCopyInto(out *TriggerParameterSource) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerParameterSource.
func (in *TriggerParameterSource) DeepCopy() *TriggerParameterSource {
	if in == nil {
		return nil
	}
	out := new(TriggerParameterSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerPolicy) DeepCopyInto(out *TriggerPolicy) {
	*out = *in
	if in.K8s != nil {
		in, out := &in.K8s, &out.K8s
		*out = new(K8SResourcePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(StatusPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerPolicy.
func (in *TriggerPolicy) DeepCopy() *TriggerPolicy {
	if in == nil {
		return nil
	}
	out := new(TriggerPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSwitch) DeepCopyInto(out *TriggerSwitch) {
	*out = *in
	if in.Any != nil {
		in, out := &in.Any, &out.Any
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.All != nil {
		in, out := &in.All, &out.All
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSwitch.
func (in *TriggerSwitch) DeepCopy() *TriggerSwitch {
	if in == nil {
		return nil
	}
	out := new(TriggerSwitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerTemplate) DeepCopyInto(out *TriggerTemplate) {
	*out = *in
	if in.Switch != nil {
		in, out := &in.Switch, &out.Switch
		*out = new(TriggerSwitch)
		(*in).DeepCopyInto(*out)
	}
	if in.K8s != nil {
		in, out := &in.K8s, &out.K8s
		*out = new(StandardK8STrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.ArgoWorkflow != nil {
		in, out := &in.ArgoWorkflow, &out.ArgoWorkflow
		*out = new(ArgoWorkflowTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(HTTPTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.AWSLambda != nil {
		in, out := &in.AWSLambda, &out.AWSLambda
		*out = new(AWSLambdaTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomTrigger != nil {
		in, out := &in.CustomTrigger, &out.CustomTrigger
		*out = new(CustomTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = new(KafkaTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.NATS != nil {
		in, out := &in.NATS, &out.NATS
		*out = new(NATSTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.Slack != nil {
		in, out := &in.Slack, &out.Slack
		*out = new(SlackTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.OpenWhisk != nil {
		in, out := &in.OpenWhisk, &out.OpenWhisk
		*out = new(OpenWhiskTrigger)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerTemplate.
func (in *TriggerTemplate) DeepCopy() *TriggerTemplate {
	if in == nil {
		return nil
	}
	out := new(TriggerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *URLArtifact) DeepCopyInto(out *URLArtifact) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new URLArtifact.
func (in *URLArtifact) DeepCopy() *URLArtifact {
	if in == nil {
		return nil
	}
	out := new(URLArtifact)
	in.DeepCopyInto(out)
	return out
}
