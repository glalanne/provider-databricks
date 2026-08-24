#!/usr/bin/env python3

import os
import sys
import glob
from pathlib import Path
from typing import List

backlist_resources: List[str] = [
    "online_store",
    "feature_engineering_kafka_config",
    "warehouses_default_warehouse_override",
    "endpoint",
    "account_setting_user_preference_v2",
    "model_serving_provisioned_throughput",  # Private Preview
    "quality_monitor_v2",  # Deprecated, use quality_monitor instead
    "feature_engineering_materialized_feature",  # Private Preview
    "apps_settings_custom_template",  # Private Preview
    "data_quality_monitor",  # Deprecated, use quality_monitor instead
    "database_database_catalog",  # Private Preview
    "enhanced_security_monitoring_setting",
    "materialized_features_feature_tag",  # Private Preview
    "app_space",  # Private Preview
    "feature_engineering_feature",  # Private Preview
    "postgres_data_api",  # Private Preview,
    "mws_private_access_settings",
    "database_synced_database_table", # Private Preview,
]

resource_groups = {
    "data_classification_catalog_config": "governance",
    "workspace_entity_tag_assignment": "tags",
    "aibi_dashboard_embedding_access_policy_setting": "settings",
    "environments_default_workspace_base_environment": "envs",
    "postgres_role": "postgres",
    "rfa_access_request_destinations": "unity",
    "postgres_synced_table": "postgres",
    "account_network_policy": "settings",
    "postgres_endpoint": "postgres",
    "compliance_security_profile_setting": "settings",
    "disable_legacy_access_setting": "settings",
    "aibi_dashboard_embedding_approved_domains_setting": "settings",
    "disable_legacy_features_setting": "settings",
    "supervisor_agent_tool": "ai",
    "environments_workspace_base_environment": "envs",
    "alert_v2": "sql",
    "disaster_recovery_failover_group": "dr",
    "postgres_project": "postgres",
    "supervisor_agent": "ai",
    "knowledge_assistant": "ai",
    "account_setting_v2": "settings",
    "postgres_cdf_config": "postgres",
    "postgres_branch": "postgres",
    "budget_policy": "billing",
    "postgres_catalog": "postgres",
    "postgres_database": "postgres",
    "workspace_network_option": "settings",
    "knowledge_assistant_knowledge_source": "ai",
    "tag_policy": "tags",
    "entity_tag_assignment": "unity",
    "secret_uc": "unity",
    "workspace_setting_v2": "settings",
    "automatic_cluster_update_setting": "settings",
    "ai_search_endpoint": "ai",
    "data_quality_refresh": "unity",
    "ai_search_index": "ai",
    "disaster_recovery_stable_url": "dr",
    "external_metadata": "unity",
    "disable_legacy_dbfs_setting": "settings",
}


def prompt_yes_no(message: str) -> bool:
    while True:
        answer = input(message).strip().lower()
        if answer in {"y", "yes"}:
            return True
        if answer in {"n", "no"}:
            return False
        print("Please answer with 'y' or 'n'.")


def prompt_non_empty(message: str) -> str:
    while True:
        value = input(message).strip()
        if value:
            return value
        print("Value cannot be empty.")


def build_config_go(base_name: str, group_name: str) -> str:
    return f"""package {base_name}

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {{
	p.AddResourceConfigurator("databricks_{base_name}", func(r *config.Resource) {{
		r.ShortGroup = "{group_name}"
	}})
}}
"""


def generate_provider_go_imports(scope: str, base_name: str) -> None:
    provider_file = Path("config") / scope / "provider.go"
    with open(provider_file, "r+", encoding="utf-8") as f:
        content = f.read()
        import_statement = (
            f'\t"github.com/glalanne/provider-databricks/config/{scope}/{base_name}"\n'
        )
        add_config_line = f"\tProviderConfiguration.AddConfig({base_name}.Configure)\n"

        if import_statement not in content:
            # Insert the import statement after the last import
            last_import_index = content.rfind("import (")
            if last_import_index != -1:
                end_of_imports = content.find(")", last_import_index)
                if end_of_imports != -1:
                    content = (
                        content[:end_of_imports]
                        + import_statement
                        + content[end_of_imports:]
                    )

        if add_config_line not in content:
            # Insert the AddConfig line before the closing brace of the Configure function
            configure_function_index = content.find("func init() {")
            if configure_function_index != -1:
                closing_brace_index = content.find("}", configure_function_index)
                if closing_brace_index != -1:
                    content = (
                        content[:closing_brace_index]
                        + add_config_line
                        + content[closing_brace_index:]
                    )

        f.seek(0)
        f.write(content)
        f.truncate()


def generate_resource_config_go(scope: str, base_name: str, group_name: str) -> bool:
    # check if folder exists in the config directory
    # if os.path.exists(f"config/{scope}/{base_name}/config.go"):
    #     print(f"{scope.capitalize()} config found: {base_name} - ✅")
    # else:
    #     print(f"{scope.capitalize()} config not found: {base_name} - ❌")

    cluster_path = Path("config") / scope / base_name
    cluster_path.mkdir(parents=True, exist_ok=True)
    config_file = cluster_path / "config.go"

    if config_file.exists() and not prompt_yes_no(
        f"{config_file} already exists. Overwrite? [y/n]: "
    ):
        print(f"Skipped existing file: {config_file}")
        return False
    else:
        config_file.write_text(
            build_config_go(base_name, group_name),
            encoding="utf-8",
        )
    return True


# usage: version_diff.py <generated resource list> <base JSON schema path> <bumped JSON schema path>
# example usage: version_diff.py config/generated.lst .work/schema.json.3.38.0 config/schema.json
if __name__ == "__main__":
    directory = ".work/databricks/databricks/docs/resources"
    files = glob.glob(f"{directory}/*.md")

    resources_to_process = []

    for file_path in files:
        base_name = Path(file_path).stem

        if base_name in backlist_resources:
            # print(f"Resource '{base_name}' is blacklisted. Skipping.")
            continue

        if not os.path.exists(
            f"config/cluster/{base_name}/config.go"
        ) or not os.path.exists(f"config/namespaced/{base_name}/config.go"):
            resources_to_process.append(base_name)
            print(f"Resource '{base_name}' needs config generation. Adding to processing list.")

    for base_name in resources_to_process:

        if base_name not in resource_groups:
            group_name = prompt_non_empty(
                f"Enter short group for '{base_name}' (e.g. clusters, jobs, secrets): "
            )
        else:
            group_name = resource_groups[base_name]

        print(f"Generating config for resource '{base_name}' with group '{group_name}'.")
        
        generate_resource_config_go("cluster", base_name, group_name)
        generate_resource_config_go("namespaced", base_name, group_name)

        generate_provider_go_imports("cluster", base_name)
        generate_provider_go_imports("namespaced", base_name)

    print("------------------------------")

    external_names = ""
    with open("config/external_name.go", "r") as file:
        external_names = file.read()

    new_external_names = []

    for base_name in resources_to_process:
        if (
            os.path.exists(f"config/namespaced/{base_name}/config.go")
            and external_names.find(f"databricks_{base_name}") == -1
        ):
            print(f"External name not found: databricks_{base_name} - ❌")
            new_external_names.append(
                f'\n\t"databricks_{base_name}":                                        config.IdentifierFromProvider,'
            )

    if new_external_names:
        print("Updating config/external_name.go with new external names...")
        # Locate the line to insert new external names
        insertion_point = external_names.find(
            "var TerraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{"
        )
        if insertion_point != -1:
            closing_brace_index = external_names.find("}", insertion_point)
            if closing_brace_index != -1:
                external_names = (
                    external_names[:closing_brace_index]
                    + "".join(new_external_names)
                    + external_names[closing_brace_index:]
                )

        with open("config/external_name.go", "w", encoding="utf-8") as file:
            file.seek(0)
            file.write(external_names)
            file.truncate()
