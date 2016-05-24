// Package innet implements spec.Network to provide functionality to analyse
// given input.
package innet

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeInNet represents the object type of the input network object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeInNet spec.ObjectType = "in-net"
)

// Config represents the configuration used to create a new input network
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

// DefaultConfig provides a default configuration to create a new input network
// object by best effort.
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

// NewInNet creates a new configured input network object.
func NewInNet(config Config) (spec.Network, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newNet := &inNet{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           newID,
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeInNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type inNet struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (in *inNet) Boot() {
	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call Boot")

	in.BootOnce.Do(func() {
	})
}

func (in *inNet) Shutdown() {
	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call Shutdown")

	in.ShutdownOnce.Do(func() {
	})
}

func (in *inNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call Trigger")

	imp, err := in.ExecNet.Trigger(imp)
	if err != nil {
		return nil, maskAny(err)
	}

	return imp, nil
}
