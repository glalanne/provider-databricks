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

type MlflowModelInitParameters struct {

	// The description of the MLflow model.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of MLflow model. Change of name triggers new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Tags for the MLflow model.
	Tags []MlflowModelTagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MlflowModelObservation struct {

	// The description of the MLflow model.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the MLflow model, the same as name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of MLflow model. Change of name triggers new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the MLflow model, the same as name.
	RegisteredModelID *string `json:"registeredModelId,omitempty" tf:"registered_model_id,omitempty"`

	// Tags for the MLflow model.
	Tags []MlflowModelTagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MlflowModelParameters struct {

	// The description of the MLflow model.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of MLflow model. Change of name triggers new resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Tags for the MLflow model.
	// +kubebuilder:validation:Optional
	Tags []MlflowModelTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MlflowModelTagsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MlflowModelTagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MlflowModelTagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// MlflowModelSpec defines the desired state of MlflowModel
type MlflowModelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MlflowModelParameters `json:"forProvider"`
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
	InitProvider MlflowModelInitParameters `json:"initProvider,omitempty"`
}

// MlflowModelStatus defines the observed state of MlflowModel.
type MlflowModelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MlflowModelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MlflowModel is the Schema for the MlflowModels API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type MlflowModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   MlflowModelSpec   `json:"spec"`
	Status MlflowModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MlflowModelList contains a list of MlflowModels
type MlflowModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MlflowModel `json:"items"`
}

// Repository type metadata.
var (
	MlflowModel_Kind             = "MlflowModel"
	MlflowModel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MlflowModel_Kind}.String()
	MlflowModel_KindAPIVersion   = MlflowModel_Kind + "." + CRDGroupVersion.String()
	MlflowModel_GroupVersionKind = CRDGroupVersion.WithKind(MlflowModel_Kind)
)

func init() {
	SchemeBuilder.Register(&MlflowModel{}, &MlflowModelList{})
}
