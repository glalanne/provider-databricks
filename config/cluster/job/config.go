package job

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_job", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields, "format")

		r.References["notebook_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["dbt_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["sql_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.sql_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["dashboard_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["power_bi_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.notebook_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.dbt_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.sql_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.dashboard_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["task.power_bi_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["for_each_task.task.notebook_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["for_each_task.task.dbt_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["for_each_task.task.sql_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["for_each_task.task.dashboard_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}

		r.References["for_each_task.task.power_bi_task.warehouse_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}
	})
}
