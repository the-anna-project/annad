package strategy

import (
	"strings"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func Test_Strategy_NewStrategy_Success(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.CLGNames = []string{"CLG", "CLG"}
	newConfig.Requestor = spec.ObjectType("requestor")

	_, err := NewStrategy(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Strategy_NewStrategy_CLGNamesError(t *testing.T) {
	newConfig := DefaultConfig()
	// CLGNames configuration is missing.
	newConfig.Requestor = spec.ObjectType("requestor")

	_, err := NewStrategy(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_NewStrategy_IDError(t *testing.T) {
	newStringMap := map[string]string{
		"string:clg-names": "[]string:CLG,CLG",
		// Note the ID configuration is missing.
		"string:requestor": "string:requestor",
	}

	newConfig := DefaultConfig()
	newConfig.StringMap = newStringMap
	_, err := NewStrategy(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_NewStrategy_RequestorError(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.CLGNames = []string{"CLG", "CLG"}
	// Requestor configuration is missing.

	_, err := NewStrategy(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_GetStringMap(t *testing.T) {
	newStringMap := map[string]string{
		"string:clg-names": "[]string:CLG,CLG",
		"string:id":        "string:id",
		"string:requestor": "string:requestor",
		"string:type":      "string:type",
	}

	newConfig := DefaultConfig()
	newConfig.StringMap = newStringMap
	newStrategy, err := NewStrategy(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Note that we only check for the first CLG name since the CLG names are
	// randomized on strategy creation. Having "CLG" given two times, should
	// always result in the first CLG name being "CLG". The second may be
	// omitted, dependening on the current randomization.
	if newStrategy.GetCLGNames()[0] != "CLG" {
		t.Fatal("expected", "CLG", "got", newStrategy.GetCLGNames()[0])
	}
	if newStrategy.GetID() != spec.ObjectID("id") {
		t.Fatal("expected", spec.ObjectID("id"), "got", newStrategy.GetID())
	}
	if newStrategy.GetRequestor() != spec.ObjectType("requestor") {
		t.Fatal("expected", spec.ObjectType("requestor"), "got", newStrategy.GetRequestor())
	}
	if newStrategy.GetType() != spec.ObjectType("type") {
		t.Fatal("expected", spec.ObjectType("type"), "got", newStrategy.GetType())
	}

	output, err := newStrategy.GetStringMap()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if len(output) != len(newStringMap) {
		t.Fatal("expected", len(newStringMap), "got", len(output))
	}
	// Note that we only check for the first CLG name since the actions are
	// randomized on strategy creation. Having "CLG" given two times, should
	// always result in the first CLG name being "CLG". The second may be
	// omitted, dependening on the current randomization.
	if !strings.Contains(output["string:clg-names"], "[]string:CLG") {
		t.Fatal("expected", true, "got", false)
	}
	if output["string:id"] != "string:id" {
		t.Fatal("expected", "string:id", "got", output["string:id"])
	}
	if output["string:requestor"] != "string:requestor" {
		t.Fatal("expected", "string:requestor", "got", output["string:requestor"])
	}
	if output["string:type"] != "string:type" {
		t.Fatal("expected", "string:type", "got", output["string:type"])
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
	newConfig.CLGNames = []string{"one", "two"}
	newConfig.Requestor = spec.ObjectType("requestor")

	for i := 0; i < 10; i++ {
		newStrategy, err := NewStrategy(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		var foundValid bool
		clgNames := strings.Join(newStrategy.GetCLGNames(), ",")
		for _, v := range valid {
			if v == clgNames {
				foundValid = true
				break
			}
		}

		if !foundValid {
			t.Fatal("expected", true, "got", false)
		}
	}
}
