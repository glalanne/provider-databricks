package job

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_job", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields, "format")

		r.References["notebook_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["dbt_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["sql_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.0.sql_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["dashboard_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["power_bi_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.0.notebook_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.0.dbt_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.0.sql_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.0.dashboard_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.0.power_bi_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["for_each_task.0.task.0.notebook_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["for_each_task.0.task.0.dbt_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["for_each_task.0.task.0.sql_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["for_each_task.0.task.0.dashboard_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["for_each_task.0.task.0.power_bi_task.0.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}
	})
}
