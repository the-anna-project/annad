package strategy

import (
	"fmt"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewStrategy(t *testing.T) spec.Strategy {
	newConfig := DefaultConfig()
	newConfig.Root = spec.CLG("Sum")
	newStrategy, err := New(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStrategy
}

func testMaybeNewStrategyWithRoot(t *testing.T, root spec.CLG) spec.Strategy {
	newConfig := DefaultConfig()
	newConfig.Root = root
	newStrategy, err := New(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStrategy
}

func testMaybeNewStrategyWithArgument(t *testing.T, argument interface{}) spec.Strategy {
	newConfig := DefaultConfig()
	newConfig.Argument = argument
	newStrategy, err := New(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStrategy
}

func Test_Strategy_New(t *testing.T) {
	// When this does not panic, the test is successfull.
	testMaybeNewStrategy(t)
}

func Test_Strategy_New_RootAndNodes(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Root = "Sum"
	newConfig.Nodes = []spec.Strategy{testMaybeNewStrategy(t)}

	_, err := New(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Strategy_New_Error_DefaultConfig(t *testing.T) {
	// DefaultConfig is empty.
	newConfig := DefaultConfig()

	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_New_Error_Argument(t *testing.T) {
	newConfig := DefaultConfig()
	// DefaultConfig is empty. Argument must not be given when Nodes is given.
	newConfig.Argument = 3
	newConfig.Nodes = []spec.Strategy{testMaybeNewStrategy(t)}

	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_New_Error_Root(t *testing.T) {
	newConfig := DefaultConfig()
	// DefaultConfig is empty. Argument must not be given when Root is given.
	newConfig.Argument = 3
	newConfig.Root = "Sum"

	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_New_Error_Nodes(t *testing.T) {
	newConfig := DefaultConfig()
	// DefaultConfig is empty. When Nodes is given, Root must be given as well.
	newConfig.Nodes = []spec.Strategy{testMaybeNewStrategy(t)}

	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_New_Error_ArgumentAndRootAndNodes(t *testing.T) {
	newConfig := DefaultConfig()
	// DefaultConfig is empty. Argument must not be given when Root and Nodes are given.
	newConfig.Argument = 3
	newConfig.Root = "Sum"
	newConfig.Nodes = []spec.Strategy{testMaybeNewStrategy(t)}

	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Strategy_SetNode(t *testing.T) {
	arg1 := testMaybeNewStrategyWithArgument(t, 3.1)
	arg2 := testMaybeNewStrategyWithArgument(t, 6.8)
	root := testMaybeNewStrategyWithRoot(t, "Sum")

	// Setting the first node should work.
	err := root.SetNode([]int{0}, arg1)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Setting the second node should work.
	err = root.SetNode([]int{1}, arg2)
	fmt.Printf("%#v\n", err)
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

	// Setting the 3. node failed in the call above. There should still only be 2
	// nodes.
	nodes := root.GetNodes()
	if len(nodes) != 2 {
		t.Fatal("expected", 2, "got", len(nodes))
	}

	// There are only two nodes. Setting a new node should require the index 2.
	// Index 7 is out of range.
	err = root.SetNode([]int{7}, arg1)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = root.Validate()
	if !IsInvalidStrategy(err) {
		t.Fatal("expected", nil, "got", err)
	}
}
