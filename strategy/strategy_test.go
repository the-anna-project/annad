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

func Test_Strategy_NewStrategy_RequestorError(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.CLGNames = []string{"CLG", "CLG"}
	// Requestor configuration is missing.

	_, err := NewStrategy(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
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
