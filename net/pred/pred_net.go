// Package prednet implements spec.Network to provide functionality to predict
// and warn about executing actions may leading to certain risks.
package prednet

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

const (
	ObjectTypePredNet spec.ObjectType = "pred-net"
)

type Config struct {
	Log     spec.Log
	Storage spec.Storage

	PatNet spec.Network
}

func DefaultConfig() Config {
	newConfig := Config{
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: storage.NewMemoryStorage(storage.DefaultMemoryStorageConfig()),

		PatNet: nil,
	}

	return newConfig
}

// NewPredNet returns a new configured pred network.
func NewPredNet(config Config) (spec.Network, error) {
	newNet := &predNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypePredNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type predNet struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (pn *predNet) Boot() {
	pn.Mutex.Lock()
	defer pn.Mutex.Unlock()

	if pn.Booted {
		return
	}
	pn.Booted = true

	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Boot")
}

func (pn *predNet) Shutdown() {
	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Shutdown")
}

func (pn *predNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Trigger")
	return imp, nil
}
