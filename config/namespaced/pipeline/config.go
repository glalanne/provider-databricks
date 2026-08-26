package pipeline

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_pipeline", func(r *config.Resource) {
		r.ShortGroup = "compute"

		r.ExternalName.OmittedFields = []string{
			"last_modified",
		}

		// Move the state field to computed so that it is not required in the spec.
		if s, ok := r.TerraformResource.Schema["state"]; ok {
			s.Optional = false
			s.Computed = true
		}
	})
}
