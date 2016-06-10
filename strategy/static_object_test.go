package strategy

import (
	"testing"
)

func Test_Strategy_Static_GetID(t *testing.T) {
	firstStrategy := testMaybeNewStatic(t)
	secondStrategy := testMaybeNewStatic(t)

	if firstStrategy.GetID() == secondStrategy.GetID() {
		t.Fatal("expected", "different IDs", "got", "equal IDs")
	}
}

func Test_Strategy_Static_GetType(t *testing.T) {
	newStrategy := testMaybeNewStatic(t)

	if newStrategy.GetType() != ObjectTypeStrategy {
		t.Fatalf("invalid object type for strategy")
	}
}
