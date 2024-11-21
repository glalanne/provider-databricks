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

type ObjectInitParameters struct {
	AddedAt *float64 `json:"addedAt,omitempty" tf:"added_at,omitempty"`

	AddedBy *string `json:"addedBy,omitempty" tf:"added_by,omitempty"`

	// Whether to enable Change Data Feed (cdf) on the shared object. When this field is set, field history_data_sharing_status can not be set.
	CdfEnabled *bool `json:"cdfEnabled,omitempty" tf:"cdf_enabled,omitempty"`

	// Description about the object.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Type of the data object, currently TABLE, SCHEMA, VOLUME, and MODEL are supported.
	DataObjectType *string `json:"dataObjectType,omitempty" tf:"data_object_type,omitempty"`

	// Whether to enable history sharing, one of: ENABLED, DISABLED. When a table has history sharing enabled, recipients can query table data by version, starting from the current table version. If not specified, clients can only query starting from the version of the object at the time it was added to the share. NOTE: The start_version should be less than or equal the current version of the object. When this field is set, field cdf_enabled can not be set.
	HistoryDataSharingStatus *string `json:"historyDataSharingStatus,omitempty" tf:"history_data_sharing_status,omitempty"`

	// Name of share. Change forces creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Partition []PartitionInitParameters `json:"partition,omitempty" tf:"partition,omitempty"`

	// A user-provided new name for the data object within the share. If this new name is not provided, the object's original name will be used as the shared_as name. The shared_as name must be unique within a Share. Change forces creation of a new resource.
	SharedAs *string `json:"sharedAs,omitempty" tf:"shared_as,omitempty"`

	// The start version associated with the object for cdf. This allows data providers to control the lowest object version that is accessible by clients.
	StartVersion *float64 `json:"startVersion,omitempty" tf:"start_version,omitempty"`

	// Status of the object, one of: ACTIVE, PERMISSION_DENIED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A user-provided new name for the data object within the share. If this new name is not provided, the object's original name will be used as the shared_as name. The shared_as name must be unique within a Share. Change forces creation of a new resource.
	StringSharedAs *string `json:"stringSharedAs,omitempty" tf:"string_shared_as,omitempty"`
}

type ObjectObservation struct {
	AddedAt *float64 `json:"addedAt,omitempty" tf:"added_at,omitempty"`

	AddedBy *string `json:"addedBy,omitempty" tf:"added_by,omitempty"`

	// Whether to enable Change Data Feed (cdf) on the shared object. When this field is set, field history_data_sharing_status can not be set.
	CdfEnabled *bool `json:"cdfEnabled,omitempty" tf:"cdf_enabled,omitempty"`

	// Description about the object.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Type of the data object, currently TABLE, SCHEMA, VOLUME, and MODEL are supported.
	DataObjectType *string `json:"dataObjectType,omitempty" tf:"data_object_type,omitempty"`

	// Whether to enable history sharing, one of: ENABLED, DISABLED. When a table has history sharing enabled, recipients can query table data by version, starting from the current table version. If not specified, clients can only query starting from the version of the object at the time it was added to the share. NOTE: The start_version should be less than or equal the current version of the object. When this field is set, field cdf_enabled can not be set.
	HistoryDataSharingStatus *string `json:"historyDataSharingStatus,omitempty" tf:"history_data_sharing_status,omitempty"`

	// Name of share. Change forces creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Partition []PartitionObservation `json:"partition,omitempty" tf:"partition,omitempty"`

	// A user-provided new name for the data object within the share. If this new name is not provided, the object's original name will be used as the shared_as name. The shared_as name must be unique within a Share. Change forces creation of a new resource.
	SharedAs *string `json:"sharedAs,omitempty" tf:"shared_as,omitempty"`

	// The start version associated with the object for cdf. This allows data providers to control the lowest object version that is accessible by clients.
	StartVersion *float64 `json:"startVersion,omitempty" tf:"start_version,omitempty"`

	// Status of the object, one of: ACTIVE, PERMISSION_DENIED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A user-provided new name for the data object within the share. If this new name is not provided, the object's original name will be used as the shared_as name. The shared_as name must be unique within a Share. Change forces creation of a new resource.
	StringSharedAs *string `json:"stringSharedAs,omitempty" tf:"string_shared_as,omitempty"`
}

type ObjectParameters struct {

	// +kubebuilder:validation:Optional
	AddedAt *float64 `json:"addedAt,omitempty" tf:"added_at,omitempty"`

	// +kubebuilder:validation:Optional
	AddedBy *string `json:"addedBy,omitempty" tf:"added_by,omitempty"`

	// Whether to enable Change Data Feed (cdf) on the shared object. When this field is set, field history_data_sharing_status can not be set.
	// +kubebuilder:validation:Optional
	CdfEnabled *bool `json:"cdfEnabled,omitempty" tf:"cdf_enabled,omitempty"`

	// Description about the object.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Type of the data object, currently TABLE, SCHEMA, VOLUME, and MODEL are supported.
	// +kubebuilder:validation:Optional
	DataObjectType *string `json:"dataObjectType" tf:"data_object_type,omitempty"`

	// Whether to enable history sharing, one of: ENABLED, DISABLED. When a table has history sharing enabled, recipients can query table data by version, starting from the current table version. If not specified, clients can only query starting from the version of the object at the time it was added to the share. NOTE: The start_version should be less than or equal the current version of the object. When this field is set, field cdf_enabled can not be set.
	// +kubebuilder:validation:Optional
	HistoryDataSharingStatus *string `json:"historyDataSharingStatus,omitempty" tf:"history_data_sharing_status,omitempty"`

	// Name of share. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Partition []PartitionParameters `json:"partition,omitempty" tf:"partition,omitempty"`

	// A user-provided new name for the data object within the share. If this new name is not provided, the object's original name will be used as the shared_as name. The shared_as name must be unique within a Share. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	SharedAs *string `json:"sharedAs,omitempty" tf:"shared_as,omitempty"`

	// The start version associated with the object for cdf. This allows data providers to control the lowest object version that is accessible by clients.
	// +kubebuilder:validation:Optional
	StartVersion *float64 `json:"startVersion,omitempty" tf:"start_version,omitempty"`

	// Status of the object, one of: ACTIVE, PERMISSION_DENIED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A user-provided new name for the data object within the share. If this new name is not provided, the object's original name will be used as the shared_as name. The shared_as name must be unique within a Share. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	StringSharedAs *string `json:"stringSharedAs,omitempty" tf:"string_shared_as,omitempty"`
}

type PartitionInitParameters struct {

	// The value of the partition column. When this value is not set, it means null value. When this field is set, field recipient_property_key can not be set.
	Value []ValueInitParameters `json:"value,omitempty" tf:"value,omitempty"`
}

type PartitionObservation struct {

	// The value of the partition column. When this value is not set, it means null value. When this field is set, field recipient_property_key can not be set.
	Value []ValueObservation `json:"value,omitempty" tf:"value,omitempty"`
}

type PartitionParameters struct {

	// The value of the partition column. When this value is not set, it means null value. When this field is set, field recipient_property_key can not be set.
	// +kubebuilder:validation:Optional
	Value []ValueParameters `json:"value,omitempty" tf:"value,omitempty"`
}

type ShareInitParameters struct {

	// Description about the object.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Time when the share was created.
	CreatedAt *float64 `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The principal that created the share.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// Name of share. Change forces creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Object []ObjectInitParameters `json:"object,omitempty" tf:"object,omitempty"`

	// User name/group name/sp application_id of the share owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	StorageRoot *string `json:"storageRoot,omitempty" tf:"storage_root,omitempty"`

	UpdatedAt *float64 `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by,omitempty"`
}

type ShareObservation struct {

	// Description about the object.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Time when the share was created.
	CreatedAt *float64 `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The principal that created the share.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// the ID of the share, the same as name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of share. Change forces creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Object []ObjectObservation `json:"object,omitempty" tf:"object,omitempty"`

	// User name/group name/sp application_id of the share owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	StorageRoot *string `json:"storageRoot,omitempty" tf:"storage_root,omitempty"`

	UpdatedAt *float64 `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by,omitempty"`
}

type ShareParameters struct {

	// Description about the object.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Time when the share was created.
	// +kubebuilder:validation:Optional
	CreatedAt *float64 `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The principal that created the share.
	// +kubebuilder:validation:Optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// Name of share. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Object []ObjectParameters `json:"object,omitempty" tf:"object,omitempty"`

	// User name/group name/sp application_id of the share owner.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// +kubebuilder:validation:Optional
	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// +kubebuilder:validation:Optional
	StorageRoot *string `json:"storageRoot,omitempty" tf:"storage_root,omitempty"`

	// +kubebuilder:validation:Optional
	UpdatedAt *float64 `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// +kubebuilder:validation:Optional
	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by,omitempty"`
}

type ValueInitParameters struct {

	// Name of share. Change forces creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to apply for the value, one of: EQUAL, LIKE
	Op *string `json:"op,omitempty" tf:"op,omitempty"`

	// The key of a Delta Sharing recipient's property. For example databricks-account-id. When this field is set, field value can not be set.
	RecipientPropertyKey *string `json:"recipientPropertyKey,omitempty" tf:"recipient_property_key,omitempty"`

	// The value of the partition column. When this value is not set, it means null value. When this field is set, field recipient_property_key can not be set.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ValueObservation struct {

	// Name of share. Change forces creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to apply for the value, one of: EQUAL, LIKE
	Op *string `json:"op,omitempty" tf:"op,omitempty"`

	// The key of a Delta Sharing recipient's property. For example databricks-account-id. When this field is set, field value can not be set.
	RecipientPropertyKey *string `json:"recipientPropertyKey,omitempty" tf:"recipient_property_key,omitempty"`

	// The value of the partition column. When this value is not set, it means null value. When this field is set, field recipient_property_key can not be set.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ValueParameters struct {

	// Name of share. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to apply for the value, one of: EQUAL, LIKE
	// +kubebuilder:validation:Optional
	Op *string `json:"op" tf:"op,omitempty"`

	// The key of a Delta Sharing recipient's property. For example databricks-account-id. When this field is set, field value can not be set.
	// +kubebuilder:validation:Optional
	RecipientPropertyKey *string `json:"recipientPropertyKey,omitempty" tf:"recipient_property_key,omitempty"`

	// The value of the partition column. When this value is not set, it means null value. When this field is set, field recipient_property_key can not be set.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// ShareSpec defines the desired state of Share
type ShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ShareParameters `json:"forProvider"`
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
	InitProvider ShareInitParameters `json:"initProvider,omitempty"`
}

// ShareStatus defines the observed state of Share.
type ShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Share is the Schema for the Shares API. ""subcategory: "Delta Sharing"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,}
type Share struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ShareSpec   `json:"spec"`
	Status ShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShareList contains a list of Shares
type ShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Share `json:"items"`
}

// Repository type metadata.
var (
	Share_Kind             = "Share"
	Share_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Share_Kind}.String()
	Share_KindAPIVersion   = Share_Kind + "." + CRDGroupVersion.String()
	Share_GroupVersionKind = CRDGroupVersion.WithKind(Share_Kind)
)

func init() {
	SchemeBuilder.Register(&Share{}, &ShareList{})
}