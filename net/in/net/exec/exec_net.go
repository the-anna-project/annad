// Package execnet implements spec.Network to execute business logic for the
// in network.
package execnet

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeInExecNet represents the object type of the input network's
	// execution network object. This is used e.g. to register itself to the
	// logger.
	ObjectTypeInExecNet spec.ObjectType = "in-exec-net"
)

// Config represents the configuration used to create a new input execution
// network object.
type Config struct {
	Log spec.Log

	CharNet spec.Network
	CtxNet  spec.Network
}

// DefaultConfig provides a default configuration to create a new input
// execution network object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Log: log.NewLog(log.DefaultConfig()),

		CharNet: nil,
		CtxNet:  nil,
	}

	return newConfig
}

// NewExecNet creates a new configured input execution network object.
func NewExecNet(config Config) (spec.Network, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newNet := &execNet{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           newID,
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeInExecNet,
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

	imp, err := en.CharNet.Trigger(imp)
	if err != nil {
		return nil, maskAny(err)
	}
	imp, err = en.CtxNet.Trigger(imp)
	if err != nil {
		return nil, maskAny(err)
	}

	return imp, nil
}
