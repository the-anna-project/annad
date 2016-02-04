package state_test

import (
	"testing"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/factory/server"
	"github.com/xh3b4sd/anna/file-system/fake"
	"github.com/xh3b4sd/anna/gateway"
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

// Test_State_004 checks that backing up state and restoring it works as
// expected. This test case is quite big, because the setup itself needs some
// code to be done. The test creates a object tree that is written to the fake
// file system implementation. Reading the state from the file system should
// result in the exactly same state we dumped before.
func Test_State_004(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway()

	// Create a new factory client and configure it with the test gateway.
	newFactoryClientConfig := factoryclient.DefaultConfig()
	newFactoryClientConfig.FactoryGateway = newFactoryGateway
	newFactoryClient := factoryclient.NewFactory(newFactoryClientConfig)

	// Create new factory server with a fake file system. This is used to create
	// new objects.
	fileSystemFake := filesystemfake.NewFileSystem()
	newFactoryServerConfig := factoryserver.DefaultConfig()
	newFactoryServerConfig.FileSystem = fileSystemFake
	newFactoryServerConfig.FactoryClient = newFactoryClient
	newFactoryServerConfig.FactoryGateway = newFactoryGateway
	newFactoryServer := factoryserver.NewFactory(newFactoryServerConfig)

	newTestState := func(objectID spec.ObjectID, objectType spec.ObjectType) spec.State {
		newStateConfig := state.DefaultConfig()
		newStateConfig.FactoryClient = newFactoryClient
		newStateConfig.FileSystem = fileSystemFake
		newStateConfig.ObjectID = objectID
		newStateConfig.ObjectType = objectType
		newState := state.NewState(newStateConfig)

		return newState
	}

	// Create a bunch of new objects.
	newCore, err := newFactoryServer.NewCore()
	if err != nil {
		t.Fatalf("FactoryServer.NewCore returned err: %#v", err)
	}
	newCore.SetState(newTestState(spec.ObjectID("core-id"), common.ObjectType.Core))
	newNetwork, err := newFactoryServer.NewNetwork()
	if err != nil {
		t.Fatalf("FactoryServer.NewNetwork returned err: %#v", err)
	}
	newNetwork.SetState(newTestState(spec.ObjectID("network-id"), common.ObjectType.Network))
	characterNeuron, err := newFactoryServer.NewCharacterNeuron()
	if err != nil {
		t.Fatalf("FactoryServer.NewCharacterNeuron returned err: %#v", err)
	}
	characterNeuron.SetState(newTestState(spec.ObjectID("character-id"), common.ObjectType.CharacterNeuron))
	firstNeuron, err := newFactoryServer.NewFirstNeuron()
	if err != nil {
		t.Fatalf("FactoryServer.NewFirstNeuron returned err: %#v", err)
	}
	firstNeuron.SetState(newTestState(spec.ObjectID("first-id"), common.ObjectType.FirstNeuron))
	jobNeuron, err := newFactoryServer.NewJobNeuron()
	if err != nil {
		t.Fatalf("FactoryServer.NewJobNeuron returned err: %#v", err)
	}
	jobNeuron.SetState(newTestState(spec.ObjectID("job-id"), common.ObjectType.JobNeuron))

	// Create a object tree by binding objects together.
	newCore.GetState().SetNetwork(newNetwork)
	newNetwork.GetState().SetNeuron(firstNeuron)
	firstNeuron.GetState().SetNeuron(jobNeuron)
	jobNeuron.GetState().SetNeuron(characterNeuron)

	// Set some well known state to some objects to test against it after
	newCore.GetState().SetBytes("test-bytes", []byte("core"))
	newNetwork.GetState().SetBytes("test-bytes", []byte("network"))
	firstNeuron.GetState().SetBytes("test-bytes", []byte("character"))

	// Write the state to the file system and load it again into a new state.
	err = newCore.GetState().Write()
	if err != nil {
		t.Fatalf("State.Write returned err: %#v", err)
	}
	newState, err := newFactoryServer.NewState(common.ObjectType.None)
	if err != nil {
		t.Fatalf("FactoryServer.NewState returned err: %#v", err)
	}
	err = newState.Read()
	if err != nil {
		t.Fatalf("State.Read returned err: %#v", err)
	}

	// Check if there are still all objects within the state.
	coreBytes, err := newState.GetBytes("test-bytes")
	if err != nil {
		t.Fatalf("State.GetBytes returned err: %#v", err)
	}
	if string(coreBytes) != "core" {
		t.Fatalf("core bytes not properly restored")
	}
	restoredNetwork, err := newState.GetNetworkByID(spec.ObjectID("network-id"))
	if err != nil {
		t.Fatalf("State.GetNetworkByID returned err: %#v", err)
	}
	networkBytes, err := restoredNetwork.GetState().GetBytes("test-bytes")
	if err != nil {
		t.Fatalf("State.GetBytes returned err: %#v", err)
	}
	if string(networkBytes) != "network" {
		t.Fatalf("network bytes not properly restored")
	}
	restoredNeuron, err := restoredNetwork.GetState().GetNeuronByID(spec.ObjectID("first-id"))
	if err != nil {
		t.Fatalf("State.GetNeuronByID returned err: %#v", err)
	}
	neuronBytes, err := restoredNeuron.GetState().GetBytes("test-bytes")
	if err != nil {
		t.Fatalf("State.GetBytes returned err: %#v", err)
	}
	if string(neuronBytes) != "character" {
		t.Fatalf("network bytes not properly restored")
	}
}
