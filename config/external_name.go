/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	// "databricks_app":                                            config.IdentifierFromProvider,
	"databricks_cluster":                                        config.IdentifierFromProvider,
	"databricks_cluster_policy":                                 config.IdentifierFromProvider,
	"databricks_credential":                                     config.IdentifierFromProvider,
	"databricks_token":                                          config.IdentifierFromProvider,
	"databricks_secret":                                         config.IdentifierFromProvider,
	"databricks_secret_scope":                                   config.IdentifierFromProvider,
	"databricks_notebook":                                       config.IdentifierFromProvider,
	"databricks_job":                                            config.IdentifierFromProvider,
	"databricks_instance_pool":                                  config.IdentifierFromProvider,
	"databricks_sql_endpoint":                                   config.IdentifierFromProvider,
	"databricks_permissions":                                    config.IdentifierFromProvider,
	"databricks_entitlements":                                   config.IdentifierFromProvider,
	"databricks_group":                                          config.IdentifierFromProvider,
	"databricks_group_member":                                   config.IdentifierFromProvider,
	"databricks_group_role":                                     config.IdentifierFromProvider,
	"databricks_ip_access_list":                                 config.IdentifierFromProvider,
	"databricks_permission_assignment":                          config.IdentifierFromProvider,
	"databricks_service_principal":                              config.IdentifierFromProvider,
	"databricks_service_principal_role":                         config.IdentifierFromProvider,
	"databricks_sql_permissions":                                config.IdentifierFromProvider,
	"databricks_grants":                                         config.IdentifierFromProvider,
	"databricks_pipeline":                                       config.IdentifierFromProvider,
	"databricks_alert":                                          config.IdentifierFromProvider,
	"databricks_query":                                          config.IdentifierFromProvider,
	"databricks_sql_alert":                                      config.IdentifierFromProvider,
	"databricks_sql_dashboard":                                  config.IdentifierFromProvider,
	"databricks_sql_global_config":                              config.IdentifierFromProvider,
	"databricks_sql_query":                                      config.IdentifierFromProvider,
	"databricks_sql_table":                                      config.IdentifierFromProvider,
	"databricks_budget":                                         config.IdentifierFromProvider,
	"databricks_git_credential":                                 config.IdentifierFromProvider,
	"databricks_catalog":                                        config.IdentifierFromProvider,
	"databricks_connection":                                     config.IdentifierFromProvider,
	"databricks_external_location":                              config.IdentifierFromProvider,
	"databricks_schema":                                         config.IdentifierFromProvider,
	"databricks_library":                                        config.IdentifierFromProvider,
	"databricks_sql_visualization":                              config.IdentifierFromProvider,
	"databricks_sql_widget":                                     config.IdentifierFromProvider,
	"databricks_provider":                                       config.IdentifierFromProvider,
	"databricks_mlflow_experiment":                              config.IdentifierFromProvider,
	"databricks_mlflow_model":                                   config.IdentifierFromProvider,
	"databricks_model_serving":                                  config.IdentifierFromProvider,
	"databricks_access_control_rule_set":                        config.IdentifierFromProvider,
	"databricks_artifact_allowlist":                             config.IdentifierFromProvider,
	"databricks_catalog_workspace_binding":                      config.IdentifierFromProvider,
	"databricks_compliance_security_profile_workspace_setting":  config.IdentifierFromProvider,
	"databricks_custom_app_integration":                         config.IdentifierFromProvider,
	"databricks_dashboard":                                      config.IdentifierFromProvider,
	"databricks_dbfs_file":                                      config.IdentifierFromProvider,
	"databricks_default_namespace_setting":                      config.IdentifierFromProvider,
	"databricks_directory":                                      config.IdentifierFromProvider,
	"databricks_enhanced_security_monitoring_workspace_setting": config.IdentifierFromProvider,
	"databricks_file":                                           config.IdentifierFromProvider,
	"databricks_global_init_script":                             config.IdentifierFromProvider,
	"databricks_grant":                                          config.IdentifierFromProvider,
	"databricks_group_instance_profile":                         config.IdentifierFromProvider,
	"databricks_instance_profile":                               config.IdentifierFromProvider,
	"databricks_lakehouse_monitor":                              config.IdentifierFromProvider,
	"databricks_metastore":                                      config.IdentifierFromProvider,
	"databricks_metastore_assignment":                           config.IdentifierFromProvider,
	"databricks_metastore_data_access":                          config.IdentifierFromProvider,
	"databricks_mount":                                          config.IdentifierFromProvider,
	"databricks_mws_credentials":                                config.IdentifierFromProvider,
	"databricks_mws_customer_managed_keys":                      config.IdentifierFromProvider,
	"databricks_mws_log_delivery":                               config.IdentifierFromProvider,
	"databricks_mws_ncc_binding":                                config.IdentifierFromProvider,
	"databricks_mws_ncc_private_endpoint_rule":                  config.IdentifierFromProvider,
	"databricks_mws_network_connectivity_config":                config.IdentifierFromProvider,
	"databricks_mws_networks":                                   config.IdentifierFromProvider,
	"databricks_mws_permission_assignment":                      config.IdentifierFromProvider,
	"databricks_mws_private_access_settings":                    config.IdentifierFromProvider,
	"databricks_mws_storage_configurations":                     config.IdentifierFromProvider,
	"databricks_mws_vpc_endpoint":                               config.IdentifierFromProvider,
	"databricks_mws_workspaces":                                 config.IdentifierFromProvider,
	"databricks_notification_destination":                       config.IdentifierFromProvider,
	"databricks_obo_token":                                      config.IdentifierFromProvider,
	"databricks_online_table":                                   config.IdentifierFromProvider,
	"databricks_quality_monitor":                                config.IdentifierFromProvider,
	"databricks_recipient":                                      config.IdentifierFromProvider,
	"databricks_registered_model":                               config.IdentifierFromProvider,
	"databricks_repo":                                           config.IdentifierFromProvider,
	"databricks_restrict_workspace_admins_setting":              config.IdentifierFromProvider,
	"databricks_secret_acl":                                     config.IdentifierFromProvider,
	"databricks_service_principal_secret":                       config.IdentifierFromProvider,
	"databricks_share":                                          config.IdentifierFromProvider,
	"databricks_storage_credential":                             config.IdentifierFromProvider,
	"databricks_system_schema":                                  config.IdentifierFromProvider,
	"databricks_user":                                           config.IdentifierFromProvider,
	"databricks_user_instance_profile":                          config.IdentifierFromProvider,
	"databricks_user_role":                                      config.IdentifierFromProvider,
	"databricks_vector_search_endpoint":                         config.IdentifierFromProvider,
	"databricks_vector_search_index":                            config.IdentifierFromProvider,
	"databricks_volume":                                         config.IdentifierFromProvider,
	"databricks_workspace_binding":                              config.IdentifierFromProvider,
	"databricks_workspace_conf":                                 config.IdentifierFromProvider,
	"databricks_workspace_file":                                 config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {

			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"

		i++
	}
	return l
}
