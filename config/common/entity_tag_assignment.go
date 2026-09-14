package common

import (
	"fmt"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
)

var EntityTagAssignmentExternalName = config.NewExternalNameFrom(
	config.IdentifierFromProvider,
	config.WithGetExternalNameFn(func(_ config.GetExternalNameFn, tfState map[string]any) (string, error) {
		values := make([]string, 3)
		for i, field := range []string{"entity_type", "entity_name", "tag_key"} {
			value, ok := tfState[field].(string)
			if !ok || value == "" {
				return "", fmt.Errorf("cannot find %q in tfstate", field)
			}
			values[i] = value
		}
		return strings.Join(values, ","), nil
	}),
)
