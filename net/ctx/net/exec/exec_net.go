// Package execnet implements spec.Network to execute business logic for the
// context network.
package execnet

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCtxExecNet represents the object type of the context network's
	// execution network object. This is used e.g. to register itself to the
	// logger.
	ObjectTypeCtxExecNet spec.ObjectType = "ctx-exec-net"
)

// Config represents the configuration used to create a new context execution
// network object.
type Config struct {
	Log spec.Log
}

// DefaultConfig provides a default configuration to create a new context
// execution network object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Log: log.NewLog(log.DefaultConfig()),
	}

	return newConfig
}

// NewExecNet creates a new configured context execution network object.
func NewExecNet(config Config) (spec.Network, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		return nil, maskAny(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		return nil, maskAny(err)
	}

	newNet := &execNet{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           newID,
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeCtxExecNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type execNet struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (en *execNet) Boot() {
	en.Log.WithTags(spec.Tags{L: "D", O: en, T: nil, V: 13}, "call Boot")

	en.BootOnce.Do(func() {
	})
}

func (en *execNet) Shutdown() {
	en.Log.WithTags(spec.Tags{L: "D", O: en, T: nil, V: 13}, "call Shutdown")

	en.ShutdownOnce.Do(func() {
	})
}

func (en *execNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	en.Log.WithTags(spec.Tags{L: "D", O: en, T: nil, V: 13}, "call Trigger")
	return imp, nil
}
