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

type MwsStorageConfigurationsInitParameters struct {

	// Account Id that could be found in the top right corner of Accounts Console
	AccountIDSecretRef v1.SecretKeySelector `json:"accountIdSecretRef" tf:"-"`

	// name of AWS S3 bucket
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// name under which this storage configuration is stored
	StorageConfigurationName *string `json:"storageConfigurationName,omitempty" tf:"storage_configuration_name,omitempty"`
}

type MwsStorageConfigurationsObservation struct {

	// name of AWS S3 bucket
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	CreationTime *float64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// Canonical unique identifier for the mws storage configurations.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) id of storage config to be used for databricks_mws_workspace resource.
	StorageConfigurationID *string `json:"storageConfigurationId,omitempty" tf:"storage_configuration_id,omitempty"`

	// name under which this storage configuration is stored
	StorageConfigurationName *string `json:"storageConfigurationName,omitempty" tf:"storage_configuration_name,omitempty"`
}

type MwsStorageConfigurationsParameters struct {

	// Account Id that could be found in the top right corner of Accounts Console
	// +kubebuilder:validation:Optional
	AccountIDSecretRef v1.SecretKeySelector `json:"accountIdSecretRef" tf:"-"`

	// name of AWS S3 bucket
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// name under which this storage configuration is stored
	// +kubebuilder:validation:Optional
	StorageConfigurationName *string `json:"storageConfigurationName,omitempty" tf:"storage_configuration_name,omitempty"`
}

// MwsStorageConfigurationsSpec defines the desired state of MwsStorageConfigurations
type MwsStorageConfigurationsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MwsStorageConfigurationsParameters `json:"forProvider"`
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
	InitProvider MwsStorageConfigurationsInitParameters `json:"initProvider,omitempty"`
}

// MwsStorageConfigurationsStatus defines the observed state of MwsStorageConfigurations.
type MwsStorageConfigurationsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MwsStorageConfigurationsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MwsStorageConfigurations is the Schema for the MwsStorageConfigurationss API. ""subcategory: "Deployment"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type MwsStorageConfigurations struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountIdSecretRef)",message="spec.forProvider.accountIdSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucketName) || (has(self.initProvider) && has(self.initProvider.bucketName))",message="spec.forProvider.bucketName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageConfigurationName) || (has(self.initProvider) && has(self.initProvider.storageConfigurationName))",message="spec.forProvider.storageConfigurationName is a required parameter"
	Spec   MwsStorageConfigurationsSpec   `json:"spec"`
	Status MwsStorageConfigurationsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MwsStorageConfigurationsList contains a list of MwsStorageConfigurationss
type MwsStorageConfigurationsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MwsStorageConfigurations `json:"items"`
}

// Repository type metadata.
var (
	MwsStorageConfigurations_Kind             = "MwsStorageConfigurations"
	MwsStorageConfigurations_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MwsStorageConfigurations_Kind}.String()
	MwsStorageConfigurations_KindAPIVersion   = MwsStorageConfigurations_Kind + "." + CRDGroupVersion.String()
	MwsStorageConfigurations_GroupVersionKind = CRDGroupVersion.WithKind(MwsStorageConfigurations_Kind)
)

func init() {
	SchemeBuilder.Register(&MwsStorageConfigurations{}, &MwsStorageConfigurationsList{})
}
