package permission_assignment

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_permission_assignment", func(r *config.Resource) {
		r.ShortGroup = "security"
	})
}
