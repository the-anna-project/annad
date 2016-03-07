// Package patnet implements spec.Network to provide functionality to find and
// interpret patterns in given input.
package patnet

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
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
	newNet := &patNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypePatNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type patNet struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (pn *patNet) Boot() {
	pn.Mutex.Lock()
	defer pn.Mutex.Unlock()

	if pn.Booted {
		return
	}
	pn.Booted = true

	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Boot")
}

func (pn *patNet) Shutdown() {
	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Shutdown")
}

func (pn *patNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Trigger")
	return imp, nil
}
