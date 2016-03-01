// Package ideanet implements spec.Network to provide functionality to bring in
// creative ideas into the inout processing and output creation.
package ideanet

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

const (
	ObjectTypeIdeaNet spec.ObjectType = "idea-net"
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

// NewIdeaNet returns a new configured idea network.
func NewIdeaNet(config Config) (spec.Network, error) {
	newNet := &ideaNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeIdeaNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type ideaNet struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (in *ideaNet) Boot() {
	in.Mutex.Lock()
	defer in.Mutex.Unlock()

	if in.Booted {
		return
	}
	in.Booted = true

	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call Boot")
}

func (in *ideaNet) Shutdown() {
	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call Shutdown")
}

func (in *ideaNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
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
