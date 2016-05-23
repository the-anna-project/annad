// Package patnet implements spec.Network to provide functionality to find and
// interpret patterns in given input.
package patnet

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypePatNet represents the object type of the pattern network object.
	// This is used e.g. to register itself to the logger.
	ObjectTypePatNet spec.ObjectType = "pat-net"
)

// Config represents the configuration used to create a new pattern network
// object.
type Config struct {
	Log     spec.Log
	Storage spec.Storage

	PatNet spec.Network
}

// DefaultConfig provides a default configuration to create a new pattern
// network object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

		PatNet: nil,
	}

	return newConfig
}

// NewPatNet creates a new configured pattern network object.
func NewPatNet(config Config) (spec.Network, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newNet := &patNet{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           newID,
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypePatNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type patNet struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (pn *patNet) Boot() {
	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Boot")

	pn.BootOnce.Do(func() {
	})
}

func (pn *patNet) Shutdown() {
	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Shutdown")

	pn.ShutdownOnce.Do(func() {
	})
}

func (pn *patNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Trigger")
	return imp, nil
}
