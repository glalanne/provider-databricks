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

type FileInitParameters struct {

	// Contents in base 64 format. Conflicts with source.
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	Md5 *string `json:"md5,omitempty" tf:"md5,omitempty"`

	// The path of the file in which you wish to save. For example, /Volumes/main/default/volume1/file.txt.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	RemoteFileModified *bool `json:"remoteFileModified,omitempty" tf:"remote_file_modified,omitempty"`

	// The full absolute path to the file. Conflicts with content_base64.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type FileObservation struct {

	// Contents in base 64 format. Conflicts with source.
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// The file size of the file that is being tracked by this resource in bytes.
	FileSize *float64 `json:"fileSize,omitempty" tf:"file_size,omitempty"`

	// Same as path.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Md5 *string `json:"md5,omitempty" tf:"md5,omitempty"`

	// The last time stamp when the file was modified
	ModificationTime *string `json:"modificationTime,omitempty" tf:"modification_time,omitempty"`

	// The path of the file in which you wish to save. For example, /Volumes/main/default/volume1/file.txt.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	RemoteFileModified *bool `json:"remoteFileModified,omitempty" tf:"remote_file_modified,omitempty"`

	// The full absolute path to the file. Conflicts with content_base64.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type FileParameters struct {

	// Contents in base 64 format. Conflicts with source.
	// +kubebuilder:validation:Optional
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// +kubebuilder:validation:Optional
	Md5 *string `json:"md5,omitempty" tf:"md5,omitempty"`

	// The path of the file in which you wish to save. For example, /Volumes/main/default/volume1/file.txt.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	RemoteFileModified *bool `json:"remoteFileModified,omitempty" tf:"remote_file_modified,omitempty"`

	// The full absolute path to the file. Conflicts with content_base64.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

// FileSpec defines the desired state of File
type FileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FileParameters `json:"forProvider"`
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
	InitProvider FileInitParameters `json:"initProvider,omitempty"`
}

// FileStatus defines the observed state of File.
type FileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// File is the Schema for the Files API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type File struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || (has(self.initProvider) && has(self.initProvider.path))",message="spec.forProvider.path is a required parameter"
	Spec   FileSpec   `json:"spec"`
	Status FileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileList contains a list of Files
type FileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []File `json:"items"`
}

// Repository type metadata.
var (
	File_Kind             = "File"
	File_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: File_Kind}.String()
	File_KindAPIVersion   = File_Kind + "." + CRDGroupVersion.String()
	File_GroupVersionKind = CRDGroupVersion.WithKind(File_Kind)
)

func init() {
	SchemeBuilder.Register(&File{}, &FileList{})
}
