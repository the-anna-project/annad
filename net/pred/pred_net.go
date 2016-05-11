// Package prednet implements spec.Network to provide functionality to predict
// and warn about executing actions may leading to certain risks.
package prednet

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypePredNet represents the object type of the prediction network
	// object. This is used e.g. to register itself to the logger.
	ObjectTypePredNet spec.ObjectType = "pred-net"
)

// Config represents the configuration used to create a new prediction network
// object.
type Config struct {
	Log     spec.Log
	Storage spec.Storage

	PatNet spec.Network
}

// DefaultConfig provides a default configuration to create a new prediction
// network object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

		PatNet: nil,
	}

	return newConfig
}

// NewPredNet creates a new configured prediction network object.
func NewPredNet(config Config) (spec.Network, error) {
	newNet := &predNet{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypePredNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type predNet struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (pn *predNet) Boot() {
	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Boot")

	pn.BootOnce.Do(func() {
	})
}

func (pn *predNet) Shutdown() {
	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Shutdown")

	pn.ShutdownOnce.Do(func() {
	})
}

func (pn *predNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	pn.Log.WithTags(spec.Tags{L: "D", O: pn, T: nil, V: 13}, "call Trigger")
	return imp, nil
}
