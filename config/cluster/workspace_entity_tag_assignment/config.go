package workspace_entity_tag_assignment

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_workspace_entity_tag_assignment", func(r *config.Resource) {
		r.ShortGroup = "tags"
	})
}
