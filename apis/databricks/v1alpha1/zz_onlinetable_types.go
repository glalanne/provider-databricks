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

type ContinuousUpdateStatusInitParameters struct {
}

type ContinuousUpdateStatusObservation struct {
	InitialPipelineSyncProgress []InitialPipelineSyncProgressObservation `json:"initialPipelineSyncProgress,omitempty" tf:"initial_pipeline_sync_progress,omitempty"`

	LastProcessedCommitVersion *float64 `json:"lastProcessedCommitVersion,omitempty" tf:"last_processed_commit_version,omitempty"`

	Timestamp *string `json:"timestamp,omitempty" tf:"timestamp,omitempty"`
}

type ContinuousUpdateStatusParameters struct {
}

type FailedStatusInitParameters struct {
}

type FailedStatusObservation struct {
	LastProcessedCommitVersion *float64 `json:"lastProcessedCommitVersion,omitempty" tf:"last_processed_commit_version,omitempty"`

	Timestamp *string `json:"timestamp,omitempty" tf:"timestamp,omitempty"`
}

type FailedStatusParameters struct {
}

type InitialPipelineSyncProgressInitParameters struct {
}

type InitialPipelineSyncProgressObservation struct {
	EstimatedCompletionTimeSeconds *float64 `json:"estimatedCompletionTimeSeconds,omitempty" tf:"estimated_completion_time_seconds,omitempty"`

	LatestVersionCurrentlyProcessing *float64 `json:"latestVersionCurrentlyProcessing,omitempty" tf:"latest_version_currently_processing,omitempty"`

	SyncProgressCompletion *float64 `json:"syncProgressCompletion,omitempty" tf:"sync_progress_completion,omitempty"`

	SyncedRowCount *float64 `json:"syncedRowCount,omitempty" tf:"synced_row_count,omitempty"`

	TotalRowCount *float64 `json:"totalRowCount,omitempty" tf:"total_row_count,omitempty"`
}

type InitialPipelineSyncProgressParameters struct {
}

type OnlineTableInitParameters struct {

	// 3-level name of the Online Table to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// object containing specification of the online table:
	Spec []OnlineTableSpecInitParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type OnlineTableObservation struct {

	// The same as the name of the online table.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// 3-level name of the Online Table to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// object containing specification of the online table:
	Spec []OnlineTableSpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`

	// object describing status of the online table:
	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`

	// Data serving REST API URL for this table.
	TableServingURL *string `json:"tableServingUrl,omitempty" tf:"table_serving_url,omitempty"`

	// The provisioning state of the online table entity in Unity Catalog. This is distinct from the state of the data synchronization pipeline (i.e. the table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it runs asynchronously).
	UnityCatalogProvisioningState *string `json:"unityCatalogProvisioningState,omitempty" tf:"unity_catalog_provisioning_state,omitempty"`
}

type OnlineTableParameters struct {

	// 3-level name of the Online Table to create.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// object containing specification of the online table:
	// +kubebuilder:validation:Optional
	Spec []OnlineTableSpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type OnlineTableSpecInitParameters struct {

	// Whether to create a full-copy pipeline -- a pipeline that stops after creates a full copy of the source table upon initialization and does not process any change data feeds (CDFs) afterwards. The pipeline can still be manually triggered afterwards, but it always perform a full copy of the source table and there are no incremental updates. This mode is useful for syncing views or tables without CDFs to online tables. Note that the full-copy pipeline only supports "triggered" scheduling policy.
	PerformFullCopy *bool `json:"performFullCopy,omitempty" tf:"perform_full_copy,omitempty"`

	// list of the columns comprising the primary key.
	PrimaryKeyColumns []*string `json:"primaryKeyColumns,omitempty" tf:"primary_key_columns,omitempty"`

	// empty block that specifies that pipeline runs continuously after generating the initial data.  Conflicts with run_triggered.
	RunContinuously []RunContinuouslyInitParameters `json:"runContinuously,omitempty" tf:"run_continuously,omitempty"`

	// empty block that specifies that pipeline stops after generating the initial data and can be triggered later (manually, through a cron job or through data triggers).
	RunTriggered []RunTriggeredInitParameters `json:"runTriggered,omitempty" tf:"run_triggered,omitempty"`

	// full name of the source table.
	SourceTableFullName *string `json:"sourceTableFullName,omitempty" tf:"source_table_full_name,omitempty"`

	// Time series key to deduplicate (tie-break) rows with the same primary key.
	TimeseriesKey *string `json:"timeseriesKey,omitempty" tf:"timeseries_key,omitempty"`
}

type OnlineTableSpecObservation struct {

	// Whether to create a full-copy pipeline -- a pipeline that stops after creates a full copy of the source table upon initialization and does not process any change data feeds (CDFs) afterwards. The pipeline can still be manually triggered afterwards, but it always perform a full copy of the source table and there are no incremental updates. This mode is useful for syncing views or tables without CDFs to online tables. Note that the full-copy pipeline only supports "triggered" scheduling policy.
	PerformFullCopy *bool `json:"performFullCopy,omitempty" tf:"perform_full_copy,omitempty"`

	// ID of the associated Delta Live Table pipeline.
	PipelineID *string `json:"pipelineId,omitempty" tf:"pipeline_id,omitempty"`

	// list of the columns comprising the primary key.
	PrimaryKeyColumns []*string `json:"primaryKeyColumns,omitempty" tf:"primary_key_columns,omitempty"`

	// empty block that specifies that pipeline runs continuously after generating the initial data.  Conflicts with run_triggered.
	RunContinuously []RunContinuouslyParameters `json:"runContinuously,omitempty" tf:"run_continuously,omitempty"`

	// empty block that specifies that pipeline stops after generating the initial data and can be triggered later (manually, through a cron job or through data triggers).
	RunTriggered []RunTriggeredParameters `json:"runTriggered,omitempty" tf:"run_triggered,omitempty"`

	// full name of the source table.
	SourceTableFullName *string `json:"sourceTableFullName,omitempty" tf:"source_table_full_name,omitempty"`

	// Time series key to deduplicate (tie-break) rows with the same primary key.
	TimeseriesKey *string `json:"timeseriesKey,omitempty" tf:"timeseries_key,omitempty"`
}

type OnlineTableSpecParameters struct {

	// Whether to create a full-copy pipeline -- a pipeline that stops after creates a full copy of the source table upon initialization and does not process any change data feeds (CDFs) afterwards. The pipeline can still be manually triggered afterwards, but it always perform a full copy of the source table and there are no incremental updates. This mode is useful for syncing views or tables without CDFs to online tables. Note that the full-copy pipeline only supports "triggered" scheduling policy.
	// +kubebuilder:validation:Optional
	PerformFullCopy *bool `json:"performFullCopy,omitempty" tf:"perform_full_copy,omitempty"`

	// list of the columns comprising the primary key.
	// +kubebuilder:validation:Optional
	PrimaryKeyColumns []*string `json:"primaryKeyColumns,omitempty" tf:"primary_key_columns,omitempty"`

	// empty block that specifies that pipeline runs continuously after generating the initial data.  Conflicts with run_triggered.
	// +kubebuilder:validation:Optional
	RunContinuously []RunContinuouslyParameters `json:"runContinuously,omitempty" tf:"run_continuously,omitempty"`

	// empty block that specifies that pipeline stops after generating the initial data and can be triggered later (manually, through a cron job or through data triggers).
	// +kubebuilder:validation:Optional
	RunTriggered []RunTriggeredParameters `json:"runTriggered,omitempty" tf:"run_triggered,omitempty"`

	// full name of the source table.
	// +kubebuilder:validation:Optional
	SourceTableFullName *string `json:"sourceTableFullName,omitempty" tf:"source_table_full_name,omitempty"`

	// Time series key to deduplicate (tie-break) rows with the same primary key.
	// +kubebuilder:validation:Optional
	TimeseriesKey *string `json:"timeseriesKey,omitempty" tf:"timeseries_key,omitempty"`
}

type ProvisioningStatusInitParameters struct {
}

type ProvisioningStatusInitialPipelineSyncProgressInitParameters struct {
}

type ProvisioningStatusInitialPipelineSyncProgressObservation struct {
	EstimatedCompletionTimeSeconds *float64 `json:"estimatedCompletionTimeSeconds,omitempty" tf:"estimated_completion_time_seconds,omitempty"`

	LatestVersionCurrentlyProcessing *float64 `json:"latestVersionCurrentlyProcessing,omitempty" tf:"latest_version_currently_processing,omitempty"`

	SyncProgressCompletion *float64 `json:"syncProgressCompletion,omitempty" tf:"sync_progress_completion,omitempty"`

	SyncedRowCount *float64 `json:"syncedRowCount,omitempty" tf:"synced_row_count,omitempty"`

	TotalRowCount *float64 `json:"totalRowCount,omitempty" tf:"total_row_count,omitempty"`
}

type ProvisioningStatusInitialPipelineSyncProgressParameters struct {
}

type ProvisioningStatusObservation struct {
	InitialPipelineSyncProgress []ProvisioningStatusInitialPipelineSyncProgressObservation `json:"initialPipelineSyncProgress,omitempty" tf:"initial_pipeline_sync_progress,omitempty"`
}

type ProvisioningStatusParameters struct {
}

type RunContinuouslyInitParameters struct {
}

type RunContinuouslyObservation struct {
}

type RunContinuouslyParameters struct {
}

type RunTriggeredInitParameters struct {
}

type RunTriggeredObservation struct {
}

type RunTriggeredParameters struct {
}

type StatusInitParameters struct {
}

type StatusObservation struct {

	// object describing status of the online table:
	ContinuousUpdateStatus []ContinuousUpdateStatusObservation `json:"continuousUpdateStatus,omitempty" tf:"continuous_update_status,omitempty"`

	// The state of the online table.
	DetailedState *string `json:"detailedState,omitempty" tf:"detailed_state,omitempty"`

	// object describing status of the online table:
	FailedStatus []FailedStatusObservation `json:"failedStatus,omitempty" tf:"failed_status,omitempty"`

	// A text description of the current state of the online table.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// object describing status of the online table:
	ProvisioningStatus []ProvisioningStatusObservation `json:"provisioningStatus,omitempty" tf:"provisioning_status,omitempty"`

	// object describing status of the online table:
	TriggeredUpdateStatus []TriggeredUpdateStatusObservation `json:"triggeredUpdateStatus,omitempty" tf:"triggered_update_status,omitempty"`
}

type StatusParameters struct {
}

type TriggeredUpdateProgressInitParameters struct {
}

type TriggeredUpdateProgressObservation struct {
	EstimatedCompletionTimeSeconds *float64 `json:"estimatedCompletionTimeSeconds,omitempty" tf:"estimated_completion_time_seconds,omitempty"`

	LatestVersionCurrentlyProcessing *float64 `json:"latestVersionCurrentlyProcessing,omitempty" tf:"latest_version_currently_processing,omitempty"`

	SyncProgressCompletion *float64 `json:"syncProgressCompletion,omitempty" tf:"sync_progress_completion,omitempty"`

	SyncedRowCount *float64 `json:"syncedRowCount,omitempty" tf:"synced_row_count,omitempty"`

	TotalRowCount *float64 `json:"totalRowCount,omitempty" tf:"total_row_count,omitempty"`
}

type TriggeredUpdateProgressParameters struct {
}

type TriggeredUpdateStatusInitParameters struct {
}

type TriggeredUpdateStatusObservation struct {
	LastProcessedCommitVersion *float64 `json:"lastProcessedCommitVersion,omitempty" tf:"last_processed_commit_version,omitempty"`

	Timestamp *string `json:"timestamp,omitempty" tf:"timestamp,omitempty"`

	TriggeredUpdateProgress []TriggeredUpdateProgressObservation `json:"triggeredUpdateProgress,omitempty" tf:"triggered_update_progress,omitempty"`
}

type TriggeredUpdateStatusParameters struct {
}

// OnlineTableSpec defines the desired state of OnlineTable
type OnlineTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OnlineTableParameters `json:"forProvider"`
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
	InitProvider OnlineTableInitParameters `json:"initProvider,omitempty"`
}

// OnlineTableStatus defines the observed state of OnlineTable.
type OnlineTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OnlineTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OnlineTable is the Schema for the OnlineTables API. ""subcategory: "Unity Catalog"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type OnlineTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   OnlineTableSpec   `json:"spec"`
	Status OnlineTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OnlineTableList contains a list of OnlineTables
type OnlineTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OnlineTable `json:"items"`
}

// Repository type metadata.
var (
	OnlineTable_Kind             = "OnlineTable"
	OnlineTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OnlineTable_Kind}.String()
	OnlineTable_KindAPIVersion   = OnlineTable_Kind + "." + CRDGroupVersion.String()
	OnlineTable_GroupVersionKind = CRDGroupVersion.WithKind(OnlineTable_Kind)
)

func init() {
	SchemeBuilder.Register(&OnlineTable{}, &OnlineTableList{})
}