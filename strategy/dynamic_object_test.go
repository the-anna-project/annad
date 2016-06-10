package strategy

import (
	"testing"
)

func Test_Strategy_Dynamic_GetID(t *testing.T) {
	firstStrategy := testMaybeNewDynamic(t, "Sum", nil)
	secondStrategy := testMaybeNewDynamic(t, "Sum", nil)

	if firstStrategy.GetID() == secondStrategy.GetID() {
		t.Fatal("expected", "different IDs", "got", "equal IDs")
	}
}

func Test_Strategy_Dynamic_GetType(t *testing.T) {
	newStrategy := testMaybeNewDynamic(t, "Sum", nil)

	if newStrategy.GetType() != ObjectTypeDynamicStrategy {
		t.Fatalf("invalid object type for strategy")
	}
}
