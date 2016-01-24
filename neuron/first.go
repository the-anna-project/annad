package neuron

import (
	"sync"

	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type FirstNeuronConfig struct {
	State spec.State `json:"state,omitempty"`
}

func DefaultFirstNeuronConfig() FirstNeuronConfig {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newDefaultJobNeuronConfig := FirstNeuronConfig{
		State: state.NewState(newStateConfig),
	}

	return newDefaultJobNeuronConfig
}

func NewFirstNeuron(config FirstNeuronConfig) spec.Neuron {
	newFirstNeuron := &firstNeuron{
		FirstNeuronConfig: config,
		Mutex:             sync.Mutex{},
	}

	return newFirstNeuron
}

type firstNeuron struct {
	FirstNeuronConfig

	Mutex sync.Mutex `json:"mutex,omitempty"`
}

func (fn *firstNeuron) GetObjectID() spec.ObjectID {
	return fn.GetState().GetObjectID()
}

func (fn *firstNeuron) GetNetwork() (spec.Network, error) {
	networks := fn.GetState().GetNetworks()

	if len(networks) != 1 {
		return nil, maskAny(invalidNetworkRelationError)
	}

	var network spec.Network
	for _, n := range fn.GetState().GetNetworks() {
		network = n
		break
	}

	return network, nil
}

func (fn *firstNeuron) GetObjectType() spec.ObjectType {
	return fn.GetState().GetObjectType()
}

func (fn *firstNeuron) GetState() spec.State {
	fn.Mutex.Lock()
	defer fn.Mutex.Unlock()
	return fn.State
}

func (fn *firstNeuron) SetState(state spec.State) {
	fn.Mutex.Lock()
	defer fn.Mutex.Unlock()
	fn.State = state
}

func (fn *firstNeuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	// Track state.
	imp.GetState().SetNeuron(fn)
	fn.GetState().SetImpulse(imp)

	neu, err := fn.GetState().GetNeuronByID(imp.GetObjectID())
	if state.IsNeuronNotFound(err) {
		// Create new job neuron with the impulse ID.
		newStateConfig := state.DefaultConfig()
		newStateConfig.Bytes["state"] = []byte{}
		newStateConfig.ObjectID = imp.GetObjectID()
		newStateConfig.ObjectType = ObjectType

		config := JobNeuronConfig{
			State: state.NewState(newStateConfig),
		}
		neu = NewJobNeuron(config)

		// Track new neuron.
		fn.GetState().SetNeuron(neu)
	} else if err != nil {
		return nil, nil, maskAny(err)
	}

	return imp, neu, nil
}
