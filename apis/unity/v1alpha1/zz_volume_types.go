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

type VolumeInitParameters struct {

	// Name of parent Catalog. Change forces creation of a new resource.
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Free-form text.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Name of the Volume
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the volume owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Name of parent Schema relative to parent Catalog. Change forces creation of a new resource.
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// Path inside an External Location. Only used for EXTERNAL Volumes. Change forces creation of a new resource.
	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// Volume type. EXTERNAL or MANAGED. Change forces creation of a new resource.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type VolumeObservation struct {

	// Name of parent Catalog. Change forces creation of a new resource.
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Free-form text.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// ID of this Unity Catalog Volume in form of <catalog>.<schema>.<name>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Volume
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the volume owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Name of parent Schema relative to parent Catalog. Change forces creation of a new resource.
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// Path inside an External Location. Only used for EXTERNAL Volumes. Change forces creation of a new resource.
	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// base file path for this Unity Catalog Volume in form of /Volumes/<catalog>/<schema>/<name>.
	VolumePath *string `json:"volumePath,omitempty" tf:"volume_path,omitempty"`

	// Volume type. EXTERNAL or MANAGED. Change forces creation of a new resource.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type VolumeParameters struct {

	// Name of parent Catalog. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Free-form text.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Name of the Volume
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the volume owner.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Name of parent Schema relative to parent Catalog. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// Path inside an External Location. Only used for EXTERNAL Volumes. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// Volume type. EXTERNAL or MANAGED. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

// VolumeSpec defines the desired state of Volume
type VolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeParameters `json:"forProvider"`
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
	InitProvider VolumeInitParameters `json:"initProvider,omitempty"`
}

// VolumeStatus defines the observed state of Volume.
type VolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Volume is the Schema for the Volumes API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type Volume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.catalogName) || (has(self.initProvider) && has(self.initProvider.catalogName))",message="spec.forProvider.catalogName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schemaName) || (has(self.initProvider) && has(self.initProvider.schemaName))",message="spec.forProvider.schemaName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.volumeType) || (has(self.initProvider) && has(self.initProvider.volumeType))",message="spec.forProvider.volumeType is a required parameter"
	Spec   VolumeSpec   `json:"spec"`
	Status VolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeList contains a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Volume `json:"items"`
}

// Repository type metadata.
var (
	Volume_Kind             = "Volume"
	Volume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Volume_Kind}.String()
	Volume_KindAPIVersion   = Volume_Kind + "." + CRDGroupVersion.String()
	Volume_GroupVersionKind = CRDGroupVersion.WithKind(Volume_Kind)
)

func init() {
	SchemeBuilder.Register(&Volume{}, &VolumeList{})
}
