package strategy

import (
	"strings"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func Test_Strategy_NewStrategy_Success(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Actions = []spec.ObjectType{spec.ObjectType("action"), spec.ObjectType("action")}
	newConfig.Requestor = spec.ObjectType("requestor")

	_, err := NewStrategy(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Strategy_NewStrategy_ActionsError(t *testing.T) {
	newConfig := DefaultConfig()
	// Action configuration is missing.
	newConfig.Requestor = spec.ObjectType("requestor")

	_, err := NewStrategy(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_NewStrategy_IDError(t *testing.T) {
	newHashMap := map[string]string{
		"actions": "action,action",
		// Note the ID configuration is missing.
		"requestor": "requestor",
	}

	newConfig := DefaultConfig()
	newConfig.HashMap = newHashMap
	_, err := NewStrategy(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_NewStrategy_RequestorError(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Actions = []spec.ObjectType{spec.ObjectType("action"), spec.ObjectType("action")}
	// Requestor configuration is missing.

	_, err := NewStrategy(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_GetHashMap(t *testing.T) {
	newHashMap := map[string]string{
		"actions":   "action,action",
		"id":        "id",
		"requestor": "requestor",
	}

	newConfig := DefaultConfig()
	newConfig.HashMap = newHashMap
	newStrategy, err := NewStrategy(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Note that we only check for the first action since the actions are
	// randomized on strategy creation. Having "action" given two times, should
	// always result in the first action being "action". The second may be
	// omitted, dependening on the current randomization.
	if newStrategy.GetActions()[0] != spec.ObjectType("action") {
		t.Fatal("expected", spec.ObjectType("action"), "got", newStrategy.GetActions()[0])
	}
	if newStrategy.GetID() != spec.ObjectID("id") {
		t.Fatal("expected", spec.ObjectID("id"), "got", newStrategy.GetID())
	}
	if newStrategy.GetRequestor() != spec.ObjectType("requestor") {
		t.Fatal("expected", spec.ObjectType("requestor"), "got", newStrategy.GetRequestor())
	}

	output := newStrategy.GetHashMap()

	if len(output) != len(newHashMap) {
		t.Fatal("expected", len(newHashMap), "got", len(output))
	}
	// Note that we only check for the first action since the actions are
	// randomized on strategy creation. Having "action" given two times, should
	// always result in the first action being "action". The second may be
	// omitted, dependening on the current randomization.
	if !strings.Contains(output["actions"], "action") {
		t.Fatal("expected", true, "got", false)
	}
	if output["id"] != "id" {
		t.Fatal("expected", "id", "got", output["id"])
	}
	if output["requestor"] != "requestor" {
		t.Fatal("expected", "requestor", "got", output["requestor"])
	}
}

// Test_Strategy_NewStrategy ensures that a specific combination of a strategy
// is created on a given seed. This test also checks that the string
// representation works as expected.
func Test_Strategy_NewStrategy(t *testing.T) {
	valid := []string{
		"one",
		"two",
		"one,one",
		"two,two",
		"one,two",
		"two,one",
	}

	newConfig := DefaultConfig()
	newConfig.Actions = []spec.ObjectType{spec.ObjectType("one"), spec.ObjectType("two")}
	newConfig.Requestor = spec.ObjectType("requestor")

	for i := 0; i < 10; i++ {
		newStrategy, err := NewStrategy(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		var foundValid bool
		s := newStrategy.ActionsToString()
		for _, v := range valid {
			if v == s {
				foundValid = true
				break
			}
		}

		if !foundValid {
			t.Fatal("expected", true, "got", false)
		}
	}
}
