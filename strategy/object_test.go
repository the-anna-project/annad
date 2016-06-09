package strategy

import (
	"testing"

	"github.com/xh3b4sd/anna/clg"
	"github.com/xh3b4sd/anna/spec"
)

func Test_Strategy_GetID(t *testing.T) {
	firstStrategy := testMaybeNewStrategy(t)
	secondStrategy := testMaybeNewStrategy(t)

	if firstStrategy.GetID() == secondStrategy.GetID() {
		t.Fatal("expected", "different IDs", "got", "equal IDs")
	}
}

func Test_Strategy_GetType(t *testing.T) {
	newStrategy := testMaybeNewStrategy(t)

	if newStrategy.GetType() != ObjectTypeStrategy {
		t.Fatalf("invalid object type for strategy")
	}
}
