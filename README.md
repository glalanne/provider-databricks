
# Provider Databricks

<div style="text-align: center;">

![CI](https://github.com/glalanne/provider-databricks/workflows/CI/badge.svg)
[![GitHub release](https://img.shields.io/github/release/glalanne/provider-databricks/all.svg)](https://github.com/glalanne/provider-databricks/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/glalanne/provider-databricks)](https://goreportcard.com/report/github.com/glalanne/provider-databricks)
[![Contributors](https://img.shields.io/github/contributors/glalanne/provider-databricks)](https://github.com/glalanne/provider-databricks/graphs/contributors)


</div>

`provider-databricks` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
[Databricks](https://registry.terraform.io/providers/databricks/databricks/latest/docs).


Most of the testing have been done on [Azure Databricks](https://azure.microsoft.com/en-ca/products/databricks)

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/glalanne/provider-databricks):
```
up ctp provider install glalanne/provider-databricks:v0.3.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-databricks
spec:
  package: glalanne/provider-databricks:v0.2.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/glalanne/provider-databricks).

## Exemples

A few exemples are provided [here](./examples/) to show case how to configure the provider, and how to use some resources


## Developing

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

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/glalanne/provider-databricks/issues).
