package aibi_dashboard_embedding_access_policy_setting

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_aibi_dashboard_embedding_access_policy_setting", func(r *config.Resource) {
		r.ShortGroup = "settings"
	})
}
