//go:build !ignore_autogenerated

/*
SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServer) DeepCopyInto(out *APIServer) {
	*out = *in
	if in.ServiceRef != nil {
		in, out := &in.ServiceRef, &out.ServiceRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.CAData != nil {
		in, out := &in.CAData, &out.CAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServer.
func (in *APIServer) DeepCopy() *APIServer {
	if in == nil {
		return nil
	}
	out := new(APIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization) DeepCopyInto(out *Authorization) {
	*out = *in
	if in.RoleBindings != nil {
		in, out := &in.RoleBindings, &out.RoleBindings
		*out = make([]RoleBinding, len(*in))
		copy(*out, *in)
	}
	if in.ProjectMemberships != nil {
		in, out := &in.ProjectMemberships, &out.ProjectMemberships
		*out = make([]ProjectMembership, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization.
func (in *Authorization) DeepCopy() *Authorization {
	if in == nil {
		return nil
	}
	out := new(Authorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCredentials) DeepCopyInto(out *ClusterCredentials) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.ShootRef != nil {
		in, out := &in.ShootRef, &out.ShootRef
		*out = new(ShootRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCredentials.
func (in *ClusterCredentials) DeepCopy() *ClusterCredentials {
	if in == nil {
		return nil
	}
	out := new(ClusterCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerConfiguration) DeepCopyInto(out *ControllerManagerConfiguration) {
	*out = *in
	in.Server.DeepCopyInto(&out.Server)
	in.Controllers.DeepCopyInto(&out.Controllers)
	out.Webhooks = in.Webhooks
	if in.HonourServiceAccountRefHostCluster != nil {
		in, out := &in.HonourServiceAccountRefHostCluster, &out.HonourServiceAccountRefHostCluster
		*out = new(bool)
		**out = **in
	}
	if in.HonourServiceAccountRefTargetCluster != nil {
		in, out := &in.HonourServiceAccountRefTargetCluster, &out.HonourServiceAccountRefTargetCluster
		*out = new(bool)
		**out = **in
	}
	if in.HonourProjectMemberships != nil {
		in, out := &in.HonourProjectMemberships, &out.HonourProjectMemberships
		*out = new(bool)
		**out = **in
	}
	if in.HonourCleanupProjectMembership != nil {
		in, out := &in.HonourCleanupProjectMembership, &out.HonourCleanupProjectMembership
		*out = new(bool)
		**out = **in
	}
	if in.LeaderElection != nil {
		in, out := &in.LeaderElection, &out.LeaderElection
		*out = new(configv1alpha1.LeaderElectionConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerConfiguration.
func (in *ControllerManagerConfiguration) DeepCopy() *ControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerControllerConfiguration) DeepCopyInto(out *ControllerManagerControllerConfiguration) {
	*out = *in
	in.Terminal.DeepCopyInto(&out.Terminal)
	out.TerminalHeartbeat = in.TerminalHeartbeat
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerControllerConfiguration.
func (in *ControllerManagerControllerConfiguration) DeepCopy() *ControllerManagerControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerWebhookConfiguration) DeepCopyInto(out *ControllerManagerWebhookConfiguration) {
	*out = *in
	out.TerminalValidation = in.TerminalValidation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerWebhookConfiguration.
func (in *ControllerManagerWebhookConfiguration) DeepCopy() *ControllerManagerWebhookConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerWebhookConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Duration) DeepCopyInto(out *Duration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Duration.
func (in *Duration) DeepCopy() *Duration {
	if in == nil {
		return nil
	}
	out := new(Duration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostCluster) DeepCopyInto(out *HostCluster) {
	*out = *in
	in.Credentials.DeepCopyInto(&out.Credentials)
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.TemporaryNamespace != nil {
		in, out := &in.TemporaryNamespace, &out.TemporaryNamespace
		*out = new(bool)
		**out = **in
	}
	in.Pod.DeepCopyInto(&out.Pod)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostCluster.
func (in *HostCluster) DeepCopy() *HostCluster {
	if in == nil {
		return nil
	}
	out := new(HostCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastError) DeepCopyInto(out *LastError) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastError.
func (in *LastError) DeepCopy() *LastError {
	if in == nil {
		return nil
	}
	out := new(LastError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastOperation) DeepCopyInto(out *LastOperation) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastOperation.
func (in *LastOperation) DeepCopy() *LastOperation {
	if in == nil {
		return nil
	}
	out := new(LastOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pod) DeepCopyInto(out *Pod) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(Container)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pod.
func (in *Pod) DeepCopy() *Pod {
	if in == nil {
		return nil
	}
	out := new(Pod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectMembership) DeepCopyInto(out *ProjectMembership) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectMembership.
func (in *ProjectMembership) DeepCopy() *ProjectMembership {
	if in == nil {
		return nil
	}
	out := new(ProjectMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBinding) DeepCopyInto(out *RoleBinding) {
	*out = *in
	out.RoleRef = in.RoleRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBinding.
func (in *RoleBinding) DeepCopy() *RoleBinding {
	if in == nil {
		return nil
	}
	out := new(RoleBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfiguration) DeepCopyInto(out *ServerConfiguration) {
	*out = *in
	if in.HealthProbes != nil {
		in, out := &in.HealthProbes, &out.HealthProbes
		*out = new(Server)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Server)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfiguration.
func (in *ServerConfiguration) DeepCopy() *ServerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountControllerConfiguration) DeepCopyInto(out *ServiceAccountControllerConfiguration) {
	*out = *in
	if in.AllowedServiceAccountNames != nil {
		in, out := &in.AllowedServiceAccountNames, &out.AllowedServiceAccountNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountControllerConfiguration.
func (in *ServiceAccountControllerConfiguration) DeepCopy() *ServiceAccountControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootRef) DeepCopyInto(out *ShootRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootRef.
func (in *ShootRef) DeepCopy() *ShootRef {
	if in == nil {
		return nil
	}
	out := new(ShootRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetCluster) DeepCopyInto(out *TargetCluster) {
	*out = *in
	in.Credentials.DeepCopyInto(&out.Credentials)
	if in.CleanupProjectMembership != nil {
		in, out := &in.CleanupProjectMembership, &out.CleanupProjectMembership
		*out = new(bool)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.TemporaryNamespace != nil {
		in, out := &in.TemporaryNamespace, &out.TemporaryNamespace
		*out = new(bool)
		**out = **in
	}
	if in.APIServerServiceRef != nil {
		in, out := &in.APIServerServiceRef, &out.APIServerServiceRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.APIServer != nil {
		in, out := &in.APIServer, &out.APIServer
		*out = new(APIServer)
		(*in).DeepCopyInto(*out)
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(Authorization)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetCluster.
func (in *TargetCluster) DeepCopy() *TargetCluster {
	if in == nil {
		return nil
	}
	out := new(TargetCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Terminal) DeepCopyInto(out *Terminal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Terminal.
func (in *Terminal) DeepCopy() *Terminal {
	if in == nil {
		return nil
	}
	out := new(Terminal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Terminal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerminalControllerConfiguration) DeepCopyInto(out *TerminalControllerConfiguration) {
	*out = *in
	if in.TokenRequestExpirationSeconds != nil {
		in, out := &in.TokenRequestExpirationSeconds, &out.TokenRequestExpirationSeconds
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerminalControllerConfiguration.
func (in *TerminalControllerConfiguration) DeepCopy() *TerminalControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(TerminalControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerminalHeartbeatControllerConfiguration) DeepCopyInto(out *TerminalHeartbeatControllerConfiguration) {
	*out = *in
	out.TimeToLive = in.TimeToLive
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerminalHeartbeatControllerConfiguration.
func (in *TerminalHeartbeatControllerConfiguration) DeepCopy() *TerminalHeartbeatControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(TerminalHeartbeatControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerminalList) DeepCopyInto(out *TerminalList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Terminal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerminalList.
func (in *TerminalList) DeepCopy() *TerminalList {
	if in == nil {
		return nil
	}
	out := new(TerminalList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerminalList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerminalSpec) DeepCopyInto(out *TerminalSpec) {
	*out = *in
	in.Host.DeepCopyInto(&out.Host)
	in.Target.DeepCopyInto(&out.Target)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerminalSpec.
func (in *TerminalSpec) DeepCopy() *TerminalSpec {
	if in == nil {
		return nil
	}
	out := new(TerminalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerminalStatus) DeepCopyInto(out *TerminalStatus) {
	*out = *in
	if in.AttachServiceAccountName != nil {
		in, out := &in.AttachServiceAccountName, &out.AttachServiceAccountName
		*out = new(string)
		**out = **in
	}
	if in.PodName != nil {
		in, out := &in.PodName, &out.PodName
		*out = new(string)
		**out = **in
	}
	if in.LastOperation != nil {
		in, out := &in.LastOperation, &out.LastOperation
		*out = new(LastOperation)
		(*in).DeepCopyInto(*out)
	}
	if in.LastError != nil {
		in, out := &in.LastError, &out.LastError
		*out = new(LastError)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerminalStatus.
func (in *TerminalStatus) DeepCopy() *TerminalStatus {
	if in == nil {
		return nil
	}
	out := new(TerminalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerminalValidatingWebhookConfiguration) DeepCopyInto(out *TerminalValidatingWebhookConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerminalValidatingWebhookConfiguration.
func (in *TerminalValidatingWebhookConfiguration) DeepCopy() *TerminalValidatingWebhookConfiguration {
	if in == nil {
		return nil
	}
	out := new(TerminalValidatingWebhookConfiguration)
	in.DeepCopyInto(out)
	return out
}
