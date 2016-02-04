package factoryserver_test

import (
	"testing"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/factory/server"
	"github.com/xh3b4sd/anna/gateway"
)

// Test_Factory_001 checks that the factory server always creates independent
// objects.
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

// Test_Factory_002 checks that the factory client always creates independent
// objects.
func Test_Factory_002(t *testing.T) {
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
