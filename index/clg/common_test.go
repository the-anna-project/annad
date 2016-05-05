package clg

import (
	"fmt"
	"reflect"
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

func Test_CLGIndex_getCLGMethodHash(t *testing.T) {
	newCLGIndex, err := NewCLGIndex(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	v := reflect.ValueOf(newCLGIndex.GetCLGCollection())

	hash, err := newCLGIndex.(*clgIndex).getCLGMethodHash(v)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	fmt.Printf("%#v\n", hash)
}
