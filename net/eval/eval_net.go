// Package evalnet implements spec.Network to provide functionality to evaluate
// strategies and decide if they should scored higher or lower.
package evalnet

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeEvalNet represents the object type of the evaluation network
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeEvalNet spec.ObjectType = "eval-net"
)

// Config represents the configuration used to create a new evaluation network
// object.
type Config struct {
	Log     spec.Log
	Storage spec.Storage

	PatNet spec.Network
}

// DefaultConfig provides a default configuration to create a new evaluation
// network object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

		PatNet: nil,
	}

	return newConfig
}

// NewEvalNet creates a new configured evaluation network object.
func NewEvalNet(config Config) (spec.Network, error) {
	newNet := &evalNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeEvalNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type evalNet struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (en *evalNet) Boot() {
	en.Mutex.Lock()
	defer en.Mutex.Unlock()

	if en.Booted {
		return
	}
	en.Booted = true

	en.Log.WithTags(spec.Tags{L: "D", O: en, T: nil, V: 13}, "call Boot")
}

func (en *evalNet) Shutdown() {
	en.Log.WithTags(spec.Tags{L: "D", O: en, T: nil, V: 13}, "call Shutdown")
}

func (en *evalNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	en.Log.WithTags(spec.Tags{L: "D", O: en, T: nil, V: 13}, "call Trigger")
	return imp, nil
}