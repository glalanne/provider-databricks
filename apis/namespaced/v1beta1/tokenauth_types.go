// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
)

// CredentialsSourceOIDCTokenRequest exchanges a service account's short-lived
// OIDC token (TokenRequest API) for an access token; no stored credentials.
const CredentialsSourceOIDCTokenRequest xpv2.CredentialsSource = "OIDCTokenRequest"

// ServiceAccountReference selects the service account the provider requests
// tokens for. The token's subject, system:serviceaccount:<namespace>:<name>,
// is what the identity provider trusts, binding the identity to the namespace.
type ServiceAccountReference struct {
	// Name of the service account.
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`
	Name string `json:"name"`

	// Namespace of the service account. Required on a ClusterProviderConfig;
	// on a ProviderConfig it defaults to, and must be, its own namespace.
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`
	// +optional
	Namespace *string `json:"namespace,omitempty"`

	// Audiences requested on the token besides the one the exchange requires
	// (api://AzureADTokenExchange for AzureAD).
	// +optional
	// +listType=set
	Audiences []string `json:"audiences,omitempty"`
}

// TokenExchangeType selects how a workload token becomes an access token
// Databricks accepts. A new type adds its settings as TokenExchange.<Type>,
// such as TokenExchange.AzureAD, allowed only with that type, and can add its
// own host rule to ProviderConfigSpec.
// +kubebuilder:validation:Enum=AzureAD
type TokenExchangeType string

const (
	// TokenExchangeAzureAD exchanges the token at Entra ID (RFC 7523 client
	// assertion) for an Azure Databricks access token.
	TokenExchangeAzureAD TokenExchangeType = "AzureAD"
)

// TokenExchange configures the exchange. Each type has an optional settings
// member, allowed only with that type.
// +kubebuilder:validation:XValidation:rule="!has(self.azureAD) || !has(self.type) || self.type == 'AzureAD'",message="azureAD is only allowed when type is AzureAD (the default)",fieldPath=".azureAD"
type TokenExchange struct {
	// Type of the exchange. Defaults to AzureAD.
	// +optional
	Type *TokenExchangeType `json:"type,omitempty"`

	// AzureAD holds settings for the AzureAD exchange.
	// +optional
	AzureAD *AzureADTokenExchange `json:"azureAD,omitempty"`
}

// AzureAuthorityHost is an Entra ID authority host the AzureAD exchange may
// send the service account's token to.
// +kubebuilder:validation:Enum="https://login.microsoftonline.com/";"https://login.microsoftonline.us/";"https://login.chinacloudapi.cn/";"https://login.partner.microsoftonline.cn/"
type AzureAuthorityHost string

const (
	// AzureAuthorityHostPublic is the public cloud's authority, and the default.
	AzureAuthorityHostPublic AzureAuthorityHost = "https://login.microsoftonline.com/"
	// AzureAuthorityHostUSGovernment is Azure Government's authority.
	AzureAuthorityHostUSGovernment AzureAuthorityHost = "https://login.microsoftonline.us/"
	// AzureAuthorityHostChina is Azure China's authority.
	AzureAuthorityHostChina AzureAuthorityHost = "https://login.chinacloudapi.cn/"
	// AzureAuthorityHostChinaPartner is Azure China's newer authority.
	AzureAuthorityHostChinaPartner AzureAuthorityHost = "https://login.partner.microsoftonline.cn/"
)

// AzureADTokenExchange holds settings for the AzureAD exchange.
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

	// AuthorityHost of the Entra ID token endpoint, with a trailing slash, for
	// the Azure public, Government or China cloud. Defaults to
	// https://login.microsoftonline.com/.
	// +optional
	AuthorityHost *AzureAuthorityHost `json:"authorityHost,omitempty"`
}
