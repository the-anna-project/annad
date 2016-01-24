package neuron

import (
	"sync"
	"time"

	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type CharacterNeuronConfig struct {
	State spec.State `json:"state,omitempty"`
}

func DefaultCharacterNeuronConfig() CharacterNeuronConfig {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newDefaultCharacterNeuronConfig := CharacterNeuronConfig{
		State: state.NewState(newStateConfig),
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

	Mutex sync.Mutex `json:"mutex,omitempty"`
}

func (cn *characterNeuron) GetObjectID() spec.ObjectID {
	return cn.GetState().GetObjectID()
}

func (cn *characterNeuron) GetNetwork() (spec.Network, error) {
	networks := cn.GetState().GetNetworks()

	if len(networks) != 1 {
		return nil, maskAny(invalidNetworkRelationError)
	}

	var network spec.Network
	for _, n := range cn.GetState().GetNetworks() {
		network = n
		break
	}

	return network, nil
}

func (cn *characterNeuron) GetObjectType() spec.ObjectType {
	return cn.GetState().GetObjectType()
}

func (cn *characterNeuron) GetState() spec.State {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	return cn.State
}

func (cn *characterNeuron) SetState(state spec.State) {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	cn.State = state
}

func (cn *characterNeuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	// Track state.
	imp.GetState().SetNeuron(cn)
	cn.GetState().SetImpulse(imp)

	time.Sleep(5 * time.Second)

	response, err := imp.GetState().GetBytes("request")
	if err != nil {
		return nil, nil, maskAny(err)
	}
	imp.GetState().SetBytes("response", response)

	return imp, nil, nil
}
