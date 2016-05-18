// Package outnet implements spec.Network to provide functionality to create
// valuable output with respect to the given input.
package outnet

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeOutNet represents the object type of the output network object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeOutNet spec.ObjectType = "out-net"
)

// Config represents the configuration used to create a new output network
// object.
type Config struct {
	Log     spec.Log
	Storage spec.Storage

	EvalNet  spec.Network
	ExecNet  spec.Network
	PatNet   spec.Network
	PredNet  spec.Network
	StratNet spec.Network
}

// DefaultConfig provides a default configuration to create a new output
// network object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

		EvalNet:  nil,
		ExecNet:  nil,
		PatNet:   nil,
		PredNet:  nil,
		StratNet: nil,
	}

	return newConfig
}

// NewOutNet creates a new configured output network object.
func NewOutNet(config Config) (spec.Network, error) {
	newNet := &outNet{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeOutNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type outNet struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (on *outNet) Boot() {
	on.Log.WithTags(spec.Tags{L: "D", O: on, T: nil, V: 13}, "call Boot")

	on.BootOnce.Do(func() {
	})
}

func (on *outNet) Shutdown() {
	on.Log.WithTags(spec.Tags{L: "D", O: on, T: nil, V: 13}, "call Shutdown")

	on.ShutdownOnce.Do(func() {
	})
}

func (on *outNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	on.Log.WithTags(spec.Tags{L: "D", O: on, T: nil, V: 13}, "call Trigger")

	// Dynamically walk impulse through the other networks.
	var err error
	for {
		imp, err = on.StratNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = on.PredNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = on.ExecNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = on.EvalNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}

		break
	}

	// Note that the impulse returned here is not actually the same as received
	// at the beginning of the call, but was manipulated during its walk through
	// the networks.
	return imp, nil
}
