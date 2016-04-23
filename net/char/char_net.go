// Package charnet implements spec.Network to provide functionality to analyse,
// connect and gather data with respect to the given input.
package charnet

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeCharNet represents the object type of the character network
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeCharNet spec.ObjectType = "char-net"
)

// Config represents the configuration used to create a new character network
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

// DefaultConfig provides a default configuration to create a new character
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

// NewCharNet creates a new configured character network object.
func NewCharNet(config Config) (spec.Network, error) {
	newNet := &charNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeCharNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type charNet struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (cn *charNet) Boot() {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()

	if cn.Booted {
		return
	}
	cn.Booted = true

	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Boot")
}

func (cn *charNet) Shutdown() {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Shutdown")
}

func (cn *charNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Trigger")

	// TODO prepare (new?) impulse (with CLGs?) for strat net

	imp, err := cn.StratNet.Trigger(imp)
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

	return imp, nil
}
