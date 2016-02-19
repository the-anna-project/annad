package state_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/factory/server"
	"github.com/xh3b4sd/anna/file-system/fake"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

func newFactoryAndFileSystem() (spec.Factory, spec.FileSystem) {
	// Create new factory server with a fake file system. This is used to create
	// new objects.
	fileSystemFake := filesystemfake.NewFileSystem()
	newFactoryServerConfig := factoryserver.DefaultConfig()
	newFactoryServerConfig.FileSystem = fileSystemFake
	newFactoryServer := factoryserver.NewFactory(newFactoryServerConfig)

	return newFactoryServer, fileSystemFake
}

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

// Test_State_003 checks that the predefined object type of a state is properly
// set.
func Test_State_004(t *testing.T) {
	objectType := spec.ObjectType("test-type")
	newConfig := state.DefaultConfig()
	newConfig.ObjectType = objectType
	newState := state.NewState(newConfig)

	if objectType != newState.GetObjectType() {
		t.Fatalf("predefined object type not properly set")
	}
}

// Test_State_005 checks that backing up state and restoring it works as
// expected. This test case is quite big, because the setup itself needs some
// code to be done. The test creates a object tree that is written to the fake
// file system implementation. Reading the state from the file system should
// result in the exactly same state we dumped before.
func Test_State_005(t *testing.T) {
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

	// Create a bunch of new objects. The first core here is the root of the
	// object tree.
	newCore, err := newFactoryServer.NewCore()
	if err != nil {
		t.Fatalf("FactoryServer.NewCore returned err: %#v", err)
	}
	newCore.SetState(newTestState(spec.ObjectID("core-id"), common.ObjectType.Core))
	anotherCore, err := newFactoryServer.NewCore()
	if err != nil {
		t.Fatalf("FactoryServer.NewCore returned err: %#v", err)
	}
	anotherCore.SetState(newTestState(spec.ObjectID("another-core-id"), common.ObjectType.Core))
	newImpulse, err := newFactoryServer.NewImpulse()
	if err != nil {
		t.Fatalf("FactoryServer.NewImpulse returned err: %#v", err)
	}
	newImpulse.SetState(newTestState(spec.ObjectID("impulse-id"), common.ObjectType.Impulse))
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
	newCore.GetState().SetCore(anotherCore)
	newCore.GetState().SetNetwork(newNetwork)
	anotherCore.GetState().SetImpulse(newImpulse)
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
	restoredCore, err := newState.GetCoreByID(spec.ObjectID("another-core-id"))
	if err != nil {
		t.Fatalf("State.GetCoreByID returned err: %#v", err)
	}
	_, err = restoredCore.GetState().GetImpulseByID(spec.ObjectID("impulse-id"))
	if err != nil {
		t.Fatalf("State.GetImpulseByID returned err: %#v", err)
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

// Test_State_006 checks that the version is proper set within the state.
func Test_State_006(t *testing.T) {
	newFactoryServer, fileSystemFake := newFactoryAndFileSystem()

	// Create new state.
	newCore, err := newFactoryServer.NewCore()
	if err != nil {
		t.Fatalf("FactoryServer.NewCore returned err: %#v", err)
	}

	// Check that state version is not set.
	if newCore.GetState().GetVersion() != "" {
		t.Fatalf("new state should not contain any version")
	}

	// Set version to verify at the end.
	newCore.GetState().SetVersion("test-version")

	// Write state to file system.
	bytes, err := json.Marshal(newCore.GetState())
	if err != nil {
		t.Fatalf("json.Marshal returned err: %#v", err)
	}
	err = fileSystemFake.WriteFile(common.DefaultStateFile, bytes, os.FileMode(0660))
	if err != nil {
		t.Fatalf("FileSystem.WriteFile returned err: %#v", err)
	}

	// Read state from file system.
	err = newCore.GetState().Read()
	if err != nil {
		t.Fatalf("State.Read returned err: %#v", err)
	}

	// Check for proper version.
	if newCore.GetState().GetVersion() != "test-version" {
		t.Fatalf("read state does not contain proper version")
	}
}

// Test_State_007 checks if getting the state's age works as expected.
func Test_State_007(t *testing.T) {
	newFactoryServer, _ := newFactoryAndFileSystem()

	// Create new state.
	newCore, err := newFactoryServer.NewCore()
	if err != nil {
		t.Fatalf("FactoryServer.NewCore returned err: %#v", err)
	}

	// Check if a non-nil duration is returned as age.
	age := newCore.GetState().GetAge()
	if age.Nanoseconds() == 0 {
		t.Fatalf("State.GetAge returned invalid duration")
	}
}

// Test_State_008 checks that setting and getting bytes works as expected.
func Test_State_008(t *testing.T) {
	newFactoryServer, _ := newFactoryAndFileSystem()

	// Create new state.
	newCore, err := newFactoryServer.NewCore()
	if err != nil {
		t.Fatalf("FactoryServer.NewCore returned err: %#v", err)
	}

	_, err = newCore.GetState().GetBytes("foo")
	if !state.IsBytesNotFound(err) {
		t.Fatalf("State.GetBytes did NOT return err")
	}
	newCore.GetState().SetBytes("foo", []byte("bar"))
	bytes, err := newCore.GetState().GetBytes("foo")
	if err != nil {
		t.Fatalf("State.GetBytes returned err: %#v", err)
	}
	if string(bytes) != "bar" {
		t.Fatalf("State.GetBytes returned invalid bytes")
	}
}

// Test_State_009 checks that setting and getting core works as expected.
func Test_State_009(t *testing.T) {
	newFactoryServer, _ := newFactoryAndFileSystem()

	// Create new state.
	newCore, err := newFactoryServer.NewCore()
	if err != nil {
		t.Fatalf("FactoryServer.NewCore returned err: %#v", err)
	}

	testCore, err := newFactoryServer.NewCore()
	if err != nil {
		t.Fatalf("FactoryServer.NewCore returned err: %#v", err)
	}
	_, err = newCore.GetState().GetCoreByID(testCore.GetObjectID())
	if !state.IsCoreNotFound(err) {
		t.Fatalf("State.GetCoreByID did NOT return err")
	}
	cores := newCore.GetState().GetCores()
	if len(cores) != 0 {
		t.Fatalf("State.GetCores should not return any core")
	}
	newCore.GetState().SetCore(testCore)
	core, err := newCore.GetState().GetCoreByID(testCore.GetObjectID())
	if err != nil {
		t.Fatalf("State.GetBytes returned err: %#v", err)
	}
	if core.GetObjectID() != testCore.GetObjectID() {
		t.Fatalf("State.GetCoreByID returned invalid core")
	}
	cores = newCore.GetState().GetCores()
	if len(cores) != 1 {
		t.Fatalf("State.GetCores should return one core")
	}
}

// Test_State_010 checks that setting and getting impulse works as expected.
func Test_State_010(t *testing.T) {
	newFactoryServer, _ := newFactoryAndFileSystem()

	// Create new state.
	newImpulse, err := newFactoryServer.NewImpulse()
	if err != nil {
		t.Fatalf("FactoryServer.NewImpulse returned err: %#v", err)
	}

	testImpulse, err := newFactoryServer.NewImpulse()
	if err != nil {
		t.Fatalf("FactoryServer.NewImpulse returned err: %#v", err)
	}
	_, err = newImpulse.GetState().GetImpulseByID(testImpulse.GetObjectID())
	if !state.IsImpulseNotFound(err) {
		t.Fatalf("State.GetImpulseByID did NOT return err")
	}
	impulses := newImpulse.GetState().GetImpulses()
	if len(impulses) != 0 {
		t.Fatalf("State.GetImpulses should not return any impulse")
	}
	newImpulse.GetState().SetImpulse(testImpulse)
	impulse, err := newImpulse.GetState().GetImpulseByID(testImpulse.GetObjectID())
	if err != nil {
		t.Fatalf("State.GetBytes returned err: %#v", err)
	}
	if impulse.GetObjectID() != testImpulse.GetObjectID() {
		t.Fatalf("State.GetImpulseByID returned invalid impulse")
	}
	impulses = newImpulse.GetState().GetImpulses()
	if len(impulses) != 1 {
		t.Fatalf("State.GetImpulses should return one impulse")
	}
}

// Test_State_011 checks that a fresh state does not contain any network.
func Test_State_011(t *testing.T) {
	newConfig := state.DefaultConfig()
	newState := state.NewState(newConfig)

	networks := newState.GetNetworks()
	if len(networks) != 0 {
		t.Fatalf("expected fresh state to not contain any network")
	}
}

// Test_State_012 checks that a fresh state returns a proper not found error
// when fetching some network.
func Test_State_012(t *testing.T) {
	newConfig := state.DefaultConfig()
	newState := state.NewState(newConfig)

	_, err := newState.GetNetworkByID(spec.ObjectID("network-id"))
	if !state.IsNetworkNotFound(err) {
		t.Fatalf("State.GetNetworks did NOT return proper err")
	}
}

// Test_State_013 checks that a fresh state does not contain any neuron.
func Test_State_013(t *testing.T) {
	newConfig := state.DefaultConfig()
	newState := state.NewState(newConfig)

	neurons := newState.GetNeurons()
	if len(neurons) != 0 {
		t.Fatalf("expected fresh state to not contain any neuron")
	}
}

// Test_State_014 checks that a fresh state returns a proper not found error
// when fetching some neuron.
func Test_State_014(t *testing.T) {
	newConfig := state.DefaultConfig()
	newState := state.NewState(newConfig)

	_, err := newState.GetNeuronByID(spec.ObjectID("neuron-id"))
	if !state.IsNeuronNotFound(err) {
		t.Fatalf("State.GetNeurons did NOT return proper err")
	}
}

// Test_State_015 checks that restoring state using an invalid state reader
// type throws a proper error.
func Test_State_015(t *testing.T) {
	newConfig := state.DefaultConfig()
	newConfig.StateReader = spec.StateType("invalid")
	newState := state.NewState(newConfig)

	err := newState.Read()
	if !state.IsInvalidStateReader(err) {
		t.Fatalf("State.Read did NOT return proper err")
	}
}

// Test_State_016 checks that restoring state using "none" state reader does
// nothing.
func Test_State_016(t *testing.T) {
	newConfig := state.DefaultConfig()
	newConfig.StateReader = common.StateType.NoneReader
	newState := state.NewState(newConfig)

	err := newState.Read()
	if err != nil {
		t.Fatalf("State.Read did return err %#v", err)
	}
}

// Test_State_017 checks that restoring state using an invalid state writer
// type throws a proper error.
func Test_State_017(t *testing.T) {
	newConfig := state.DefaultConfig()
	newConfig.StateWriter = spec.StateType("invalid")
	newState := state.NewState(newConfig)

	err := newState.Write()
	if !state.IsInvalidStateWriter(err) {
		t.Fatalf("State.Read did NOT return proper err")
	}
}

// Test_State_020 checks that restoring state using "none" state writer does
// nothing.
func Test_State_020(t *testing.T) {
	newConfig := state.DefaultConfig()
	newConfig.StateWriter = common.StateType.NoneWriter
	newState := state.NewState(newConfig)

	err := newState.Write()
	if err != nil {
		t.Fatalf("State.Read did return err %#v", err)
	}
}
