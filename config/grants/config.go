package grants

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_grants", func(r *config.Resource) {
		r.ShortGroup = "unity"
		// (glalanne) Renaming this resource to avoid a colision with databricks_grant
		// because controller-gen would just evict the CRD generation because
		// the plural of grant is grants and creates a conflict with databricks_grants
		r.Kind = "GrantMap"
	})
}
