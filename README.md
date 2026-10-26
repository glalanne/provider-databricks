
# Provider Databricks

<div style="text-align: center;">

![CI](https://github.com/glalanne/provider-databricks/workflows/CI/badge.svg)
[![GitHub release](https://img.shields.io/github/release/glalanne/provider-databricks/all.svg)](https://github.com/glalanne/provider-databricks/releases)
[![Contributors](https://img.shields.io/github/contributors/glalanne/provider-databricks)](https://github.com/glalanne/provider-databricks/graphs/contributors)


</div>

`provider-databricks` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
[Databricks](https://registry.terraform.io/providers/databricks/databricks/latest/docs).


Most of the testing have been done on [Azure Databricks](https://azure.microsoft.com/en-ca/products/databricks)

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/lalanne/provider-databricks):
```
up ctp provider install xpkg.upbound.io/lalanne/provider-databricks:v2.5.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-databricks
spec:
  package: xpkg.upbound.io/lalanne/provider-databricks:v2.5.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/glalanne/provider-databricks).

## Exemples

A few exemples are provided [here](./examples/) to show case how to configure the provider, and how to use some resources


## Developing

Run unit and API round-trip tests:
```console
make test
```

Library resources retain the Terraform Plugin SDKv2 implementation used by the
generated controllers. Provider configuration restores that implementation when
the upstream provider omits it in favor of Plugin Framework, for both runtime
and code generation, preserving the list-shaped blocks in both API versions.
No `USE_SDK_V2_RESOURCES` environment override is required.

Run code-generation pipeline:
```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

Local deployment:

```console
make local-deploy
```

End to end tests:

```console
# Ensure that you have a .datasource.yaml use the sample for instructions
# Ensure that you have a UPTEST_CLOUD_CREDENTIALS env define with the creds you want to use.

make e2e
```

### Offline diff client

The internal diff server uses a separate SDKv2 and Framework client setup.
It never reads
credential Secrets, requests service account tokens, runs authentication CLIs,
or loads Databricks environment variables or profiles. Authentication uses a
local placeholder PAT, and host metadata discovery is disabled. Normal
reconciliation continues to use the regular authenticated client.

Start the server without a kubeconfig or Kubernetes cluster:

```console
provider internal diff-server --address 127.0.0.1:9099
```

The caller supplies ProviderConfigs and reference objects through the request's
`kubernetes_object_store`. Namespaced resources can use either a local
ProviderConfig or a ClusterProviderConfig. Framework configuration receives the
same isolated, egress-guarded client as SDKv2 configuration.

ProviderConfigs are still resolved using the supplied Kubernetes client. Set
`spec.host` and, for account-level resources, `spec.accountID` on the
ProviderConfig when these affect a diff; values stored only in credential
Secrets are intentionally unavailable. Without a host, the client uses a
reserved `.invalid` workspace hostname.

Outbound requests from the client, including derived workspace and account
clients, are blocked. Diffs that require an API call return Upjet's
diff-computation-not-supported error rather than fabricated API responses.
This setup is for offline diff computation only, not reconciliation.

Run the client unit tests:

```console
go test ./internal/clients
```

The provider integration test builds the binary, starts it without Kubernetes
configuration, and exercises real gRPC plans for both scopes and Terraform
execution modes:

```console
go test ./cmd/provider -run TestDiffServerWithoutKubernetes
```

The egress regression tests run in fresh subprocesses with an HTTP/HTTPS
recording proxy, without replacing `http.DefaultTransport`. Positive controls
verify that normal SDK discovery, authentication, API calls, API-dependent
`CustomizeDiff` hooks, and separately constructed HTTP transports reach the
proxy. A separate recording listener covers loopback URLs, which Go excludes
from proxying. Offline setup, authentication, diffs, and derived clients must
produce zero recorded requests. Neither listener forwards requests. This tests
the exercised HTTP paths, not an OS-level network sandbox.

### BudgetPolicy identifier handling

Both cluster-scoped and namespaced BudgetPolicy resources treat a Terraform
Read state with a missing, null, or empty `policy_id` as absent, even if the
state object itself is non-null. Unknown or malformed identifiers fail
reconciliation instead of being treated as permission to create a policy.
An existing non-empty identifier continues to be used as the external name.

Before reconciling a resource without an identifier, check whether its policy
already exists in Databricks. If it does, set `crossplane.io/external-name` to
the real policy ID to import it rather than creating another policy.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/glalanne/provider-databricks/issues).
