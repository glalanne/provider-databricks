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

type PermissionAssignmentInitParameters struct {

	// The list of workspace permissions to assign to the principal:
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Databricks ID of the user, service principal, or group. The principal ID can be retrieved using the account-level SCIM API, or using databricks_user, databricks_service_principal or databricks_group data sources with account API (and has to be an account admin).
	PrincipalID *float64 `json:"principalId,omitempty" tf:"principal_id,omitempty"`
}

type PermissionAssignmentObservation struct {

	// ID of the permission assignment - same as principal_id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The list of workspace permissions to assign to the principal:
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Databricks ID of the user, service principal, or group. The principal ID can be retrieved using the account-level SCIM API, or using databricks_user, databricks_service_principal or databricks_group data sources with account API (and has to be an account admin).
	PrincipalID *float64 `json:"principalId,omitempty" tf:"principal_id,omitempty"`
}

type PermissionAssignmentParameters struct {

	// The list of workspace permissions to assign to the principal:
	// +kubebuilder:validation:Optional
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Databricks ID of the user, service principal, or group. The principal ID can be retrieved using the account-level SCIM API, or using databricks_user, databricks_service_principal or databricks_group data sources with account API (and has to be an account admin).
	// +kubebuilder:validation:Optional
	PrincipalID *float64 `json:"principalId,omitempty" tf:"principal_id,omitempty"`
}

// PermissionAssignmentSpec defines the desired state of PermissionAssignment
type PermissionAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionAssignmentParameters `json:"forProvider"`
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
	InitProvider PermissionAssignmentInitParameters `json:"initProvider,omitempty"`
}

// PermissionAssignmentStatus defines the observed state of PermissionAssignment.
type PermissionAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PermissionAssignment is the Schema for the PermissionAssignments API. ""subcategory: "Security"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type PermissionAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permissions) || (has(self.initProvider) && has(self.initProvider.permissions))",message="spec.forProvider.permissions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principalId) || (has(self.initProvider) && has(self.initProvider.principalId))",message="spec.forProvider.principalId is a required parameter"
	Spec   PermissionAssignmentSpec   `json:"spec"`
	Status PermissionAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionAssignmentList contains a list of PermissionAssignments
type PermissionAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PermissionAssignment `json:"items"`
}

// Repository type metadata.
var (
	PermissionAssignment_Kind             = "PermissionAssignment"
	PermissionAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PermissionAssignment_Kind}.String()
	PermissionAssignment_KindAPIVersion   = PermissionAssignment_Kind + "." + CRDGroupVersion.String()
	PermissionAssignment_GroupVersionKind = CRDGroupVersion.WithKind(PermissionAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&PermissionAssignment{}, &PermissionAssignmentList{})
}
