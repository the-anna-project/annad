package clg

import (
	"testing"
)

func Test_Index_GetType(t *testing.T) {
	newConfig := DefaultConfig()
	newIndex, err := NewIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newIndex.GetType() != ObjectTypeIndex {
		t.Fatal("expected", ObjectTypeIndex, "got", newIndex.GetType())
	}
}

func Test_Index_GetID(t *testing.T) {
	firstIndex, err := NewIndex(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	secondIndex, err := NewIndex(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if firstIndex.GetID() == secondIndex.GetID() {
		t.Fatal("expected", false, "got", true)
	}
}
