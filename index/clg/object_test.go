package clg

import (
	"testing"
)

func Test_Index_GetType(t *testing.T) {
	newConfig := DefaultIndexConfig()
	newIndex, err := NewIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newIndex.GetType() != ObjectTypeCLGIndex {
		t.Fatal("expected", ObjectTypeCLGIndex, "got", newIndex.GetType())
	}
}

func Test_Index_GetID(t *testing.T) {
	newConfig := DefaultIndexConfig()
	firstIndex, err := NewIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newConfig = DefaultIndexConfig()
	secondIndex, err := NewIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if firstIndex.GetID() == secondIndex.GetID() {
		t.Fatal("expected", false, "got", true)
	}
}
