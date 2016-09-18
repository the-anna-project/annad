package input

import (
	"reflect"
	"testing"

	redigo "github.com/garyburd/redigo/redis"
	"github.com/rafaeljusto/redigomock"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/context"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
	"github.com/xh3b4sd/anna/storage/redis"
)

type testErrorIDFactory struct{}

// New is only a test implementation of spec.IDFactory to do nothing but
// returning some error we can check against.
func (f *testErrorIDFactory) New() (spec.ObjectID, error) {
	return "", maskAny(invalidConfigError)
}

func (f *testErrorIDFactory) WithType(idType spec.IDType) (spec.ObjectID, error) {
	return "", nil
}

func testMustNewErrorIDFactory(t *testing.T) spec.IDFactory {
	return &testErrorIDFactory{}
}

type testIDFactory struct{}

// New is only a test implementation of spec.IDFactory to do nothing but
// returning some error we can check against.
func (f *testIDFactory) New() (spec.ObjectID, error) {
	return "new-ID", nil
}

func (f *testIDFactory) WithType(idType spec.IDType) (spec.ObjectID, error) {
	return "", nil
}

func testMustNewIDFactory(t *testing.T) spec.IDFactory {
	return &testIDFactory{}
}

func testMustNewStorageWithConn(t *testing.T, c redigo.Conn) spec.Storage {
	newStorage, err := redis.NewStorage(redis.DefaultStorageConfigWithConn(c))
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStorage
}

func testMustNewNetworkPayload(t *testing.T, ctx spec.Context, input string) spec.NetworkPayload {
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(input)}
	newNetworkPayloadConfig.Destination = "destination"
	newNetworkPayloadConfig.Sources = []spec.ObjectID{"source"}
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newNetworkPayload
}

func Test_CLG_Input_KnownInputSequence(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newStorage := memory.MustNew()

	// Create record for the test input.
	informationID := "123"
	newInput := "test input"
	informationIDKey := key.NewCLGKey("information-sequence:%s:information-id", newInput)
	err := newStorage.Set(informationIDKey, informationID)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).GeneralStorage = newStorage

	// Execute CLG.
	calculatedNetworkPayload, err := newCLG.Calculate(testMustNewNetworkPayload(t, newCtx, newInput))
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Check the calculated payload. The interface of the input CLG only returns
	// an error. This error is filtered to be handled during the call to
	// Calculate. Thus it is removed from the calculated payload. Anyway there is
	// the original context be obtained as first argument within the network
	// payload.
	args := calculatedNetworkPayload.GetArgs()
	if len(args) != 1 {
		t.Fatal("expected", 1, "got", len(args))
	}

	// Check if the information ID was set to the context.
	injectedInformationID := newCtx.GetInformationID()
	if informationID != injectedInformationID {
		t.Fatal("expected", informationID, "got", injectedInformationID)
	}
}

func Test_CLG_Input_UnknownInputSequence(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newStorage := memory.MustNew()

	// Note we do not create a record for the test input. This test is about an
	// unknown input sequence.
	newInput := "test input"

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).GeneralStorage = newStorage

	// Execute CLG.
	calculatedNetworkPayload, err := newCLG.Calculate(testMustNewNetworkPayload(t, newCtx, newInput))
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Check the calculated payload. The interface of the input CLG only returns
	// an error. This error is filtered to be handled during the call to
	// Calculate. Thus it is removed from the calculated payload. Anyway there is
	// the original context be obtained as first argument within the network
	// payload.
	args := calculatedNetworkPayload.GetArgs()
	if len(args) != 1 {
		t.Fatal("expected", 1, "got", len(args))
	}

	// Check if the information ID was set to the context.
	injectedInformationID := newCtx.GetInformationID()
	if injectedInformationID == "" {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_CLG_Input_DataProperlyStored(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newStorage := memory.MustNew()
	newIDFactory := testMustNewIDFactory(t)

	// Note we do not create a record for the test input. This test is about an
	// unknown input sequence.
	newInput := "test input"
	// Our test ID factory always returns the same ID. That way we are able to
	// check for the ID being used during the test.
	newID, err := newIDFactory.New()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).GeneralStorage = newStorage
	newCLG.(*clg).IDFactory = newIDFactory

	// Execute CLG.
	calculatedNetworkPayload, err := newCLG.Calculate(testMustNewNetworkPayload(t, newCtx, newInput))
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Check the calculated payload. The interface of the input CLG only returns
	// an error. This error is filtered to be handled during the call to
	// Calculate. Thus it is removed from the calculated payload. Anyway there is
	// the original context be obtained as first argument within the network
	// payload.
	args := calculatedNetworkPayload.GetArgs()
	if len(args) != 1 {
		t.Fatal("expected", 1, "got", len(args))
	}

	informationIDKey := key.NewCLGKey("information-sequence:%s:information-id", newInput)
	storedID, err := newStorage.Get(informationIDKey)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if storedID != string(newID) {
		t.Fatal("expected", newID, "got", storedID)
	}

	informationSequenceKey := key.NewCLGKey("information-id:%s:information-sequence", newID)
	storedInput, err := newStorage.Get(informationSequenceKey)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if storedInput != newInput {
		t.Fatal("expected", newInput, "got", storedInput)
	}
}

func Test_CLG_Input_IDFactoryError(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newStorage := memory.MustNew()
	newIDFactory := testMustNewErrorIDFactory(t)

	// Note we do not create a record for the test input. This test is about an
	// unknown input sequence.
	newInput := "test input"

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).IDFactory = newIDFactory
	newCLG.(*clg).GeneralStorage = newStorage

	// Execute CLG.
	_, err := newCLG.Calculate(testMustNewNetworkPayload(t, newCtx, newInput))
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_Input_SetInformationIDError(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newIDFactory := testMustNewIDFactory(t)

	// Prepare the storage connection to fake a returned error.
	newInput := "test input"
	informationIDKey := key.NewCLGKey("information-sequence:%s:information-id", newInput)
	// Our test ID factory always returns the same ID. That way we are able to
	// check for the ID being used during the test.
	newID, err := newIDFactory.New()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	c := redigomock.NewConn()
	c.Command("GET", "prefix:"+informationIDKey).ExpectError(redigo.ErrNil)
	c.Command("SET", "prefix:"+informationIDKey, string(newID)).ExpectError(invalidConfigError)
	newStorage := testMustNewStorageWithConn(t, c)

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).GeneralStorage = newStorage
	newCLG.(*clg).IDFactory = newIDFactory

	// Execute CLG.
	_, err = newCLG.Calculate(testMustNewNetworkPayload(t, newCtx, newInput))
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_Input_SetInformationSequenceError(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newIDFactory := testMustNewIDFactory(t)

	// Prepare the storage connection to fake a returned error.
	newInput := "test input"
	informationIDKey := key.NewCLGKey("information-sequence:%s:information-id", newInput)
	// Our test ID factory always returns the same ID. That way we are able to
	// check for the ID being used during the test.
	newID, err := newIDFactory.New()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	informationSequenceKey := key.NewCLGKey("information-id:%s:information-sequence", newID)

	c := redigomock.NewConn()
	c.Command("GET", "prefix:"+informationIDKey).ExpectError(redigo.ErrNil)
	c.Command("SET", "prefix:"+informationIDKey, string(newID)).Expect("OK")
	c.Command("SET", "prefix:"+informationSequenceKey, newInput).ExpectError(invalidConfigError)
	newStorage := testMustNewStorageWithConn(t, c)

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).GeneralStorage = newStorage
	newCLG.(*clg).IDFactory = newIDFactory

	// Execute CLG.
	_, err = newCLG.Calculate(testMustNewNetworkPayload(t, newCtx, newInput))
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_Input_GetInformationIDError(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()

	newInput := "test input"
	informationIDKey := key.NewCLGKey("information-sequence:%s:information-id", newInput)

	// Prepare the storage connection to fake a returned error.
	c := redigomock.NewConn()
	c.Command("GET", "prefix:"+informationIDKey).ExpectError(invalidConfigError)
	newStorage := testMustNewStorageWithConn(t, c)

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).GeneralStorage = newStorage

	// Execute CLG.
	_, err := newCLG.Calculate(testMustNewNetworkPayload(t, newCtx, newInput))
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}
