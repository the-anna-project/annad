package characterneuron

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
	newStateConfig.ObjectType = common.ObjectType.CharacterNeuron

	newNeuronConfig := Config{
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		State:         state.NewState(newStateConfig),
	}

	return newNeuronConfig
}

func NewNeuron(config Config) spec.Neuron {
	newNeuron := &neuron{
		Config: config,
		Mutex:  sync.Mutex{},
	}

	return newNeuron
}

type neuron struct {
	Config

	Mutex sync.Mutex `json:"-"`
}

// TODO
func (n *neuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Trigger")

	return imp, nil, nil
}
