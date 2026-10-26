// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
// +kubebuilder:validation:XValidation:rule="self.credentials.source != 'OIDCTokenRequest' || has(self.host)",message="host is required when credentials.source is OIDCTokenRequest"
// +kubebuilder:validation:XValidation:rule=`self.credentials.source != 'OIDCTokenRequest' || !has(self.host) || (has(self.credentials.tokenExchange) && has(self.credentials.tokenExchange.type) && self.credentials.tokenExchange.type != 'AzureAD') || self.host.matches(r'^https://([a-z0-9]([a-z0-9-]*[a-z0-9])?\.)+(azuredatabricks\.net|databricks\.azure\.us|databricks\.azure\.cn)/?$')`,message="host must be an https URL on an Azure Databricks domain (azuredatabricks.net, databricks.azure.us or databricks.azure.cn), with no port or path, when tokenExchange.type is AzureAD (the default)",fieldPath=".host"
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	Credentials ProviderCredentials `json:"credentials"`

	// Host is the Databricks workspace URL (or the accounts console URL for
	// account-level operations) the provider calls, and sends the access token
	// to. Required when Credentials.Source is OIDCTokenRequest. With the
	// AzureAD token exchange (the default), it must be an https URL on an
	// Azure Databricks domain (azuredatabricks.net, databricks.azure.us or
	// databricks.azure.cn), with no port or path.
	// +optional
	Host *string `json:"host,omitempty"`

	// AccountID is the Databricks account ID (a GUID), for account-level
	// operations through the accounts console host (for example
	// https://accounts.azuredatabricks.net). Used when Credentials.Source is
	// OIDCTokenRequest; leave unset for a workspace host.
	// +kubebuilder:validation:Pattern=`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	// +optional
	AccountID *string `json:"accountID,omitempty"`
}

// ProviderCredentials required to authenticate.
// +kubebuilder:validation:XValidation:rule="self.source != 'OIDCTokenRequest' || (has(self.serviceAccountRef) && has(self.serviceAccountRef.namespace))",message="serviceAccountRef with a namespace is required when source is OIDCTokenRequest"
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=None;Secret;UserAssignedManagedIdentity;SystemAssignedManagedIdentity;OIDCTokenFile;Upbound;Filesystem;OIDCTokenRequest
	Source xpv2.CredentialsSource `json:"source"`

	// ServiceAccountRef selects the service account the provider requests
	// tokens for when Source is OIDCTokenRequest.
	// +optional
	ServiceAccountRef *ServiceAccountReference `json:"serviceAccountRef,omitempty"`

	// TokenExchange configures how the service account's token is exchanged
	// for an access token when Source is OIDCTokenRequest. Defaults to type
	// AzureAD.
	// +optional
	TokenExchange *TokenExchange `json:"tokenExchange,omitempty"`

	xpv2.CommonCredentialSelectors `json:",inline"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv2.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures the Databricks provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:scope=Cluster,categories={crossplane,providerconfig,databricks}
// +kubebuilder:storageversion
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// +kubebuilder:object:root=true

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CONFIG-NAME",type="string",JSONPath=".providerConfigRef.name"
// +kubebuilder:printcolumn:name="RESOURCE-KIND",type="string",JSONPath=".resourceRef.kind"
// +kubebuilder:printcolumn:name="RESOURCE-NAME",type="string",JSONPath=".resourceRef.name"
// Please replace `PROVIDER-NAME` with your actual provider name, like `aws`, `azure`, `gcp`, `alibaba`
// +kubebuilder:resource:scope=Cluster,categories={crossplane,providerconfig,databricks}
// +kubebuilder:storageversion
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv2.ProviderConfigUsage `json:",inline"`
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}
