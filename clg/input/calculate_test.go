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

type testIDFactory struct{}

// New is only a test implementation of spec.IDFactory to do nothing but
// returning some error we can check against.
func (f *testIDFactory) New() (spec.ObjectID, error) {
	return "", maskAny(invalidConfigError)
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

func Test_CLG_Input_KnownInputSequence(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newStorage := memory.MustNew()

	// Create record for the test input.
	informationID := "123"
	input := "test input"
	informationIDKey := key.NewCLGKey("information-sequence:%s:information-id", input)
	err := newStorage.Set(informationIDKey, informationID)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).GeneralStorage = newStorage

	// Execute CLG.
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(newCtx), reflect.ValueOf(input)}
	newNetworkPayloadConfig.Destination = "destination"
	newNetworkPayloadConfig.Sources = []spec.ObjectID{"source"}
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	calculatedNetworkPayload, err := newCLG.Calculate(newNetworkPayload)
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
	input := "test input"

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).GeneralStorage = newStorage

	// Execute CLG.
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(newCtx), reflect.ValueOf(input)}
	newNetworkPayloadConfig.Destination = "destination"
	newNetworkPayloadConfig.Sources = []spec.ObjectID{"source"}
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	calculatedNetworkPayload, err := newCLG.Calculate(newNetworkPayload)
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

func Test_CLG_Input_IDFactoryError(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newStorage := memory.MustNew()
	newIDFactory := testMustNewIDFactory(t)

	// Note we do not create a record for the test input. This test is about an
	// unknown input sequence.
	input := "test input"

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).IDFactory = newIDFactory
	newCLG.(*clg).GeneralStorage = newStorage

	// Execute CLG.
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(newCtx), reflect.ValueOf(input)}
	newNetworkPayloadConfig.Destination = "destination"
	newNetworkPayloadConfig.Sources = []spec.ObjectID{"source"}
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	_, err = newCLG.Calculate(newNetworkPayload)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_CLG_Input_SetInformationIDError(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()

	// Prepare the storage connection to fake a returned error.
	input := "test input"
	informationIDKey := key.NewCLGKey("information-sequence:%s:information-id", input)

	c := redigomock.NewConn()
	c.Command("GET", "prefix:"+informationIDKey).ExpectError(redigo.ErrNil)
	c.Command("SET").ExpectError(invalidConfigError)
	newStorage := testMustNewStorageWithConn(t, c)

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).GeneralStorage = newStorage

	// Execute CLG.
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(newCtx), reflect.ValueOf(input)}
	newNetworkPayloadConfig.Destination = "destination"
	newNetworkPayloadConfig.Sources = []spec.ObjectID{"source"}
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	_, err = newCLG.Calculate(newNetworkPayload)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_CLG_Input_UnknownCLGTree(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newStorage := memory.MustNew()

	// Note we do not create a record for the clg tree ID. This test is about an
	// unknown clg tree.
	informationID := "123"
	input := "test input"
	informationIDKey := key.NewCLGKey("information-sequence:%s:information-id", input)
	err := newStorage.Set(informationIDKey, informationID)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).GeneralStorage = newStorage

	// Execute CLG.
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(newCtx), reflect.ValueOf(input)}
	newNetworkPayloadConfig.Destination = "destination"
	newNetworkPayloadConfig.Sources = []spec.ObjectID{"source"}
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	calculatedNetworkPayload, err := newCLG.Calculate(newNetworkPayload)
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

	// Check if clg tree ID was set to the context.
	injectedCLGTreeID := newCtx.GetCLGTreeID()
	if injectedCLGTreeID != "" {
		t.Fatal("expected", "", "got", injectedCLGTreeID)
	}
}

func Test_CLG_Input_GetInformationIDError(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()

	input := "test input"
	informationIDKey := key.NewCLGKey("information-sequence:%s:information-id", input)

	// Prepare the storage connection to fake a returned error.
	c := redigomock.NewConn()
	c.Command("GET", "prefix:"+informationIDKey).ExpectError(invalidConfigError)
	newStorage := testMustNewStorageWithConn(t, c)

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).GeneralStorage = newStorage

	// Execute CLG.
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(newCtx), reflect.ValueOf(input)}
	newNetworkPayloadConfig.Destination = "destination"
	newNetworkPayloadConfig.Sources = []spec.ObjectID{"source"}
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	_, err = newCLG.Calculate(newNetworkPayload)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}
