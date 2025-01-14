package vector_search_endpoint

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_vector_search_endpoint", func(r *config.Resource) {
		r.ShortGroup = "mosaic"
	})
}
