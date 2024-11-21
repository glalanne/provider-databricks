package custom_app_integration

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_custom_app_integration", func(r *config.Resource) {
		r.ShortGroup = "apps"
	})
}
