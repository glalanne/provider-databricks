//go:build generate
// +build generate

/*
Copyright 2021 Upbound Inc.
*/

// NOTE: See the below link for details on what is happening here.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

// Remove existing CRDs
//go:generate rm -rf ../package/crds

// Remove generated files
//go:generate bash -c "find . -iname 'zz_*' ! -iname 'zz_generated.managed*.go' -delete"
//go:generate bash -c "find . -type d -empty -delete"
//go:generate bash -c "find ../internal/controller -iname 'zz_*' -delete"
//go:generate bash -c "find ../internal/controller -type d -empty -delete"
//go:generate bash -c "find ../cmd/provider -name 'zz_*' -type f -delete"
//go:generate bash -c "find ../cmd/provider -type d -maxdepth 1 -mindepth 1 -empty -delete"
//go:generate rm -rf ../examples-generated

// Generate documentation from Terraform docs.
//go:generate go run github.com/crossplane/upjet/v2/cmd/scraper -n databricks/databricks -r ../.work/databricks/databricks/docs/resources -o ../config/provider-metadata.yaml

// Run Upjet generator
//go:generate go run ../cmd/generator/main.go ..

// Generate deepcopy methodsets and CRD manifests
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=./... crd:allowDangerousTypes=true,crdVersions=v1 output:artifacts:config=../package/crds

// Generate crossplane-runtime methodsets (resource.Claim, etc)
//go:generate go run -tags generate github.com/crossplane/crossplane-tools/cmd/angryjet generate-methodsets --header-file=../hack/boilerplate.go.txt ./...

// Run upjet's transformer for the generated resolvers to get rid of the cross
// API-group imports and to prevent import cycles
// go:generate go run github.com/crossplane/upjet/v2/cmd/resolver -g databricks.crossplane.io -a github.com/glalanne/provider-databricks/internal/apis -o databricks.databricks.crossplane.io=databricks.crossplane.io -s -p ../apis/cluster/...
// go:generate go run github.com/crossplane/upjet/v2/cmd/resolver -g databricks.m.crossplane.io -a github.com/glalanne/provider-databricks/internal/apis -o databricks.databricks.m.crossplane.io=databricks.m.crossplane.io -s -p ../apis/namespaced/...

package generate

import (
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen" //nolint:typecheck

	_ "github.com/crossplane/crossplane-tools/cmd/angryjet" //nolint:typecheck

	_ "github.com/crossplane/upjet/v2/cmd/scraper"
)
