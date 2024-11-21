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

type IPAccessListInitParameters struct {

	// Boolean true or false indicating whether this list should be active.  Defaults to true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A string list of IP addresses and CIDR ranges.
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// This is the display name for the given IP ACL List.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Can only be "ALLOW" or "BLOCK".
	ListType *string `json:"listType,omitempty" tf:"list_type,omitempty"`
}

type IPAccessListObservation struct {

	// Boolean true or false indicating whether this list should be active.  Defaults to true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Canonical unique identifier for the IP Access List, same as list_id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A string list of IP addresses and CIDR ranges.
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// This is the display name for the given IP ACL List.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Can only be "ALLOW" or "BLOCK".
	ListType *string `json:"listType,omitempty" tf:"list_type,omitempty"`
}

type IPAccessListParameters struct {

	// Boolean true or false indicating whether this list should be active.  Defaults to true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A string list of IP addresses and CIDR ranges.
	// +kubebuilder:validation:Optional
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// This is the display name for the given IP ACL List.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Can only be "ALLOW" or "BLOCK".
	// +kubebuilder:validation:Optional
	ListType *string `json:"listType,omitempty" tf:"list_type,omitempty"`
}

// IPAccessListSpec defines the desired state of IPAccessList
type IPAccessListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPAccessListParameters `json:"forProvider"`
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
	InitProvider IPAccessListInitParameters `json:"initProvider,omitempty"`
}

// IPAccessListStatus defines the observed state of IPAccessList.
type IPAccessListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPAccessListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IPAccessList is the Schema for the IPAccessLists API. ""subcategory: "Security"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type IPAccessList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipAddresses) || (has(self.initProvider) && has(self.initProvider.ipAddresses))",message="spec.forProvider.ipAddresses is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.label) || (has(self.initProvider) && has(self.initProvider.label))",message="spec.forProvider.label is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.listType) || (has(self.initProvider) && has(self.initProvider.listType))",message="spec.forProvider.listType is a required parameter"
	Spec   IPAccessListSpec   `json:"spec"`
	Status IPAccessListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPAccessListList contains a list of IPAccessLists
type IPAccessListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPAccessList `json:"items"`
}

// Repository type metadata.
var (
	IPAccessList_Kind             = "IPAccessList"
	IPAccessList_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPAccessList_Kind}.String()
	IPAccessList_KindAPIVersion   = IPAccessList_Kind + "." + CRDGroupVersion.String()
	IPAccessList_GroupVersionKind = CRDGroupVersion.WithKind(IPAccessList_Kind)
)

func init() {
	SchemeBuilder.Register(&IPAccessList{}, &IPAccessListList{})
}
