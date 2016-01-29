package state_test

import (
	"testing"

	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

// Test_State_001 checks that reading and writing bytes to a state works as
// expected.
func Test_State_001(t *testing.T) {
	newState := state.NewState(state.DefaultConfig())

	_, err := newState.GetBytes("foo")
	if !state.IsBytesNotFound(err) {
		t.Fatalf("State.GetBytes did NOT return proper error")
	}

	newState.SetBytes("foo", []byte("bar"))

	bytes, err := newState.GetBytes("foo")
	if err != nil {
		t.Fatalf("State.GetBytes did return error: %#v", err)
	}
	if string(bytes) != "bar" {
		t.Fatalf("State.GetBytes did return wrong result: %s", bytes)
	}
}

// Test_State_002 checks that object IDs of different states are NOT equal.
// original state.
func Test_State_002(t *testing.T) {
	firstState := state.NewState(state.DefaultConfig())
	secondState := state.NewState(state.DefaultConfig())

	if firstState.GetObjectID() == secondState.GetObjectID() {
		t.Fatalf("object ID of first state and second state is equal")
	}
}

// Test_State_003 checks that the predefined object ID of a state is properly
// set.
func Test_State_003(t *testing.T) {
	objectID := spec.ObjectID("test-id")
	newConfig := state.DefaultConfig()
	newConfig.ObjectID = objectID
	newState := state.NewState(newConfig)

	if objectID != newState.GetObjectID() {
		t.Fatalf("predefined object ID not properly set")
	}
}
