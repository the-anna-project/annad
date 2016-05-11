// Package respnet implements spec.Network to provide functionality to create
// valuable responses with respect to all gathered information in preceding
// networks.
package respnet

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeRespNet represents the object type of the response network
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeRespNet spec.ObjectType = "resp-net"
)

// Config represents the configuration used to create a new response network
// object.
type Config struct {
	FactoryClient spec.Factory
	Log           spec.Log
	Storage       spec.Storage

	EvalNet  spec.Network
	ExecNet  spec.Network
	PatNet   spec.Network
	PredNet  spec.Network
	StratNet spec.Network
}

// DefaultConfig provides a default configuration to create a new response
// network object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		Storage:       memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

		EvalNet:  nil,
		ExecNet:  nil,
		PatNet:   nil,
		PredNet:  nil,
		StratNet: nil,
	}

	return newConfig
}

// NewRespNet creates a new configured response network object.
func NewRespNet(config Config) (spec.Network, error) {
	newNet := &respNet{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeRespNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type respNet struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (rn *respNet) Boot() {
	rn.Log.WithTags(spec.Tags{L: "D", O: rn, T: nil, V: 13}, "call Boot")

	rn.BootOnce.Do(func() {
	})
}

func (rn *respNet) Shutdown() {
	rn.Log.WithTags(spec.Tags{L: "D", O: rn, T: nil, V: 13}, "call Shutdown")

	rn.ShutdownOnce.Do(func() {
	})
}

func (rn *respNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	rn.Log.WithTags(spec.Tags{L: "D", O: rn, T: nil, V: 13}, "call Trigger")

	// Dynamically walk impulse through the other networks.
	var err error
	for {
		imp, err = rn.StratNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = rn.PredNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = rn.ExecNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = rn.EvalNet.Trigger(imp)
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
