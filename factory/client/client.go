// Package factoryclient implements spec.Factory and provides a decentralized
// service for general object creation.
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
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeFactoryClient,
	}

	newFactory.Log.Register(newFactory.GetType())

	return newFactory
}

type factoryClient struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (fc *factoryClient) Boot() {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call Boot")

	fc.BootOnce.Do(func() {
	})
}

func (fc *factoryClient) NewImpulse() (spec.Impulse, error) {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call NewImpulse")

	output, err := forwardSignal(fc.FactoryGateway, common.ObjectTypeImpulse, nil)
	if err != nil {
		return nil, maskAny(err)
	}

	return output.(spec.Impulse), nil
}

func (fc *factoryClient) Shutdown() {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call Shutdown")

	fc.ShutdownOnce.Do(func() {
		fc.FactoryGateway.Close()
	})
}
