// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
)

// CredentialsSourceOIDCTokenRequest makes the provider request a
// short-lived OIDC token for a Kubernetes service account (TokenRequest API)
// and exchange it for an access token, instead of reading stored credentials.
const CredentialsSourceOIDCTokenRequest xpv2.CredentialsSource = "OIDCTokenRequest"

// ServiceAccountReference selects the Kubernetes service account the provider
// requests tokens for when Credentials.Source is OIDCTokenRequest. The
// token's subject (system:serviceaccount:<namespace>:<name>) is what the
// identity provider trusts, which binds the identity to that namespace.
type ServiceAccountReference struct {
	// Name of the service account.
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`
	Name string `json:"name"`

	// Namespace of the service account. Required: this ProviderConfig is
	// cluster-scoped, so there is no namespace to default to.
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`
	// +optional
	Namespace *string `json:"namespace,omitempty"`

	// Audiences requested on the token in addition to the one the token
	// exchange requires (api://AzureADTokenExchange for AzureAD).
	// +optional
	// +listType=set
	Audiences []string `json:"audiences,omitempty"`
}

// TokenExchangeType selects how a workload token is exchanged for an access
// token Databricks accepts.
// +kubebuilder:validation:Enum=AzureAD
type TokenExchangeType string

const (
	// TokenExchangeAzureAD exchanges the token at Entra ID (RFC 7523 client
	// assertion) for an access token for the Azure Databricks application.
	TokenExchangeAzureAD TokenExchangeType = "AzureAD"
)

// TokenExchange configures how a workload token is exchanged for an access
// token. Each type has its own optional member for type-specific settings.
type TokenExchange struct {
	// Type of the exchange. Defaults to AzureAD when omitted.
	// +optional
	Type *TokenExchangeType `json:"type,omitempty"`

	// AzureAD holds settings specific to the AzureAD exchange.
	// +optional
	AzureAD *AzureADTokenExchange `json:"azureAD,omitempty"`
}

// AzureAuthorityHost is an Entra ID authority host the AzureAD exchange may
// send the service account's token to.
// +kubebuilder:validation:Enum="https://login.microsoftonline.com/";"https://login.microsoftonline.us/";"https://login.chinacloudapi.cn/";"https://login.partner.microsoftonline.cn/"
type AzureAuthorityHost string

const (
	// AzureAuthorityHostPublic is the Entra ID authority for the Azure public
	// cloud, and the default.
	AzureAuthorityHostPublic AzureAuthorityHost = "https://login.microsoftonline.com/"
	// AzureAuthorityHostUSGovernment is the Entra ID authority for Azure
	// Government.
	AzureAuthorityHostUSGovernment AzureAuthorityHost = "https://login.microsoftonline.us/"
	// AzureAuthorityHostChina is the Entra ID authority for Azure China.
	AzureAuthorityHostChina AzureAuthorityHost = "https://login.chinacloudapi.cn/"
	// AzureAuthorityHostChinaPartner is Azure China's newer Entra ID authority.
	AzureAuthorityHostChinaPartner AzureAuthorityHost = "https://login.partner.microsoftonline.cn/"
)

// AzureADTokenExchange holds settings specific to the AzureAD exchange.
type AzureADTokenExchange struct {
	// ClientID of the managed identity (or app registration) whose federated
	// credential trusts the service account, as a GUID. Defaults to the
	// service account's azure.workload.identity/client-id annotation.
	// +kubebuilder:validation:Pattern=`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	// +optional
	ClientID *string `json:"clientID,omitempty"`

	// TenantID of the identity, as a GUID. Defaults to the service account's
	// azure.workload.identity/tenant-id annotation.
	// +kubebuilder:validation:Pattern=`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	// +optional
	TenantID *string `json:"tenantID,omitempty"`

	// AuthorityHost of the Entra ID token endpoint: one of the Entra ID
	// authorities for the Azure public, Government or China clouds, with a
	// trailing slash. Defaults to https://login.microsoftonline.com/.
	// +optional
	AuthorityHost *AzureAuthorityHost `json:"authorityHost,omitempty"`
}
