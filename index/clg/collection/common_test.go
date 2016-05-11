package collection

import (
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeFatalCase(t *testing.T, index int, err error) {
	if err != nil {
		t.Fatal("case", index+1, "expected", nil, "got", err)
	}
}

func testMaybeNewCollection(t *testing.T) spec.CLGCollection {
	newCollection, err := New(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newCollection
}
