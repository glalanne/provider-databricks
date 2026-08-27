package cluster

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_cluster", configureResource)
}

func configureResource(r *config.Resource) {
	r.ShortGroup = "compute"

	if librarySchema, ok := r.TerraformResource.Schema["library"]; ok && librarySchema.Set != nil {
		libraryHash := librarySchema.Set
		librarySchema.Set = func(value any) int {
			return libraryHash(normalizeLibraryForHash(value))
		}
	}
}

func normalizeLibraryForHash(value any) any {
	library, ok := value.(map[string]any)
	if !ok {
		return value
	}
	normalized := make(map[string]any, len(library))
	for key, field := range library {
		normalized[key] = field
	}
	for _, key := range []string{"pypi", "maven", "cran"} {
		if values, ok := normalized[key].([]any); ok && len(values) == 1 && values[0] == nil {
			normalized[key] = []any{}
		}
	}
	return normalized
}
