package clg

import (
	"testing"
)

func Test_Index_GetType(t *testing.T) {
	newIndex := testMaybeNewIndex(t)

	if newIndex.GetType() != ObjectTypeCLGIndex {
		t.Fatal("expected", ObjectTypeCLGIndex, "got", newIndex.GetType())
	}
}

func Test_Index_GetID(t *testing.T) {
	firstIndex := testMaybeNewIndex(t)
	secondIndex := testMaybeNewIndex(t)

	if firstIndex.GetID() == secondIndex.GetID() {
		t.Fatal("expected", false, "got", true)
	}
}
