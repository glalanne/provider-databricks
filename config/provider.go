/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	"context"
	_ "embed"
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	"github.com/crossplane/upjet/v2/pkg/types/name"
	"github.com/glalanne/provider-databricks/config/access_control_rule_set"
	"github.com/glalanne/provider-databricks/config/alert"
	"github.com/glalanne/provider-databricks/config/artifact_allowlist"
	"github.com/glalanne/provider-databricks/config/budget"
	"github.com/glalanne/provider-databricks/config/catalog"
	"github.com/glalanne/provider-databricks/config/catalog_workspace_binding"
	"github.com/glalanne/provider-databricks/config/cluster"
	"github.com/glalanne/provider-databricks/config/cluster_policy"
	"github.com/glalanne/provider-databricks/config/compliance_security_profile_workspace_setting"
	"github.com/glalanne/provider-databricks/config/connection"
	"github.com/glalanne/provider-databricks/config/credential"
	"github.com/glalanne/provider-databricks/config/custom_app_integration"
	"github.com/glalanne/provider-databricks/config/dashboard"
	"github.com/glalanne/provider-databricks/config/dbfs_file"
	"github.com/glalanne/provider-databricks/config/default_namespace_setting"
	"github.com/glalanne/provider-databricks/config/directory"
	"github.com/glalanne/provider-databricks/config/enhanced_security_monitoring_workspace_setting"
	"github.com/glalanne/provider-databricks/config/entitlements"
	"github.com/glalanne/provider-databricks/config/external_location"
	"github.com/glalanne/provider-databricks/config/file"
	"github.com/glalanne/provider-databricks/config/git_credential"
	"github.com/glalanne/provider-databricks/config/global_init_script"
	"github.com/glalanne/provider-databricks/config/grant"
	"github.com/glalanne/provider-databricks/config/grants"
	"github.com/glalanne/provider-databricks/config/group"
	"github.com/glalanne/provider-databricks/config/group_instance_profile"
	"github.com/glalanne/provider-databricks/config/group_member"
	"github.com/glalanne/provider-databricks/config/group_role"
	"github.com/glalanne/provider-databricks/config/instance_pool"
	"github.com/glalanne/provider-databricks/config/instance_profile"
	"github.com/glalanne/provider-databricks/config/ip_access_list"
	"github.com/glalanne/provider-databricks/config/job"
	"github.com/glalanne/provider-databricks/config/lakehouse_monitor"
	"github.com/glalanne/provider-databricks/config/library"
	"github.com/glalanne/provider-databricks/config/metastore"
	"github.com/glalanne/provider-databricks/config/metastore_assignment"
	"github.com/glalanne/provider-databricks/config/metastore_data_access"
	"github.com/glalanne/provider-databricks/config/mlflow_experiment"
	"github.com/glalanne/provider-databricks/config/mlflow_model"
	"github.com/glalanne/provider-databricks/config/mlflow_webhook"
	"github.com/glalanne/provider-databricks/config/model_serving"
	"github.com/glalanne/provider-databricks/config/mount"
	"github.com/glalanne/provider-databricks/config/mws_credentials"
	"github.com/glalanne/provider-databricks/config/mws_customer_managed_keys"
	"github.com/glalanne/provider-databricks/config/mws_log_delivery"
	"github.com/glalanne/provider-databricks/config/mws_ncc_binding"
	"github.com/glalanne/provider-databricks/config/mws_ncc_private_endpoint_rule"
	"github.com/glalanne/provider-databricks/config/mws_network_connectivity_config"
	"github.com/glalanne/provider-databricks/config/mws_networks"
	"github.com/glalanne/provider-databricks/config/mws_permission_assignment"
	"github.com/glalanne/provider-databricks/config/mws_private_access_settings"
	"github.com/glalanne/provider-databricks/config/mws_storage_configurations"
	"github.com/glalanne/provider-databricks/config/mws_vpc_endpoint"
	"github.com/glalanne/provider-databricks/config/mws_workspaces"
	"github.com/glalanne/provider-databricks/config/notebook"
	"github.com/glalanne/provider-databricks/config/notification_destination"
	"github.com/glalanne/provider-databricks/config/obo_token"
	"github.com/glalanne/provider-databricks/config/online_table"
	"github.com/glalanne/provider-databricks/config/permission_assignment"
	"github.com/glalanne/provider-databricks/config/permissions"
	"github.com/glalanne/provider-databricks/config/pipeline"
	"github.com/glalanne/provider-databricks/config/provider"
	"github.com/glalanne/provider-databricks/config/quality_monitor"
	"github.com/glalanne/provider-databricks/config/query"
	"github.com/glalanne/provider-databricks/config/recipient"
	"github.com/glalanne/provider-databricks/config/registered_model"
	"github.com/glalanne/provider-databricks/config/repo"
	"github.com/glalanne/provider-databricks/config/restrict_workspace_admins_setting"
	"github.com/glalanne/provider-databricks/config/schema"
	"github.com/glalanne/provider-databricks/config/secret"
	"github.com/glalanne/provider-databricks/config/secret_acl"
	"github.com/glalanne/provider-databricks/config/secret_scope"
	"github.com/glalanne/provider-databricks/config/service_principal"
	"github.com/glalanne/provider-databricks/config/service_principal_role"
	"github.com/glalanne/provider-databricks/config/service_principal_secret"
	"github.com/glalanne/provider-databricks/config/share"
	"github.com/glalanne/provider-databricks/config/sql_alert"
	"github.com/glalanne/provider-databricks/config/sql_dashboard"
	"github.com/glalanne/provider-databricks/config/sql_endpoint"
	"github.com/glalanne/provider-databricks/config/sql_global_config"
	"github.com/glalanne/provider-databricks/config/sql_permissions"
	"github.com/glalanne/provider-databricks/config/sql_query"
	"github.com/glalanne/provider-databricks/config/sql_table"
	"github.com/glalanne/provider-databricks/config/sql_visualization"
	"github.com/glalanne/provider-databricks/config/sql_widget"
	"github.com/glalanne/provider-databricks/config/storage_credential"
	"github.com/glalanne/provider-databricks/config/system_schema"
	"github.com/glalanne/provider-databricks/config/token"
	"github.com/glalanne/provider-databricks/config/user"
	"github.com/glalanne/provider-databricks/config/user_instance_profile"
	"github.com/glalanne/provider-databricks/config/user_role"
	"github.com/glalanne/provider-databricks/config/vector_search_endpoint"
	"github.com/glalanne/provider-databricks/config/vector_search_index"
	"github.com/glalanne/provider-databricks/config/volume"
	"github.com/glalanne/provider-databricks/config/workspace_binding"
	"github.com/glalanne/provider-databricks/config/workspace_conf"
	"github.com/glalanne/provider-databricks/config/workspace_file"
	tfjson "github.com/hashicorp/terraform-json"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/pkg/errors"
)

const (
	resourcePrefix = "databricks"
	modulePath     = "github.com/glalanne/provider-databricks"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

func getProviderSchema(s string) (*tfschema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &tfschema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(_ context.Context, sdkProvider *tfschema.Provider, generationProvider bool) (*ujconfig.Provider, error) {

	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		// ujconfig.WithShortName("databricks"),
		ujconfig.WithRootGroup("databricks.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// app.Configure,
		cluster.Configure,
		cluster_policy.Configure,
		credential.Configure,
		token.Configure,
		secret.Configure,
		secret_scope.Configure,
		notebook.Configure,
		job.Configure,
		instance_pool.Configure,
		permissions.Configure,
		sql_endpoint.Configure,
		entitlements.Configure,
		group.Configure,
		group_member.Configure,
		group_role.Configure,
		ip_access_list.Configure,
		permission_assignment.Configure,
		service_principal.Configure,
		service_principal_role.Configure,
		sql_permissions.Configure,
		grants.Configure,
		pipeline.Configure,
		alert.Configure,
		query.Configure,
		sql_alert.Configure,
		sql_dashboard.Configure,
		sql_global_config.Configure,
		sql_query.Configure,
		sql_table.Configure,
		budget.Configure,
		git_credential.Configure,
		catalog.Configure,
		connection.Configure,
		external_location.Configure,
		schema.Configure,
		library.Configure,
		sql_visualization.Configure,
		sql_widget.Configure,
		provider.Configure,
		mlflow_experiment.Configure,
		mlflow_model.Configure,
		mlflow_webhook.Configure,
		model_serving.Configure,
		access_control_rule_set.Configure,
		artifact_allowlist.Configure,
		catalog_workspace_binding.Configure,
		compliance_security_profile_workspace_setting.Configure,
		custom_app_integration.Configure,
		dashboard.Configure,
		dbfs_file.Configure,
		default_namespace_setting.Configure,
		directory.Configure,
		enhanced_security_monitoring_workspace_setting.Configure,
		file.Configure,
		global_init_script.Configure,
		grant.Configure,
		group_instance_profile.Configure,
		instance_profile.Configure,
		lakehouse_monitor.Configure,
		metastore.Configure,
		metastore_assignment.Configure,
		metastore_data_access.Configure,
		mount.Configure,
		mws_credentials.Configure,
		mws_customer_managed_keys.Configure,
		mws_log_delivery.Configure,
		mws_ncc_binding.Configure,
		mws_ncc_private_endpoint_rule.Configure,
		mws_network_connectivity_config.Configure,
		mws_networks.Configure,
		mws_permission_assignment.Configure,
		mws_private_access_settings.Configure,
		mws_storage_configurations.Configure,
		mws_vpc_endpoint.Configure,
		mws_workspaces.Configure,
		notification_destination.Configure,
		obo_token.Configure,
		online_table.Configure,
		quality_monitor.Configure,
		recipient.Configure,
		registered_model.Configure,
		repo.Configure,
		restrict_workspace_admins_setting.Configure,
		secret_acl.Configure,
		service_principal_secret.Configure,
		share.Configure,
		storage_credential.Configure,
		system_schema.Configure,
		user.Configure,
		user_instance_profile.Configure,
		user_role.Configure,
		vector_search_endpoint.Configure,
		vector_search_index.Configure,
		volume.Configure,
		workspace_binding.Configure,
		workspace_conf.Configure,
		workspace_file.Configure,
	} {
		configure(pc)
	}

	// Rename resources to make it more pleasing to the eye
	for _, r := range pc.Resources {

		parts := strings.Split(r.Name, "_")
		if len(parts) > 1 {
			r.ShortGroup = resourcePrefix
			r.Kind = name.NewFromSnake(strings.Join(parts[1:], "_")).Camel
		}

	}

	pc.ConfigureResources()
	return pc, nil
}

// GetProviderNamespaced returns provider configuration for namespace-scoped resources
func GetProviderNamespaced(_ context.Context, sdkProvider *tfschema.Provider, generationProvider bool) (*ujconfig.Provider, error) {
	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		// ujconfig.WithShortName("databricks"),
		ujconfig.WithRootGroup("databricks.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// app.Configure,
		cluster.Configure,
		cluster_policy.Configure,
		credential.Configure,
		token.Configure,
		secret.Configure,
		secret_scope.Configure,
		notebook.Configure,
		job.Configure,
		instance_pool.Configure,
		permissions.Configure,
		sql_endpoint.Configure,
		entitlements.Configure,
		group.Configure,
		group_member.Configure,
		group_role.Configure,
		ip_access_list.Configure,
		permission_assignment.Configure,
		service_principal.Configure,
		service_principal_role.Configure,
		sql_permissions.Configure,
		grants.Configure,
		pipeline.Configure,
		alert.Configure,
		query.Configure,
		sql_alert.Configure,
		sql_dashboard.Configure,
		sql_global_config.Configure,
		sql_query.Configure,
		sql_table.Configure,
		budget.Configure,
		git_credential.Configure,
		catalog.Configure,
		connection.Configure,
		external_location.Configure,
		schema.Configure,
		library.Configure,
		sql_visualization.Configure,
		sql_widget.Configure,
		provider.Configure,
		mlflow_experiment.Configure,
		mlflow_model.Configure,
		mlflow_webhook.Configure,
		model_serving.Configure,
		access_control_rule_set.Configure,
		artifact_allowlist.Configure,
		catalog_workspace_binding.Configure,
		compliance_security_profile_workspace_setting.Configure,
		custom_app_integration.Configure,
		dashboard.Configure,
		dbfs_file.Configure,
		default_namespace_setting.Configure,
		directory.Configure,
		enhanced_security_monitoring_workspace_setting.Configure,
		file.Configure,
		global_init_script.Configure,
		grant.Configure,
		group_instance_profile.Configure,
		instance_profile.Configure,
		lakehouse_monitor.Configure,
		metastore.Configure,
		metastore_assignment.Configure,
		metastore_data_access.Configure,
		mount.Configure,
		mws_credentials.Configure,
		mws_customer_managed_keys.Configure,
		mws_log_delivery.Configure,
		mws_ncc_binding.Configure,
		mws_ncc_private_endpoint_rule.Configure,
		mws_network_connectivity_config.Configure,
		mws_networks.Configure,
		mws_permission_assignment.Configure,
		mws_private_access_settings.Configure,
		mws_storage_configurations.Configure,
		mws_vpc_endpoint.Configure,
		mws_workspaces.Configure,
		notification_destination.Configure,
		obo_token.Configure,
		online_table.Configure,
		quality_monitor.Configure,
		recipient.Configure,
		registered_model.Configure,
		repo.Configure,
		restrict_workspace_admins_setting.Configure,
		secret_acl.Configure,
		service_principal_secret.Configure,
		share.Configure,
		storage_credential.Configure,
		system_schema.Configure,
		user.Configure,
		user_instance_profile.Configure,
		user_role.Configure,
		vector_search_endpoint.Configure,
		vector_search_index.Configure,
		volume.Configure,
		workspace_binding.Configure,
		workspace_conf.Configure,
		workspace_file.Configure,
	} {
		configure(pc)
	}

	// Rename resources to make it more pleasing to the eye
	for _, r := range pc.Resources {

		parts := strings.Split(r.Name, "_")
		if len(parts) > 1 {
			r.ShortGroup = resourcePrefix
			r.Kind = name.NewFromSnake(strings.Join(parts[1:], "_")).Camel
		}

	}

	pc.ConfigureResources()
	return pc, nil
}
