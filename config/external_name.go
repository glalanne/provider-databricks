/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

var (
	oldSingletonListAPIs = []string{
		"databricks_cluster",
		"databricks_cluster_policy",
		"databricks_credential",
		"databricks_token",
		"databricks_secret",
		"databricks_secret_scope",
		"databricks_notebook",
		"databricks_job",
		"databricks_instance_pool",
		"databricks_sql_endpoint",
		"databricks_permissions",
		"databricks_entitlements",
		"databricks_group",
		"databricks_group_member",
		"databricks_group_role",
		"databricks_ip_access_list",
		"databricks_permission_assignment",
		"databricks_service_principal",
		"databricks_service_principal_role",
		"databricks_sql_permissions",
		"databricks_grants",
		"databricks_pipeline",
		"databricks_alert",
		"databricks_query",
		"databricks_sql_alert",
		"databricks_sql_dashboard",
		"databricks_sql_global_config",
		"databricks_sql_query",
		"databricks_sql_table",
		"databricks_budget",
		"databricks_git_credential",
		"databricks_catalog",
		"databricks_connection",
		"databricks_external_location",
		"databricks_schema",
		"databricks_library",
		"databricks_sql_visualization",
		"databricks_sql_widget",
		"databricks_provider",
		"databricks_mlflow_experiment",
		"databricks_mlflow_model",
		"databricks_model_serving",
		"databricks_access_control_rule_set",
		"databricks_artifact_allowlist",
		"databricks_catalog_workspace_binding",
		"databricks_compliance_security_profile_workspace_setting",
		"databricks_custom_app_integration",
		"databricks_dashboard",
		"databricks_dbfs_file",
		"databricks_default_namespace_setting",
		"databricks_directory",
		"databricks_enhanced_security_monitoring_workspace_setting",
		"databricks_file",
		"databricks_global_init_script",
		"databricks_grant",
		"databricks_group_instance_profile",
		"databricks_instance_profile",
		"databricks_lakehouse_monitor",
		"databricks_metastore",
		"databricks_metastore_assignment",
		"databricks_metastore_data_access",
		"databricks_mount",
		"databricks_mlflow_webhook",
		"databricks_mws_credentials",
		"databricks_mws_customer_managed_keys",
		"databricks_mws_log_delivery",
		"databricks_mws_ncc_binding",
		"databricks_mws_ncc_private_endpoint_rule",
		"databricks_mws_network_connectivity_config",
		"databricks_mws_networks",
		"databricks_mws_permission_assignment",
		"databricks_mws_private_access_settings",
		"databricks_mws_storage_configurations",
		"databricks_mws_vpc_endpoint",
		"databricks_mws_workspaces",
		"databricks_notification_destination",
		"databricks_obo_token",
		"databricks_online_table",
		"databricks_quality_monitor",
		"databricks_recipient",
		"databricks_registered_model",
		"databricks_repo",
		"databricks_restrict_workspace_admins_setting",
		"databricks_secret_acl",
		"databricks_service_principal_secret",
		"databricks_storage_credential",
		"databricks_system_schema",
		"databricks_user",
		"databricks_user_instance_profile",
		"databricks_user_role",
		"databricks_vector_search_endpoint",
		"databricks_vector_search_index",
		"databricks_volume",
		"databricks_workspace_binding",
		"databricks_workspace_conf",
		"databricks_workspace_file",
		"databricks_app",
		"databricks_share",
		"databricks_service_principal_federation_policy",
		"databricks_account_federation_policy",
		"databricks_database_instance",
		"databricks_policy_info",
	}
)

// TerraformPluginSDKExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the no-fork
// architecture for this provider.
var TerraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"databricks_cluster":                                           config.IdentifierFromProvider,
	"databricks_cluster_policy":                                    config.IdentifierFromProvider,
	"databricks_credential":                                        config.IdentifierFromProvider,
	"databricks_token":                                             config.IdentifierFromProvider,
	"databricks_secret":                                            config.IdentifierFromProvider,
	"databricks_secret_scope":                                      config.IdentifierFromProvider,
	"databricks_notebook":                                          config.IdentifierFromProvider,
	"databricks_job":                                               config.IdentifierFromProvider,
	"databricks_instance_pool":                                     config.IdentifierFromProvider,
	"databricks_sql_endpoint":                                      config.IdentifierFromProvider,
	"databricks_permissions":                                       config.IdentifierFromProvider,
	"databricks_entitlements":                                      config.IdentifierFromProvider,
	"databricks_group":                                             config.IdentifierFromProvider,
	"databricks_group_member":                                      config.IdentifierFromProvider,
	"databricks_group_role":                                        config.IdentifierFromProvider,
	"databricks_ip_access_list":                                    config.IdentifierFromProvider,
	"databricks_permission_assignment":                             config.IdentifierFromProvider,
	"databricks_service_principal":                                 config.IdentifierFromProvider,
	"databricks_service_principal_role":                            config.IdentifierFromProvider,
	"databricks_sql_permissions":                                   config.IdentifierFromProvider,
	"databricks_grants":                                            config.IdentifierFromProvider,
	"databricks_pipeline":                                          config.IdentifierFromProvider,
	"databricks_alert":                                             config.IdentifierFromProvider,
	"databricks_query":                                             config.IdentifierFromProvider,
	"databricks_sql_alert":                                         config.IdentifierFromProvider,
	"databricks_sql_dashboard":                                     config.IdentifierFromProvider,
	"databricks_sql_global_config":                                 config.IdentifierFromProvider,
	"databricks_sql_query":                                         config.IdentifierFromProvider,
	"databricks_sql_table":                                         config.IdentifierFromProvider,
	"databricks_budget":                                            config.IdentifierFromProvider,
	"databricks_git_credential":                                    config.IdentifierFromProvider,
	"databricks_catalog":                                           config.IdentifierFromProvider,
	"databricks_connection":                                        config.IdentifierFromProvider,
	"databricks_external_location":                                 config.IdentifierFromProvider,
	"databricks_schema":                                            config.IdentifierFromProvider,
	"databricks_library":                                           config.IdentifierFromProvider,
	"databricks_sql_visualization":                                 config.IdentifierFromProvider,
	"databricks_sql_widget":                                        config.IdentifierFromProvider,
	"databricks_provider":                                          config.IdentifierFromProvider,
	"databricks_mlflow_experiment":                                 config.IdentifierFromProvider,
	"databricks_mlflow_model":                                      config.IdentifierFromProvider,
	"databricks_model_serving":                                     config.IdentifierFromProvider,
	"databricks_access_control_rule_set":                           config.IdentifierFromProvider,
	"databricks_artifact_allowlist":                                config.IdentifierFromProvider,
	"databricks_catalog_workspace_binding":                         config.IdentifierFromProvider,
	"databricks_compliance_security_profile_workspace_setting":     config.IdentifierFromProvider,
	"databricks_custom_app_integration":                            config.IdentifierFromProvider,
	"databricks_dashboard":                                         config.IdentifierFromProvider,
	"databricks_dbfs_file":                                         config.IdentifierFromProvider,
	"databricks_default_namespace_setting":                         config.IdentifierFromProvider,
	"databricks_directory":                                         config.IdentifierFromProvider,
	"databricks_enhanced_security_monitoring_workspace_setting":    config.IdentifierFromProvider,
	"databricks_file":                                              config.IdentifierFromProvider,
	"databricks_global_init_script":                                config.IdentifierFromProvider,
	"databricks_grant":                                             config.IdentifierFromProvider,
	"databricks_group_instance_profile":                            config.IdentifierFromProvider,
	"databricks_instance_profile":                                  config.IdentifierFromProvider,
	"databricks_lakehouse_monitor":                                 config.IdentifierFromProvider,
	"databricks_metastore":                                         config.IdentifierFromProvider,
	"databricks_metastore_assignment":                              config.IdentifierFromProvider,
	"databricks_metastore_data_access":                             config.IdentifierFromProvider,
	"databricks_mount":                                             config.IdentifierFromProvider,
	"databricks_mlflow_webhook":                                    config.IdentifierFromProvider,
	"databricks_mws_credentials":                                   config.IdentifierFromProvider,
	"databricks_mws_customer_managed_keys":                         config.IdentifierFromProvider,
	"databricks_mws_log_delivery":                                  config.IdentifierFromProvider,
	"databricks_mws_ncc_binding":                                   config.IdentifierFromProvider,
	"databricks_mws_ncc_private_endpoint_rule":                     config.IdentifierFromProvider,
	"databricks_mws_network_connectivity_config":                   config.IdentifierFromProvider,
	"databricks_mws_networks":                                      config.IdentifierFromProvider,
	"databricks_mws_permission_assignment":                         config.IdentifierFromProvider,
	"databricks_mws_private_access_settings":                       config.IdentifierFromProvider,
	"databricks_mws_storage_configurations":                        config.IdentifierFromProvider,
	"databricks_mws_vpc_endpoint":                                  config.IdentifierFromProvider,
	"databricks_mws_workspaces":                                    config.IdentifierFromProvider,
	"databricks_notification_destination":                          config.IdentifierFromProvider,
	"databricks_obo_token":                                         config.IdentifierFromProvider,
	"databricks_online_table":                                      config.IdentifierFromProvider,
	"databricks_quality_monitor":                                   config.IdentifierFromProvider,
	"databricks_recipient":                                         config.IdentifierFromProvider,
	"databricks_registered_model":                                  config.IdentifierFromProvider,
	"databricks_repo":                                              config.IdentifierFromProvider,
	"databricks_restrict_workspace_admins_setting":                 config.IdentifierFromProvider,
	"databricks_secret_acl":                                        config.IdentifierFromProvider,
	"databricks_service_principal_secret":                          config.IdentifierFromProvider,
	"databricks_storage_credential":                                config.IdentifierFromProvider,
	"databricks_system_schema":                                     config.IdentifierFromProvider,
	"databricks_user":                                              config.IdentifierFromProvider,
	"databricks_user_instance_profile":                             config.IdentifierFromProvider,
	"databricks_user_role":                                         config.IdentifierFromProvider,
	"databricks_vector_search_endpoint":                            config.IdentifierFromProvider,
	"databricks_vector_search_index":                               config.IdentifierFromProvider,
	"databricks_volume":                                            config.IdentifierFromProvider,
	"databricks_workspace_binding":                                 config.IdentifierFromProvider,
	"databricks_workspace_conf":                                    config.IdentifierFromProvider,
	"databricks_workspace_file":                                    config.IdentifierFromProvider,
	"databricks_data_classification_catalog_config":                config.IdentifierFromProvider,
	"databricks_workspace_entity_tag_assignment":                   config.IdentifierFromProvider,
	"databricks_aibi_dashboard_embedding_access_policy_setting":    config.IdentifierFromProvider,
	"databricks_environments_default_workspace_base_environment":   config.IdentifierFromProvider,
	"databricks_postgres_role":                                     config.IdentifierFromProvider,
	"databricks_rfa_access_request_destinations":                   config.IdentifierFromProvider,
	"databricks_postgres_synced_table":                             config.IdentifierFromProvider,
	"databricks_account_network_policy":                            config.IdentifierFromProvider,
	"databricks_postgres_endpoint":                                 config.IdentifierFromProvider,
	"databricks_compliance_security_profile_setting":               config.IdentifierFromProvider,
	"databricks_disable_legacy_access_setting":                     config.IdentifierFromProvider,
	"databricks_aibi_dashboard_embedding_approved_domains_setting": config.IdentifierFromProvider,
	"databricks_disable_legacy_features_setting":                   config.IdentifierFromProvider,
	"databricks_supervisor_agent_tool":                             config.IdentifierFromProvider,
	"databricks_environments_workspace_base_environment":           config.IdentifierFromProvider,
	"databricks_alert_v2":                                          config.IdentifierFromProvider,
	"databricks_disaster_recovery_failover_group":                  config.IdentifierFromProvider,
	"databricks_postgres_project":                                  config.IdentifierFromProvider,
	"databricks_supervisor_agent":                                  config.IdentifierFromProvider,
	"databricks_knowledge_assistant":                               config.IdentifierFromProvider,
	"databricks_account_setting_v2":                                config.IdentifierFromProvider,
	"databricks_postgres_cdf_config":                               config.IdentifierFromProvider,
	"databricks_postgres_branch":                                   config.IdentifierFromProvider,
	"databricks_budget_policy":                                     config.IdentifierFromProvider,
	"databricks_postgres_catalog":                                  config.IdentifierFromProvider,
	"databricks_postgres_database":                                 config.IdentifierFromProvider,
	"databricks_workspace_network_option":                          config.IdentifierFromProvider,
	"databricks_knowledge_assistant_knowledge_source":              config.IdentifierFromProvider,
	"databricks_tag_policy":                                        config.IdentifierFromProvider,
	"databricks_entity_tag_assignment":                             config.IdentifierFromProvider,
	"databricks_secret_uc":                                         config.IdentifierFromProvider,
	"databricks_workspace_setting_v2":                              config.IdentifierFromProvider,
	"databricks_automatic_cluster_update_setting":                  config.IdentifierFromProvider,
	"databricks_ai_search_endpoint":                                config.IdentifierFromProvider,
	"databricks_data_quality_refresh":                              config.IdentifierFromProvider,
	"databricks_ai_search_index":                                   config.IdentifierFromProvider,
	"databricks_disaster_recovery_stable_url":                      config.IdentifierFromProvider,
	"databricks_external_metadata":                                 config.IdentifierFromProvider,
	"databricks_disable_legacy_dbfs_setting":                       config.IdentifierFromProvider,
}

var TerraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{
	"databricks_app":   config.IdentifierFromProvider,
	"databricks_share": config.IdentifierFromProvider,
	"databricks_service_principal_federation_policy": config.IdentifierFromProvider,
	"databricks_account_federation_policy":           config.IdentifierFromProvider,
	"databricks_database_instance":                   config.IdentifierFromProvider,
	"databricks_policy_info":                         config.IdentifierFromProvider,
}

var CLIReconciledExternalNameConfigs = map[string]config.ExternalName{}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := TerraformPluginSDKExternalNameConfigs[r.Name]; ok {

			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(TerraformPluginSDKExternalNameConfigs))
	i := 0
	for name := range TerraformPluginSDKExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"

		i++
	}
	return l
}

// ResourceConfigurator applies all external name configs
// listed in the table TerraformPluginSDKExternalNameConfigs and
// CLIReconciledExternalNameConfigs and sets the version
// of those resources to v1beta1. For those resource in
// TerraformPluginSDKExternalNameConfigs, it also sets
// config.Resource.UseNoForkClient to `true`.
func ResourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// if configured both for the no-fork and CLI based architectures,
		// no-fork configuration prevails
		e, configured := TerraformPluginSDKExternalNameConfigs[r.Name]
		if !configured {
			e, configured = CLIReconciledExternalNameConfigs[r.Name]
		}
		if !configured {
			return
		}

		r.Version = "v1beta1"
		r.ExternalName = e
	}
}
