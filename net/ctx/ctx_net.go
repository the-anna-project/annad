// Package ctxnet implements spec.Network to provide functionality to analyse,
// collect contextual information and enrich a request's scope with important
// data.
package ctxnet

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeCtxNet represents the object type of the context network object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeCtxNet spec.ObjectType = "ctx-net"
)

// Config represents the configuration used to create a new context network
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

// DefaultConfig provides a default configuration to create a new context
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

// NewCtxNet creates a new configured context network object.
func NewCtxNet(config Config) (spec.Network, error) {
	newNet := &ctxNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeCtxNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type ctxNet struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (cn *ctxNet) Boot() {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()

	if cn.Booted {
		return
	}
	cn.Booted = true

	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Boot")
}

func (cn *ctxNet) Shutdown() {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Shutdown")
}

func (cn *ctxNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Trigger")

	// Dynamically walk impulse through the other networks.
	var err error
	for {
		imp, err = cn.StratNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = cn.PredNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = cn.ExecNet.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = cn.EvalNet.Trigger(imp)
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
