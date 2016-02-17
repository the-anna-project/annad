package factoryserver_test

import (
	"testing"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/factory/server"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/spec"
)

// Test_Factory_001 checks that the factory server always creates independent
// cores.
func Test_Factory_001(t *testing.T) {
	newServer := factoryserver.NewFactory(factoryserver.DefaultConfig())

	firstCore, err := newServer.NewCore()
	if err != nil {
		t.Fatalf("FactoryServer.NewCore returned err: %#v", err)
	}

	secondCore, err := newServer.NewCore()
	if err != nil {
		t.Fatalf("FactoryServer.NewCore returned err: %#v", err)
	}

	if firstCore.GetObjectID() == secondCore.GetObjectID() {
		t.Fatalf("object ID of first core and second core is equal")
	}
}

// Test_Factory_002 checks that the factory server always creates independent
// impulses.
func Test_Factory_002(t *testing.T) {
	newServer := factoryserver.NewFactory(factoryserver.DefaultConfig())

	firstImpulse, err := newServer.NewImpulse()
	if err != nil {
		t.Fatalf("FactoryServer.NewImpulse returned err: %#v", err)
	}

	secondImpulse, err := newServer.NewImpulse()
	if err != nil {
		t.Fatalf("FactoryServer.NewImpulse returned err: %#v", err)
	}

	if firstImpulse.GetObjectID() == secondImpulse.GetObjectID() {
		t.Fatalf("object ID of first impulse and second impulse is equal")
	}
}

// Test_Factory_003 checks that the factory client always creates independent
// objects.
func Test_Factory_003(t *testing.T) {
	// Create new test gateway that all participants can use.
	newFactoryGateway := gateway.NewGateway()

	// Create a new factory server and configure it with the test gateway.
	newServerConfig := factoryserver.DefaultConfig()
	newServerConfig.FactoryGateway = newFactoryGateway
	newServer := factoryserver.NewFactory(newServerConfig)

	// Create a new factory client and configure it with the test gateway.
	newClientConfig := factoryclient.DefaultConfig()
	newClientConfig.FactoryGateway = newFactoryGateway
	newClient := factoryclient.NewFactory(newClientConfig)

	firstCore, err := newServer.NewCore()
	if err != nil {
		t.Fatalf("FactoryServer.NewCore returned err: %#v", err)
	}

	secondCore, err := newClient.NewCore()
	if err != nil {
		t.Fatalf("Factory.NewCore returned err: %#v", err)
	}

	if firstCore.GetObjectID() == secondCore.GetObjectID() {
		t.Fatalf("object ID of first core and second core is equal")
	}
}

// Test_Factory_004 checks that the factory server returns its proper object
// type.
func Test_Factory_004(t *testing.T) {
	newServer := factoryserver.NewFactory(factoryserver.DefaultConfig())

	object, ok := newServer.(spec.Object)
	if !ok {
		t.Fatalf("factory server does not implement spec.Object")
	}
	if object.GetObjectType() != common.ObjectType.FactoryServer {
		t.Fatalf("invalid object tyoe of factory server")
	}
}
