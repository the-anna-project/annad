package neuron

import (
	"fmt"
	"sync"

	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type JobNeuronConfig struct {
	State spec.State `json:"state,omitempty"`
}

func DefaultJobNeuronConfig() JobNeuronConfig {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newDefaultJobNeuronConfig := JobNeuronConfig{
		State: state.NewState(newStateConfig),
	}

	return newDefaultJobNeuronConfig
}

func NewJobNeuron(config JobNeuronConfig) spec.Neuron {
	newJobNeuron := &jobNeuron{
		JobNeuronConfig: config,
		Mutex:           sync.Mutex{},
	}

	return newJobNeuron
}

type jobNeuron struct {
	JobNeuronConfig

	Mutex sync.Mutex `json:"mutex,omitempty"`
}

func (jn *jobNeuron) GetObjectID() spec.ObjectID {
	return jn.GetState().GetObjectID()
}

func (jn *jobNeuron) GetNetwork() (spec.Network, error) {
	networks := jn.GetState().GetNetworks()

	if len(networks) != 1 {
		return nil, maskAny(invalidNetworkRelationError)
	}

	var network spec.Network
	for _, n := range jn.GetState().GetNetworks() {
		network = n
		break
	}

	return network, nil
}

func (jn *jobNeuron) GetObjectType() spec.ObjectType {
	return jn.GetState().GetObjectType()
}

func (jn *jobNeuron) GetState() spec.State {
	jn.Mutex.Lock()
	defer jn.Mutex.Unlock()
	return jn.State
}

func (jn *jobNeuron) SetState(state spec.State) {
	jn.Mutex.Lock()
	defer jn.Mutex.Unlock()
	jn.State = state
}

func (jn *jobNeuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	// Track state.
	imp.GetState().SetNeuron(jn)
	jn.GetState().SetImpulse(imp)

	bytes, err := jn.GetState().GetBytes("state")
	if err != nil {
		return nil, nil, maskAny(err)
	}

	switch string(bytes) {
	case "":
		// Create new impulse to process it asynchronously.
		go func(imp spec.Impulse) {
			jn.GetState().SetBytes("state", []byte("in-progress"))

			neu := NewCharacterNeuron(DefaultCharacterNeuronConfig())
			jn.GetState().SetNeuron(neu)

			imp, _, err = imp.WalkThrough(neu)
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				return
			}

			response, err := imp.GetState().GetBytes("response")
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				return
			}
			jn.GetState().SetBytes("response", response)
			jn.GetState().SetBytes("state", []byte("finished"))
		}(imp)

		// Return the impulse ID to signal a job registration.
		imp.GetState().SetBytes("response", []byte(imp.GetState().GetObjectID()))

		return imp, nil, nil
	case "in-progress":
		// Return to keep waiting.
		return imp, nil, nil
	case "finished":
		// Return response.
		response, err := jn.GetState().GetBytes("response")
		if err != nil {
			return nil, nil, maskAny(err)
		}

		imp.GetState().SetBytes("response", response)

		return imp, nil, nil
	}

	return nil, nil, maskAny(invalidImpulseStateError)
}
