package permissions

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_permissions", func(r *config.Resource) {
		r.ShortGroup = "security"
		r.References["sql_endpoint_id"] = config.Reference{
			TerraformName: "databricks_sql_endpoint",
		}
		r.References["cluster_id"] = config.Reference{
			TerraformName: "databricks_cluster",
		}
		r.References["cluster_policy_id"] = config.Reference{
			TerraformName: "databricks_cluster_policy",
		}
		r.References["instance_pool_id"] = config.Reference{
			TerraformName: "databricks_instance_pool",
		}
		r.References["job_id"] = config.Reference{
			TerraformName: "databricks_job",
		}
		r.References["notebook_id"] = config.Reference{
			TerraformName: "databricks_notebook",
		}
		r.References["pipeline_id"] = config.Reference{
			TerraformName: "databricks_pipeline",
		}
		r.References["sql_alert_id"] = config.Reference{
			TerraformName: "databricks_sql_alert",
		}
		r.References["sql_dashboard_id"] = config.Reference{
			TerraformName: "databricks_sql_dashboard",
		}
		r.References["sql_query_id"] = config.Reference{
			TerraformName: "databricks_sql_query",
		}
		r.References["access_control.service_principal_name"] = config.Reference{
			TerraformName: "databricks_service_principal",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractParamPath(\"application_id\", false)",
		}
	})
}
