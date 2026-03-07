package job

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_job", func(r *config.Resource) {
		r.ShortGroup = "compute"

		r.References["notebook_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["dbt_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}
	})
}
