package neuron

import (
	"sync"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type CharacterNeuronConfig struct {
	FactoryClient spec.Factory `json:"-"`

	Log spec.Log `json:"-"`

	State spec.State `json:"state,omitempty"`
}

func DefaultCharacterNeuronConfig() CharacterNeuronConfig {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = common.ObjectType.CharacterNeuron

	newDefaultCharacterNeuronConfig := CharacterNeuronConfig{
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		State:         state.NewState(newStateConfig),
	}

	return newDefaultCharacterNeuronConfig
}

func NewCharacterNeuron(config CharacterNeuronConfig) spec.Neuron {
	newCharacterNeuron := &characterNeuron{
		CharacterNeuronConfig: config,
		Mutex: sync.Mutex{},
	}

	return newCharacterNeuron
}

type characterNeuron struct {
	CharacterNeuronConfig

	Mutex sync.Mutex `json:"-"`
}

// TODO
func (cn *characterNeuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	cn.Log.V(11).Debugf("call CharacterNeuron.Trigger")

	return imp, nil, nil
}
