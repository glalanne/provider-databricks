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

type CatalogInitParameters struct {

	// User-supplied free-form text.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// For Foreign Catalogs: the name of the connection to an external data source. Changes forces creation of a new resource.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Whether predictive optimization should be enabled for this object and objects under it. Can be ENABLE, DISABLE or INHERIT
	EnablePredictiveOptimization *string `json:"enablePredictiveOptimization,omitempty" tf:"enable_predictive_optimization,omitempty"`

	// Delete catalog regardless of its contents.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Whether the catalog is accessible from all workspaces or a specific set of workspaces. Can be ISOLATED or OPEN. Setting the catalog to ISOLATED will automatically allow access from the current workspace.
	IsolationMode *string `json:"isolationMode,omitempty" tf:"isolation_mode,omitempty"`

	// ID of the parent metastore.
	MetastoreID *string `json:"metastoreId,omitempty" tf:"metastore_id,omitempty"`

	// Name of Catalog relative to parent metastore.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// For Foreign Catalogs: the name of the entity from an external data source that maps to a catalog. For example, the database name in a PostgreSQL server.
	// +mapType=granular
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// Username/groupname/sp application_id of the catalog owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Extensible Catalog properties.
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// For Delta Sharing Catalogs: the name of the delta sharing provider. Change forces creation of a new resource.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// For Delta Sharing Catalogs: the name of the share under the share provider. Change forces creation of a new resource.
	ShareName *string `json:"shareName,omitempty" tf:"share_name,omitempty"`

	// Managed location of the catalog. Location in cloud storage where data for managed tables will be stored. If not specified, the location will default to the metastore root location. Change forces creation of a new resource.
	StorageRoot *string `json:"storageRoot,omitempty" tf:"storage_root,omitempty"`
}

type CatalogObservation struct {

	// User-supplied free-form text.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// For Foreign Catalogs: the name of the connection to an external data source. Changes forces creation of a new resource.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Whether predictive optimization should be enabled for this object and objects under it. Can be ENABLE, DISABLE or INHERIT
	EnablePredictiveOptimization *string `json:"enablePredictiveOptimization,omitempty" tf:"enable_predictive_optimization,omitempty"`

	// Delete catalog regardless of its contents.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// ID of this catalog - same as the name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the catalog is accessible from all workspaces or a specific set of workspaces. Can be ISOLATED or OPEN. Setting the catalog to ISOLATED will automatically allow access from the current workspace.
	IsolationMode *string `json:"isolationMode,omitempty" tf:"isolation_mode,omitempty"`

	// ID of the parent metastore.
	MetastoreID *string `json:"metastoreId,omitempty" tf:"metastore_id,omitempty"`

	// Name of Catalog relative to parent metastore.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// For Foreign Catalogs: the name of the entity from an external data source that maps to a catalog. For example, the database name in a PostgreSQL server.
	// +mapType=granular
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// Username/groupname/sp application_id of the catalog owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Extensible Catalog properties.
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// For Delta Sharing Catalogs: the name of the delta sharing provider. Change forces creation of a new resource.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// For Delta Sharing Catalogs: the name of the share under the share provider. Change forces creation of a new resource.
	ShareName *string `json:"shareName,omitempty" tf:"share_name,omitempty"`

	// Managed location of the catalog. Location in cloud storage where data for managed tables will be stored. If not specified, the location will default to the metastore root location. Change forces creation of a new resource.
	StorageRoot *string `json:"storageRoot,omitempty" tf:"storage_root,omitempty"`
}

type CatalogParameters struct {

	// User-supplied free-form text.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// For Foreign Catalogs: the name of the connection to an external data source. Changes forces creation of a new resource.
	// +kubebuilder:validation:Optional
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Whether predictive optimization should be enabled for this object and objects under it. Can be ENABLE, DISABLE or INHERIT
	// +kubebuilder:validation:Optional
	EnablePredictiveOptimization *string `json:"enablePredictiveOptimization,omitempty" tf:"enable_predictive_optimization,omitempty"`

	// Delete catalog regardless of its contents.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Whether the catalog is accessible from all workspaces or a specific set of workspaces. Can be ISOLATED or OPEN. Setting the catalog to ISOLATED will automatically allow access from the current workspace.
	// +kubebuilder:validation:Optional
	IsolationMode *string `json:"isolationMode,omitempty" tf:"isolation_mode,omitempty"`

	// ID of the parent metastore.
	// +kubebuilder:validation:Optional
	MetastoreID *string `json:"metastoreId,omitempty" tf:"metastore_id,omitempty"`

	// Name of Catalog relative to parent metastore.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// For Foreign Catalogs: the name of the entity from an external data source that maps to a catalog. For example, the database name in a PostgreSQL server.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// Username/groupname/sp application_id of the catalog owner.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Extensible Catalog properties.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// For Delta Sharing Catalogs: the name of the delta sharing provider. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// For Delta Sharing Catalogs: the name of the share under the share provider. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	ShareName *string `json:"shareName,omitempty" tf:"share_name,omitempty"`

	// Managed location of the catalog. Location in cloud storage where data for managed tables will be stored. If not specified, the location will default to the metastore root location. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	StorageRoot *string `json:"storageRoot,omitempty" tf:"storage_root,omitempty"`
}

// CatalogSpec defines the desired state of Catalog
type CatalogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogParameters `json:"forProvider"`
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
	InitProvider CatalogInitParameters `json:"initProvider,omitempty"`
}

// CatalogStatus defines the observed state of Catalog.
type CatalogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Catalog is the Schema for the Catalogs API. ""subcategory: "Unity Catalog"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type Catalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   CatalogSpec   `json:"spec"`
	Status CatalogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogList contains a list of Catalogs
type CatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Catalog `json:"items"`
}

// Repository type metadata.
var (
	Catalog_Kind             = "Catalog"
	Catalog_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Catalog_Kind}.String()
	Catalog_KindAPIVersion   = Catalog_Kind + "." + CRDGroupVersion.String()
	Catalog_GroupVersionKind = CRDGroupVersion.WithKind(Catalog_Kind)
)

func init() {
	SchemeBuilder.Register(&Catalog{}, &CatalogList{})
}
