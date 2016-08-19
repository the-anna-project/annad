package collection

import (
	"testing"
)

func Test_Collection_GetID(t *testing.T) {
	firstCollection := testMaybeNewCollection(t)
	secondCollection := testMaybeNewCollection(t)

	if firstCollection.GetID() == secondCollection.GetID() {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_Collection_GetType(t *testing.T) {
	newCollection := testMaybeNewCollection(t)
	objectType := newCollection.GetType()

	if objectType != ObjectTypeCLGCollection {
		t.Fatal("expected", ObjectTypeCLGCollection, "got", objectType)
	}
}
