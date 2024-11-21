//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSecurityProfileWorkspaceInitParameters) DeepCopyInto(out *ComplianceSecurityProfileWorkspaceInitParameters) {
	*out = *in
	if in.ComplianceStandards != nil {
		in, out := &in.ComplianceStandards, &out.ComplianceStandards
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSecurityProfileWorkspaceInitParameters.
func (in *ComplianceSecurityProfileWorkspaceInitParameters) DeepCopy() *ComplianceSecurityProfileWorkspaceInitParameters {
	if in == nil {
		return nil
	}
	out := new(ComplianceSecurityProfileWorkspaceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSecurityProfileWorkspaceObservation) DeepCopyInto(out *ComplianceSecurityProfileWorkspaceObservation) {
	*out = *in
	if in.ComplianceStandards != nil {
		in, out := &in.ComplianceStandards, &out.ComplianceStandards
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSecurityProfileWorkspaceObservation.
func (in *ComplianceSecurityProfileWorkspaceObservation) DeepCopy() *ComplianceSecurityProfileWorkspaceObservation {
	if in == nil {
		return nil
	}
	out := new(ComplianceSecurityProfileWorkspaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSecurityProfileWorkspaceParameters) DeepCopyInto(out *ComplianceSecurityProfileWorkspaceParameters) {
	*out = *in
	if in.ComplianceStandards != nil {
		in, out := &in.ComplianceStandards, &out.ComplianceStandards
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSecurityProfileWorkspaceParameters.
func (in *ComplianceSecurityProfileWorkspaceParameters) DeepCopy() *ComplianceSecurityProfileWorkspaceParameters {
	if in == nil {
		return nil
	}
	out := new(ComplianceSecurityProfileWorkspaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSecurityProfileWorkspaceSetting) DeepCopyInto(out *ComplianceSecurityProfileWorkspaceSetting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSecurityProfileWorkspaceSetting.
func (in *ComplianceSecurityProfileWorkspaceSetting) DeepCopy() *ComplianceSecurityProfileWorkspaceSetting {
	if in == nil {
		return nil
	}
	out := new(ComplianceSecurityProfileWorkspaceSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComplianceSecurityProfileWorkspaceSetting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSecurityProfileWorkspaceSettingInitParameters) DeepCopyInto(out *ComplianceSecurityProfileWorkspaceSettingInitParameters) {
	*out = *in
	if in.ComplianceSecurityProfileWorkspace != nil {
		in, out := &in.ComplianceSecurityProfileWorkspace, &out.ComplianceSecurityProfileWorkspace
		*out = make([]ComplianceSecurityProfileWorkspaceInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSecurityProfileWorkspaceSettingInitParameters.
func (in *ComplianceSecurityProfileWorkspaceSettingInitParameters) DeepCopy() *ComplianceSecurityProfileWorkspaceSettingInitParameters {
	if in == nil {
		return nil
	}
	out := new(ComplianceSecurityProfileWorkspaceSettingInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSecurityProfileWorkspaceSettingList) DeepCopyInto(out *ComplianceSecurityProfileWorkspaceSettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComplianceSecurityProfileWorkspaceSetting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSecurityProfileWorkspaceSettingList.
func (in *ComplianceSecurityProfileWorkspaceSettingList) DeepCopy() *ComplianceSecurityProfileWorkspaceSettingList {
	if in == nil {
		return nil
	}
	out := new(ComplianceSecurityProfileWorkspaceSettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComplianceSecurityProfileWorkspaceSettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSecurityProfileWorkspaceSettingObservation) DeepCopyInto(out *ComplianceSecurityProfileWorkspaceSettingObservation) {
	*out = *in
	if in.ComplianceSecurityProfileWorkspace != nil {
		in, out := &in.ComplianceSecurityProfileWorkspace, &out.ComplianceSecurityProfileWorkspace
		*out = make([]ComplianceSecurityProfileWorkspaceObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSecurityProfileWorkspaceSettingObservation.
func (in *ComplianceSecurityProfileWorkspaceSettingObservation) DeepCopy() *ComplianceSecurityProfileWorkspaceSettingObservation {
	if in == nil {
		return nil
	}
	out := new(ComplianceSecurityProfileWorkspaceSettingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSecurityProfileWorkspaceSettingParameters) DeepCopyInto(out *ComplianceSecurityProfileWorkspaceSettingParameters) {
	*out = *in
	if in.ComplianceSecurityProfileWorkspace != nil {
		in, out := &in.ComplianceSecurityProfileWorkspace, &out.ComplianceSecurityProfileWorkspace
		*out = make([]ComplianceSecurityProfileWorkspaceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSecurityProfileWorkspaceSettingParameters.
func (in *ComplianceSecurityProfileWorkspaceSettingParameters) DeepCopy() *ComplianceSecurityProfileWorkspaceSettingParameters {
	if in == nil {
		return nil
	}
	out := new(ComplianceSecurityProfileWorkspaceSettingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSecurityProfileWorkspaceSettingSpec) DeepCopyInto(out *ComplianceSecurityProfileWorkspaceSettingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSecurityProfileWorkspaceSettingSpec.
func (in *ComplianceSecurityProfileWorkspaceSettingSpec) DeepCopy() *ComplianceSecurityProfileWorkspaceSettingSpec {
	if in == nil {
		return nil
	}
	out := new(ComplianceSecurityProfileWorkspaceSettingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSecurityProfileWorkspaceSettingStatus) DeepCopyInto(out *ComplianceSecurityProfileWorkspaceSettingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSecurityProfileWorkspaceSettingStatus.
func (in *ComplianceSecurityProfileWorkspaceSettingStatus) DeepCopy() *ComplianceSecurityProfileWorkspaceSettingStatus {
	if in == nil {
		return nil
	}
	out := new(ComplianceSecurityProfileWorkspaceSettingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNamespaceSetting) DeepCopyInto(out *DefaultNamespaceSetting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNamespaceSetting.
func (in *DefaultNamespaceSetting) DeepCopy() *DefaultNamespaceSetting {
	if in == nil {
		return nil
	}
	out := new(DefaultNamespaceSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultNamespaceSetting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNamespaceSettingInitParameters) DeepCopyInto(out *DefaultNamespaceSettingInitParameters) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = make([]NamespaceInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNamespaceSettingInitParameters.
func (in *DefaultNamespaceSettingInitParameters) DeepCopy() *DefaultNamespaceSettingInitParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultNamespaceSettingInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNamespaceSettingList) DeepCopyInto(out *DefaultNamespaceSettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DefaultNamespaceSetting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNamespaceSettingList.
func (in *DefaultNamespaceSettingList) DeepCopy() *DefaultNamespaceSettingList {
	if in == nil {
		return nil
	}
	out := new(DefaultNamespaceSettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultNamespaceSettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNamespaceSettingObservation) DeepCopyInto(out *DefaultNamespaceSettingObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = make([]NamespaceObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNamespaceSettingObservation.
func (in *DefaultNamespaceSettingObservation) DeepCopy() *DefaultNamespaceSettingObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultNamespaceSettingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNamespaceSettingParameters) DeepCopyInto(out *DefaultNamespaceSettingParameters) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = make([]NamespaceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNamespaceSettingParameters.
func (in *DefaultNamespaceSettingParameters) DeepCopy() *DefaultNamespaceSettingParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultNamespaceSettingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNamespaceSettingSpec) DeepCopyInto(out *DefaultNamespaceSettingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNamespaceSettingSpec.
func (in *DefaultNamespaceSettingSpec) DeepCopy() *DefaultNamespaceSettingSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultNamespaceSettingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultNamespaceSettingStatus) DeepCopyInto(out *DefaultNamespaceSettingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultNamespaceSettingStatus.
func (in *DefaultNamespaceSettingStatus) DeepCopy() *DefaultNamespaceSettingStatus {
	if in == nil {
		return nil
	}
	out := new(DefaultNamespaceSettingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedSecurityMonitoringWorkspaceInitParameters) DeepCopyInto(out *EnhancedSecurityMonitoringWorkspaceInitParameters) {
	*out = *in
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedSecurityMonitoringWorkspaceInitParameters.
func (in *EnhancedSecurityMonitoringWorkspaceInitParameters) DeepCopy() *EnhancedSecurityMonitoringWorkspaceInitParameters {
	if in == nil {
		return nil
	}
	out := new(EnhancedSecurityMonitoringWorkspaceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedSecurityMonitoringWorkspaceObservation) DeepCopyInto(out *EnhancedSecurityMonitoringWorkspaceObservation) {
	*out = *in
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedSecurityMonitoringWorkspaceObservation.
func (in *EnhancedSecurityMonitoringWorkspaceObservation) DeepCopy() *EnhancedSecurityMonitoringWorkspaceObservation {
	if in == nil {
		return nil
	}
	out := new(EnhancedSecurityMonitoringWorkspaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedSecurityMonitoringWorkspaceParameters) DeepCopyInto(out *EnhancedSecurityMonitoringWorkspaceParameters) {
	*out = *in
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedSecurityMonitoringWorkspaceParameters.
func (in *EnhancedSecurityMonitoringWorkspaceParameters) DeepCopy() *EnhancedSecurityMonitoringWorkspaceParameters {
	if in == nil {
		return nil
	}
	out := new(EnhancedSecurityMonitoringWorkspaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedSecurityMonitoringWorkspaceSetting) DeepCopyInto(out *EnhancedSecurityMonitoringWorkspaceSetting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedSecurityMonitoringWorkspaceSetting.
func (in *EnhancedSecurityMonitoringWorkspaceSetting) DeepCopy() *EnhancedSecurityMonitoringWorkspaceSetting {
	if in == nil {
		return nil
	}
	out := new(EnhancedSecurityMonitoringWorkspaceSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnhancedSecurityMonitoringWorkspaceSetting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedSecurityMonitoringWorkspaceSettingInitParameters) DeepCopyInto(out *EnhancedSecurityMonitoringWorkspaceSettingInitParameters) {
	*out = *in
	if in.EnhancedSecurityMonitoringWorkspace != nil {
		in, out := &in.EnhancedSecurityMonitoringWorkspace, &out.EnhancedSecurityMonitoringWorkspace
		*out = make([]EnhancedSecurityMonitoringWorkspaceInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedSecurityMonitoringWorkspaceSettingInitParameters.
func (in *EnhancedSecurityMonitoringWorkspaceSettingInitParameters) DeepCopy() *EnhancedSecurityMonitoringWorkspaceSettingInitParameters {
	if in == nil {
		return nil
	}
	out := new(EnhancedSecurityMonitoringWorkspaceSettingInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedSecurityMonitoringWorkspaceSettingList) DeepCopyInto(out *EnhancedSecurityMonitoringWorkspaceSettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnhancedSecurityMonitoringWorkspaceSetting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedSecurityMonitoringWorkspaceSettingList.
func (in *EnhancedSecurityMonitoringWorkspaceSettingList) DeepCopy() *EnhancedSecurityMonitoringWorkspaceSettingList {
	if in == nil {
		return nil
	}
	out := new(EnhancedSecurityMonitoringWorkspaceSettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnhancedSecurityMonitoringWorkspaceSettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedSecurityMonitoringWorkspaceSettingObservation) DeepCopyInto(out *EnhancedSecurityMonitoringWorkspaceSettingObservation) {
	*out = *in
	if in.EnhancedSecurityMonitoringWorkspace != nil {
		in, out := &in.EnhancedSecurityMonitoringWorkspace, &out.EnhancedSecurityMonitoringWorkspace
		*out = make([]EnhancedSecurityMonitoringWorkspaceObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedSecurityMonitoringWorkspaceSettingObservation.
func (in *EnhancedSecurityMonitoringWorkspaceSettingObservation) DeepCopy() *EnhancedSecurityMonitoringWorkspaceSettingObservation {
	if in == nil {
		return nil
	}
	out := new(EnhancedSecurityMonitoringWorkspaceSettingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedSecurityMonitoringWorkspaceSettingParameters) DeepCopyInto(out *EnhancedSecurityMonitoringWorkspaceSettingParameters) {
	*out = *in
	if in.EnhancedSecurityMonitoringWorkspace != nil {
		in, out := &in.EnhancedSecurityMonitoringWorkspace, &out.EnhancedSecurityMonitoringWorkspace
		*out = make([]EnhancedSecurityMonitoringWorkspaceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedSecurityMonitoringWorkspaceSettingParameters.
func (in *EnhancedSecurityMonitoringWorkspaceSettingParameters) DeepCopy() *EnhancedSecurityMonitoringWorkspaceSettingParameters {
	if in == nil {
		return nil
	}
	out := new(EnhancedSecurityMonitoringWorkspaceSettingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedSecurityMonitoringWorkspaceSettingSpec) DeepCopyInto(out *EnhancedSecurityMonitoringWorkspaceSettingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedSecurityMonitoringWorkspaceSettingSpec.
func (in *EnhancedSecurityMonitoringWorkspaceSettingSpec) DeepCopy() *EnhancedSecurityMonitoringWorkspaceSettingSpec {
	if in == nil {
		return nil
	}
	out := new(EnhancedSecurityMonitoringWorkspaceSettingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedSecurityMonitoringWorkspaceSettingStatus) DeepCopyInto(out *EnhancedSecurityMonitoringWorkspaceSettingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedSecurityMonitoringWorkspaceSettingStatus.
func (in *EnhancedSecurityMonitoringWorkspaceSettingStatus) DeepCopy() *EnhancedSecurityMonitoringWorkspaceSettingStatus {
	if in == nil {
		return nil
	}
	out := new(EnhancedSecurityMonitoringWorkspaceSettingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceInitParameters) DeepCopyInto(out *NamespaceInitParameters) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceInitParameters.
func (in *NamespaceInitParameters) DeepCopy() *NamespaceInitParameters {
	if in == nil {
		return nil
	}
	out := new(NamespaceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceObservation) DeepCopyInto(out *NamespaceObservation) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceObservation.
func (in *NamespaceObservation) DeepCopy() *NamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(NamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceParameters) DeepCopyInto(out *NamespaceParameters) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceParameters.
func (in *NamespaceParameters) DeepCopy() *NamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(NamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictWorkspaceAdminsInitParameters) DeepCopyInto(out *RestrictWorkspaceAdminsInitParameters) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictWorkspaceAdminsInitParameters.
func (in *RestrictWorkspaceAdminsInitParameters) DeepCopy() *RestrictWorkspaceAdminsInitParameters {
	if in == nil {
		return nil
	}
	out := new(RestrictWorkspaceAdminsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictWorkspaceAdminsObservation) DeepCopyInto(out *RestrictWorkspaceAdminsObservation) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictWorkspaceAdminsObservation.
func (in *RestrictWorkspaceAdminsObservation) DeepCopy() *RestrictWorkspaceAdminsObservation {
	if in == nil {
		return nil
	}
	out := new(RestrictWorkspaceAdminsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictWorkspaceAdminsParameters) DeepCopyInto(out *RestrictWorkspaceAdminsParameters) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictWorkspaceAdminsParameters.
func (in *RestrictWorkspaceAdminsParameters) DeepCopy() *RestrictWorkspaceAdminsParameters {
	if in == nil {
		return nil
	}
	out := new(RestrictWorkspaceAdminsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictWorkspaceAdminsSetting) DeepCopyInto(out *RestrictWorkspaceAdminsSetting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictWorkspaceAdminsSetting.
func (in *RestrictWorkspaceAdminsSetting) DeepCopy() *RestrictWorkspaceAdminsSetting {
	if in == nil {
		return nil
	}
	out := new(RestrictWorkspaceAdminsSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RestrictWorkspaceAdminsSetting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictWorkspaceAdminsSettingInitParameters) DeepCopyInto(out *RestrictWorkspaceAdminsSettingInitParameters) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.RestrictWorkspaceAdmins != nil {
		in, out := &in.RestrictWorkspaceAdmins, &out.RestrictWorkspaceAdmins
		*out = make([]RestrictWorkspaceAdminsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictWorkspaceAdminsSettingInitParameters.
func (in *RestrictWorkspaceAdminsSettingInitParameters) DeepCopy() *RestrictWorkspaceAdminsSettingInitParameters {
	if in == nil {
		return nil
	}
	out := new(RestrictWorkspaceAdminsSettingInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictWorkspaceAdminsSettingList) DeepCopyInto(out *RestrictWorkspaceAdminsSettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RestrictWorkspaceAdminsSetting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictWorkspaceAdminsSettingList.
func (in *RestrictWorkspaceAdminsSettingList) DeepCopy() *RestrictWorkspaceAdminsSettingList {
	if in == nil {
		return nil
	}
	out := new(RestrictWorkspaceAdminsSettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RestrictWorkspaceAdminsSettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictWorkspaceAdminsSettingObservation) DeepCopyInto(out *RestrictWorkspaceAdminsSettingObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RestrictWorkspaceAdmins != nil {
		in, out := &in.RestrictWorkspaceAdmins, &out.RestrictWorkspaceAdmins
		*out = make([]RestrictWorkspaceAdminsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictWorkspaceAdminsSettingObservation.
func (in *RestrictWorkspaceAdminsSettingObservation) DeepCopy() *RestrictWorkspaceAdminsSettingObservation {
	if in == nil {
		return nil
	}
	out := new(RestrictWorkspaceAdminsSettingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictWorkspaceAdminsSettingParameters) DeepCopyInto(out *RestrictWorkspaceAdminsSettingParameters) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.RestrictWorkspaceAdmins != nil {
		in, out := &in.RestrictWorkspaceAdmins, &out.RestrictWorkspaceAdmins
		*out = make([]RestrictWorkspaceAdminsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SettingName != nil {
		in, out := &in.SettingName, &out.SettingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictWorkspaceAdminsSettingParameters.
func (in *RestrictWorkspaceAdminsSettingParameters) DeepCopy() *RestrictWorkspaceAdminsSettingParameters {
	if in == nil {
		return nil
	}
	out := new(RestrictWorkspaceAdminsSettingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictWorkspaceAdminsSettingSpec) DeepCopyInto(out *RestrictWorkspaceAdminsSettingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictWorkspaceAdminsSettingSpec.
func (in *RestrictWorkspaceAdminsSettingSpec) DeepCopy() *RestrictWorkspaceAdminsSettingSpec {
	if in == nil {
		return nil
	}
	out := new(RestrictWorkspaceAdminsSettingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictWorkspaceAdminsSettingStatus) DeepCopyInto(out *RestrictWorkspaceAdminsSettingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictWorkspaceAdminsSettingStatus.
func (in *RestrictWorkspaceAdminsSettingStatus) DeepCopy() *RestrictWorkspaceAdminsSettingStatus {
	if in == nil {
		return nil
	}
	out := new(RestrictWorkspaceAdminsSettingStatus)
	in.DeepCopyInto(out)
	return out
}
