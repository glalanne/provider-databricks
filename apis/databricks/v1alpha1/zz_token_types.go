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

type TokenInitParameters struct {
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	CreationTime *float64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	ExpiryTime *float64 `json:"expiryTime,omitempty" tf:"expiry_time,omitempty"`

	LifetimeSeconds *float64 `json:"lifetimeSeconds,omitempty" tf:"lifetime_seconds,omitempty"`

	TokenID *string `json:"tokenId,omitempty" tf:"token_id,omitempty"`
}

type TokenObservation struct {
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	CreationTime *float64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	ExpiryTime *float64 `json:"expiryTime,omitempty" tf:"expiry_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LifetimeSeconds *float64 `json:"lifetimeSeconds,omitempty" tf:"lifetime_seconds,omitempty"`

	TokenID *string `json:"tokenId,omitempty" tf:"token_id,omitempty"`
}

type TokenParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Optional
	CreationTime *float64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// +kubebuilder:validation:Optional
	ExpiryTime *float64 `json:"expiryTime,omitempty" tf:"expiry_time,omitempty"`

	// +kubebuilder:validation:Optional
	LifetimeSeconds *float64 `json:"lifetimeSeconds,omitempty" tf:"lifetime_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	TokenID *string `json:"tokenId,omitempty" tf:"token_id,omitempty"`
}

// TokenSpec defines the desired state of Token
type TokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TokenParameters `json:"forProvider"`
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
	InitProvider TokenInitParameters `json:"initProvider,omitempty"`
}

// TokenStatus defines the observed state of Token.
type TokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Token is the Schema for the Tokens API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TokenSpec   `json:"spec"`
	Status            TokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TokenList contains a list of Tokens
type TokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Token `json:"items"`
}

// Repository type metadata.
var (
	Token_Kind             = "Token"
	Token_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Token_Kind}.String()
	Token_KindAPIVersion   = Token_Kind + "." + CRDGroupVersion.String()
	Token_GroupVersionKind = CRDGroupVersion.WithKind(Token_Kind)
)

func init() {
	SchemeBuilder.Register(&Token{}, &TokenList{})
}
