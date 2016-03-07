// FactoryClient implements spec.Factory and provides a decentralized service for
// general object creation.
package factoryclient

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/common"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeFactoryClient represents the object type of the factory client
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeFactoryClient spec.ObjectType = "factory-client"
)

// Config represents the configuration used to create a new factory client
// object.
type Config struct {
	FactoryGateway spec.Gateway

	Log spec.Log
}

// DefaultConfig provides a default configuration to create a new factory
// client object by best effort.
func DefaultConfig() Config {
	config := Config{
		FactoryGateway: gateway.NewGateway(gateway.DefaultConfig()),
		Log:            log.NewLog(log.DefaultConfig()),
	}

	return config
}

// NewFactory creates a new configured factory client object.
func NewFactory(config Config) spec.Factory {
	newFactory := &factoryClient{
		Booted: false,
		Closed: false,
		Closer: make(chan struct{}, 1),
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeFactoryClient,
	}

	newFactory.Log.Register(newFactory.GetType())

	return newFactory
}

type factoryClient struct {
	Config

	Booted bool
	Closed bool
	Closer chan struct{}
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (fc *factoryClient) Boot() {
	fc.Mutex.Lock()
	defer fc.Mutex.Unlock()

	if fc.Booted {
		return
	}
	fc.Booted = true

	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call Boot")
}

func (fc *factoryClient) NewImpulse() (spec.Impulse, error) {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call NewImpulse")

	output, err := forwardSignal(fc.FactoryGateway, common.ObjectTypeImpulse, fc.Closer)
	if err != nil {
		return nil, maskAny(err)
	}

	return output.(spec.Impulse), nil
}

func (fc *factoryClient) Shutdown() {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call Shutdown")

	fc.Mutex.Lock()
	defer fc.Mutex.Unlock()

	fc.FactoryGateway.Close()

	if !fc.Closed {
		fc.Closer <- struct{}{}
		fc.Closed = true
	}
}
