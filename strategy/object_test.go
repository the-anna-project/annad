package strategy

import (
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func Test_Strategy_GetID(t *testing.T) {
	newFirstConfig := DefaultConfig()
	newFirstConfig.CLGNames = []string{"CLG"}
	newFirstConfig.Requestor = spec.ObjectType("requestor")
	firstStrategy, err := NewStrategy(newFirstConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newSecondConfig := DefaultConfig()
	newSecondConfig.CLGNames = []string{"CLG"}
	newSecondConfig.Requestor = spec.ObjectType("requestor")
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
	newConfig.CLGNames = []string{"CLG"}
	newConfig.Requestor = spec.ObjectType("requestor")

	newStrategy, err := NewStrategy(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newStrategy.GetType() != ObjectTypeStrategy {
		t.Fatalf("invalid object type for strategy")
	}
}
