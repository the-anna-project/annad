package strategy

import (
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func Test_Strategy_NewStrategy_Error(t *testing.T) {
	newFirstConfig := DefaultConfig()
	_, err := NewStrategy(newFirstConfig)
	if !IsInvalidActions(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_GetID(t *testing.T) {
	newFirstConfig := DefaultConfig()
	newFirstConfig.Actions = []spec.ObjectType{
		spec.ObjectType("test"),
	}
	firstStrategy, err := NewStrategy(newFirstConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newSecondConfig := DefaultConfig()
	newSecondConfig.Actions = []spec.ObjectType{
		spec.ObjectType("test"),
	}
	secondStrategy, err := NewStrategy(newSecondConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if firstStrategy.GetID() == secondStrategy.GetID() {
		t.Fatal("expected", "different IDs", "got", "equal IDs")
	}
}

func Test_Strategy_GetType(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Actions = []spec.ObjectType{
		spec.ObjectType("test"),
	}

	newStrategy, err := NewStrategy(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newStrategy.GetType() != ObjectTypeStrategy {
		t.Fatalf("invalid object tyoe of factory client")
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
	newConfig.Actions = []spec.ObjectType{
		spec.ObjectType("one"),
		spec.ObjectType("two"),
	}

	for i := 0; i < 10; i++ {
		newStrategy, err := NewStrategy(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		var foundValid bool
		s := newStrategy.String()
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
