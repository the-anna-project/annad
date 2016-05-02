// Package factoryserver implements spec.Factory and provides a centralized
// service for general object creation accessible via a gateway.
package factoryserver

import (
	"sync"

	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeFactoryServer represents the object type of the factory server
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeFactoryServer spec.ObjectType = "factory-server"
)

// Config represents the configuration used to create a new factory server
// object.
type Config struct {
	// Dependencies.
	FactoryGateway spec.Gateway
	Log            spec.Log
	TextGateway    spec.Gateway
}

// DefaultConfig provides a default configuration to create a new factory
// server object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		FactoryGateway: gateway.NewGateway(gateway.DefaultConfig()),
		Log:            log.NewLog(log.DefaultConfig()),
		TextGateway:    gateway.NewGateway(gateway.DefaultConfig()),
	}

	return newConfig
}

// NewFactory creates a new configured factory server object.
func NewFactory(config Config) spec.Factory {
	newFactory := &factoryServer{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeFactoryServer,
	}

	newFactory.Log.Register(newFactory.GetType())

	return newFactory
}

type factoryServer struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (fs *factoryServer) Boot() {
	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call Boot")

	fs.BootOnce.Do(func() {
		go fs.FactoryGateway.Listen(fs.gatewayListener, nil)
	})
}

func (fs *factoryServer) NewImpulse() (spec.Impulse, error) {
	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call NewImpulse")

	newConfig := impulse.DefaultConfig()
	newConfig.Log = fs.Log
	newImpulse, err := impulse.NewImpulse(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newImpulse, nil
}

func (fs *factoryServer) Shutdown() {
	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call Shutdown")

	fs.ShutdownOnce.Do(func() {
		fs.FactoryGateway.Close()
		fs.TextGateway.Close()
	})
}
