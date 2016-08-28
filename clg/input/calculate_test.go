package input

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/context"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

func Test_CLG_Input_KnownInputSequence(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newStorage := memory.MustNew()

	// Create record for the test input.
	informationID := "123"
	input := "test input"
	informationIDKey := key.NewCLGKey(newCLG, "input-sequence:information-id:%s", input)
	err := newStorage.Set(informationIDKey, informationID)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Create record for the clg tree ID.
	clgTreeID := "456"
	clgTreeIDKey := key.NewCLGKey(newCLG, "information-id:clg-tree-id:%s", informationID)
	err = newStorage.Set(clgTreeIDKey, clgTreeID)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).Storage = newStorage

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

	// Check calculated payload. The input CLG only returns an error. This error
	// is filtered to be handled during the call to Calculate. Thus it is removed
	// from the calculated payload.
	args := calculatedNetworkPayload.GetArgs()
	if len(args) != 0 {
		t.Fatal("expected", 0, "got", len(args))
	}

	// Check if clg tree ID was set to the context.
	injectedCLGTreeID := newCtx.GetCLGTreeID()
	if clgTreeID != injectedCLGTreeID {
		t.Fatal("expected", clgTreeID, "got", injectedCLGTreeID)
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
	newCLG.(*clg).Storage = newStorage

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

	// Check calculated payload. The input CLG only returns an error. This error
	// is filtered to be handled during the call to Calculate. Thus it is removed
	// from the calculated payload.
	args := calculatedNetworkPayload.GetArgs()
	if len(args) != 0 {
		t.Fatal("expected", 0, "got", len(args))
	}

	// Check if clg tree ID was set to the context.
	injectedCLGTreeID := newCtx.GetCLGTreeID()
	if injectedCLGTreeID != "" {
		t.Fatal("expected", "", "got", injectedCLGTreeID)
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
	informationIDKey := key.NewCLGKey(newCLG, "input-sequence:information-id:%s", input)
	err := newStorage.Set(informationIDKey, informationID)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).Storage = newStorage

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

	// Check calculated payload. The input CLG only returns an error. This error
	// is filtered to be handled during the call to Calculate. Thus it is removed
	// from the calculated payload.
	args := calculatedNetworkPayload.GetArgs()
	if len(args) != 0 {
		t.Fatal("expected", 0, "got", len(args))
	}

	// Check if clg tree ID was set to the context.
	injectedCLGTreeID := newCtx.GetCLGTreeID()
	if injectedCLGTreeID != "" {
		t.Fatal("expected", "", "got", injectedCLGTreeID)
	}
}
