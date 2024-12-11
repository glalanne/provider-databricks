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

type ClusterPolicyInitParameters struct {

	// Policy definition: JSON document expressed in Databricks Policy Definition Language. Cannot be used with policy_family_id
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	// Additional human-readable description of the cluster policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Libraries []LibrariesInitParameters `json:"libraries,omitempty" tf:"libraries,omitempty"`

	// Maximum number of clusters allowed per user. When omitted, there is no limit. If specified, value must be greater than zero.
	MaxClustersPerUser *float64 `json:"maxClustersPerUser,omitempty" tf:"max_clusters_per_user,omitempty"`

	// the name of the built-in cluster policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// settings to override in the built-in cluster policy.
	PolicyFamilyDefinitionOverrides *string `json:"policyFamilyDefinitionOverrides,omitempty" tf:"policy_family_definition_overrides,omitempty"`

	// the ID of the cluster policy family used for built-in cluster policy.
	PolicyFamilyID *string `json:"policyFamilyId,omitempty" tf:"policy_family_id,omitempty"`
}

type ClusterPolicyObservation struct {

	// Policy definition: JSON document expressed in Databricks Policy Definition Language. Cannot be used with policy_family_id
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	// Additional human-readable description of the cluster policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Canonical unique identifier for the cluster policy. This is equal to policy_id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Libraries []LibrariesObservation `json:"libraries,omitempty" tf:"libraries,omitempty"`

	// Maximum number of clusters allowed per user. When omitted, there is no limit. If specified, value must be greater than zero.
	MaxClustersPerUser *float64 `json:"maxClustersPerUser,omitempty" tf:"max_clusters_per_user,omitempty"`

	// the name of the built-in cluster policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// settings to override in the built-in cluster policy.
	PolicyFamilyDefinitionOverrides *string `json:"policyFamilyDefinitionOverrides,omitempty" tf:"policy_family_definition_overrides,omitempty"`

	// the ID of the cluster policy family used for built-in cluster policy.
	PolicyFamilyID *string `json:"policyFamilyId,omitempty" tf:"policy_family_id,omitempty"`

	// Canonical unique identifier for the cluster policy.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`
}

type ClusterPolicyParameters struct {

	// Policy definition: JSON document expressed in Databricks Policy Definition Language. Cannot be used with policy_family_id
	// +kubebuilder:validation:Optional
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	// Additional human-readable description of the cluster policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Libraries []LibrariesParameters `json:"libraries,omitempty" tf:"libraries,omitempty"`

	// Maximum number of clusters allowed per user. When omitted, there is no limit. If specified, value must be greater than zero.
	// +kubebuilder:validation:Optional
	MaxClustersPerUser *float64 `json:"maxClustersPerUser,omitempty" tf:"max_clusters_per_user,omitempty"`

	// the name of the built-in cluster policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// settings to override in the built-in cluster policy.
	// +kubebuilder:validation:Optional
	PolicyFamilyDefinitionOverrides *string `json:"policyFamilyDefinitionOverrides,omitempty" tf:"policy_family_definition_overrides,omitempty"`

	// the ID of the cluster policy family used for built-in cluster policy.
	// +kubebuilder:validation:Optional
	PolicyFamilyID *string `json:"policyFamilyId,omitempty" tf:"policy_family_id,omitempty"`
}

type LibrariesCranInitParameters struct {
	Package *string `json:"package,omitempty" tf:"package,omitempty"`

	Repo *string `json:"repo,omitempty" tf:"repo,omitempty"`
}

type LibrariesCranObservation struct {
	Package *string `json:"package,omitempty" tf:"package,omitempty"`

	Repo *string `json:"repo,omitempty" tf:"repo,omitempty"`
}

type LibrariesCranParameters struct {

	// +kubebuilder:validation:Optional
	Package *string `json:"package" tf:"package,omitempty"`

	// +kubebuilder:validation:Optional
	Repo *string `json:"repo,omitempty" tf:"repo,omitempty"`
}

type LibrariesInitParameters struct {
	Cran []LibrariesCranInitParameters `json:"cran,omitempty" tf:"cran,omitempty"`

	Egg *string `json:"egg,omitempty" tf:"egg,omitempty"`

	Jar *string `json:"jar,omitempty" tf:"jar,omitempty"`

	Maven []LibrariesMavenInitParameters `json:"maven,omitempty" tf:"maven,omitempty"`

	Pypi []LibrariesPypiInitParameters `json:"pypi,omitempty" tf:"pypi,omitempty"`

	Requirements *string `json:"requirements,omitempty" tf:"requirements,omitempty"`

	Whl *string `json:"whl,omitempty" tf:"whl,omitempty"`
}

type LibrariesMavenInitParameters struct {
	Coordinates *string `json:"coordinates,omitempty" tf:"coordinates,omitempty"`

	Exclusions []*string `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	Repo *string `json:"repo,omitempty" tf:"repo,omitempty"`
}

type LibrariesMavenObservation struct {
	Coordinates *string `json:"coordinates,omitempty" tf:"coordinates,omitempty"`

	Exclusions []*string `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	Repo *string `json:"repo,omitempty" tf:"repo,omitempty"`
}

type LibrariesMavenParameters struct {

	// +kubebuilder:validation:Optional
	Coordinates *string `json:"coordinates" tf:"coordinates,omitempty"`

	// +kubebuilder:validation:Optional
	Exclusions []*string `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// +kubebuilder:validation:Optional
	Repo *string `json:"repo,omitempty" tf:"repo,omitempty"`
}

type LibrariesObservation struct {
	Cran []LibrariesCranObservation `json:"cran,omitempty" tf:"cran,omitempty"`

	Egg *string `json:"egg,omitempty" tf:"egg,omitempty"`

	Jar *string `json:"jar,omitempty" tf:"jar,omitempty"`

	Maven []LibrariesMavenObservation `json:"maven,omitempty" tf:"maven,omitempty"`

	Pypi []LibrariesPypiObservation `json:"pypi,omitempty" tf:"pypi,omitempty"`

	Requirements *string `json:"requirements,omitempty" tf:"requirements,omitempty"`

	Whl *string `json:"whl,omitempty" tf:"whl,omitempty"`
}

type LibrariesParameters struct {

	// +kubebuilder:validation:Optional
	Cran []LibrariesCranParameters `json:"cran,omitempty" tf:"cran,omitempty"`

	// +kubebuilder:validation:Optional
	Egg *string `json:"egg,omitempty" tf:"egg,omitempty"`

	// +kubebuilder:validation:Optional
	Jar *string `json:"jar,omitempty" tf:"jar,omitempty"`

	// +kubebuilder:validation:Optional
	Maven []LibrariesMavenParameters `json:"maven,omitempty" tf:"maven,omitempty"`

	// +kubebuilder:validation:Optional
	Pypi []LibrariesPypiParameters `json:"pypi,omitempty" tf:"pypi,omitempty"`

	// +kubebuilder:validation:Optional
	Requirements *string `json:"requirements,omitempty" tf:"requirements,omitempty"`

	// +kubebuilder:validation:Optional
	Whl *string `json:"whl,omitempty" tf:"whl,omitempty"`
}

type LibrariesPypiInitParameters struct {
	Package *string `json:"package,omitempty" tf:"package,omitempty"`

	Repo *string `json:"repo,omitempty" tf:"repo,omitempty"`
}

type LibrariesPypiObservation struct {
	Package *string `json:"package,omitempty" tf:"package,omitempty"`

	Repo *string `json:"repo,omitempty" tf:"repo,omitempty"`
}

type LibrariesPypiParameters struct {

	// +kubebuilder:validation:Optional
	Package *string `json:"package" tf:"package,omitempty"`

	// +kubebuilder:validation:Optional
	Repo *string `json:"repo,omitempty" tf:"repo,omitempty"`
}

// ClusterPolicySpec defines the desired state of ClusterPolicy
type ClusterPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterPolicyParameters `json:"forProvider"`
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
	InitProvider ClusterPolicyInitParameters `json:"initProvider,omitempty"`
}

// ClusterPolicyStatus defines the observed state of ClusterPolicy.
type ClusterPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClusterPolicy is the Schema for the ClusterPolicys API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type ClusterPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterPolicySpec   `json:"spec"`
	Status            ClusterPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterPolicyList contains a list of ClusterPolicys
type ClusterPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterPolicy `json:"items"`
}

// Repository type metadata.
var (
	ClusterPolicy_Kind             = "ClusterPolicy"
	ClusterPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterPolicy_Kind}.String()
	ClusterPolicy_KindAPIVersion   = ClusterPolicy_Kind + "." + CRDGroupVersion.String()
	ClusterPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ClusterPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterPolicy{}, &ClusterPolicyList{})
}
