package credential

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_credential", func(r *config.Resource) {
		r.ShortGroup = "unity"
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields,
			"created_at",
			"created_by",
			"full_name",
			"metastore_id",
			"updated_at",
			"updated_by",
			"used_for_managed_storage",
			"aws_iam_role.external_id",
			"aws_iam_role.unity_catalog_iam_arn",
			"azure_managed_identity.credential_id",
			"databricks_gcp_service_account.credential_id",
			"databricks_gcp_service_account.email",
			"databricks_gcp_service_account.private_key_id",
		)
	})
}
