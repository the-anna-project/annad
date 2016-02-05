// Package impulse implementes spec.Impulse. An impulse can walk through any
// spec.Core, spec.Network and spec.Neuron. Concrete implementations and their
// dynamic state decide about the way an impulse is going, resulting in
// behaviour.
package impulse

import (
	"sync"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type Config struct {
	FactoryClient spec.Factory `json:"-"`

	Log spec.Log `json:"-"`

	State spec.State `json:"state,omitempty"`
}

func DefaultConfig() Config {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = common.ObjectType.Impulse

	newConfig := Config{
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		State:         state.NewState(newStateConfig),
	}

	return newConfig
}

func NewImpulse(config Config) spec.Impulse {
	newImpulse := &impulse{
		Config: config,
		Mutex:  sync.Mutex{},
	}

	return newImpulse
}

type impulse struct {
	Config

	Mutex sync.Mutex `json:"-"`
}

func (i *impulse) WalkThrough(neu spec.Neuron) (spec.Impulse, spec.Neuron, error) {
	i.Log.V(11).Debugf("call Impulse.WalkThrough")

	var err error
	var imp spec.Impulse = i

	// Process the further walk through of the impulse dynamically. The loop is a
	// good alternative for real recursion.
	for {
		imp, neu, err = neu.Trigger(imp)
		if err != nil {
			return nil, nil, maskAny(err)
		}

		if neu == nil {
			// As soon as a neuron has decided to not forward an impulse to any
			// further neuron, the impulse went its way all through the whole
			// network. So we break here to return the impulse.
			break
		}
	}

	return imp, neu, nil
}
