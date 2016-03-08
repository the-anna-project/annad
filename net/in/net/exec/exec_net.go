// Package execnet implements spec.Network to execute business logic for the
// in network.
package execnet

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
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
	newNet := &execNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeInExecNet,
	}

	newNet.Log.Register(newNet.GetType())

	return newNet, nil
}

type execNet struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (en *execNet) Boot() {
	en.Mutex.Lock()
	defer en.Mutex.Unlock()

	if en.Booted {
		return
	}
	en.Booted = true

	en.Log.WithTags(spec.Tags{L: "D", O: en, T: nil, V: 13}, "call Boot")
}

func (en *execNet) Shutdown() {
	en.Log.WithTags(spec.Tags{L: "D", O: en, T: nil, V: 13}, "call Shutdown")
}

func (en *execNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	en.Log.WithTags(spec.Tags{L: "D", O: en, T: nil, V: 13}, "call Trigger")
	return imp, nil
}
