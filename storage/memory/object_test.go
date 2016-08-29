package memory

import (
	"testing"
)

func Test_Storage_GetID(t *testing.T) {
	firstStorage := MustNew()
	secondStorage := MustNew()

	if firstStorage.GetID() == secondStorage.GetID() {
		t.Fatal("expected", "different IDs", "got", "equal IDs")
	}
}

func Test_Storage_GetType(t *testing.T) {
	newStorage := MustNew()

	if newStorage.GetType() != ObjectType {
		t.Fatal("expected", ObjectType, "got", newStorage.GetType())
	}
}
