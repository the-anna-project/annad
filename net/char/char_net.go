// Package charnet implements spec.Network to provide functionality to analyse,
// connect and gather data with respect to the given input.
package charnet

import (
	"sync"

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
	Log     spec.Log
	Storage spec.Storage

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

// NewCharNet creates a new configured character network object.
func NewCharNet(config Config) (spec.Network, error) {
	newNet := &charNet{
		Config: config,

		BootOnce:     sync.Once{},
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeCharNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type charNet struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (cn *charNet) Boot() {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Boot")

	cn.BootOnce.Do(func() {
	})
}

func (cn *charNet) Shutdown() {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Shutdown")

	cn.ShutdownOnce.Do(func() {
	})
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
