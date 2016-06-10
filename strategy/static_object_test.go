package strategy

import (
	"testing"
)

func Test_Strategy_Static_GetID(t *testing.T) {
	firstStrategy := testMaybeNewStatic(t, nil)
	secondStrategy := testMaybeNewStatic(t, nil)

	if firstStrategy.GetID() == secondStrategy.GetID() {
		t.Fatal("expected", "different IDs", "got", "equal IDs")
	}
}

func Test_Strategy_Static_GetType(t *testing.T) {
	newStrategy := testMaybeNewStatic(t, nil)

	if newStrategy.GetType() != ObjectTypeStaticStrategy {
		t.Fatalf("invalid object type for strategy")
	}
}
