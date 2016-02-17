package factoryclient_test

import (
	"testing"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/core"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/spec"
)

// Test_Factory_001 checks that the factory client always creates independent
// cores.
func Test_Factory_001(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway()

	newClientConfig := factoryclient.DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := factoryclient.NewFactory(newClientConfig)

	var objectID spec.ObjectID
	go func() {
		newSignal, err := newFactoryGateway.ReceiveSignal()
		if err != nil {
			t.Fatalf("Gateway.ReceiveSignal returned err: %#v", err)
		}
		responder, err := newSignal.GetResponder()
		if err != nil {
			t.Fatalf("Gateway.GetResponder returned err: %#v", err)
		}

		newConfig := core.DefaultConfig()
		newCore := core.NewCore(newConfig)
		objectID = newCore.GetObjectID()
		newSignal.SetObject("response", newCore)

		responder <- newSignal
	}()

	core, err := newClient.NewCore()
	if err != nil {
		t.Fatalf("FactoryClient.NewCore returned err: %#v", err)
	}

	if objectID != core.GetObjectID() {
		t.Fatalf("Factory.NewCore returned wrong core")
	}
}

// Test_Factory_002 checks that the factory client always creates independent
// impulses.
func Test_Factory_002(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway()

	newClientConfig := factoryclient.DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := factoryclient.NewFactory(newClientConfig)

	var objectID spec.ObjectID
	go func() {
		newSignal, err := newFactoryGateway.ReceiveSignal()
		if err != nil {
			t.Fatalf("Gateway.ReceiveSignal returned err: %#v", err)
		}
		responder, err := newSignal.GetResponder()
		if err != nil {
			t.Fatalf("Gateway.GetResponder returned err: %#v", err)
		}

		newConfig := impulse.DefaultConfig()
		newImpulse := impulse.NewImpulse(newConfig)
		objectID = newImpulse.GetObjectID()
		newSignal.SetObject("response", newImpulse)

		responder <- newSignal
	}()

	impulse, err := newClient.NewImpulse()
	if err != nil {
		t.Fatalf("FactoryClient.NewImpulse returned err: %#v", err)
	}

	if objectID != impulse.GetObjectID() {
		t.Fatalf("Factory.NewImpulse returned wrong impulse")
	}
}

// Test_Factory_003 checks that the factory client returns its proper object
// type.
func Test_Factory_003(t *testing.T) {
	newClient := factoryclient.NewFactory(factoryclient.DefaultConfig())

	object, ok := newClient.(spec.Object)
	if !ok {
		t.Fatalf("factory client does not implement spec.Object")
	}
	if object.GetObjectType() != common.ObjectType.FactoryClient {
		t.Fatalf("invalid object tyoe of factory client")
	}
}
