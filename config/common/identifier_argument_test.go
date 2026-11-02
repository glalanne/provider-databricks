package common

import "testing"

func TestIdentifierArgumentExternalName(t *testing.T) {
	en := IdentifierArgumentExternalName("policy_id")

	base := map[string]any{"policy_name": "p"}
	en.SetIdentifierArgumentFn(base, "abc")
	if base["policy_id"] != "abc" || base["policy_name"] != "p" {
		t.Fatalf("policy_id not set: %v", base)
	}

	empty := map[string]any{"policy_name": "p"}
	en.SetIdentifierArgumentFn(empty, "")
	if len(empty) != 1 || empty["policy_name"] != "p" {
		t.Fatalf("empty external name must leave base unchanged: %v", empty)
	}

	got, err := en.GetExternalNameFn(base)
	if err != nil || got != "abc" {
		t.Fatalf("round trip: got %q, err %v", got, err)
	}
	for _, st := range []map[string]any{{}, {"policy_id": nil}, {"policy_id": ""}} {
		if _, err := en.GetExternalNameFn(st); err == nil {
			t.Fatalf("expected error for state %v", st)
		}
	}

	if len(en.TFPluginFrameworkOptions.ComputedIdentifierAttributes) != 1 {
		t.Fatal("policy_id must be a computed identifier attribute")
	}
}
