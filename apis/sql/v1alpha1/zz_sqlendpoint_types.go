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

type ChannelInitParameters struct {
	DbsqlVersion *string `json:"dbsqlVersion,omitempty" tf:"dbsql_version,omitempty"`

	// Name of the Databricks SQL release channel. Possible values are: CHANNEL_NAME_PREVIEW and CHANNEL_NAME_CURRENT. Default is CHANNEL_NAME_CURRENT.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ChannelObservation struct {
	DbsqlVersion *string `json:"dbsqlVersion,omitempty" tf:"dbsql_version,omitempty"`

	// Name of the Databricks SQL release channel. Possible values are: CHANNEL_NAME_PREVIEW and CHANNEL_NAME_CURRENT. Default is CHANNEL_NAME_CURRENT.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ChannelParameters struct {

	// +kubebuilder:validation:Optional
	DbsqlVersion *string `json:"dbsqlVersion,omitempty" tf:"dbsql_version,omitempty"`

	// Name of the Databricks SQL release channel. Possible values are: CHANNEL_NAME_PREVIEW and CHANNEL_NAME_CURRENT. Default is CHANNEL_NAME_CURRENT.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CustomTagsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomTagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomTagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type FailureReasonInitParameters struct {
}

type FailureReasonObservation struct {
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FailureReasonParameters struct {
}

type HealthInitParameters struct {
}

type HealthObservation struct {
	Details *string `json:"details,omitempty" tf:"details,omitempty"`

	FailureReason []FailureReasonObservation `json:"failureReason,omitempty" tf:"failure_reason,omitempty"`

	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Summary *string `json:"summary,omitempty" tf:"summary,omitempty"`
}

type HealthParameters struct {
}

type OdbcParamsInitParameters struct {
}

type OdbcParamsObservation struct {

	// Name of the SQL warehouse. Must be unique.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type OdbcParamsParameters struct {
}

type SQLEndpointInitParameters struct {

	// Time in minutes until an idle SQL warehouse terminates all clusters and stops. This field is optional. The default is 120, set to 0 to disable the auto stop.
	AutoStopMins *float64 `json:"autoStopMins,omitempty" tf:"auto_stop_mins,omitempty"`

	// block, consisting of following fields:
	Channel []ChannelInitParameters `json:"channel,omitempty" tf:"channel,omitempty"`

	// The size of the clusters allocated to the endpoint: "2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large".
	ClusterSize *string `json:"clusterSize,omitempty" tf:"cluster_size,omitempty"`

	// ID of the data source for this endpoint. This is used to bind an Databricks SQL query to an endpoint.
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// Whether to enable Photon. This field is optional and is enabled by default.
	EnablePhoton *bool `json:"enablePhoton,omitempty" tf:"enable_photon,omitempty"`

	// Whether this SQL warehouse is a serverless endpoint. See below for details about the default values. To avoid ambiguity, especially for organizations with many workspaces, Databricks recommends that you always set this field explicitly.
	EnableServerlessCompute *bool `json:"enableServerlessCompute,omitempty" tf:"enable_serverless_compute,omitempty"`

	InstanceProfileArn *string `json:"instanceProfileArn,omitempty" tf:"instance_profile_arn,omitempty"`

	// Maximum number of clusters available when a SQL warehouse is running. This field is required. If multi-cluster load balancing is not enabled, this is default to 1.
	MaxNumClusters *float64 `json:"maxNumClusters,omitempty" tf:"max_num_clusters,omitempty"`

	// Minimum number of clusters available when a SQL warehouse is running. The default is 1.
	MinNumClusters *float64 `json:"minNumClusters,omitempty" tf:"min_num_clusters,omitempty"`

	// Name of the SQL warehouse. Must be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The spot policy to use for allocating instances to clusters: COST_OPTIMIZED or RELIABILITY_OPTIMIZED. This field is optional. Default is COST_OPTIMIZED.
	SpotInstancePolicy *string `json:"spotInstancePolicy,omitempty" tf:"spot_instance_policy,omitempty"`

	// Databricks tags all endpoint resources with these tags.
	Tags []TagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// SQL warehouse type. See for AWS or Azure. Set to PRO or CLASSIC. If the field enable_serverless_compute has the value true either explicitly or through the default logic (see that field above for details), the default is PRO, which is required for serverless SQL warehouses. Otherwise, the default is CLASSIC.
	WarehouseType *string `json:"warehouseType,omitempty" tf:"warehouse_type,omitempty"`
}

type SQLEndpointObservation struct {

	// Time in minutes until an idle SQL warehouse terminates all clusters and stops. This field is optional. The default is 120, set to 0 to disable the auto stop.
	AutoStopMins *float64 `json:"autoStopMins,omitempty" tf:"auto_stop_mins,omitempty"`

	// block, consisting of following fields:
	Channel []ChannelObservation `json:"channel,omitempty" tf:"channel,omitempty"`

	// The size of the clusters allocated to the endpoint: "2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large".
	ClusterSize *string `json:"clusterSize,omitempty" tf:"cluster_size,omitempty"`

	// The username of the user who created the endpoint.
	CreatorName *string `json:"creatorName,omitempty" tf:"creator_name,omitempty"`

	// ID of the data source for this endpoint. This is used to bind an Databricks SQL query to an endpoint.
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// Whether to enable Photon. This field is optional and is enabled by default.
	EnablePhoton *bool `json:"enablePhoton,omitempty" tf:"enable_photon,omitempty"`

	// Whether this SQL warehouse is a serverless endpoint. See below for details about the default values. To avoid ambiguity, especially for organizations with many workspaces, Databricks recommends that you always set this field explicitly.
	EnableServerlessCompute *bool `json:"enableServerlessCompute,omitempty" tf:"enable_serverless_compute,omitempty"`

	// Health status of the endpoint.
	Health []HealthObservation `json:"health,omitempty" tf:"health,omitempty"`

	// the unique ID of the SQL warehouse.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceProfileArn *string `json:"instanceProfileArn,omitempty" tf:"instance_profile_arn,omitempty"`

	// JDBC connection string.
	JdbcURL *string `json:"jdbcUrl,omitempty" tf:"jdbc_url,omitempty"`

	// Maximum number of clusters available when a SQL warehouse is running. This field is required. If multi-cluster load balancing is not enabled, this is default to 1.
	MaxNumClusters *float64 `json:"maxNumClusters,omitempty" tf:"max_num_clusters,omitempty"`

	// Minimum number of clusters available when a SQL warehouse is running. The default is 1.
	MinNumClusters *float64 `json:"minNumClusters,omitempty" tf:"min_num_clusters,omitempty"`

	// Name of the SQL warehouse. Must be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The current number of clusters used by the endpoint.
	NumActiveSessions *float64 `json:"numActiveSessions,omitempty" tf:"num_active_sessions,omitempty"`

	// The current number of clusters used by the endpoint.
	NumClusters *float64 `json:"numClusters,omitempty" tf:"num_clusters,omitempty"`

	// ODBC connection params: odbc_params.hostname, odbc_params.path, odbc_params.protocol, and odbc_params.port.
	OdbcParams []OdbcParamsObservation `json:"odbcParams,omitempty" tf:"odbc_params,omitempty"`

	// The spot policy to use for allocating instances to clusters: COST_OPTIMIZED or RELIABILITY_OPTIMIZED. This field is optional. Default is COST_OPTIMIZED.
	SpotInstancePolicy *string `json:"spotInstancePolicy,omitempty" tf:"spot_instance_policy,omitempty"`

	// The current state of the endpoint.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Databricks tags all endpoint resources with these tags.
	Tags []TagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	// SQL warehouse type. See for AWS or Azure. Set to PRO or CLASSIC. If the field enable_serverless_compute has the value true either explicitly or through the default logic (see that field above for details), the default is PRO, which is required for serverless SQL warehouses. Otherwise, the default is CLASSIC.
	WarehouseType *string `json:"warehouseType,omitempty" tf:"warehouse_type,omitempty"`
}

type SQLEndpointParameters struct {

	// Time in minutes until an idle SQL warehouse terminates all clusters and stops. This field is optional. The default is 120, set to 0 to disable the auto stop.
	// +kubebuilder:validation:Optional
	AutoStopMins *float64 `json:"autoStopMins,omitempty" tf:"auto_stop_mins,omitempty"`

	// block, consisting of following fields:
	// +kubebuilder:validation:Optional
	Channel []ChannelParameters `json:"channel,omitempty" tf:"channel,omitempty"`

	// The size of the clusters allocated to the endpoint: "2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large".
	// +kubebuilder:validation:Optional
	ClusterSize *string `json:"clusterSize,omitempty" tf:"cluster_size,omitempty"`

	// ID of the data source for this endpoint. This is used to bind an Databricks SQL query to an endpoint.
	// +kubebuilder:validation:Optional
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// Whether to enable Photon. This field is optional and is enabled by default.
	// +kubebuilder:validation:Optional
	EnablePhoton *bool `json:"enablePhoton,omitempty" tf:"enable_photon,omitempty"`

	// Whether this SQL warehouse is a serverless endpoint. See below for details about the default values. To avoid ambiguity, especially for organizations with many workspaces, Databricks recommends that you always set this field explicitly.
	// +kubebuilder:validation:Optional
	EnableServerlessCompute *bool `json:"enableServerlessCompute,omitempty" tf:"enable_serverless_compute,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceProfileArn *string `json:"instanceProfileArn,omitempty" tf:"instance_profile_arn,omitempty"`

	// Maximum number of clusters available when a SQL warehouse is running. This field is required. If multi-cluster load balancing is not enabled, this is default to 1.
	// +kubebuilder:validation:Optional
	MaxNumClusters *float64 `json:"maxNumClusters,omitempty" tf:"max_num_clusters,omitempty"`

	// Minimum number of clusters available when a SQL warehouse is running. The default is 1.
	// +kubebuilder:validation:Optional
	MinNumClusters *float64 `json:"minNumClusters,omitempty" tf:"min_num_clusters,omitempty"`

	// Name of the SQL warehouse. Must be unique.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The spot policy to use for allocating instances to clusters: COST_OPTIMIZED or RELIABILITY_OPTIMIZED. This field is optional. Default is COST_OPTIMIZED.
	// +kubebuilder:validation:Optional
	SpotInstancePolicy *string `json:"spotInstancePolicy,omitempty" tf:"spot_instance_policy,omitempty"`

	// Databricks tags all endpoint resources with these tags.
	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// SQL warehouse type. See for AWS or Azure. Set to PRO or CLASSIC. If the field enable_serverless_compute has the value true either explicitly or through the default logic (see that field above for details), the default is PRO, which is required for serverless SQL warehouses. Otherwise, the default is CLASSIC.
	// +kubebuilder:validation:Optional
	WarehouseType *string `json:"warehouseType,omitempty" tf:"warehouse_type,omitempty"`
}

type TagsInitParameters struct {

	// Databricks tags all endpoint resources with these tags.
	CustomTags []CustomTagsInitParameters `json:"customTags,omitempty" tf:"custom_tags,omitempty"`
}

type TagsObservation struct {

	// Databricks tags all endpoint resources with these tags.
	CustomTags []CustomTagsObservation `json:"customTags,omitempty" tf:"custom_tags,omitempty"`
}

type TagsParameters struct {

	// Databricks tags all endpoint resources with these tags.
	// +kubebuilder:validation:Optional
	CustomTags []CustomTagsParameters `json:"customTags,omitempty" tf:"custom_tags,omitempty"`
}

// SQLEndpointSpec defines the desired state of SQLEndpoint
type SQLEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLEndpointParameters `json:"forProvider"`
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
	InitProvider SQLEndpointInitParameters `json:"initProvider,omitempty"`
}

// SQLEndpointStatus defines the observed state of SQLEndpoint.
type SQLEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SQLEndpoint is the Schema for the SQLEndpoints API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type SQLEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusterSize) || (has(self.initProvider) && has(self.initProvider.clusterSize))",message="spec.forProvider.clusterSize is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SQLEndpointSpec   `json:"spec"`
	Status SQLEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLEndpointList contains a list of SQLEndpoints
type SQLEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLEndpoint `json:"items"`
}

// Repository type metadata.
var (
	SQLEndpoint_Kind             = "SQLEndpoint"
	SQLEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLEndpoint_Kind}.String()
	SQLEndpoint_KindAPIVersion   = SQLEndpoint_Kind + "." + CRDGroupVersion.String()
	SQLEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(SQLEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLEndpoint{}, &SQLEndpointList{})
}
