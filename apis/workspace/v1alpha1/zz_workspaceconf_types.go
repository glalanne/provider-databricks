// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type WorkspaceConfInitParameters struct {

	// Key-value map of strings that represent workspace configuration. Upon resource deletion, properties that start with enable or enforce will be reset to false value, regardless of initial default one.
	// +mapType=granular
	CustomConfig map[string]*string `json:"customConfig,omitempty" tf:"custom_config,omitempty"`
}

type WorkspaceConfObservation struct {

	// Key-value map of strings that represent workspace configuration. Upon resource deletion, properties that start with enable or enforce will be reset to false value, regardless of initial default one.
	// +mapType=granular
	CustomConfig map[string]*string `json:"customConfig,omitempty" tf:"custom_config,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WorkspaceConfParameters struct {

	// Key-value map of strings that represent workspace configuration. Upon resource deletion, properties that start with enable or enforce will be reset to false value, regardless of initial default one.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	CustomConfig map[string]*string `json:"customConfig,omitempty" tf:"custom_config,omitempty"`
}

// WorkspaceConfSpec defines the desired state of WorkspaceConf
type WorkspaceConfSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceConfParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider WorkspaceConfInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceConfStatus defines the observed state of WorkspaceConf.
type WorkspaceConfStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceConfObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WorkspaceConf is the Schema for the WorkspaceConfs API. ""subcategory: "Workspace"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,}
type WorkspaceConf struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceConfSpec   `json:"spec"`
	Status            WorkspaceConfStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceConfList contains a list of WorkspaceConfs
type WorkspaceConfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceConf `json:"items"`
}

// Repository type metadata.
var (
	WorkspaceConf_Kind             = "WorkspaceConf"
	WorkspaceConf_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkspaceConf_Kind}.String()
	WorkspaceConf_KindAPIVersion   = WorkspaceConf_Kind + "." + CRDGroupVersion.String()
	WorkspaceConf_GroupVersionKind = CRDGroupVersion.WithKind(WorkspaceConf_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkspaceConf{}, &WorkspaceConfList{})
}