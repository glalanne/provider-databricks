# Authentication

This guide describes authentication implemented by this Crossplane Databricks
provider. It distinguishes **where credentials come from**
(`spec.credentials.source`) from **how Databricks authenticates them**
(`auth_type` in the credential JSON).

Authentication does not grant authorization: the resulting identity must also
have access to the target Databricks workspace or account and permissions for
the resources being managed.

## Choosing a method

| Method | Credential source | Use when |
| --- | --- | --- |
| Kubernetes service-account OIDC exchange | `OIDCTokenRequest` | Recommended passwordless option for Azure Databricks, especially when different namespaces need different identities. |
| Databricks OAuth machine-to-machine | `Secret` or `Filesystem` | A Databricks service principal has an OAuth client ID and secret. |
| Personal access token (PAT) | `Secret` or `Filesystem` | A workspace accepts PATs; useful for initial setup or environments without federation. You manage token rotation. |
| Azure service-principal client secret | `Secret` or `Filesystem` | An Entra application/service principal authenticates to Azure Databricks with a secret. |
| Azure managed identity via the Databricks SDK | `Secret` or `Filesystem` | The provider has access to a compatible Azure managed-identity endpoint. |
| Google credentials or service-account impersonation | `Secret` or `Filesystem` | Managing Databricks on Google Cloud with the SDK's Google authentication methods. |
| SDK environment/profile authentication | `Secret` or `Filesystem` containing at least `{}` | The provider pod supplies the SDK's authentication configuration. Identity is generally shared at pod level. |
| Mounted OIDC token | `OIDCTokenFile` | Legacy provider path; see the compatibility warning below before using it. |
| Upbound-injected token | `Upbound` | Legacy provider path for an Upbound-projected token; see the compatibility warning below. |
| Named managed-identity sources | `UserAssignedManagedIdentity`, `SystemAssignedManagedIdentity` | Legacy provider paths; prefer the SDK configuration or `OIDCTokenRequest` described below. |

`None` is accepted by the CRD but **does not currently work as a normal
reconciliation authentication source**: the credential extractor returns no
bytes, and this provider then tries to parse those bytes as JSON. It is not a
shortcut for SDK environment authentication.

There is no credential source literally named `OIDCToken`, `Environment`, or
`InjectedIdentity` in the current CRD's allowed values. Some inherited field
comments and error messages mention `InjectedIdentity`; use the actual source
values documented here.

## Account-level versus workspace-level authentication

Databricks account APIs and workspace APIs use different endpoints and
authorization scopes. Choose the endpoint required by the managed resource,
not simply the endpoint where the identity can log in.

| Setting | Workspace-level | Account-level |
| --- | --- | --- |
| Target | Resources within a particular workspace, such as notebooks. | Account resources, such as account-wide identities and workspace provisioning. |
| Host | The workspace URL, for example `https://adb-1234567890123456.7.azuredatabricks.net`. | The accounts endpoint for the cloud: `https://accounts.cloud.databricks.com` (AWS), `https://accounts.azuredatabricks.net` (Azure public cloud), or `https://accounts.gcp.databricks.com` (Google Cloud). |
| Account ID | Normally omitted for workspace API authentication. | Supply the Databricks account ID, not an Azure subscription ID or workspace ID. |
| Identity permissions | The identity needs workspace access and permissions for the managed resources. | The identity needs account-level roles/permissions for the managed resources; workspace access alone is insufficient. |
| Authentication | A method supported by that workspace, including PAT when enabled. | Use account-supported authentication, such as Databricks OAuth M2M or, on Azure, the Entra token exchange described below. |

> **Warning:** Databricks personal access tokens (PATs) do **not** work for
> account-level authentication. They authenticate workspace API requests only.
> Adding an account ID or changing the host to an accounts endpoint does not
> turn a workspace PAT into an account credential.

For Secret/Filesystem credentials, configure `host` and `account_id` **inside
the credential JSON**. For example, account-level Databricks OAuth M2M on Azure:

```json
{
  "host": "https://accounts.azuredatabricks.net",
  "account_id": "<databricks-account-id>",
  "auth_type": "oauth-m2m",
  "client_id": "<databricks-service-principal-client-id>",
  "client_secret": "<databricks-oauth-secret>"
}
```

For `OIDCTokenRequest`, configure top-level `spec.host` and `spec.accountID`
instead; see [the account configuration example](#explicit-identity-settings-and-sharedaccount-configuration).
Although that implementation internally selects `auth_type: pat` to send a
bearer token, its credential is an **Entra access token, not a Databricks PAT**.
The PAT restriction above does not preclude this Azure account-level flow.

Use separate ProviderConfigs for account and workspace operations, and point
each resource's `providerConfigRef` at the appropriate configuration. The same
service principal can be used where authorized, but its account permissions
and workspace permissions must be granted separately.

Databricks account/workspace scope is independent of Kubernetes scope:
`ClusterProviderConfig` means a cluster-scoped Kubernetes configuration, **not**
an account-level Databricks connection. Either ProviderConfig kind can target
an account or a workspace.

## ProviderConfig scope and resource references

| Managed resource scope | ProviderConfig API | Configuration kind |
| --- | --- | --- |
| Namespaced | `databricks.m.crossplane.io/v1beta1` | Namespaced `ProviderConfig` or cluster-scoped `ClusterProviderConfig` |
| Legacy cluster-scoped | `databricks.crossplane.io/v1beta1` | Cluster-scoped `ProviderConfig` |

A namespaced resource references a local config as follows:

```yaml
spec:
  providerConfigRef:
    kind: ProviderConfig
    name: default
```

Use `kind: ClusterProviderConfig` for a shared config. A legacy cluster-scoped
resource uses `providerConfigRef: {name: default}`, without `kind`.

- A namespaced ProviderConfig reads credential Secrets from its own namespace.
- A ClusterProviderConfig or legacy cluster-scoped ProviderConfig needs an
  explicit namespace in its Secret or service-account reference.
- Shared ClusterProviderConfigs are not namespace-specific identity boundaries.
  Use namespaced ProviderConfigs when identities should be isolated per tenant,
  and control who can edit ProviderConfigs, service accounts, RBAC, and trust
  relationships.

## Secret-backed credentials

The selected Secret key contains a **JSON object whose values are strings**,
not a raw token and not a `.databrickscfg` INI file. Use `stringData` when
writing YAML; Kubernetes stores it as Secret data.

Example using a PAT for workspace-level authentication only:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: databricks-credentials
  namespace: team-a
type: Opaque
stringData:
  credentials: |
    {
      "host": "https://adb-1234567890123456.7.azuredatabricks.net",
      "auth_type": "pat",
      "token": "<databricks-personal-access-token>"
    }
---
apiVersion: databricks.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
  namespace: team-a
spec:
  credentials:
    source: Secret
    secretRef:
      name: databricks-credentials
      key: credentials
```

For a ClusterProviderConfig, omit `metadata.namespace`, change `kind`, and add
`credentials.secretRef.namespace: team-a`.

Replace the JSON with one of the following alternatives. Do not combine
unrelated credential sets; conflicting SDK authentication fields can cause
configuration errors.

### Databricks OAuth machine-to-machine

```json
{
  "host": "https://your-workspace.cloud.databricks.com",
  "auth_type": "oauth-m2m",
  "client_id": "<databricks-service-principal-client-id>",
  "client_secret": "<databricks-oauth-secret>"
}
```

Use a Databricks OAuth secret, not an Entra application secret. For account
operations, use the appropriate accounts console hostname and add
`"account_id": "<databricks-account-id>"`. The principal needs the corresponding
account-level permissions.

### Azure service principal with a client secret

```json
{
  "host": "https://adb-1234567890123456.7.azuredatabricks.net",
  "auth_type": "azure-client-secret",
  "azure_client_id": "<entra-application-client-id>",
  "azure_client_secret": "<entra-client-secret>",
  "azure_tenant_id": "<entra-tenant-id>"
}
```

An optional `azure_workspace_resource_id` identifies the Azure workspace
resource. `azure_environment` selects the SDK's Azure cloud environment.
These are different from the generic `client_id`/`client_secret` keys used by
Databricks OAuth.

### Azure managed identity through the Databricks SDK

```json
{
  "host": "https://adb-1234567890123456.7.azuredatabricks.net",
  "auth_type": "azure-msi",
  "azure_use_msi": "true",
  "azure_client_id": "<user-assigned-managed-identity-client-id>"
}
```

Omit `azure_client_id` to use the available system-assigned identity. The
managed-identity endpoint must be reachable from the provider's runtime.
Simply running on AKS or annotating a service account does not make this path
equivalent to OIDC federation. For Kubernetes service-account federation, prefer
`OIDCTokenRequest`.

### Google credentials

For a Google service-account credential JSON, the value of `google_credentials`
is itself a JSON-encoded **string**:

```json
{
  "host": "https://your-workspace.gcp.databricks.com",
  "auth_type": "google-credentials",
  "google_credentials": "<JSON-encoded Google service-account credentials>"
}
```

The placeholder above must be replaced with the actual credential JSON,
properly escaped as a string. Do not put an object directly in that field.
Alternatively, `auth_type: "google-id"` and
`google_service_account: "<service-account-email>"` select the SDK's Google
service-account authentication path. Supply the ambient Google credentials and
impersonation permissions that method requires in the provider runtime.

### Forwarded JSON fields

The provider forwards only these credential JSON keys:

```text
host, account_id, auth_type, token,
client_id, client_secret,
azure_workspace_resource_id, azure_use_msi,
azure_client_id, azure_client_secret, azure_tenant_id, azure_environment,
google_credentials, google_service_account,
subscription_id, tenant_id, environment
```

Unrecognized JSON keys are not forwarded. The last three generic keys are
legacy compatibility fields, not substitutes for the SDK's Azure-prefixed
fields.

For `Secret` and `Filesystem`, put the host and account ID in the JSON.
Top-level `spec.host` and `spec.accountID` are used by `OIDCTokenRequest` and
offline planning; the normal Secret/Filesystem path does not merge them into
its credential configuration.

## Filesystem credentials and ambient SDK authentication

`Filesystem` reads the same JSON format from a path inside the provider
container:

```yaml
spec:
  credentials:
    source: Filesystem
    fs:
      path: /var/run/databricks/credentials.json
```

Mount that file using the provider's DeploymentRuntimeConfig. Files on the
operator's workstation are not visible to the container.

To let the Databricks SDK use environment variables or its profile file,
provide valid JSON such as `{}` through a Secret or file, then configure the
provider pod with the SDK's settings. Examples include `DATABRICKS_HOST`,
`DATABRICKS_TOKEN`, `DATABRICKS_CLIENT_ID`, `DATABRICKS_CLIENT_SECRET`,
`DATABRICKS_AUTH_TYPE`, and a mounted `.databrickscfg` selected with
`DATABRICKS_CONFIG_FILE`/`DATABRICKS_CONFIG_PROFILE`.

For SDK-native file OIDC federation, the current SDK exposes the `file-oidc`
auth type and `DATABRICKS_OIDC_TOKEN_FILEPATH`. This is distinct from this
provider's legacy `OIDCTokenFile` source: it uses SDK environment configuration
and requires a Databricks federation policy that trusts the token issuer,
subject, and audience. Its additional file-path key cannot be passed through
the credential JSON allowlist above. SDK-specific methods remain subject to
their own runtime prerequisites and the embedded SDK version.

The SDK also supports `env-oidc`: `DATABRICKS_OIDC_TOKEN_ENV` names the
environment variable containing the OIDC token, rather than containing the
token itself. Supply the host and client ID and set
`DATABRICKS_AUTH_TYPE=env-oidc`. Like `file-oidc`, it requires a matching
Databricks federation policy. An environment variable is not automatically
rotated inside an already running pod; prefer a rotating token file when
appropriate.

Authentication CLI methods require the corresponding CLI and login state in
the runtime image. They are not a recommended production Kubernetes setup, and
the standard image does not install Azure CLI, Google Cloud CLI, or Databricks
CLI. Prefer explicit credentials or token exchange.

## OIDC token authentication: what it does and when to use it

OIDC federation replaces a stored client secret with a signed, short-lived
workload token. The identity provider verifies that token against an explicitly
configured issuer, subject, and audience, then issues an access token for the
target API. An OIDC token is an assertion used in an exchange, not generally a
Databricks API bearer token by itself.

This provider exposes **two different source options**:

| Option | Who obtains the OIDC token? | Typical use |
| --- | --- | --- |
| `OIDCTokenRequest` | The provider calls Kubernetes TokenRequest for a selected service account. | Passwordless Azure Databricks with per-namespace or explicitly shared identities. |
| `OIDCTokenFile` | An external runtime/webhook projects a token file into the provider pod. | Legacy pod-level federation; subject to the compatibility limitation below. |

### Recommended: `OIDCTokenRequest`

#### What happens during reconciliation

1. Resolve `credentials.serviceAccountRef`.
2. Read Entra client/tenant IDs from `credentials.tokenExchange.azureAD`, or
   fall back independently to the service account's
   `azure.workload.identity/client-id` and
   `azure.workload.identity/tenant-id` annotations.
3. Request a short-lived Kubernetes JWT for that service account. The requested
   lifetime is 600 seconds; its subject is
   `system:serviceaccount:<namespace>:<name>`.
4. Include `api://AzureADTokenExchange` in the requested audiences.
   `serviceAccountRef.audiences` adds audiences; it does not replace this one.
5. Exchange the JWT at Entra ID using a client-credentials grant with a JWT
   client assertion. The requested scope is
   `2ff814a6-3304-4ab8-85cb-cd0e6f879c1d/.default` (Azure Databricks).
6. Configure Databricks with the resulting access token as a bearer token.
   Internally `auth_type` is `pat`; **this does not create or store a Databricks
   PAT**. The value is the Entra access token.

Tokens are cached in process and refreshed after approximately half their
lifetime, when fewer than five minutes remain, or after one hour of reuse,
whichever threshold is reached first. Concurrent requests share an exchange.
Revoking RBAC or federation trust does not invalidate an already issued access
token immediately. An in-flight asynchronous operation retains the token it
started with.

#### When to use it

- Azure Databricks automation without long-lived client secrets.
- Multi-tenant clusters where each namespace should use its own Entra identity.
- Platforms where workloads should opt in by granting narrowly scoped
  service-account TokenRequest permissions.
- Shared platform identities through ClusterProviderConfig, when sharing is
  intentional.

The only implemented `tokenExchange.type` is **`AzureAD`**, which is also the
default. This source is not currently an AWS/GCP or generic Databricks OIDC
exchange implementation.

#### Prerequisites

1. The Kubernetes cluster exposes a service-account OIDC issuer and signing
   keys that Entra ID can validate.
2. An Entra application or managed identity has a federated identity credential
   with:
   - Issuer: the cluster's exact service-account issuer URL.
   - Subject: `system:serviceaccount:team-a:databricks` for the example below.
   - Audience: `api://AzureADTokenExchange`.
3. That identity is provisioned as a Databricks service principal and granted
   the required workspace/account permissions.
4. The **provider pod's service account** may create tokens for the selected
   service account and, when using annotation fallbacks, read that account.
   These are different service accounts; use the actual provider runtime
   service-account name, ideally stabilized through DeploymentRuntimeConfig.
5. The provider can reach the Kubernetes API, the chosen Entra authority, and
   the Azure Databricks host.

No workload-identity mutating webhook or projected token volume is required
for this path: the provider requests the token through the Kubernetes API.
Service-account annotations alone do not establish federation trust.

#### Namespaced example and opt-in RBAC

Replace the GUIDs, host, and provider runtime service-account subject.
Create the namespace separately if it does not exist.

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: databricks
  namespace: team-a
  annotations:
    azure.workload.identity/client-id: "11111111-1111-1111-1111-111111111111"
    azure.workload.identity/tenant-id: "22222222-2222-2222-2222-222222222222"
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: databricks-tokenrequest
  namespace: team-a
rules:
  - apiGroups: [""]
    resources: ["serviceaccounts/token"]
    resourceNames: ["databricks"]
    verbs: ["create"]
  - apiGroups: [""]
    resources: ["serviceaccounts"]
    resourceNames: ["databricks"]
    verbs: ["get"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: databricks-tokenrequest
  namespace: team-a
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: databricks-tokenrequest
subjects:
  - kind: ServiceAccount
    name: provider-databricks
    namespace: crossplane-system
---
apiVersion: databricks.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
  namespace: team-a
spec:
  credentials:
    source: OIDCTokenRequest
    serviceAccountRef:
      name: databricks
    tokenExchange:
      type: AzureAD
  host: https://adb-1234567890123456.7.azuredatabricks.net
```

The provider pins `serviceAccountRef.namespace` to the namespaced
ProviderConfig's namespace and rejects a different namespace. The `get` rule
can be omitted if both IDs are explicitly set under `tokenExchange.azureAD`,
since annotation discovery is then unnecessary. Token creation remains
required.

#### Explicit identity settings and shared/account configuration

For shared access, place the service account and its Role/RoleBinding in a
platform-controlled namespace and reference it explicitly:

```yaml
apiVersion: databricks.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: account
spec:
  credentials:
    source: OIDCTokenRequest
    serviceAccountRef:
      name: databricks-account
      namespace: platform
    tokenExchange:
      type: AzureAD
      azureAD:
        clientID: "11111111-1111-1111-1111-111111111111"
        tenantID: "22222222-2222-2222-2222-222222222222"
        authorityHost: https://login.microsoftonline.com/
  host: https://accounts.azuredatabricks.net
  accountID: "33333333-3333-3333-3333-333333333333"
```

For a legacy cluster-scoped ProviderConfig, use
`apiVersion: databricks.crossplane.io/v1beta1` and `kind: ProviderConfig`;
`serviceAccountRef.namespace` is still required.

Explicit IDs belong under `credentials.tokenExchange.azureAD` for this
source, **not** at top-level `spec.clientID`/`spec.tenantID`.

AzureAD exchange hosts must be HTTPS URLs under `azuredatabricks.net`,
`databricks.azure.us`, or `databricks.azure.cn`, with no port or path other
than an optional trailing slash. The account ID must be a GUID.
Allowed authority hosts are:

- `https://login.microsoftonline.com/` (default).
- `https://login.microsoftonline.us/`.
- `https://login.chinacloudapi.cn/`.
- `https://login.partner.microsoftonline.cn/`.

Use the authority matching your tenant/cloud. Allowlisting these hosts is not
a guarantee that every resource has been tested in every sovereign cloud.

Complete examples:
[namespaced ProviderConfig](../examples/namespaced/providerconfig-oidctokenrequest.yaml),
[ClusterProviderConfig](../examples/namespaced/clusterproviderconfig-oidctokenrequest.yaml),
and [legacy ProviderConfig](../examples/cluster/providerconfig-oidctokenrequest.yaml).

### Legacy mounted-token option: `OIDCTokenFile`

This path expects a token already projected into the provider pod and requires
top-level `subscriptionID`, `tenantID`, and `clientID`. File selection is:

1. `spec.oidcTokenFilePath`, if specified.
2. `AZURE_FEDERATED_TOKEN_FILE` in the provider process.
3. `/var/run/secrets/azure/tokens/azure-identity-token`.

Its intended use is pod-level federation: the runtime provides the token and
its rotation, and a matching federated identity credential must trust that
token's issuer/subject/audience.

**Current compatibility limitation:** this branch sets legacy Terraform keys
`use_oidc`, `oidc_token_file_path`, `subscription_id`, and `tenant_id`. These
are not the current Databricks SDK's native file-OIDC configuration fields.
The branch's configuration-building tests do not establish end-to-end
authentication support with the embedded Databricks provider. Prefer
`OIDCTokenRequest` for Azure, or deliberately configure SDK-native federation
through the environment as described above. Do not assume this legacy source
works simply because the CRD accepts it.

### Legacy `Upbound` and named managed-identity sources

`Upbound` uses the same legacy OIDC setup but fixes the token path to
`/var/run/secrets/upbound.io/provider/token`. It requires top-level
`subscriptionID`, `tenantID`, and `clientID`, and a runtime that actually
projects that token. It has the same Databricks SDK compatibility limitation
as `OIDCTokenFile`.

`UserAssignedManagedIdentity` and `SystemAssignedManagedIdentity` both call the
same helper. It requires top-level `subscriptionID` and `tenantID`, forwards
optional `clientID`, `msiEndpoint`, and `environment`, and sets `use_msi`.
The source name itself does not enforce whether `clientID` is present. These
generic keys differ from Databricks SDK keys such as `azure_use_msi` and
`azure_client_id`. Prefer the explicit `azure-msi` JSON configuration above
when using an actual managed-identity endpoint.

## Troubleshooting

| Symptom | Check |
| --- | --- |
| `cannot unmarshal databricks credentials as JSON` | The Secret/file contains a JSON object with string values, not a bare token or INI file. `source: None` currently reaches this error too. |
| `cannot get referenced ProviderConfig` | Correct API group, namespace, name, and `providerConfigRef.kind` for the resource scope. |
| `no client ID` / `no tenant ID` | For TokenRequest, set the nested AzureAD fields or the selected service account's annotations. |
| TokenRequest `Forbidden` | RoleBinding names the actual provider pod service account; Role allows `create` on `serviceaccounts/token` for the selected account in its namespace. |
| Service-account `get` is forbidden | Grant the narrow `get` permission, or explicitly configure both AzureAD IDs. |
| `AADSTS700213` / no matching federated identity | Compare issuer, namespace/name subject, audience, client ID, and tenant with the Entra federated identity credential. |
| Token exchange succeeds but Databricks denies access | Provision the identity in the workspace/account and grant the appropriate permissions; check the host and account ID. |
| Authentication conflicts | Remove unrelated credentials from JSON, environment variables, and profiles; choose an explicit compatible `auth_type`. |
| Legacy identity source fails | Do not confuse accepted CRD values with current SDK support. Use TokenRequest or SDK-native authentication. |

Never commit real credential material, JWTs, or access tokens. Restrict Secret
access and TokenRequest permissions, rotate stored secrets, and avoid dumping
credential payloads when collecting diagnostics.

## Offline diff server

Offline planning is deliberately different from reconciliation: it uses local
placeholder credentials, does not read credential Secrets or request/exchange
service-account tokens, and blocks API requests. It does not validate whether
your real authentication settings or Databricks permissions work.
See [the offline diff documentation](../README.md#offline-diff-client).

## Implementation references

- [ProviderConfig schema](../apis/namespaced/v1beta1/types.go).
- [Token exchange schema](../apis/namespaced/v1beta1/tokenauth_types.go).
- [Credential selection and JSON forwarding](../internal/clients/databricks.go).
- [Service-account exchange wiring](../internal/clients/oidctokenrequest.go).
- [Token exchange and refresh cache](../internal/clients/tokenauth.go).