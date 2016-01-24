package network

import (
	"sync"

	"github.com/xh3b4sd/anna/neuron"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type Config struct {
	State spec.State `json:"state,omitempty"`
}

const (
	ObjectType spec.ObjectType = "network"
)

func DefaultConfig() Config {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newConfig := Config{
		State: state.NewState(newStateConfig),
	}

	return newConfig
}

func NewNetwork(config Config) spec.Network {
	newNetwork := &network{
		Config: config,
		Mutex:  sync.Mutex{},
	}

	return newNetwork
}

type network struct {
	Config

	Mutex sync.Mutex `json:"mutex,omitempty"`
}

func (n *network) GetObjectID() spec.ObjectID {
	return n.GetState().GetObjectID()
}

func (n *network) GetObjectType() spec.ObjectType {
	return n.GetState().GetObjectType()
}

func (n *network) GetState() spec.State {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	return n.State
}

func (n *network) SetState(state spec.State) {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	n.State = state
}

func (n *network) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	// Track state.
	imp.GetState().SetNetwork(n)
	n.GetState().SetImpulse(imp)

	// Get first neuron.
	var err error
	var neu spec.Neuron
	neurons := n.GetState().GetNeurons()
	if len(neurons) == 0 {
		neu = neuron.NewFirstNeuron(neuron.DefaultFirstNeuronConfig())

		// Track state.
		neu.GetState().SetNetwork(n)
		n.GetState().SetNeuron(neu)
	} else {
		for _, n := range neurons {
			// We just want the very first neuron here. So we just break after the
			// first is iteration.
			neu = n
			break
		}
	}

	imp, _, err = imp.WalkThrough(neu)
	if err != nil {
		return nil, maskAny(err)
	}

	return imp, nil
}
