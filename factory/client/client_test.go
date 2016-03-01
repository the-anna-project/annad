package factoryclient

import (
	"testing"

	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/spec"
)

// Test_FactoryClient_NewImpulse checks that the factory client always creates
// proper impulses.
func Test_FactoryClient_NewImpulse(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	newClientConfig := DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := NewFactory(newClientConfig)

	var objectID spec.ObjectID
	go func() {
		listener := func(newSignal spec.Signal) (spec.Signal, error) {
			newConfig := impulse.DefaultConfig()
			newImpulse, err := impulse.NewImpulse(newConfig)
			if err != nil {
				t.Fatalf("impulse.NewImpulse returned err: %#v", err)
			}
			objectID = newImpulse.GetID()

			newSignal.SetOutput(newImpulse)

			return newSignal, nil
		}

		newFactoryGateway.Listen(listener, nil)
	}()

	impulse, err := newClient.NewImpulse()
	if err != nil {
		t.Fatalf("NewImpulse returned err: %#v", err)
	}

	if objectID != impulse.GetID() {
		t.Fatalf("Factory.NewImpulse returned wrong impulse")
	}
}

// Test_FactoryClient_GetType checks that the factory client returns its proper
// object type.
func Test_FactoryClient_GetType(t *testing.T) {
	newClient := NewFactory(DefaultConfig())

	object, ok := newClient.(spec.Object)
	if !ok {
		t.Fatalf("factory client does not implement spec.Object")
	}
	if object.GetType() != ObjectTypeFactoryClient {
		t.Fatalf("invalid object tyoe of factory client")
	}
}

// Test_FactoryClient_GetID checks that always independent factory clients are
// created.
func Test_FactoryClient_GetID(t *testing.T) {
	firstClient := NewFactory(DefaultConfig())
	firstObject, _ := firstClient.(spec.Object)
	secondClient := NewFactory(DefaultConfig())
	secondObject, _ := secondClient.(spec.Object)

	if firstObject.GetID() == secondObject.GetID() {
		t.Fatalf("IDs of factory clients are equal")
	}
}

// Test_FactoryClient_Shutdown_Single checks that the factory client properly
// shuts down.
func Test_FactoryClient_Shutdown_Single(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	newClientConfig := DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := NewFactory(newClientConfig)

	var objectID spec.ObjectID
	go func() {
		listener := func(newSignal spec.Signal) (spec.Signal, error) {
			newConfig := impulse.DefaultConfig()
			newImpulse, err := impulse.NewImpulse(newConfig)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			objectID = newImpulse.GetID()

			newSignal.SetOutput(newImpulse)

			return newSignal, nil
		}

		newFactoryGateway.Listen(listener, nil)
	}()

	newClient.Shutdown()

	_, err := newClient.NewImpulse()
	if !gateway.IsGatewayClosed(err) {
		t.Fatalf("NewImpulse did NOT return proper error")
	}

	if objectID != "" {
		t.Fatalf("factory client did not close factory gateway properly")
	}
}

// Test_FactoryClient_Shutdown_Multiple checks that shutting down multiple
// times causes no problems.
func Test_FactoryClient_Shutdown_Multiple(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	newClientConfig := DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := NewFactory(newClientConfig)

	var objectID spec.ObjectID
	go func() {
		listener := func(newSignal spec.Signal) (spec.Signal, error) {
			newConfig := impulse.DefaultConfig()
			newImpulse, err := impulse.NewImpulse(newConfig)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			objectID = newImpulse.GetID()

			newSignal.SetOutput(newImpulse)

			return newSignal, nil
		}

		newFactoryGateway.Listen(listener, nil)
	}()

	newClient.Shutdown()
	newClient.Shutdown()
	newClient.Shutdown()

	_, err := newClient.NewImpulse()
	if !gateway.IsGatewayClosed(err) {
		t.Fatalf("NewImpulse did NOT return proper error")
	}

	if objectID != "" {
		t.Fatalf("factory client did not close factory gateway properly")
	}
}
