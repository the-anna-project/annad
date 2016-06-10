package strategy

import (
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewDynamic(t *testing.T, root spec.CLG, nodes []spec.Strategy) spec.Strategy {
	newConfig := DefaultDynamicConfig()
	newConfig.Root = root
	newConfig.Nodes = nodes
	newStrategy, err := NewDynamic(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStrategy
}

func Test_Strategy_Dynamic_New(t *testing.T) {
	// When this does not panic, the test is successfull.
	testMaybeNewDynamic(t, "Sum", nil)
}

func Test_Strategy_Dynamic_New_RootAndNodes(t *testing.T) {
	newConfig := DefaultDynamicConfig()
	newConfig.Root = "Sum"
	newConfig.Nodes = []spec.Strategy{testMaybeNewDynamic(t, "Sum", nil)}

	_, err := NewDynamic(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Strategy_Dynamic_New_Error_DefaultConfig(t *testing.T) {
	// DefaultDynamicConfig is empty.
	newConfig := DefaultDynamicConfig()

	_, err := NewDynamic(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_Dynamic_New_Error_Nodes(t *testing.T) {
	newConfig := DefaultDynamicConfig()
	// DefaultDynamicConfig is empty. When Nodes is given, Root must be given as
	// well.
	newConfig.Nodes = []spec.Strategy{testMaybeNewDynamic(t, "Sum", nil)}

	_, err := NewDynamic(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_Dynamic_SetNode(t *testing.T) {
	arg1 := testMaybeNewStatic(t, 3.1)
	arg2 := testMaybeNewStatic(t, 6.8)
	root := testMaybeNewDynamic(t, "Sum", nil)

	// Setting the first node should work.
	err := root.SetNode([]int{0}, arg1)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Setting the second node should work.
	err = root.SetNode([]int{1}, arg2)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Now that 2 nodes are set the strategy should be valid, because Sum
	// requires two inpit arguments.
	err = root.Validate()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Overwriting the first node with itself should work.
	err = root.SetNode([]int{0}, arg1)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = root.Validate()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Overwriting the first node with another node should work.
	err = root.SetNode([]int{0}, arg2)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = root.Validate()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Set yourself as a node should throw an error because of a circular
	// strategy.
	err = root.SetNode([]int{2}, root)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = root.Validate()
	if !IsInvalidStrategy(err) {
		t.Fatal("expected", nil, "got", err)
	}

	// Removing the faulty node should work.
	err = root.RemoveNode([]int{2})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = root.Validate()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// The third node was removed in the call above. There should be 2 nodes again.
	if len(root.(*dynamic).Nodes) != 2 {
		t.Fatal("expected", 2, "got", len(root.(*dynamic).Nodes))
	}

	// There are only two nodes. Setting a new node should require the index 2.
	// Index 7 is out of range.
	err = root.SetNode([]int{7}, arg1)
	if !IsIndexOutOfRange(err) {
		t.Fatal("expected", true, "got", false)
	}
}
