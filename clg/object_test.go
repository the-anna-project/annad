package clg

import (
	"testing"
)

func Test_Index_GetType(t *testing.T) {
	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newCLGIndex.GetType() != ObjectTypeCLGIndex {
		t.Fatal("expected", ObjectTypeCLGIndex, "got", newCLGIndex.GetType())
	}
}

func Test_Index_GetID(t *testing.T) {
	newConfig := DefaultConfig()
	firstCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newConfig = DefaultConfig()
	secondCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if firstCLGIndex.GetID() == secondCLGIndex.GetID() {
		t.Fatal("expected", false, "got", true)
	}
}
