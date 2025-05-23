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

type PrivilegeAssignmentsInitParameters struct {

	// display_name for a databricks_group or databricks_user, application_id for a databricks_service_principal.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// set of available privilege names in upper case.
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`
}

type PrivilegeAssignmentsObservation struct {

	// display_name for a databricks_group or databricks_user, application_id for a databricks_service_principal.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// set of available privilege names in upper case.
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`
}

type PrivilegeAssignmentsParameters struct {

	// display_name for a databricks_group or databricks_user, application_id for a databricks_service_principal.
	// +kubebuilder:validation:Optional
	Principal *string `json:"principal" tf:"principal,omitempty"`

	// set of available privilege names in upper case.
	// +kubebuilder:validation:Optional
	// +listType=set
	Privileges []*string `json:"privileges" tf:"privileges,omitempty"`
}

type SQLPermissionsInitParameters struct {

	// (Boolean) If this access control for using an anonymous function. Defaults to false.
	AnonymousFunction *bool `json:"anonymousFunction,omitempty" tf:"anonymous_function,omitempty"`

	// (Boolean) If this access control for reading/writing any file. Defaults to false.
	AnyFile *bool `json:"anyFile,omitempty" tf:"any_file,omitempty"`

	// (Boolean) If this access control for the entire catalog. Defaults to false.
	Catalog *bool `json:"catalog,omitempty" tf:"catalog,omitempty"`

	// Id of an existing databricks_cluster, where the appropriate GRANT/REVOKE commands are executed. This cluster must have the appropriate data security mode (USER_ISOLATION or LEGACY_TABLE_ACL specified).
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Name of the database. Has a default value of default.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	PrivilegeAssignments []PrivilegeAssignmentsInitParameters `json:"privilegeAssignments,omitempty" tf:"privilege_assignments,omitempty"`

	// Name of the table. Can be combined with the database.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// Name of the view. Can be combined with the database.
	View *string `json:"view,omitempty" tf:"view,omitempty"`
}

type SQLPermissionsObservation struct {

	// (Boolean) If this access control for using an anonymous function. Defaults to false.
	AnonymousFunction *bool `json:"anonymousFunction,omitempty" tf:"anonymous_function,omitempty"`

	// (Boolean) If this access control for reading/writing any file. Defaults to false.
	AnyFile *bool `json:"anyFile,omitempty" tf:"any_file,omitempty"`

	// (Boolean) If this access control for the entire catalog. Defaults to false.
	Catalog *bool `json:"catalog,omitempty" tf:"catalog,omitempty"`

	// Id of an existing databricks_cluster, where the appropriate GRANT/REVOKE commands are executed. This cluster must have the appropriate data security mode (USER_ISOLATION or LEGACY_TABLE_ACL specified).
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Name of the database. Has a default value of default.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PrivilegeAssignments []PrivilegeAssignmentsObservation `json:"privilegeAssignments,omitempty" tf:"privilege_assignments,omitempty"`

	// Name of the table. Can be combined with the database.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// Name of the view. Can be combined with the database.
	View *string `json:"view,omitempty" tf:"view,omitempty"`
}

type SQLPermissionsParameters struct {

	// (Boolean) If this access control for using an anonymous function. Defaults to false.
	// +kubebuilder:validation:Optional
	AnonymousFunction *bool `json:"anonymousFunction,omitempty" tf:"anonymous_function,omitempty"`

	// (Boolean) If this access control for reading/writing any file. Defaults to false.
	// +kubebuilder:validation:Optional
	AnyFile *bool `json:"anyFile,omitempty" tf:"any_file,omitempty"`

	// (Boolean) If this access control for the entire catalog. Defaults to false.
	// +kubebuilder:validation:Optional
	Catalog *bool `json:"catalog,omitempty" tf:"catalog,omitempty"`

	// Id of an existing databricks_cluster, where the appropriate GRANT/REVOKE commands are executed. This cluster must have the appropriate data security mode (USER_ISOLATION or LEGACY_TABLE_ACL specified).
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Name of the database. Has a default value of default.
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// +kubebuilder:validation:Optional
	PrivilegeAssignments []PrivilegeAssignmentsParameters `json:"privilegeAssignments,omitempty" tf:"privilege_assignments,omitempty"`

	// Name of the table. Can be combined with the database.
	// +kubebuilder:validation:Optional
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// Name of the view. Can be combined with the database.
	// +kubebuilder:validation:Optional
	View *string `json:"view,omitempty" tf:"view,omitempty"`
}

// SQLPermissionsSpec defines the desired state of SQLPermissions
type SQLPermissionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLPermissionsParameters `json:"forProvider"`
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
	InitProvider SQLPermissionsInitParameters `json:"initProvider,omitempty"`
}

// SQLPermissionsStatus defines the observed state of SQLPermissions.
type SQLPermissionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLPermissionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SQLPermissions is the Schema for the SQLPermissionss API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type SQLPermissions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLPermissionsSpec   `json:"spec"`
	Status            SQLPermissionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPermissionsList contains a list of SQLPermissionss
type SQLPermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLPermissions `json:"items"`
}

// Repository type metadata.
var (
	SQLPermissions_Kind             = "SQLPermissions"
	SQLPermissions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLPermissions_Kind}.String()
	SQLPermissions_KindAPIVersion   = SQLPermissions_Kind + "." + CRDGroupVersion.String()
	SQLPermissions_GroupVersionKind = CRDGroupVersion.WithKind(SQLPermissions_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLPermissions{}, &SQLPermissionsList{})
}
