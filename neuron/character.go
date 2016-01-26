package neuron

import (
	"sync"
	"time"

	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type CharacterNeuronConfig struct {
	Log spec.Log `json:"-"`

	States map[string]spec.State `json:"states,omitempty"`
}

func DefaultCharacterNeuronConfig() CharacterNeuronConfig {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newDefaultCharacterNeuronConfig := CharacterNeuronConfig{
		Log: log.NewLog(log.DefaultConfig()),
		States: map[string]spec.State{
			"default": state.NewState(newStateConfig),
		},
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

func (cn *characterNeuron) Copy() spec.Neuron {
	characterNeuronCopy := *cn

	for key, state := range characterNeuronCopy.States {
		characterNeuronCopy.States[key] = state.Copy()
	}

	return &characterNeuronCopy
}

func (cn *characterNeuron) GetObjectID() spec.ObjectID {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()

	return cn.States["default"].GetObjectID()
}

func (cn *characterNeuron) GetObjectType() spec.ObjectType {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()

	return cn.States["default"].GetObjectType()
}

func (cn *characterNeuron) GetState(key string) (spec.State, error) {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()

	if state, ok := cn.States[key]; ok {
		return state, nil
	}

	return nil, maskAny(stateNotFoundError)
}

func (cn *characterNeuron) SetState(key string, state spec.State) {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	cn.States[key] = state
}

// TODO
func (cn *characterNeuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	// Track state.
	impState, err := imp.GetState("default")
	if err != nil {
		return nil, nil, maskAny(err)
	}
	impState.SetNeuron(cn)
	neuronState, err := cn.GetState("default")
	if err != nil {
		return nil, nil, maskAny(err)
	}
	neuronState.SetImpulse(imp)

	time.Sleep(5 * time.Second)

	response, err := impState.GetBytes("request")
	if err != nil {
		return nil, nil, maskAny(err)
	}
	impState.SetBytes("response", response)

	return imp, nil, nil
}
