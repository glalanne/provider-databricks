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

type SQLGlobalConfigInitParameters struct {

	// Data access configuration for databricks_sql_endpoint, such as configuration for an external Hive metastore, Hadoop Filesystem configuration, etc.  Please note that the list of supported configuration properties is limited, so refer to the documentation for a full list.  Apply will fail if you're specifying not permitted configuration.
	// +mapType=granular
	DataAccessConfig map[string]*string `json:"dataAccessConfig,omitempty" tf:"data_access_config,omitempty"`

	EnableServerlessCompute *bool `json:"enableServerlessCompute,omitempty" tf:"enable_serverless_compute,omitempty"`

	// used to access GCP services, such as Cloud Storage, from databricks_sql_endpoint. Please note that this parameter is only for GCP, and will generate an error if used on other clouds.
	GoogleServiceAccount *string `json:"googleServiceAccount,omitempty" tf:"google_service_account,omitempty"`

	// databricks_instance_profile used to access storage from databricks_sql_endpoint. Please note that this parameter is only for AWS, and will generate an error if used on other clouds.
	InstanceProfileArn *string `json:"instanceProfileArn,omitempty" tf:"instance_profile_arn,omitempty"`

	// SQL Configuration Parameters let you override the default behavior for all sessions with all endpoints.
	// +mapType=granular
	SQLConfigParams map[string]*string `json:"sqlConfigParams,omitempty" tf:"sql_config_params,omitempty"`

	// The policy for controlling access to datasets. Default value: DATA_ACCESS_CONTROL, consult documentation for list of possible values
	SecurityPolicy *string `json:"securityPolicy,omitempty" tf:"security_policy,omitempty"`
}

type SQLGlobalConfigObservation struct {

	// Data access configuration for databricks_sql_endpoint, such as configuration for an external Hive metastore, Hadoop Filesystem configuration, etc.  Please note that the list of supported configuration properties is limited, so refer to the documentation for a full list.  Apply will fail if you're specifying not permitted configuration.
	// +mapType=granular
	DataAccessConfig map[string]*string `json:"dataAccessConfig,omitempty" tf:"data_access_config,omitempty"`

	EnableServerlessCompute *bool `json:"enableServerlessCompute,omitempty" tf:"enable_serverless_compute,omitempty"`

	// used to access GCP services, such as Cloud Storage, from databricks_sql_endpoint. Please note that this parameter is only for GCP, and will generate an error if used on other clouds.
	GoogleServiceAccount *string `json:"googleServiceAccount,omitempty" tf:"google_service_account,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// databricks_instance_profile used to access storage from databricks_sql_endpoint. Please note that this parameter is only for AWS, and will generate an error if used on other clouds.
	InstanceProfileArn *string `json:"instanceProfileArn,omitempty" tf:"instance_profile_arn,omitempty"`

	// SQL Configuration Parameters let you override the default behavior for all sessions with all endpoints.
	// +mapType=granular
	SQLConfigParams map[string]*string `json:"sqlConfigParams,omitempty" tf:"sql_config_params,omitempty"`

	// The policy for controlling access to datasets. Default value: DATA_ACCESS_CONTROL, consult documentation for list of possible values
	SecurityPolicy *string `json:"securityPolicy,omitempty" tf:"security_policy,omitempty"`
}

type SQLGlobalConfigParameters struct {

	// Data access configuration for databricks_sql_endpoint, such as configuration for an external Hive metastore, Hadoop Filesystem configuration, etc.  Please note that the list of supported configuration properties is limited, so refer to the documentation for a full list.  Apply will fail if you're specifying not permitted configuration.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DataAccessConfig map[string]*string `json:"dataAccessConfig,omitempty" tf:"data_access_config,omitempty"`

	// +kubebuilder:validation:Optional
	EnableServerlessCompute *bool `json:"enableServerlessCompute,omitempty" tf:"enable_serverless_compute,omitempty"`

	// used to access GCP services, such as Cloud Storage, from databricks_sql_endpoint. Please note that this parameter is only for GCP, and will generate an error if used on other clouds.
	// +kubebuilder:validation:Optional
	GoogleServiceAccount *string `json:"googleServiceAccount,omitempty" tf:"google_service_account,omitempty"`

	// databricks_instance_profile used to access storage from databricks_sql_endpoint. Please note that this parameter is only for AWS, and will generate an error if used on other clouds.
	// +kubebuilder:validation:Optional
	InstanceProfileArn *string `json:"instanceProfileArn,omitempty" tf:"instance_profile_arn,omitempty"`

	// SQL Configuration Parameters let you override the default behavior for all sessions with all endpoints.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	SQLConfigParams map[string]*string `json:"sqlConfigParams,omitempty" tf:"sql_config_params,omitempty"`

	// The policy for controlling access to datasets. Default value: DATA_ACCESS_CONTROL, consult documentation for list of possible values
	// +kubebuilder:validation:Optional
	SecurityPolicy *string `json:"securityPolicy,omitempty" tf:"security_policy,omitempty"`
}

// SQLGlobalConfigSpec defines the desired state of SQLGlobalConfig
type SQLGlobalConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLGlobalConfigParameters `json:"forProvider"`
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
	InitProvider SQLGlobalConfigInitParameters `json:"initProvider,omitempty"`
}

// SQLGlobalConfigStatus defines the observed state of SQLGlobalConfig.
type SQLGlobalConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLGlobalConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SQLGlobalConfig is the Schema for the SQLGlobalConfigs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type SQLGlobalConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLGlobalConfigSpec   `json:"spec"`
	Status            SQLGlobalConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLGlobalConfigList contains a list of SQLGlobalConfigs
type SQLGlobalConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLGlobalConfig `json:"items"`
}

// Repository type metadata.
var (
	SQLGlobalConfig_Kind             = "SQLGlobalConfig"
	SQLGlobalConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLGlobalConfig_Kind}.String()
	SQLGlobalConfig_KindAPIVersion   = SQLGlobalConfig_Kind + "." + CRDGroupVersion.String()
	SQLGlobalConfig_GroupVersionKind = CRDGroupVersion.WithKind(SQLGlobalConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLGlobalConfig{}, &SQLGlobalConfigList{})
}
