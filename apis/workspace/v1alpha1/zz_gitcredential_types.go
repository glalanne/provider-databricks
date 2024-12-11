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

type GitCredentialInitParameters struct {

	// specify if settings need to be enforced - right now, Databricks allows only single Git credential, so if it's already configured, the apply operation will fail.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// case insensitive name of the Git provider.  Following values are supported right now (could be a subject for a change, consult Git Credentials API documentation): gitHub, gitHubEnterprise, bitbucketCloud, bitbucketServer, azureDevOpsServices, gitLab, gitLabEnterpriseEdition, awsCodeCommit.
	GitProvider *string `json:"gitProvider,omitempty" tf:"git_provider,omitempty"`

	// user name at Git provider.
	GitUsername *string `json:"gitUsername,omitempty" tf:"git_username,omitempty"`

	// The personal access token used to authenticate to the corresponding Git provider. If value is not provided, it's sourced from the first environment variable of GITHUB_TOKEN, GITLAB_TOKEN, or AZDO_PERSONAL_ACCESS_TOKEN, that has a non-empty value.
	PersonalAccessToken *string `json:"personalAccessToken,omitempty" tf:"personal_access_token,omitempty"`
}

type GitCredentialObservation struct {

	// specify if settings need to be enforced - right now, Databricks allows only single Git credential, so if it's already configured, the apply operation will fail.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// case insensitive name of the Git provider.  Following values are supported right now (could be a subject for a change, consult Git Credentials API documentation): gitHub, gitHubEnterprise, bitbucketCloud, bitbucketServer, azureDevOpsServices, gitLab, gitLabEnterpriseEdition, awsCodeCommit.
	GitProvider *string `json:"gitProvider,omitempty" tf:"git_provider,omitempty"`

	// user name at Git provider.
	GitUsername *string `json:"gitUsername,omitempty" tf:"git_username,omitempty"`

	// identifier of specific Git credential
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The personal access token used to authenticate to the corresponding Git provider. If value is not provided, it's sourced from the first environment variable of GITHUB_TOKEN, GITLAB_TOKEN, or AZDO_PERSONAL_ACCESS_TOKEN, that has a non-empty value.
	PersonalAccessToken *string `json:"personalAccessToken,omitempty" tf:"personal_access_token,omitempty"`
}

type GitCredentialParameters struct {

	// specify if settings need to be enforced - right now, Databricks allows only single Git credential, so if it's already configured, the apply operation will fail.
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// case insensitive name of the Git provider.  Following values are supported right now (could be a subject for a change, consult Git Credentials API documentation): gitHub, gitHubEnterprise, bitbucketCloud, bitbucketServer, azureDevOpsServices, gitLab, gitLabEnterpriseEdition, awsCodeCommit.
	// +kubebuilder:validation:Optional
	GitProvider *string `json:"gitProvider,omitempty" tf:"git_provider,omitempty"`

	// user name at Git provider.
	// +kubebuilder:validation:Optional
	GitUsername *string `json:"gitUsername,omitempty" tf:"git_username,omitempty"`

	// The personal access token used to authenticate to the corresponding Git provider. If value is not provided, it's sourced from the first environment variable of GITHUB_TOKEN, GITLAB_TOKEN, or AZDO_PERSONAL_ACCESS_TOKEN, that has a non-empty value.
	// +kubebuilder:validation:Optional
	PersonalAccessToken *string `json:"personalAccessToken,omitempty" tf:"personal_access_token,omitempty"`
}

// GitCredentialSpec defines the desired state of GitCredential
type GitCredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GitCredentialParameters `json:"forProvider"`
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
	InitProvider GitCredentialInitParameters `json:"initProvider,omitempty"`
}

// GitCredentialStatus defines the observed state of GitCredential.
type GitCredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GitCredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GitCredential is the Schema for the GitCredentials API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type GitCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.gitProvider) || (has(self.initProvider) && has(self.initProvider.gitProvider))",message="spec.forProvider.gitProvider is a required parameter"
	Spec   GitCredentialSpec   `json:"spec"`
	Status GitCredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GitCredentialList contains a list of GitCredentials
type GitCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitCredential `json:"items"`
}

// Repository type metadata.
var (
	GitCredential_Kind             = "GitCredential"
	GitCredential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GitCredential_Kind}.String()
	GitCredential_KindAPIVersion   = GitCredential_Kind + "." + CRDGroupVersion.String()
	GitCredential_GroupVersionKind = CRDGroupVersion.WithKind(GitCredential_Kind)
)

func init() {
	SchemeBuilder.Register(&GitCredential{}, &GitCredentialList{})
}
