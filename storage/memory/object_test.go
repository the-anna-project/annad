package memory

import (
	"testing"
)

func Test_Storage_GetID(t *testing.T) {
	firstStorage, err := NewStorage(DefaultStorageConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	secondStorage, err := NewStorage(DefaultStorageConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if firstStorage.GetID() == secondStorage.GetID() {
		t.Fatal("expected", "different IDs", "got", "equal IDs")
	}
}

func Test_Storage_GetType(t *testing.T) {
	newStorage, err := NewStorage(DefaultStorageConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newStorage.GetType() != ObjectType {
		t.Fatal("expected", ObjectType, "got", newStorage.GetType())
	}
}
