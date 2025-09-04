package restrict_workspace_admins_setting

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_restrict_workspace_admins_setting", func(r *config.Resource) {
		r.ShortGroup = "settings"
	})
}
