package mlflow_webhook

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_mlflow_webhook", func(r *config.Resource) {
		r.ShortGroup = "mlflow"
	})
}
