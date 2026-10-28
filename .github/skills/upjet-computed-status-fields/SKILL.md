---
name: upjet-computed-status-fields
description: "Use when: deciding whether an Upjet Optional + Computed field should remain in spec with late initialization disabled or become status-only; diagnosing stale observed values in spec, commit_hash diffs, or Databricks omitempty schema inconsistencies. Covers runtime-schema safety, field-specific late-init exclusions, status-only eligibility, regeneration, migration, and behavioral validation."
---

# Upjet Computed Status Fields

Use this workflow to distinguish unwanted late initialization from an intentional
status-only API change. Do not assume that a server-assigned value must be removed
from spec, or that a Terraform diff mentioning it proves a schema defect.

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

## Decide Which Problem To Solve First

| Requirement or evidence | Preferred action |
| --- | --- |
| A valid user input is defaulted by the server | Keep Optional + Computed and normal late initialization unless there is a specific reason to exclude it. |
| An observed value is copied into spec and later becomes stale; status-only is not required | Keep the upstream schema and exclude that Terraform field path from late initialization. |
| The field must not be accepted as desired input in the Crossplane API | Consider computed-only, but only after all status-only safety checks below pass. |
| Observation fails with `has omitempty, but is not optional` after a schema override | Remove the incompatible runtime override. Late-init exclusion alone cannot repair a failing response conversion. |
| The only evidence is a `Diff detected` log | Inspect desired parameters, observed state, and provider diff logic before choosing either change. Preserve unrelated legitimate diffs. |

Prefer the smallest behavior change that meets the actual requirement. Do not
introduce a generation/runtime schema split just to avoid late initialization.
Neither method is a general instruction to ignore an explicitly configured value.

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

5. Inspect the affected object's `spec.forProvider`, `spec.initProvider`, and source
   manifests or compositions. Determine whether the value was supplied explicitly
   or copied by late initialization. A stale concrete `New` value with
   `NewComputed: false` can indicate a configured value, not merely a new observed
   output. Reproduce the diff with the field omitted versus configured.

6. Trace the pinned runtime provider implementation, not only the exported JSON
   schema. In this repository, `config/provider.go` invokes resource
   configurators for both generation and runtime. For SDK resources, Upjet uses
   the configured Terraform resource for refresh, diff, apply, and state
   reconstruction. Changing its schema may also change a map captured by
   upstream callbacks.

## Exclude A Field From Late Initialization

Use this when the problem is observed state becoming unwanted desired state and
removing the field from the API is not a requirement.

For Repo, preserve the upstream Optional + Computed schema and configure:

```go
p.AddResourceConfigurator("databricks_repo", func(r *config.Resource) {
    r.ShortGroup = "workspace"
    r.LateInitializer.IgnoredFields = append(
        r.LateInitializer.IgnoredFields,
        "commit_hash",
    )
})
```

Use Terraform field paths, not JSON names such as `commitHash`. The generator
maps these to canonical names in the generated late initializer. Apply the same
exclusion in both `config/cluster/repo/config.go` and
`config/namespaced/repo/config.go`.

This preserves the optional parameter fields and observed status. It prevents
future copying of this field into spec without disabling late initialization of
other fields. It does not erase existing values, prevent explicit configuration,
or suppress Terraform diffs for explicitly supplied values.

Remove stale values from both spec sections and their source manifests or
compositions. Otherwise an existing value or a GitOps/composition writer can
continue supplying the stale desired value.

### Repo Failure And Why This Is The Preferred Fix

The Databricks Repo response struct declares:

```go
HeadCommitID string `json:"head_commit_id,omitempty" tf:"computed,alias:commit_hash"`
```

The provider derives Optional from `omitempty` and Computed from the Terraform
tag. During response conversion, `StructToData`/`iterFields` also requires that
this `omitempty` field remain Optional. Mutating its live schema to
`Optional=false` causes:

```text
inconsistency: commit_hash has omitempty, but is not optional
```

This happens during observation before late initialization. Therefore:

- Late-init exclusion cannot fix that error while the runtime override remains.
- Removing the override and excluding `commit_hash` from late initialization
  addresses the runtime error and prevents future stale hash initialization.
- With the original SDK schema, an omitted hash does not cause a hash diff in
  isolated SDK planning; a configured stale hash does. Legitimate `tag` changes
  must still be planned.

## When A Status-Only Override Is Allowed

All of the following must be established before changing the flags:

1. **API requirement:** status-only is intentional and necessary, rather than an
   assumed remedy for unwanted late initialization.
2. **Input semantics:** the field is not a meaningful supported desired input.
   An optional create-time choice or server-defaulted input must not be removed
   merely because it is returned by the service.
3. **Runtime compatibility:** the pinned provider's create, read, update, diff,
   and conversion code tolerates computed-only. Check captured schema maps,
   struct-tag validation, custom diff logic, and state reconstruction. Output-only
   documentation alone is not sufficient evidence of runtime safety.
4. **API compatibility:** assess existing manifests and stored resources that
   contain the parameter. Removing a spec field changes the CRD contract and may
   cause pruning or rejection. Document the migration.
5. **Behavioral evidence:** verify response conversion and reconciliation, not
   only generated field placement or successful compilation.

Do not blanket-convert Optional + Computed fields. Do not copy another provider's
override without checking this provider's runtime invariants.

### Apply A Safe Resource-Scoped Override

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

This changes the configured Terraform schema, **not just generated types**.
Shared configurators also execute at runtime. The snippet is appropriate only
after the eligibility and compatibility checks above; it is not a safe recipe
for Repo's `commit_hash`. Explicitly retaining Computed makes the intended
computed-only state unambiguous.

If the repository exposes the resource in both scopes, apply the identical resource-specific override in both locations:

```text
config/cluster/<resource>/config.go
config/namespaced/<resource>/config.go
```

Do not edit `apis/**/zz_*`, `package/crds/**`, examples-generated output, or controller files. They are generated artifacts and the next generation would overwrite manual changes.

### Generation-Only Adaptations Are Exceptional

If status-only is required but upstream runtime code cannot tolerate the schema
change and upstream cannot be changed, a computed-only override confined to an
independent generation schema is a possible custom adaptation, not the default
fix or a verified general Upjet recipe.

Prove schema ownership/isolation and preserve the original runtime schema without
temporarily mutating and resetting shared state. Test both scopes/versions,
status round-tripping, parameters, create/read/update, and restart reconstruction.
Public providers using different generation/runtime schema sources are related
architectural precedents, not proof that arbitrary flag divergence is safe.

## Find Other Candidates

Start from observed symptoms rather than bulk-changing every `Optional + Computed` field. Good candidates have all of these properties:

- The field appears in `spec.forProvider` and `status.atProvider`.
- It is `Optional + Computed` in `config/schema.json`.
- The remote API/provider assigns the value and does not accept it as a meaningful desired input.
- There is an explicit need to remove it from the API, and runtime compatibility can be verified.

If only stale late-initialized state is causing a perpetual diff, prefer the
late-initialization exclusion workflow instead.

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

2. Validate the chosen contract in both scopes and every served API version:

   - **Late-init exclusion:** the field remains in Parameters, InitParameters,
     and Observation as dictated by the original schema. The generated
     `LateInitialize` includes a name filter for the field. Other fields still
     late-initialize.
   - **Status-only:** the affected top-level field is absent from both
     Parameters and InitParameters, while remaining in Observation.

   For example, inspect the generated types:

   ```sh
   rg -n -C 3 'type Pipeline(InitParameters|Parameters|Observation)|State \*string' \
     apis/cluster/compute/v1beta1/zz_pipeline_types.go
   ```

3. Validate the CRD. For status-only, the field must be absent from top-level
   `spec.forProvider` and `spec.initProvider`, and present under
   `status.atProvider`. For late-init exclusion of an Optional + Computed field,
   expect it to remain present in all three locations:

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

   Status-only expected result for each served version: `forProvider: null`,
   `initProvider: null`, and a schema object at `atProvider`. For late-init
   exclusion, expect a schema object at all three paths.

4. Test behavior using the pinned dependencies:

   - Resource configuration preserves the intended runtime schema flags.
   - Actual upstream response conversion succeeds with the configured schema.
   - An omitted hash has no hash diff; a legitimate branch/tag change still has
     a diff. Do not delete unrelated diff attributes to make the test pass.
   - The generated late initializer does not copy the excluded value, while
     unrelated fields still initialize and the hash remains observable.
   - Existing supplied values are not silently removed by late-init exclusion.
   - For schema adaptations, validate state reconstruction/restart and real
     create/read/update behavior. Clearly distinguish offline component tests
     from live reconciliation evidence.

   Run the narrow affected tests and normal repository validation. For Repo,
   `go test ./config -run '^TestRepo' -count=1` covers the regression tests.

5. Review the generated diff. Expected changes are limited to the resource configurator and generated APIs, CRDs, examples, and controllers that the repository normally updates during generation.

## Conditional Pipeline Example

The pipeline configurators contain a computed-only override for the top-level
`state` field. Its generated placement illustrates a status-only change; the
existence of that override is not proof of runtime safety for other fields:

```go
if field, ok := r.TerraformResource.Schema["state"]; ok {
    field.Optional = false
    field.Computed = true
}
```

After regeneration, `PipelineObservation` keeps `State`, while `PipelineParameters` and `PipelineInitParameters` do not. The top-level CRD field moves from `spec.forProvider.state` to `status.atProvider.state`. Nested fields such as `latestUpdates.state` are independent schema paths and are not changed by this override.

## Primary References

- [Upjet resource configuration](https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md):
  late-initialization exclusions and schema-to-API mapping.
- [Upjet type builder](https://github.com/crossplane/upjet/blob/main/pkg/types/builder.go):
  `IsObservation` distinguishes computed-only from Optional + Computed.
- [Upjet SDK reconciler](https://github.com/crossplane/upjet/blob/main/pkg/controller/external_tfpluginsdk.go):
  refresh, parameter diffing, late initialization, and state reconstruction.
- [AWS schema overrides](https://github.com/crossplane-contrib/provider-upjet-aws/blob/main/config/overrides.go):
  computed-only precedent, not a Databricks compatibility guarantee.

Use the repository's pinned module versions when verifying runtime behavior;
current upstream branches may differ.
