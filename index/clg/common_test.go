package clg

import (
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeFatalCase(t *testing.T, index int, err error) {
	if err != nil {
		t.Fatal("case", index+1, "expected", nil, "got", err)
	}
}

func testMaybeNewCLGCollection(t *testing.T) spec.CLGCollection {
	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newCLGIndex.GetCLGCollection()
}
