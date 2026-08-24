package knowledge_assistant_knowledge_source

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_knowledge_assistant_knowledge_source", func(r *config.Resource) {
		r.ShortGroup = "ai"
	})
}
