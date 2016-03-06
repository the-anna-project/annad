package factoryserver

import (
	"testing"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/spec"
)

// Test_FactoryServer_NewImpulse checks that the factory server always creates
// independent impulses.
func Test_FactoryServer_NewImpulse(t *testing.T) {
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

// Test_FactoryServer_NewImpulse_WithClient checks that the factory client
// always creates independent impulses.
//
// Further checks that booting works.
// Further checks that shutting down works.
func Test_FactoryServer_NewImpulse_WithClient(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway(gateway.DefaultConfig())

	// Create a new factory server and configure it with the test gateway.
	newServerConfig := DefaultConfig()
	newServerConfig.FactoryGateway = newFactoryGateway
	newServer := NewFactory(newServerConfig)

	// Check multiple boots works.
	newServer.Boot()
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

	// Check multiple shutdowns works.
	newServer.Shutdown()
	newClient.Shutdown()
}

// Test_FactoryServer_GetType checks that the factory server returns its proper
// object type.
func Test_FactoryServer_GetType(t *testing.T) {
	newServer := NewFactory(DefaultConfig())

	object, ok := newServer.(spec.Object)
	if !ok {
		t.Fatalf("factory server does not implement spec.Object")
	}
	if object.GetType() != ObjectTypeFactoryServer {
		t.Fatalf("invalid object tyoe of factory server")
	}
}

// Test_FactoryServer_GetID checks that the factory server returns its proper
// object ID.
func Test_FactoryClient_GetID(t *testing.T) {
	firstClient := NewFactory(DefaultConfig())
	firstObject, _ := firstClient.(spec.Object)
	secondClient := NewFactory(DefaultConfig())
	secondObject, _ := secondClient.(spec.Object)

	if firstObject.GetID() == secondObject.GetID() {
		t.Fatalf("IDs of factory servers are equal")
	}
}

// Test_FactoryServer_NoObjectType checks that the factory server returns a
// proper error in case no object type is requested.
func Test_FactoryServer_NoObjectType(t *testing.T) {
	// Create a new factory server and configure it with the test gateway.
	newServerConfig := DefaultConfig()
	newServer := NewFactory(newServerConfig)
	fs := newServer.(*factoryServer)

	newConfig := gateway.DefaultSignalConfig()
	newConfig.Input = nil
	newSignal := gateway.NewSignal(newConfig)

	_, err := fs.gatewayListener(newSignal)
	if !IsInvalidFactoryGatewayRequest(err) {
		t.Fatalf("FactoryServer did NOT return proper error")
	}
}

// Test_FactoryServer_InvalidObjectType checks that the factory server returns
// a proper error in case an invalid object type is requested.
func Test_FactoryServer_InvalidObjectType(t *testing.T) {
	// Create a new factory server and configure it with the test gateway.
	newServerConfig := DefaultConfig()
	newServer := NewFactory(newServerConfig)
	fs := newServer.(*factoryServer)

	newConfig := gateway.DefaultSignalConfig()
	newConfig.Input = spec.ObjectType("invalid")
	newSignal := gateway.NewSignal(newConfig)

	_, err := fs.gatewayListener(newSignal)
	if !IsInvalidFactoryGatewayRequest(err) {
		t.Fatalf("FactoryServer did NOT return proper error")
	}
}
