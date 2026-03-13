/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	"context"
	_ "embed"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	uname "github.com/crossplane/upjet/v2/pkg/types/name"
	tfjson "github.com/hashicorp/terraform-json"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/glalanne/provider-databricks/config/cluster"
	"github.com/glalanne/provider-databricks/config/namespaced"

	"github.com/pkg/errors"
)

const (
	resourcePrefix = "databricks"
	modulePath     = "github.com/glalanne/provider-databricks"
)

type generationMode int

const (
	generationModeV1Beta1 generationMode = iota
	generationModeV1Alpha1Legacy
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
func GetProvider(_ context.Context, fwProvider fwprovider.Provider, sdkProvider *tfschema.Provider, generationProvider bool) (*config.Provider, error) {
	pc, err := getProviderWithMode(fwProvider, sdkProvider, generationProvider, "databricks.crossplane.io", generationModeV1Beta1)
	if err != nil {
		return nil, err
	}
	for _, configure := range cluster.ProviderConfiguration {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc, nil
}

// GetProviderV1Alpha1Legacy returns provider configuration for generating
// the legacy v1alpha1 APIs (array/singleton-list shape).
func GetProviderV1Alpha1Legacy(_ context.Context, fwProvider fwprovider.Provider, sdkProvider *tfschema.Provider, generationProvider bool) (*config.Provider, error) {
	pc, err := getProviderWithMode(fwProvider, sdkProvider, generationProvider, "databricks.crossplane.io", generationModeV1Alpha1Legacy)
	if err != nil {
		return nil, err
	}
	for _, configure := range cluster.ProviderConfiguration {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc, nil
}

// GetProviderNamespaced returns provider configuration for namespace-scoped resources
func GetProviderNamespaced(_ context.Context, fwProvider fwprovider.Provider, sdkProvider *tfschema.Provider, generationProvider bool) (*config.Provider, error) {
	pc, err := getProviderWithMode(fwProvider, sdkProvider, generationProvider, "databricks.m.crossplane.io", generationModeV1Beta1)
	if err != nil {
		return nil, err
	}
	for _, configure := range namespaced.ProviderConfiguration {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc, nil
}

// GetProviderNamespacedV1Alpha1Legacy returns namespaced provider
// configuration for generating the legacy v1alpha1 APIs.
func GetProviderNamespacedV1Alpha1Legacy(_ context.Context, fwProvider fwprovider.Provider, sdkProvider *tfschema.Provider, generationProvider bool) (*config.Provider, error) {
	pc, err := getProviderWithMode(fwProvider, sdkProvider, generationProvider, "databricks.m.crossplane.io", generationModeV1Alpha1Legacy)
	if err != nil {
		return nil, err
	}
	for _, configure := range namespaced.ProviderConfiguration {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc, nil
}

func getProviderWithMode(fwProvider fwprovider.Provider, sdkProvider *tfschema.Provider, generationProvider bool, rootGroup string, mode generationMode) (*config.Provider, error) {
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

	providerOpts := []config.ProviderOption{
		config.WithRootGroup(rootGroup),
		config.WithIncludeList(CLIReconciledResourceList()),
		config.WithTerraformPluginSDKIncludeList(TerraformPluginSDKResourceList()),
		config.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),
		config.WithDefaultResourceOptions(ResourceConfigurator()),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
		config.WithFeaturesPackage("internal/features"),
		config.WithTerraformProvider(sdkProvider),
		config.WithTerraformPluginFrameworkProvider(fwProvider),
	}
	if mode == generationModeV1Beta1 {
		providerOpts = append(providerOpts, config.WithSchemaTraversers(&config.SingletonListEmbedder{}))
	}
	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata), providerOpts...)

	// Rename resources to make it more pleasing to the eye
	for _, r := range pc.Resources {
		parts := strings.Split(r.Name, "_")
		if len(parts) > 1 {
			r.ShortGroup = resourcePrefix
			r.Kind = uname.NewFromSnake(strings.Join(parts[1:], "_")).Camel
		}
	}

	if mode == generationModeV1Alpha1Legacy {
		setLegacyV1Alpha1(pc)
	} else {
		bumpVersionsWithEmbeddedLists(pc)
	}

	return pc, nil
}

// CLIReconciledResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the TF CLI based architecture.
func CLIReconciledResourceList() []string {
	l := make([]string, len(CLIReconciledExternalNameConfigs))
	i := 0
	for name := range CLIReconciledExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// TerraformPluginSDKResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the no-fork architecture.
func TerraformPluginSDKResourceList() []string {
	l := make([]string, len(TerraformPluginSDKExternalNameConfigs))
	i := 0
	for name := range TerraformPluginSDKExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

func TerraformPluginFrameworkResourceList() []string {
	l := make([]string, len(TerraformPluginFrameworkExternalNameConfigs))
	i := 0
	for name := range TerraformPluginFrameworkExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

func bumpVersionsWithEmbeddedLists(pc *config.Provider) {
	for name, r := range pc.Resources {
		r := r
		paths := r.CRDListConversionPaths()
		r.Version = "v1beta1"
		r.PreviousVersions = []string{"v1alpha1"}
		// Keep storage on the singleton API version so v1alpha1 can be
		// deprecated in a later release.
		r.SetCRDStorageVersion(r.Version)
		// Use v1alpha1 as conversion hub for legacy array-shaped APIs.
		r.SetCRDHubVersion("v1alpha1")
		r.Conversions = []conversion.Conversion{
			conversion.NewIdentityConversionExpandPaths(conversion.AllVersions, conversion.AllVersions, conversion.DefaultPathPrefixes(), paths...),
			conversion.NewSingletonListConversion("v1alpha1", "v1beta1", conversion.DefaultPathPrefixes(), paths, conversion.ToEmbeddedObject),
			conversion.NewSingletonListConversion("v1beta1", "v1alpha1", conversion.DefaultPathPrefixes(), paths, conversion.ToSingletonList),
		}
		if err := r.SetDeprecatedVersion("v1alpha1", config.VersionDeprecation{
			Warning: "This API version is deprecated. Please migrate to v1beta1.",
		}); err != nil {
			panic(err)
		}
		r.TerraformConversions = []config.TerraformConversion{
			config.NewTFSingletonConversion(),
		}
		pc.Resources[name] = r
	}
}

func setLegacyV1Alpha1(pc *config.Provider) {
	for name, r := range pc.Resources {
		r := r
		r.Version = "v1alpha1"
		r.PreviousVersions = nil
		r.SetCRDStorageVersion(r.Version)
		r.Conversions = nil
		r.TerraformConversions = nil
		pc.Resources[name] = r
	}
}
