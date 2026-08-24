package workspace_network_option

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_workspace_network_option", func(r *config.Resource) {
		r.ShortGroup = "settings"
	})
}
