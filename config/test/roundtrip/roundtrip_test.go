// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// package roundtrip contains API roundtrip tests
// for multi-version managed resource APIs.
package roundtrip

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/apitesting/roundtrip"
	"github.com/databricks/terraform-provider-databricks/xpprovider"
	"k8s.io/apimachinery/pkg/runtime"

	clusterapis "github.com/glalanne/provider-databricks/apis/cluster"
	namespacedapis "github.com/glalanne/provider-databricks/apis/namespaced"
	"github.com/glalanne/provider-databricks/config"
)

// TestRoundTrip configures and invokes the API roundtrip tests.
// For each registered API, a serialization roundtrip test and
// conversion roundtrip test is invoked, with the given fuzzer
// configurations.
func TestRoundTrip(t *testing.T) {
	fwProvider, sdkProvider, err := xpprovider.GetProvider(t.Context())
	if err != nil {
		t.Fatalf("GetProvider: %s", err)
	}
	provider, err := config.GetProvider(t.Context(), fwProvider, sdkProvider, false)
	if err != nil {
		t.Fatalf("GetProvider: %s", err)
	}
	providerNamespaced, err := config.GetProviderNamespaced(t.Context(), fwProvider, sdkProvider, false)
	if err != nil {
		t.Fatalf("GetProviderNamespaced: %s", err)
	}

	testScheme := runtime.NewScheme()
	if err := clusterapis.AddToScheme(testScheme); err != nil {
		t.Fatalf("cluster-scoped apis AddToScheme: %s", err)
	}
	if err := namespacedapis.AddToScheme(testScheme); err != nil {
		t.Fatalf("namespaced apis AddToScheme: %s", err)
	}

	rt, err := roundtrip.NewRoundTripTest(provider, providerNamespaced, testScheme,
		roundtrip.WithFuzzerConfig(
			roundtrip.FuzzerIterations(10),
			roundtrip.FuzzerNilChance(0)),
		roundtrip.WithFuzzerConfig(
			roundtrip.FuzzerIterations(30),
			roundtrip.FuzzerNilChance(0.3)),
		roundtrip.WithComparisonOptions(
			roundtrip.EquateEmptyAndSingleZeroSlice(),
			roundtrip.EquateNilAndZeroValuePtr(),
		),
	)
	if err != nil {
		t.Fatalf("NewRoundTripTest: %s", err)
	}

	t.Run("TestSerializationRoundtrip", func(t *testing.T) {
		rt.TestSerializationRoundtrip(t)
	})

	t.Run("TestConversionRoundtrip", func(t *testing.T) {
		rt.TestConversionRoundtrip(t)
	})
}