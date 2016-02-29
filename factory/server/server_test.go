package factoryserver

import (
	"testing"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/spec"
)

// Test_FactoryServer_001 checks that the factory server always creates
// independent cores.
func Test_FactoryServer_001(t *testing.T) {
	newServer := NewFactory(DefaultConfig())

	firstCore, err := newServer.NewCore()
	if err != nil {
		t.Fatalf("NewCore returned err: %#v", err)
	}

	secondCore, err := newServer.NewCore()
	if err != nil {
		t.Fatalf("NewCore returned err: %#v", err)
	}

	if firstCore.GetID() == secondCore.GetID() {
		t.Fatalf("object ID of first core and second core is equal")
	}
}

// Test_FactoryServer_002 checks that the factory server always creates
// independent impulses.
func Test_FactoryServer_002(t *testing.T) {
	newServer := NewFactory(DefaultConfig())

	firstImpulse, err := newServer.NewImpulse()
	if err != nil {
		t.Fatalf("NewImpulse returned err: %#v", err)
	}

	secondImpulse, err := newServer.NewImpulse()
	if err != nil {
		t.Fatalf("NewImpulse returned err: %#v", err)
	}

	if firstImpulse.GetID() == secondImpulse.GetID() {
		t.Fatalf("object ID of first impulse and second impulse is equal")
	}
}

// Test_FactoryServer_003 checks that the factory client always creates
// independent cores.
func Test_FactoryServer_003(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	// Create a new factory server and configure it with the test gateway.
	newServerConfig := DefaultConfig()
	newServerConfig.FactoryGateway = newFactoryGateway
	newServer := NewFactory(newServerConfig)
	newServer.Boot()

	// Create a new factory client and configure it with the test gateway.
	newClientConfig := factoryclient.DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := factoryclient.NewFactory(newClientConfig)

	firstCore, err := newServer.NewCore()
	if err != nil {
		t.Fatalf("NewCore returned err: %#v", err)
	}

	secondCore, err := newClient.NewCore()
	if err != nil {
		t.Fatalf("Factory.NewCore returned err: %#v", err)
	}

	if firstCore.GetID() == secondCore.GetID() {
		t.Fatalf("object ID of first core and second core is equal")
	}

	newServer.Shutdown()
	newClient.Shutdown()
}

// Test_FactoryServer_004 checks that the factory client always creates
// independent impulses.
func Test_FactoryServer_004(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	// Create a new factory server and configure it with the test gateway.
	newServerConfig := DefaultConfig()
	newServerConfig.FactoryGateway = newFactoryGateway
	newServer := NewFactory(newServerConfig)
	newServer.Boot()

	// Create a new factory client and configure it with the test gateway.
	newClientConfig := factoryclient.DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := factoryclient.NewFactory(newClientConfig)

	firstImpulse, err := newServer.NewImpulse()
	if err != nil {
		t.Fatalf("NewImpulse returned err: %#v", err)
	}

	secondImpulse, err := newClient.NewImpulse()
	if err != nil {
		t.Fatalf("Factory.NewImpulse returned err: %#v", err)
	}

	if firstImpulse.GetID() == secondImpulse.GetID() {
		t.Fatalf("object ID of first core and second core is equal")
	}

	newServer.Shutdown()
	newClient.Shutdown()
}

// Test_FactoryServer_005 checks that the factory client always creates
// independent redis storages.
func Test_FactoryServer_005(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	// Create a new factory server and configure it with the test gateway.
	newServerConfig := DefaultConfig()
	newServerConfig.FactoryGateway = newFactoryGateway
	newServer := NewFactory(newServerConfig)
	newServer.Boot()

	// Create a new factory client and configure it with the test gateway.
	newClientConfig := factoryclient.DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := factoryclient.NewFactory(newClientConfig)

	firstRedisStorage, err := newServer.NewRedisStorage()
	if err != nil {
		t.Fatalf("NewRedisStorage returned err: %#v", err)
	}

	secondRedisStorage, err := newClient.NewRedisStorage()
	if err != nil {
		t.Fatalf("Factory.NewRedisStorage returned err: %#v", err)
	}

	if firstRedisStorage.GetID() == secondRedisStorage.GetID() {
		t.Fatalf("object ID of first core and second core is equal")
	}

	newServer.Shutdown()
	newClient.Shutdown()
}

// Test_FactoryServer_006 checks that the factory client always creates
// independent strategy networks.
func Test_FactoryServer_006(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	// Create a new factory server and configure it with the test gateway.
	newServerConfig := DefaultConfig()
	newServerConfig.FactoryGateway = newFactoryGateway
	newServer := NewFactory(newServerConfig)
	newServer.Boot()

	// Create a new factory client and configure it with the test gateway.
	newClientConfig := factoryclient.DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := factoryclient.NewFactory(newClientConfig)

	firstStrategyNetwork, err := newServer.NewStrategyNetwork()
	if err != nil {
		t.Fatalf("NewStrategyNetwork returned err: %#v", err)
	}

	secondStrategyNetwork, err := newClient.NewStrategyNetwork()
	if err != nil {
		t.Fatalf("Factory.NewStrategyNetwork returned err: %#v", err)
	}

	if firstStrategyNetwork.GetID() == secondStrategyNetwork.GetID() {
		t.Fatalf("object ID of first core and second core is equal")
	}

	newServer.Shutdown()
	newClient.Shutdown()
}

// Test_FactoryServer_007 checks that the factory server returns its proper
// object type.
func Test_FactoryServer_007(t *testing.T) {
	newServer := NewFactory(DefaultConfig())

	object, ok := newServer.(spec.Object)
	if !ok {
		t.Fatalf("factory server does not implement spec.Object")
	}
	if object.GetType() != ObjectTypeFactoryServer {
		t.Fatalf("invalid object tyoe of factory server")
	}
}

// Test_FactoryServer_008 checks that always independent factory clients are
// created.
func Test_FactoryClient_008(t *testing.T) {
	firstClient := NewFactory(DefaultConfig())
	firstObject, _ := firstClient.(spec.Object)
	secondClient := NewFactory(DefaultConfig())
	secondObject, _ := secondClient.(spec.Object)

	if firstObject.GetID() == secondObject.GetID() {
		t.Fatalf("IDs of factory servers are equal")
	}
}

// Test_FactoryServer_009 checks that the factory server returns a proper error
// in case no object type is requested.
func Test_FactoryServer_009(t *testing.T) {
	// Create a new factory server and configure it with the test gateway.
	newServerConfig := DefaultConfig()
	newServer := NewFactory(newServerConfig)
	fs := newServer.(*factoryServer)

	newConfig := gateway.DefaultSignalConfig()
	newSignal := gateway.NewSignal(newConfig)
	newSignal.SetInput(nil)

	_, err := fs.gatewayListener(newSignal)
	if !IsInvalidFactoryGatewayRequest(err) {
		t.Fatalf("FactoryServer did NOT return proper error")
	}
}

// Test_FactoryServer_010 checks that the factory server returns a proper error
// in case an invalid object type is requested.
func Test_FactoryServer_010(t *testing.T) {
	// Create a new factory server and configure it with the test gateway.
	newServerConfig := DefaultConfig()
	newServer := NewFactory(newServerConfig)
	fs := newServer.(*factoryServer)

	newConfig := gateway.DefaultSignalConfig()
	newSignal := gateway.NewSignal(newConfig)
	newSignal.SetInput(spec.ObjectType("invalid"))

	_, err := fs.gatewayListener(newSignal)
	if !IsInvalidFactoryGatewayRequest(err) {
		t.Fatalf("FactoryServer did NOT return proper error")
	}
}
