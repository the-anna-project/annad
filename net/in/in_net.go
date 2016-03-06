// Package innet implements spec.Network to provide functionality to analyse
// given input.
package innet

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	ObjectTypeInNet spec.ObjectType = "in-net"
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
		Storage:       memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

		EvalNet:  nil,
		ExecNet:  nil,
		PatNet:   nil,
		PredNet:  nil,
		StratNet: nil,
	}

	return newConfig
}

// NewInNet returns a new configured input network.
func NewInNet(config Config) (spec.Network, error) {
	newNet := &inNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeInNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type inNet struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (in *inNet) Boot() {
	in.Mutex.Lock()
	defer in.Mutex.Unlock()

	if in.Booted {
		return
	}
	in.Booted = true

	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call Boot")
}

func (in *inNet) Shutdown() {
	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call Shutdown")
}

func (in *inNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call Trigger")

	// Dynamically walk impulse through the other networks.
	var err error
	for {
		imp, err = in.StratNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = in.PredNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = in.ExecNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = in.EvalNet.Trigger(imp)
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
