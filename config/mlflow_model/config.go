package mlflow_model

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_mlflow_model", func(r *config.Resource) {
		r.ShortGroup = "mlflow"
	})
}
