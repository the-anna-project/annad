package clg

import (
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewIndex(t *testing.T) spec.CLGIndex {
	newIndexConfig, err := DefaultIndexConfig()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newIndex, err := NewIndex(newIndexConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newIndex
}
