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

type ServicePrincipalInitParameters struct {

	// identifier for use in databricks_access_control_rule_set, e.g. servicePrincipals/00000000-0000-0000-0000-000000000000.
	ACLPrincipalID *string `json:"aclPrincipalId,omitempty" tf:"acl_principal_id,omitempty"`

	// Either service principal is active or not. True by default, but can be set to false in case of service principal deactivation with preserving service principal assets.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Allow the service principal to have cluster create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and cluster_id argument. Everyone without allow_cluster_create argument set, but with permission to use Cluster Policy would be able to create clusters, but within the boundaries of that specific policy.
	AllowClusterCreate *bool `json:"allowClusterCreate,omitempty" tf:"allow_cluster_create,omitempty"`

	// Allow the service principal to have instance pool create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and instance_pool_id argument.
	AllowInstancePoolCreate *bool `json:"allowInstancePoolCreate,omitempty" tf:"allow_instance_pool_create,omitempty"`

	// managed service principals this value is auto-generated.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// This is a field to allow the group to have access to Databricks SQL feature through databricks_sql_endpoint.
	DatabricksSQLAccess *bool `json:"databricksSqlAccess,omitempty" tf:"databricks_sql_access,omitempty"`

	// Deactivate the service principal when deleting the resource, rather than deleting the service principal entirely. Defaults to true when the provider is configured at the account-level and false when configured at the workspace-level. This flag is exclusive to force_delete_repos and force_delete_home_dir flags.
	DisableAsUserDeletion *bool `json:"disableAsUserDeletion,omitempty" tf:"disable_as_user_deletion,omitempty"`

	// This is an alias for the service principal and can be the full name of the service principal.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// ID of the service principal in an external identity provider.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// This functionality is experimental and is designed to simplify corner cases, like Azure Active Directory synchronisation.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// This flag determines whether the service principal's home directory is deleted when the user is deleted. It will have no impact when in the accounts SCIM API. False by default.
	ForceDeleteHomeDir *bool `json:"forceDeleteHomeDir,omitempty" tf:"force_delete_home_dir,omitempty"`

	// This flag determines whether the service principal's repo directory is deleted when the user is deleted. It will have no impact when in the accounts SCIM API. False by default.
	ForceDeleteRepos *bool `json:"forceDeleteRepos,omitempty" tf:"force_delete_repos,omitempty"`

	// Home folder of the service principal, e.g. /Users/00000000-0000-0000-0000-000000000000.
	Home *string `json:"home,omitempty" tf:"home,omitempty"`

	// Personal Repos location of the service principal, e.g. /Repos/00000000-0000-0000-0000-000000000000.
	Repos *string `json:"repos,omitempty" tf:"repos,omitempty"`

	// This is a field to allow the group to have access to Databricks Workspace.
	WorkspaceAccess *bool `json:"workspaceAccess,omitempty" tf:"workspace_access,omitempty"`
}

type ServicePrincipalObservation struct {

	// identifier for use in databricks_access_control_rule_set, e.g. servicePrincipals/00000000-0000-0000-0000-000000000000.
	ACLPrincipalID *string `json:"aclPrincipalId,omitempty" tf:"acl_principal_id,omitempty"`

	// Either service principal is active or not. True by default, but can be set to false in case of service principal deactivation with preserving service principal assets.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Allow the service principal to have cluster create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and cluster_id argument. Everyone without allow_cluster_create argument set, but with permission to use Cluster Policy would be able to create clusters, but within the boundaries of that specific policy.
	AllowClusterCreate *bool `json:"allowClusterCreate,omitempty" tf:"allow_cluster_create,omitempty"`

	// Allow the service principal to have instance pool create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and instance_pool_id argument.
	AllowInstancePoolCreate *bool `json:"allowInstancePoolCreate,omitempty" tf:"allow_instance_pool_create,omitempty"`

	// managed service principals this value is auto-generated.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// This is a field to allow the group to have access to Databricks SQL feature through databricks_sql_endpoint.
	DatabricksSQLAccess *bool `json:"databricksSqlAccess,omitempty" tf:"databricks_sql_access,omitempty"`

	// Deactivate the service principal when deleting the resource, rather than deleting the service principal entirely. Defaults to true when the provider is configured at the account-level and false when configured at the workspace-level. This flag is exclusive to force_delete_repos and force_delete_home_dir flags.
	DisableAsUserDeletion *bool `json:"disableAsUserDeletion,omitempty" tf:"disable_as_user_deletion,omitempty"`

	// This is an alias for the service principal and can be the full name of the service principal.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// ID of the service principal in an external identity provider.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// This functionality is experimental and is designed to simplify corner cases, like Azure Active Directory synchronisation.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// This flag determines whether the service principal's home directory is deleted when the user is deleted. It will have no impact when in the accounts SCIM API. False by default.
	ForceDeleteHomeDir *bool `json:"forceDeleteHomeDir,omitempty" tf:"force_delete_home_dir,omitempty"`

	// This flag determines whether the service principal's repo directory is deleted when the user is deleted. It will have no impact when in the accounts SCIM API. False by default.
	ForceDeleteRepos *bool `json:"forceDeleteRepos,omitempty" tf:"force_delete_repos,omitempty"`

	// Home folder of the service principal, e.g. /Users/00000000-0000-0000-0000-000000000000.
	Home *string `json:"home,omitempty" tf:"home,omitempty"`

	// Canonical unique identifier for the service principal.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Personal Repos location of the service principal, e.g. /Repos/00000000-0000-0000-0000-000000000000.
	Repos *string `json:"repos,omitempty" tf:"repos,omitempty"`

	// This is a field to allow the group to have access to Databricks Workspace.
	WorkspaceAccess *bool `json:"workspaceAccess,omitempty" tf:"workspace_access,omitempty"`
}

type ServicePrincipalParameters struct {

	// identifier for use in databricks_access_control_rule_set, e.g. servicePrincipals/00000000-0000-0000-0000-000000000000.
	// +kubebuilder:validation:Optional
	ACLPrincipalID *string `json:"aclPrincipalId,omitempty" tf:"acl_principal_id,omitempty"`

	// Either service principal is active or not. True by default, but can be set to false in case of service principal deactivation with preserving service principal assets.
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Allow the service principal to have cluster create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and cluster_id argument. Everyone without allow_cluster_create argument set, but with permission to use Cluster Policy would be able to create clusters, but within the boundaries of that specific policy.
	// +kubebuilder:validation:Optional
	AllowClusterCreate *bool `json:"allowClusterCreate,omitempty" tf:"allow_cluster_create,omitempty"`

	// Allow the service principal to have instance pool create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and instance_pool_id argument.
	// +kubebuilder:validation:Optional
	AllowInstancePoolCreate *bool `json:"allowInstancePoolCreate,omitempty" tf:"allow_instance_pool_create,omitempty"`

	// managed service principals this value is auto-generated.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// This is a field to allow the group to have access to Databricks SQL feature through databricks_sql_endpoint.
	// +kubebuilder:validation:Optional
	DatabricksSQLAccess *bool `json:"databricksSqlAccess,omitempty" tf:"databricks_sql_access,omitempty"`

	// Deactivate the service principal when deleting the resource, rather than deleting the service principal entirely. Defaults to true when the provider is configured at the account-level and false when configured at the workspace-level. This flag is exclusive to force_delete_repos and force_delete_home_dir flags.
	// +kubebuilder:validation:Optional
	DisableAsUserDeletion *bool `json:"disableAsUserDeletion,omitempty" tf:"disable_as_user_deletion,omitempty"`

	// This is an alias for the service principal and can be the full name of the service principal.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// ID of the service principal in an external identity provider.
	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// This functionality is experimental and is designed to simplify corner cases, like Azure Active Directory synchronisation.
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// This flag determines whether the service principal's home directory is deleted when the user is deleted. It will have no impact when in the accounts SCIM API. False by default.
	// +kubebuilder:validation:Optional
	ForceDeleteHomeDir *bool `json:"forceDeleteHomeDir,omitempty" tf:"force_delete_home_dir,omitempty"`

	// This flag determines whether the service principal's repo directory is deleted when the user is deleted. It will have no impact when in the accounts SCIM API. False by default.
	// +kubebuilder:validation:Optional
	ForceDeleteRepos *bool `json:"forceDeleteRepos,omitempty" tf:"force_delete_repos,omitempty"`

	// Home folder of the service principal, e.g. /Users/00000000-0000-0000-0000-000000000000.
	// +kubebuilder:validation:Optional
	Home *string `json:"home,omitempty" tf:"home,omitempty"`

	// Personal Repos location of the service principal, e.g. /Repos/00000000-0000-0000-0000-000000000000.
	// +kubebuilder:validation:Optional
	Repos *string `json:"repos,omitempty" tf:"repos,omitempty"`

	// This is a field to allow the group to have access to Databricks Workspace.
	// +kubebuilder:validation:Optional
	WorkspaceAccess *bool `json:"workspaceAccess,omitempty" tf:"workspace_access,omitempty"`
}

// ServicePrincipalSpec defines the desired state of ServicePrincipal
type ServicePrincipalSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServicePrincipalParameters `json:"forProvider"`
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
	InitProvider ServicePrincipalInitParameters `json:"initProvider,omitempty"`
}

// ServicePrincipalStatus defines the observed state of ServicePrincipal.
type ServicePrincipalStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServicePrincipalObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServicePrincipal is the Schema for the ServicePrincipals API. ""subcategory: "Security"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type ServicePrincipal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicePrincipalSpec   `json:"spec"`
	Status            ServicePrincipalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicePrincipalList contains a list of ServicePrincipals
type ServicePrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicePrincipal `json:"items"`
}

// Repository type metadata.
var (
	ServicePrincipal_Kind             = "ServicePrincipal"
	ServicePrincipal_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServicePrincipal_Kind}.String()
	ServicePrincipal_KindAPIVersion   = ServicePrincipal_Kind + "." + CRDGroupVersion.String()
	ServicePrincipal_GroupVersionKind = CRDGroupVersion.WithKind(ServicePrincipal_Kind)
)

func init() {
	SchemeBuilder.Register(&ServicePrincipal{}, &ServicePrincipalList{})
}
