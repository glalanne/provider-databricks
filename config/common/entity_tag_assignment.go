package common

import (
	"fmt"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
)

func CompositeExternalName(fields ...string) config.ExternalName {
	return config.NewExternalNameFrom(
		config.IdentifierFromProvider,
		config.WithGetExternalNameFn(func(_ config.GetExternalNameFn, tfState map[string]any) (string, error) {
			values := make([]string, len(fields))
			for i, field := range fields {
				value, ok := tfState[field]
				if !ok || value == nil {
					return "", fmt.Errorf("cannot find %q in tfstate", field)
				}
				values[i] = fmt.Sprint(value)
				if values[i] == "" {
					return "", fmt.Errorf("cannot find %q in tfstate", field)
				}
			}
			return strings.Join(values, ","), nil
		}),
	)
}

var EntityTagAssignmentExternalName = CompositeExternalName("entity_type", "entity_name", "tag_key")

var DataQualityRefreshExternalName = CompositeExternalName("object_type", "object_id", "refresh_id")

var RFAAccessRequestDestinationsExternalName = CompositeExternalName("securable_type", "full_name")

var ServicePrincipalFederationPolicyExternalName = CompositeExternalName("service_principal_id", "policy_id")

var WorkspaceEntityTagAssignmentExternalName = CompositeExternalName("entity_type", "entity_id", "tag_key")
