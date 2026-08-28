---
name: upjet-sdk-state-reconstruction
description: "Use when: diagnosing Crossplane provider restart reconciliation loops, repeated 'Instance state not found in cache, reconstructing...' logs, Upjet Terraform Plugin SDKv2 state reconstruction failures, or Databricks Terraform TypeSet hash panics caused by optional nested blocks. Covers finding and fixing other resources with the same symptom."
---

# Upjet SDK State Reconstruction

Use this workflow when an Upjet-based provider logs `Instance state not found in cache, reconstructing...` repeatedly after its pod restarts.

A single message per managed resource after restart is normal: Upjet's `OperationTrackerStore` is in-memory and is empty in a new pod. A message on every reconciliation, especially with exponential requeues, indicates that reconstruction cannot produce a usable Terraform `InstanceState` or that observation fails immediately afterward.

## Scope

This repository uses Upjet v2's Terraform Plugin SDKv2 async connector for generated resources. The connector reconstructs Terraform state from `status.atProvider` and resource annotations in `Connect` before calling `Observe`.

Relevant dependency paths:

- `pkg/controller/external_tfpluginsdk.go`: reconstructs state in `Connect`.
- `pkg/controller/nofork_store.go`: caches state by Kubernetes resource UID.
- `pkg/controller/external_async_tfpluginsdk.go`: wraps the synchronous SDK connector for async operations.

Generated controllers use this connector through `NewTerraformPluginSDKAsyncConnector`.

## Triage

1. Confirm the message repeats for the same UID after a single provider restart:

   ```sh
   kubectl -n <provider-namespace> logs <provider-pod> --since=30m | \
     grep -E 'Instance state not found|Observing the external resource|Cannot connect|Cannot observe'
   ```

   A healthy resource has one reconstruction log followed by `Observing the external resource` and then an up-to-date or normal reconcile result.

2. Check that the managed resource has a stable UID, a non-empty external name, and an observed ID:

   ```sh
   kubectl get <resource> <name> -o json | jq '{
     uid: .metadata.uid,
     externalName: .metadata.annotations["crossplane.io/external-name"],
     observedID: .status.atProvider.id,
     conditions: .status.conditions
   }'
   ```

   If UID or external name changes, fix that lifecycle/identity problem first.

3. Interpret the logs:

   - `Observing the external resource` follows reconstruction: state reconstruction succeeded. Investigate `RefreshWithoutUpgrade`, provider credentials, API errors, or diff/status processing.
   - No observe log and retries follow an exponential pattern: the failure is inside `Connect` after the reconstruction message. This can be a panic in Terraform schema state conversion; controller-runtime may recover it without printing a useful managed-resource condition.

4. Reproduce from the live object before changing the provider. Create a short, temporary Go program outside the repository that:

   - unmarshals `kubectl get <resource> <name> -o json` into the generated managed-resource Go type;
   - obtains the configured `config.Resource` from `config.GetProvider` or `config.GetProviderNamespaced`;
   - calls `GetObservation()`;
   - calls `ApplyTFConversions(state, config.ToTerraform)`;
   - assigns `state["id"] = managed.GetID()`;
   - calls `schema.JSONMapToStateValue(state, resource.TerraformResource.CoreConfigSchema())`;
   - calls `resource.TerraformResource.ShimInstanceStateFromValue(stateValue)`.

   A panic or error here is a local, deterministic reproduction of the restart failure. Remove the temporary program after diagnosis.

## Known Compute Cluster Failure

`databricks_cluster` uses a custom Terraform `TypeSet` hash for `library`:

```go
s.SchemaPath("library").Schema.Set = func(value any) int {
    lib := libraries.NewLibraryFromInstanceState(value)
    return schema.HashString(lib.String())
}
```

The generated Crossplane observation can retain empty optional nested blocks, for example:

```yaml
library:
  - cran: {}
    maven: {}
    pypi:
      package: cleo==2.0.0
      repo: ""
```

During Upjet reconstruction, the Terraform SDK can flatten empty singleton blocks to `[]any{nil}`. The Databricks hash implementation assumes `pypi[0]`, `maven[0]`, and `cran[0]` are maps and panics with:

```text
panic: interface conversion: interface {} is nil, not map[string]interface {}
.../libraries.NewLibraryFromInstanceState
.../clusters.ClusterSpec.CustomizeSchema.func1
```

The result is repeated reconstruction without `Observing the external resource`.

## Fix Pattern

Apply a resource-specific schema wrapper in both `config/cluster/<resource>/config.go` and `config/namespaced/<resource>/config.go` when both resource scopes exist.

For a custom `TypeSet` hash that does not tolerate `[]any{nil}`, wrap the hash and normalize only semantically empty singleton blocks:

```go
if librarySchema, ok := r.TerraformResource.Schema["library"]; ok && librarySchema.Set != nil {
    libraryHash := librarySchema.Set
    librarySchema.Set = func(value any) int {
        return libraryHash(normalizeLibraryForHash(value))
    }
}
```

Rules for normalization:

- Preserve every non-empty value exactly; the provider's original hash remains authoritative.
- Copy the top-level map before modifying it. Do not mutate the object supplied by the Terraform SDK.
- Convert only `[]any{nil}` to `[]any{}` for optional nested blocks known to be semantically absent.
- Do not globally change Terraform SDK or Upjet behavior.
- Keep the fix resource-specific unless multiple affected schemas share an identical safe normalization requirement.

For provider-assigned IDs where the Terraform resource schema does not include `id`, add only a computed `id` schema field in the resource configurator:

```go
if _, ok := r.TerraformResource.Schema["id"]; !ok {
    r.TerraformResource.Schema["id"] = &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
    }
}
```

This allows Upjet's existing reconstruction path to carry `state["id"]` into the top-level `terraform.InstanceState.ID`, which `OperationTrackerStore.HasState()` requires.

## Finding Similar Resources

Search the Terraform provider dependency for custom set hashes and unsafe assertions:

```sh
grep -RniE '\.Set = func|\.SetHash|\.\(map\[string\]any\)|\.\(map\[string\]interface\{\}\)' \
  "$(go env GOMODCACHE)/github.com/glalanne/terraform-provider-databricks@*"
```

Prioritize resources that have all of these characteristics:

- They use the Terraform Plugin SDKv2 connector.
- They define a `TypeSet` with a custom hash function.
- The hash type-asserts nested list elements to maps without checking for nil.
- Their generated `status.atProvider` exposes optional nested blocks as empty objects.
- They reproduce the repeated post-restart reconstruction symptom.

Do not apply this fix merely because a resource has a custom set hash. Reproduce using its live `status.atProvider` first.

## Tests And Validation

Add a focused test alongside the resource configurator. It should use the actual Terraform resource schema from the Databricks provider and include the problematic status shape. Assert that:

1. `schema.JSONMapToStateValue` succeeds.
2. `ShimInstanceStateFromValue` does not panic.
3. The returned `InstanceState.ID` equals the resource external name.

Run focused validation before broad suites:

```sh
go test ./config/cluster/<resource> ./config/namespaced/<resource>
```

Then rebuild, redeploy, and restart the provider. Verify:

```sh
kubectl -n <provider-namespace> logs <provider-pod> --since=10m | \
  grep -E 'Instance state not found|Observing the external resource'
```

Success criteria:

- One reconstruction message for the affected object after a restart.
- `Observing the external resource` immediately follows.
- No exponential reconciliation retries.
- Subsequent poll reconciliations do not log reconstruction for that UID.

If reconstruction persists after Observe succeeds, inspect the refreshed Terraform state: `OperationTrackerStore.HasState()` requires a non-nil state with non-empty top-level `InstanceState.ID`.
