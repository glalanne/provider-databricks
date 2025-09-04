---
title: Quickstart
weight: 1
---

# Quickstart

This guide walks through the process to install the Databricks Provider.

To use provider, make sure Crossplane is already deployed into your Kubernetes cluster, then install the `Provider`, apply a `ProviderConfig`, and create a *managed resource* in Databricks via Kubernetes.


## Install the Databricks provider-databricks

Install the provider-databricks into the Kubernetes cluster with a Kubernetes configuration file.


```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-databricks
spec:
  package: xpkg.upbound.io/lalanne/provider-databricks:<version>
```

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.   

```shell
NAME                            INSTALLED   HEALTHY   PACKAGE                                                     AGE
provider-databricks             True        True      xpkg.upbound.io/lalanne/provider-databricks:v2.0.0          42s
```

It may take up to 5 minutes to report `HEALTHY`.

If you are going to use your own registry please check [Install Providers in an offline environment](https://docs.upbound.io/providers/provider-families/#installing-a-provider-family:~:text=services%20to%20install.-,Install%20Providers%20in%20an%20offline%20environment,-View%20the%20installed)

## Create a Kubernetes secret
The provider requires credentials to create and manage Databricks resources.


### Create a Kubernetes secret with the Databricks credentials JSON file
Use `kubectl create secret -n crossplane-system` to generate the Kubernetes secret object inside the Kubernetes cluster.

`kubectl create secret generic databricks-secret -n crossplane-system --from-file=credentials=./dbx-credentials.json`

dbx-credentials.json using a PAT
```json
{
    "host": "https://adb-1234567890.15.azuredatabricks.net",
    "token": "dapi2222222222222222222222-2",
    "auth_type": "pat"
}
```

dbx-credentials.json using a service account in Azure
```json
{
    "host": "https://adb-1234567890.15.azuredatabricks.net",
    "azure_client_id": "dfa6f769-4bde-4ea4-8fb9-ec47294024d3",
    "azure_client_secret": "1._SECR8T42.ECNy4w2JBddd",
    "azure_tenant_id": "000000-1cfa-49c2-a24f-000b00000",
    "auth_type": "azure-client-secret"
}
```

View the secret with `kubectl describe secret`
```shell
$ kubectl describe secret databricks-secret -n crossplane-system
Name:         databricks-secret
Namespace:    crossplane-system
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
credentials:  629 bytes
```
## Create a ProviderConfig
Create a `ProviderConfig` Kubernetes configuration file to attach the Databricks credentials to the installed `provider-databricks`.

```yaml
apiVersion: databricks.crossplane.io/v1beta1
metadata:
  name: default
kind: ProviderConfig
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: databricks-secret
      key: credentials
```

Apply this configuration with `kubectl apply -f`.

**Note:** the `Providerconfig` value `spec.secretRef.name` must match the `name` of the secret in `kubectl get secrets -n crossplane-system` and `spec.secretRef.key` must match the value in the `Data` section of the secret.

Verify the `ProviderConfig` with `kubectl describe providerconfigs`. 

```yaml
$ kubectl describe providerconfigs
Name:         default
Namespace:
API Version:  databricks.crossplane.io/v1beta1
Kind:         ProviderConfig
# Output truncated
Spec:
  Credentials:
    Secret Ref:
      Key:        credentials
      Name:       databricks-secret
      Namespace:  crossplane-system
    Source:       Secret
```
## Create a ClusterProviderConfig
Create a `ClusterProviderConfig` Kubernetes configuration file to attach the Databricks credentials to the installed `provider-databricks`.

```yaml
apiVersion: databricks.m.crossplane.io/v1beta1
metadata:
  name: default
kind: ClusterProviderConfig
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: databricks-secret
      key: credentials
```

Apply this configuration with `kubectl apply -f`.

**Note:** the `Providerconfig` value `spec.secretRef.name` must match the `name` of the secret in `kubectl get secrets -n crossplane-system` and `spec.secretRef.key` must match the value in the `Data` section of the secret.

Verify the `ProviderConfig` with `kubectl describe providerconfigs`. 

```yaml
$ kubectl describe providerconfigs
Name:         default
Namespace:
API Version:  databricks.crossplane.io/v1beta1
Kind:         ProviderConfig
# Output truncated
Spec:
  Credentials:
    Secret Ref:
      Key:        credentials
      Name:       databricks-secret
      Namespace:  crossplane-system
    Source:       Secret
```

## Create a managed resource
Create a managed resource to verify the `provider-databricks` is functioning. 

This example creates a Databricks Cluster resource in the Workspace.

```yaml
apiVersion: compute.databricks.crossplane.io/v1alpha1
kind: Cluster
metadata:
  annotations:
    meta.upbound.io/example-id: compute/v1alpha1/cluster
  labels:
    testing.upbound.io/example-name: shared_autoscaling
  name: shared-autoscaling-test
spec:
  forProvider:
    autoscale:
    - maxWorkers: 2
      minWorkers: 1
    autoterminationMinutes: 20
    nodeTypeId: Standard_D4plds_v6
    sparkVersion: 16.0.x-scala2.12
    clusterName: shared-autoscaling-test
  providerConfigRef:
    name: default

```

**Note:** the `spec.providerConfigRef.name` must match the `ProviderConfig` `metadata.name` value.

Apply this configuration with `kubectl apply -f`.

Use `kubectl get managed` to verify resource group creation.

```shell
$ kubectl get managed
NAME                                                               READY   SYNCED   EXTERNAL-NAME          AGE
cluster.compute.databricks.crossplane.io/shared-autoscaling-test   True     True    1209-040405-o14upcvf   2d15h
```

Provider created the Cluster when the values `READY` and `SYNCED` are `True`.

If the `READY` or `SYNCED` are blank or `False` use `kubectl describe` to understand why.
