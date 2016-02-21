package factoryclient

import (
	"testing"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/core"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/network/strategy"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

// Test_FactoryClient_001 checks that the factory client always creates proper
// cores.
func Test_FactoryClient_001(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	newClientConfig := DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := NewFactory(newClientConfig)

	var objectID spec.ObjectID
	go func() {
		listener := func(newSignal spec.Signal) (spec.Signal, error) {
			newConfig := core.DefaultConfig()
			newCore := core.NewCore(newConfig)
			objectID = newCore.GetID()

			newSignal.SetOutput(newCore)

			return newSignal, nil
		}

		newFactoryGateway.Listen(listener, nil)
	}()

	core, err := newClient.NewCore()
	if err != nil {
		t.Fatalf("NewCore returned err: %#v", err)
	}

	if objectID != core.GetID() {
		t.Fatalf("Factory.NewCore returned wrong core")
	}
}

// Test_FactoryClient_002 checks that the factory client always creates proper
// impulses.
func Test_FactoryClient_002(t *testing.T) {
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

// Test_FactoryClient_003 checks that the factory client always creates proper
// redis storages.
func Test_FactoryClient_003(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	newClientConfig := DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := NewFactory(newClientConfig)

	var objectID spec.ObjectID
	go func() {
		listener := func(newSignal spec.Signal) (spec.Signal, error) {
			newConfig := storage.DefaultRedisStorageConfig()
			newRedisStorage := storage.NewRedisStorage(newConfig)

			objectID = newRedisStorage.GetID()

			newSignal.SetOutput(newRedisStorage)

			return newSignal, nil
		}

		newFactoryGateway.Listen(listener, nil)
	}()

	newRedisStorage, err := newClient.NewRedisStorage()
	if err != nil {
		t.Fatalf("NewImpulse returned err: %#v", err)
	}

	if objectID != newRedisStorage.GetID() {
		t.Fatalf("Factory.NewRedisStorage returned wrong redis storage")
	}
}

// Test_FactoryClient_004 checks that the factory client always creates proper
// strategy networks.
func Test_FactoryClient_004(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	newClientConfig := DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := NewFactory(newClientConfig)

	var objectID spec.ObjectID
	go func() {
		listener := func(newSignal spec.Signal) (spec.Signal, error) {
			newConfig := strategynetwork.DefaultNetworkConfig()
			newStrategyNetwork, err := strategynetwork.NewNetwork(newConfig)
			if err != nil {
				t.Fatalf("impulse.NewStrategyNetwork returned err: %#v", err)
			}

			objectID = newStrategyNetwork.GetID()

			newSignal.SetOutput(newStrategyNetwork)

			return newSignal, nil
		}

		newFactoryGateway.Listen(listener, nil)
	}()

	newStrategyNetwork, err := newClient.NewStrategyNetwork()
	if err != nil {
		t.Fatalf("NewImpulse returned err: %#v", err)
	}

	if objectID != newStrategyNetwork.GetID() {
		t.Fatalf("Factory.NewRedisStorage returned wrong strategy network")
	}
}

// Test_FactoryClient_005 checks that the factory client returns its proper object
// type.
func Test_FactoryClient_005(t *testing.T) {
	newClient := NewFactory(DefaultConfig())

	object, ok := newClient.(spec.Object)
	if !ok {
		t.Fatalf("factory client does not implement spec.Object")
	}
	if object.GetType() != common.ObjectType.FactoryClient {
		t.Fatalf("invalid object tyoe of factory client")
	}
}

// Test_FactoryClient_006 checks that always independent factory clients are
// created.
func Test_FactoryClient_006(t *testing.T) {
	firstClient := NewFactory(DefaultConfig())
	firstObject, _ := firstClient.(spec.Object)
	secondClient := NewFactory(DefaultConfig())
	secondObject, _ := secondClient.(spec.Object)

	if firstObject.GetID() == secondObject.GetID() {
		t.Fatalf("IDs of factory clients are equal")
	}
}

// Test_FactoryClient_007 checks that the factory client properly shuts down.
func Test_FactoryClient_007(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	newClientConfig := DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := NewFactory(newClientConfig)

	var objectID spec.ObjectID
	go func() {
		listener := func(newSignal spec.Signal) (spec.Signal, error) {
			newConfig := core.DefaultConfig()
			newCore := core.NewCore(newConfig)
			objectID = newCore.GetID()

			newSignal.SetOutput(newCore)

			return newSignal, nil
		}

		newFactoryGateway.Listen(listener, nil)
	}()

	newClient.Shutdown()

	_, err := newClient.NewCore()
	if !gateway.IsGatewayClosed(err) {
		t.Fatalf("NewCore did NOT return proper error")
	}

	if objectID != "" {
		t.Fatalf("factory client did not close factory gateway properly")
	}
}

// Test_FactoryClient_008 checks that shutting down multiple times causes no problems.
func Test_FactoryClient_008(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	newClientConfig := DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := NewFactory(newClientConfig)

	var objectID spec.ObjectID
	go func() {
		listener := func(newSignal spec.Signal) (spec.Signal, error) {
			newConfig := core.DefaultConfig()
			newCore := core.NewCore(newConfig)
			objectID = newCore.GetID()

			newSignal.SetOutput(newCore)

			return newSignal, nil
		}

		newFactoryGateway.Listen(listener, nil)
	}()

	newClient.Shutdown()
	newClient.Shutdown()
	newClient.Shutdown()

	_, err := newClient.NewCore()
	if !gateway.IsGatewayClosed(err) {
		t.Fatalf("NewCore did NOT return proper error")
	}

	if objectID != "" {
		t.Fatalf("factory client did not close factory gateway properly")
	}
}
