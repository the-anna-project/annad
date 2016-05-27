package redisstorage

import (
	"testing"
)

func Test_Storage_GetID(t *testing.T) {
	firstStorage := testMaybeNewStorage(t)
	secondStorage := testMaybeNewStorage(t)

	if firstStorage.GetID() == secondStorage.GetID() {
		t.Fatal("expected", "different IDs", "got", "equal IDs")
	}
}

func Test_Storage_GetType(t *testing.T) {
	newStorage := testMaybeNewStorage(t)

	if newStorage.GetType() != ObjectTypeRedisStorage {
		t.Fatal("expected", ObjectTypeRedisStorage, "got", newStorage.GetType())
	}
}
