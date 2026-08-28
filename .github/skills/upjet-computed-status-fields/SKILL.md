---
name: upjet-computed-status-fields
description: "Use when: an Upjet-generated Crossplane resource exposes an API-assigned Terraform Optional + Computed field in spec.forProvider that should be status.atProvider-only; for example databricks_pipeline state or last_modified. Covers diagnosing the schema flags, finding similar resources, applying a resource-scoped override, regenerating APIs and CRDs, and validating the result."
---

# Upjet Computed Status Fields

Use this workflow for a generated resource field that is present in `spec.forProvider` but is assigned by Databricks and should be reported only in `status.atProvider`.

Examples include `databricks_pipeline.state` and `databricks_pipeline.last_modified`.

## Why A Computed Field Can Be In Spec

In Terraform Plugin SDK schemas, `Computed` does not by itself determine whether a generated Crossplane field appears in `spec`.

| Terraform schema flags | Upjet-generated location |
| --- | --- |
| `Required` | `spec.forProvider` |
| `Optional` | `spec.forProvider` |
| `Computed` only | `status.atProvider` |
| `Optional` and `Computed` | Both `spec.forProvider` and `status.atProvider` |

Terraform uses `Optional + Computed` for a value that the user may set but the provider may default, populate, or change. Upjet exposes it in both parameter and observation types accordingly.

A field belongs in status only when it is not a supported desired input for the Crossplane API. Do not move a field merely because Databricks returns it during observation.

## Triage A Suspected Field

1. Inspect the resource schema in `config/schema.json`. Replace the resource type and field name as appropriate:

   ```sh
   jq '
     .provider_schemas["registry.terraform.io/databricks/databricks"]
     .resource_schemas["databricks_pipeline"]
     .block.attributes["state"]
   ' config/schema.json
   ```

2. Confirm the distinction between the relevant schema flags:

   - `"optional": true, "computed": true`: expected in both spec and status.
   - `"computed": true` without `optional` or `required`: expected only in status.
   - `"required": true`: it must remain in spec.

3. Inspect the generated type and CRD before changing configuration. Check the top-level resource field, rather than unrelated nested fields with the same name:

   ```sh
   rg -n -C 3 'type Pipeline(InitParameters|Parameters|Observation)|State \*string' \
     apis/cluster/compute/v1beta1/zz_pipeline_types.go

   rg -n -C 3 '^                  state:' \
     package/crds/compute.databricks.crossplane.io_pipelines.yaml
   ```

4. Check whether the field is genuinely API-assigned and should not be sent to Terraform on create or update. Use the upstream Terraform resource docs and resource implementation where the provider schema alone is ambiguous.

## Apply A Resource-Scoped Override

Add the override in the resource configurator, for example `config/cluster/pipeline/config.go`:

```go
p.AddResourceConfigurator("databricks_pipeline", func(r *config.Resource) {
    // Existing resource configuration.

    if field, ok := r.TerraformResource.Schema["state"]; ok {
        field.Optional = false
        field.Computed = true
    }
})
```

This changes the generator's in-memory Terraform schema before it emits the Crossplane API. The original schema entry may already have `Computed: true`; explicitly retaining it makes the intended computed-only state unambiguous.

If the repository exposes the resource in both scopes, apply the identical resource-specific override in both locations:

```text
config/cluster/<resource>/config.go
config/namespaced/<resource>/config.go
```

Do not edit `apis/**/zz_*`, `package/crds/**`, examples-generated output, or controller files. They are generated artifacts and the next generation would overwrite manual changes.

## Find Other Candidates

Start from observed symptoms rather than bulk-changing every `Optional + Computed` field. Good candidates have all of these properties:

- The field appears in `spec.forProvider` and `status.atProvider`.
- It is `Optional + Computed` in `config/schema.json`.
- The remote API/provider assigns the value and does not accept it as a meaningful desired input.
- Supplying it can cause a perpetual diff, invalid update, or an unnecessarily misleading spec surface.

List `Optional + Computed` top-level attributes for a resource:

```sh
jq -r '
  .provider_schemas["registry.terraform.io/databricks/databricks"]
  .resource_schemas["databricks_pipeline"]
  .block.attributes
  | to_entries[]
  | select(.value.optional == true and .value.computed == true)
  | .key
' config/schema.json
```

For each candidate, inspect its upstream documentation and implementation before overriding it. Fields that are user-controllable but provider-defaulted, such as a create-time option with a service-selected default, should remain `Optional + Computed` and therefore in spec.

Nested fields require extra care. The override above changes only the top-level `r.TerraformResource.Schema` entry. Do not infer that a nested `state` field is the same as a resource's top-level `state` attribute.

## Regenerate And Validate

1. Run the repository's normal generation target after the configuration change: 
   ```sh
   make generate
   ```
   Consult `make help` if the target is not known in the current checkout.

2. Validate the generated types. The affected top-level field must be absent from both `*Parameters` and `*InitParameters`, while remaining in `*Observation`:

   ```sh
   rg -n -C 3 'type Pipeline(InitParameters|Parameters|Observation)|State \*string' \
     apis/cluster/compute/v1beta1/zz_pipeline_types.go
   ```

3. Validate the CRD. The field must be absent from top-level `spec.forProvider` and `spec.initProvider`, and present under `status.atProvider`:

   ```sh
   yq '
     .spec.versions[]
     | {
         version: .name,
         forProvider: .schema.openAPIV3Schema.properties.spec.properties.forProvider.properties.state,
         initProvider: .schema.openAPIV3Schema.properties.spec.properties.initProvider.properties.state,
         atProvider: .schema.openAPIV3Schema.properties.status.properties.atProvider.properties.state
       }
   ' package/crds/compute.databricks.crossplane.io_pipelines.yaml
   ```

   Expected result for each served version: `forProvider: null`, `initProvider: null`, and a schema object at `atProvider`.

4. Build or test the narrow affected package when one exists, then run the normal provider generation/build validation required by the repository.

5. Review the generated diff. Expected changes are limited to the resource configurator and generated APIs, CRDs, examples, and controllers that the repository normally updates during generation.

## Pipeline Example

For `databricks_pipeline.state`, set its top-level Terraform schema entry to computed-only in both pipeline configurators:

```go
if field, ok := r.TerraformResource.Schema["state"]; ok {
    field.Optional = false
    field.Computed = true
}
```

After regeneration, `PipelineObservation` keeps `State`, while `PipelineParameters` and `PipelineInitParameters` do not. The top-level CRD field moves from `spec.forProvider.state` to `status.atProvider.state`. Nested fields such as `latestUpdates.state` are independent schema paths and are not changed by this override.
