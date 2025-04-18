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

type ColumnInitParameters struct {

	// User-supplied free-form text. Changing the comment is not currently supported on the VIEW table type.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Whether the field is an identity column. Can be default, always, or unset. It is unset by default.
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// Name of table relative to parent catalog and schema. Change forces the creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether field is nullable (Default: true)
	Nullable *bool `json:"nullable,omitempty" tf:"nullable,omitempty"`

	// Column type spec (with metadata) as SQL text. Not supported for VIEW table_type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	TypeJSON *string `json:"typeJson,omitempty" tf:"type_json,omitempty"`
}

type ColumnObservation struct {

	// User-supplied free-form text. Changing the comment is not currently supported on the VIEW table type.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Whether the field is an identity column. Can be default, always, or unset. It is unset by default.
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// Name of table relative to parent catalog and schema. Change forces the creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether field is nullable (Default: true)
	Nullable *bool `json:"nullable,omitempty" tf:"nullable,omitempty"`

	// Column type spec (with metadata) as SQL text. Not supported for VIEW table_type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	TypeJSON *string `json:"typeJson,omitempty" tf:"type_json,omitempty"`
}

type ColumnParameters struct {

	// User-supplied free-form text. Changing the comment is not currently supported on the VIEW table type.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Whether the field is an identity column. Can be default, always, or unset. It is unset by default.
	// +kubebuilder:validation:Optional
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// Name of table relative to parent catalog and schema. Change forces the creation of a new resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Whether field is nullable (Default: true)
	// +kubebuilder:validation:Optional
	Nullable *bool `json:"nullable,omitempty" tf:"nullable,omitempty"`

	// Column type spec (with metadata) as SQL text. Not supported for VIEW table_type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	TypeJSON *string `json:"typeJson,omitempty" tf:"type_json,omitempty"`
}

type SQLTableInitParameters struct {

	// Name of parent catalog. Change forces the creation of a new resource.
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// All table CRUD operations must be executed on a running cluster or SQL warehouse. If a cluster_id is specified, it will be used to execute SQL commands to manage this table. Conflicts with warehouse_id.
	// +crossplane:generate:reference:type=github.com/glalanne/provider-databricks/apis/compute/v1alpha1.Cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster in compute to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in compute to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// a subset of columns to liquid cluster the table by. For automatic clustering, set cluster_keys to ["AUTO"]. To turn off clustering, set it to ["NONE"]. Conflicts with partitions.
	ClusterKeys []*string `json:"clusterKeys,omitempty" tf:"cluster_keys,omitempty"`

	Column []ColumnInitParameters `json:"column,omitempty" tf:"column,omitempty"`

	// User-supplied free-form text. Changing the comment is not currently supported on the VIEW table type.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// External tables are supported in multiple data source formats. The string constants identifying these formats are DELTA, CSV, JSON, AVRO, PARQUET, ORC, and TEXT. Change forces the creation of a new resource. Not supported for MANAGED tables or VIEW.
	DataSourceFormat *string `json:"dataSourceFormat,omitempty" tf:"data_source_format,omitempty"`

	// Name of table relative to parent catalog and schema. Change forces the creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of user defined table options. Change forces creation of a new resource.
	// +mapType=granular
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// User name/group name/sp application_id of the table owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// a subset of columns to partition the table by. Change forces the creation of a new resource. Conflicts with cluster_keys.
	Partitions []*string `json:"partitions,omitempty" tf:"partitions,omitempty"`

	// A map of table properties.
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Name of parent Schema relative to parent Catalog. Change forces the creation of a new resource.
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// For EXTERNAL Tables only: the name of storage credential to use. Change forces the creation of a new resource.
	StorageCredentialName *string `json:"storageCredentialName,omitempty" tf:"storage_credential_name,omitempty"`

	// URL of storage location for Table data . Not supported for VIEW or MANAGED table_type.
	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// Distinguishes a view vs. managed/external Table. MANAGED, EXTERNAL, or VIEW. Change forces the creation of a new resource.
	TableType *string `json:"tableType,omitempty" tf:"table_type,omitempty"`

	// SQL text defining the view (for table_type == "VIEW"). Not supported for MANAGED or EXTERNAL table_type.
	ViewDefinition *string `json:"viewDefinition,omitempty" tf:"view_definition,omitempty"`

	// All table CRUD operations must be executed on a running cluster or SQL warehouse. If a warehouse_id is specified, that SQL warehouse will be used to execute SQL commands to manage this table. Conflicts with cluster_id.
	// +crossplane:generate:reference:type=github.com/glalanne/provider-databricks/apis/sql/v1alpha1.SQLEndpoint
	WarehouseID *string `json:"warehouseId,omitempty" tf:"warehouse_id,omitempty"`

	// Reference to a SQLEndpoint in sql to populate warehouseId.
	// +kubebuilder:validation:Optional
	WarehouseIDRef *v1.Reference `json:"warehouseIdRef,omitempty" tf:"-"`

	// Selector for a SQLEndpoint in sql to populate warehouseId.
	// +kubebuilder:validation:Optional
	WarehouseIDSelector *v1.Selector `json:"warehouseIdSelector,omitempty" tf:"-"`
}

type SQLTableObservation struct {

	// Name of parent catalog. Change forces the creation of a new resource.
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// All table CRUD operations must be executed on a running cluster or SQL warehouse. If a cluster_id is specified, it will be used to execute SQL commands to manage this table. Conflicts with warehouse_id.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// a subset of columns to liquid cluster the table by. For automatic clustering, set cluster_keys to ["AUTO"]. To turn off clustering, set it to ["NONE"]. Conflicts with partitions.
	ClusterKeys []*string `json:"clusterKeys,omitempty" tf:"cluster_keys,omitempty"`

	Column []ColumnObservation `json:"column,omitempty" tf:"column,omitempty"`

	// User-supplied free-form text. Changing the comment is not currently supported on the VIEW table type.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// External tables are supported in multiple data source formats. The string constants identifying these formats are DELTA, CSV, JSON, AVRO, PARQUET, ORC, and TEXT. Change forces the creation of a new resource. Not supported for MANAGED tables or VIEW.
	DataSourceFormat *string `json:"dataSourceFormat,omitempty" tf:"data_source_format,omitempty"`

	// A map of table properties.
	// +mapType=granular
	EffectiveProperties map[string]*string `json:"effectiveProperties,omitempty" tf:"effective_properties,omitempty"`

	// ID of this table in the form of <catalog_name>.<schema_name>.<name>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of table relative to parent catalog and schema. Change forces the creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of user defined table options. Change forces creation of a new resource.
	// +mapType=granular
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// User name/group name/sp application_id of the table owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// a subset of columns to partition the table by. Change forces the creation of a new resource. Conflicts with cluster_keys.
	Partitions []*string `json:"partitions,omitempty" tf:"partitions,omitempty"`

	// A map of table properties.
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Name of parent Schema relative to parent Catalog. Change forces the creation of a new resource.
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// For EXTERNAL Tables only: the name of storage credential to use. Change forces the creation of a new resource.
	StorageCredentialName *string `json:"storageCredentialName,omitempty" tf:"storage_credential_name,omitempty"`

	// URL of storage location for Table data . Not supported for VIEW or MANAGED table_type.
	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// Distinguishes a view vs. managed/external Table. MANAGED, EXTERNAL, or VIEW. Change forces the creation of a new resource.
	TableType *string `json:"tableType,omitempty" tf:"table_type,omitempty"`

	// SQL text defining the view (for table_type == "VIEW"). Not supported for MANAGED or EXTERNAL table_type.
	ViewDefinition *string `json:"viewDefinition,omitempty" tf:"view_definition,omitempty"`

	// All table CRUD operations must be executed on a running cluster or SQL warehouse. If a warehouse_id is specified, that SQL warehouse will be used to execute SQL commands to manage this table. Conflicts with cluster_id.
	WarehouseID *string `json:"warehouseId,omitempty" tf:"warehouse_id,omitempty"`
}

type SQLTableParameters struct {

	// Name of parent catalog. Change forces the creation of a new resource.
	// +kubebuilder:validation:Optional
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// All table CRUD operations must be executed on a running cluster or SQL warehouse. If a cluster_id is specified, it will be used to execute SQL commands to manage this table. Conflicts with warehouse_id.
	// +crossplane:generate:reference:type=github.com/glalanne/provider-databricks/apis/compute/v1alpha1.Cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster in compute to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in compute to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// a subset of columns to liquid cluster the table by. For automatic clustering, set cluster_keys to ["AUTO"]. To turn off clustering, set it to ["NONE"]. Conflicts with partitions.
	// +kubebuilder:validation:Optional
	ClusterKeys []*string `json:"clusterKeys,omitempty" tf:"cluster_keys,omitempty"`

	// +kubebuilder:validation:Optional
	Column []ColumnParameters `json:"column,omitempty" tf:"column,omitempty"`

	// User-supplied free-form text. Changing the comment is not currently supported on the VIEW table type.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// External tables are supported in multiple data source formats. The string constants identifying these formats are DELTA, CSV, JSON, AVRO, PARQUET, ORC, and TEXT. Change forces the creation of a new resource. Not supported for MANAGED tables or VIEW.
	// +kubebuilder:validation:Optional
	DataSourceFormat *string `json:"dataSourceFormat,omitempty" tf:"data_source_format,omitempty"`

	// Name of table relative to parent catalog and schema. Change forces the creation of a new resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of user defined table options. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// User name/group name/sp application_id of the table owner.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// a subset of columns to partition the table by. Change forces the creation of a new resource. Conflicts with cluster_keys.
	// +kubebuilder:validation:Optional
	Partitions []*string `json:"partitions,omitempty" tf:"partitions,omitempty"`

	// A map of table properties.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Name of parent Schema relative to parent Catalog. Change forces the creation of a new resource.
	// +kubebuilder:validation:Optional
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// For EXTERNAL Tables only: the name of storage credential to use. Change forces the creation of a new resource.
	// +kubebuilder:validation:Optional
	StorageCredentialName *string `json:"storageCredentialName,omitempty" tf:"storage_credential_name,omitempty"`

	// URL of storage location for Table data . Not supported for VIEW or MANAGED table_type.
	// +kubebuilder:validation:Optional
	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// Distinguishes a view vs. managed/external Table. MANAGED, EXTERNAL, or VIEW. Change forces the creation of a new resource.
	// +kubebuilder:validation:Optional
	TableType *string `json:"tableType,omitempty" tf:"table_type,omitempty"`

	// SQL text defining the view (for table_type == "VIEW"). Not supported for MANAGED or EXTERNAL table_type.
	// +kubebuilder:validation:Optional
	ViewDefinition *string `json:"viewDefinition,omitempty" tf:"view_definition,omitempty"`

	// All table CRUD operations must be executed on a running cluster or SQL warehouse. If a warehouse_id is specified, that SQL warehouse will be used to execute SQL commands to manage this table. Conflicts with cluster_id.
	// +crossplane:generate:reference:type=github.com/glalanne/provider-databricks/apis/sql/v1alpha1.SQLEndpoint
	// +kubebuilder:validation:Optional
	WarehouseID *string `json:"warehouseId,omitempty" tf:"warehouse_id,omitempty"`

	// Reference to a SQLEndpoint in sql to populate warehouseId.
	// +kubebuilder:validation:Optional
	WarehouseIDRef *v1.Reference `json:"warehouseIdRef,omitempty" tf:"-"`

	// Selector for a SQLEndpoint in sql to populate warehouseId.
	// +kubebuilder:validation:Optional
	WarehouseIDSelector *v1.Selector `json:"warehouseIdSelector,omitempty" tf:"-"`
}

// SQLTableSpec defines the desired state of SQLTable
type SQLTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLTableParameters `json:"forProvider"`
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
	InitProvider SQLTableInitParameters `json:"initProvider,omitempty"`
}

// SQLTableStatus defines the observed state of SQLTable.
type SQLTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SQLTable is the Schema for the SQLTables API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type SQLTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.catalogName) || (has(self.initProvider) && has(self.initProvider.catalogName))",message="spec.forProvider.catalogName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schemaName) || (has(self.initProvider) && has(self.initProvider.schemaName))",message="spec.forProvider.schemaName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tableType) || (has(self.initProvider) && has(self.initProvider.tableType))",message="spec.forProvider.tableType is a required parameter"
	Spec   SQLTableSpec   `json:"spec"`
	Status SQLTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLTableList contains a list of SQLTables
type SQLTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLTable `json:"items"`
}

// Repository type metadata.
var (
	SQLTable_Kind             = "SQLTable"
	SQLTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLTable_Kind}.String()
	SQLTable_KindAPIVersion   = SQLTable_Kind + "." + CRDGroupVersion.String()
	SQLTable_GroupVersionKind = CRDGroupVersion.WithKind(SQLTable_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLTable{}, &SQLTableList{})
}
