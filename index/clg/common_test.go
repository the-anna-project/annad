package clg

import (
	"fmt"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeFatalCase(t *testing.T, index int, err error) {
	if err != nil {
		t.Fatal("case", index+1, "expected", nil, "got", err)
	}
}

func testMaybeNewCLGCollection(t *testing.T) spec.CLGCollection {
	return testMaybeNewCLGIndex(t).GetCLGCollection()
}

func testMaybeNewCLGIndex(t *testing.T) spec.CLGIndex {
	newCLGIndex, err := NewCLGIndex(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newCLGIndex
}

// TODO
func Test_CLGIndex_getCLGMethodHash(t *testing.T) {
	hash, err := testMaybeNewCLGIndex(t).(*clgIndex).getCLGMethodHash("clgName", "clgBody")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	fmt.Printf("%#v\n", hash)
}
