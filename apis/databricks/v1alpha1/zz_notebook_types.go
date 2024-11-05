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

type NotebookInitParameters struct {
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	Md5 *string `json:"md5,omitempty" tf:"md5,omitempty"`

	ObjectID *float64 `json:"objectId,omitempty" tf:"object_id,omitempty"`

	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type NotebookObservation struct {
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	Md5 *string `json:"md5,omitempty" tf:"md5,omitempty"`

	ObjectID *float64 `json:"objectId,omitempty" tf:"object_id,omitempty"`

	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	WorkspacePath *string `json:"workspacePath,omitempty" tf:"workspace_path,omitempty"`
}

type NotebookParameters struct {

	// +kubebuilder:validation:Optional
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// +kubebuilder:validation:Optional
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// +kubebuilder:validation:Optional
	Md5 *string `json:"md5,omitempty" tf:"md5,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectID *float64 `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

// NotebookSpec defines the desired state of Notebook
type NotebookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotebookParameters `json:"forProvider"`
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
	InitProvider NotebookInitParameters `json:"initProvider,omitempty"`
}

// NotebookStatus defines the observed state of Notebook.
type NotebookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotebookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Notebook is the Schema for the Notebooks API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type Notebook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || (has(self.initProvider) && has(self.initProvider.path))",message="spec.forProvider.path is a required parameter"
	Spec   NotebookSpec   `json:"spec"`
	Status NotebookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotebookList contains a list of Notebooks
type NotebookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notebook `json:"items"`
}

// Repository type metadata.
var (
	Notebook_Kind             = "Notebook"
	Notebook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Notebook_Kind}.String()
	Notebook_KindAPIVersion   = Notebook_Kind + "." + CRDGroupVersion.String()
	Notebook_GroupVersionKind = CRDGroupVersion.WithKind(Notebook_Kind)
)

func init() {
	SchemeBuilder.Register(&Notebook{}, &NotebookList{})
}
