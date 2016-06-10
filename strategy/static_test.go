package strategy

import (
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewStatic(t *testing.T, argument interface{}) spec.Strategy {
	newConfig := DefaultStaticConfig()
	newConfig.Argument = argument
	newStrategy, err := NewStatic(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStrategy
}

func Test_Strategy_Static_NewStatic(t *testing.T) {
	// When this does not panic, the test is successfull.
	testMaybeNewStatic(t, "foo")
}

func Test_Strategy_Static_SetNode(t *testing.T) {
	newStrategy := testMaybeNewStatic(t, "foo")
	arg1 := testMaybeNewStatic(t, "bar")

	err := newStrategy.SetNode([]int{0}, arg1)
	if !IsNotSettable(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_Static_RemoveNode(t *testing.T) {
	newStrategy := testMaybeNewStatic(t, "foo")

	err := newStrategy.RemoveNode([]int{0})
	if !IsNotRemovable(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_Static_Validate(t *testing.T) {
	newStrategy := testMaybeNewStatic(t, "foo")
	err := newStrategy.Validate()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newStrategy = testMaybeNewStatic(t, nil)
	err = newStrategy.Validate()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}
