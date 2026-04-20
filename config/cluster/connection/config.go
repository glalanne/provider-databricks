package connection

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_connection", func(r *config.Resource) {
		r.ShortGroup = "unity"
		if s, ok := r.TerraformResource.Schema["provisioning_info"]; ok {
			s.Optional = false
			s.Computed = true
		}
	})
}
