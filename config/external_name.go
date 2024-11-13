/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"databricks_cluster":                config.IdentifierFromProvider,
	"databricks_cluster_policy":         config.IdentifierFromProvider,
	"databricks_token":                  config.IdentifierFromProvider,
	"databricks_secret":                 config.IdentifierFromProvider,
	"databricks_secret_scope":           config.IdentifierFromProvider,
	"databricks_notebook":               config.IdentifierFromProvider,
	"databricks_job":                    config.IdentifierFromProvider,
	"databricks_instance_pool":          config.IdentifierFromProvider,
	"databricks_sql_endpoint":           config.IdentifierFromProvider,
	"databricks_permissions":            config.IdentifierFromProvider,
	"databricks_entitlements":           config.IdentifierFromProvider,
	"databricks_group":                  config.IdentifierFromProvider,
	"databricks_group_member":           config.IdentifierFromProvider,
	"databricks_group_role":             config.IdentifierFromProvider,
	"databricks_ip_access_list":         config.IdentifierFromProvider,
	"databricks_permission_assignment":  config.IdentifierFromProvider,
	"databricks_service_principal":      config.IdentifierFromProvider,
	"databricks_service_principal_role": config.IdentifierFromProvider,
	"databricks_sql_permissions":        config.IdentifierFromProvider,
	"databricks_grants":                 config.IdentifierFromProvider,
	"databricks_pipeline":               config.IdentifierFromProvider,
	"databricks_alert":                  config.IdentifierFromProvider,
	"databricks_query":                  config.IdentifierFromProvider,
	"databricks_sql_alert":              config.IdentifierFromProvider,
	"databricks_sql_dashboard":          config.IdentifierFromProvider,
	"databricks_sql_global_config":      config.IdentifierFromProvider,
	"databricks_sql_query":              config.IdentifierFromProvider,
	"databricks_budget":                 config.IdentifierFromProvider,
	"databricks_git_credential":         config.IdentifierFromProvider,
	"databricks_catalog":                config.IdentifierFromProvider,
	"databricks_connection":             config.IdentifierFromProvider,
	"databricks_external_location":      config.IdentifierFromProvider,
	"databricks_schema":                 config.IdentifierFromProvider,
	"databricks_sql_table":              config.IdentifierFromProvider,
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
