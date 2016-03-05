// Package respnet implements spec.Network to provide functionality to create
// valuable responses with respect to all gathered information in preceeding
// networks.
package respnet

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

const (
	ObjectTypeRespNet spec.ObjectType = "resp-net"
)

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

func DefaultConfig() Config {
	newConfig := Config{
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		Storage:       storage.NewMemoryStorage(storage.DefaultMemoryStorageConfig()),

		EvalNet:  nil,
		ExecNet:  nil,
		PatNet:   nil,
		PredNet:  nil,
		StratNet: nil,
	}

	return newConfig
}

// NewRespNet returns a new configured resp network.
func NewRespNet(config Config) (spec.Network, error) {
	newNet := &respNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeRespNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type respNet struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (rn *respNet) Boot() {
	rn.Mutex.Lock()
	defer rn.Mutex.Unlock()

	if rn.Booted {
		return
	}
	rn.Booted = true

	rn.Log.WithTags(spec.Tags{L: "D", O: rn, T: nil, V: 13}, "call Boot")
}

func (rn *respNet) Shutdown() {
	rn.Log.WithTags(spec.Tags{L: "D", O: rn, T: nil, V: 13}, "call Shutdown")
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
