// Package execnet implements spec.Network to execute business logic for the
// character network.
package execnet

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	ObjectTypeCharExecNet spec.ObjectType = "char-exec-net"
)

type Config struct {
	Log spec.Log
}

func DefaultConfig() Config {
	newConfig := Config{
		Log: log.NewLog(log.DefaultConfig()),
	}

	return newConfig
}

// NewExecNet returns a new configured exec network.
func NewExecNet(config Config) (spec.Network, error) {
	newNet := &execNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeCharExecNet,
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
