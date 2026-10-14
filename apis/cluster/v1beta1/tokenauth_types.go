// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
)

// CredentialsSourceServiceAccount makes the provider request a
// short-lived token for a Kubernetes service account (TokenRequest API) and
// exchange it for an access token, instead of reading stored credentials.
const CredentialsSourceServiceAccount xpv2.CredentialsSource = "ServiceAccount"

// ServiceAccountReference selects the Kubernetes service account the provider
// requests tokens for when Credentials.Source is ServiceAccount. The
// token's subject (system:serviceaccount:<namespace>:<name>) is what the
// identity provider trusts, which binds the identity to that namespace.
type ServiceAccountReference struct {
	// Name of the service account.
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// Namespace of the service account. Required: this ProviderConfig is
	// cluster-scoped, so there is no namespace to default to.
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

// AzureADTokenExchange holds settings specific to the AzureAD exchange.
type AzureADTokenExchange struct {
	// ClientID of the managed identity (or app registration) whose federated
	// credential trusts the service account. Defaults to the service
	// account's azure.workload.identity/client-id annotation.
	// +optional
	ClientID *string `json:"clientID,omitempty"`

	// TenantID of the identity. Defaults to the service account's
	// azure.workload.identity/tenant-id annotation.
	// +optional
	TenantID *string `json:"tenantID,omitempty"`

	// AuthorityHost of the Entra ID token endpoint. Defaults to
	// https://login.microsoftonline.com/.
	// +optional
	AuthorityHost *string `json:"authorityHost,omitempty"`
}
