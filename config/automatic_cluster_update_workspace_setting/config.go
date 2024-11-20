package automatic_cluster_update_workspace_setting

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_automatic_cluster_update_workspace_setting", func(r *config.Resource) {
		r.ShortGroup = "databricks"
	})
}
