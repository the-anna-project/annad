package neuron

import (
	"fmt"
	"sync"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type JobNeuronConfig struct {
	Log spec.Log `json:"-"`

	States map[string]spec.State `json:"states,omitempty"`
}

func DefaultJobNeuronConfig() JobNeuronConfig {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newDefaultJobNeuronConfig := JobNeuronConfig{
		Log: log.NewLog(log.DefaultConfig()),
		States: map[string]spec.State{
			common.DefaultStateKey: state.NewState(newStateConfig),
		},
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

	Mutex sync.Mutex `json:"-"`
}

func (jn *jobNeuron) Copy() spec.Neuron {
	jobNeuronCopy := *jn

	for key, state := range jobNeuronCopy.States {
		jobNeuronCopy.States[key] = state.Copy()
	}

	return &jobNeuronCopy
}

func (jn *jobNeuron) GetObjectID() spec.ObjectID {
	jn.Mutex.Lock()
	defer jn.Mutex.Unlock()

	return jn.States[common.DefaultStateKey].GetObjectID()
}

func (jn *jobNeuron) GetObjectType() spec.ObjectType {
	jn.Mutex.Lock()
	defer jn.Mutex.Unlock()

	return jn.States[common.DefaultStateKey].GetObjectType()
}

func (jn *jobNeuron) GetState(key string) (spec.State, error) {
	jn.Mutex.Lock()
	defer jn.Mutex.Unlock()

	if state, ok := jn.States[key]; ok {
		return state, nil
	}

	return nil, maskAny(stateNotFoundError)
}

func (jn *jobNeuron) SetState(key string, state spec.State) {
	jn.Mutex.Lock()
	defer jn.Mutex.Unlock()
	jn.States[key] = state
}

func (jn *jobNeuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	// Track state.
	impState, err := imp.GetState(common.DefaultStateKey)
	if err != nil {
		return nil, nil, maskAny(err)
	}
	impState.SetNeuron(jn)
	neuronState, err := jn.GetState(common.DefaultStateKey)
	if err != nil {
		return nil, nil, maskAny(err)
	}
	neuronState.SetImpulse(imp)

	bytes, err := neuronState.GetBytes("state")
	if err != nil {
		return nil, nil, maskAny(err)
	}

	switch string(bytes) {
	case "":
		// Create new impulse to process it asynchronously.
		go func(imp spec.Impulse) {
			defaultNeuronState, err := jn.GetState(common.DefaultStateKey)
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				return
			}
			defaultNeuronState.SetBytes("state", []byte("in-progress"))

			initNeuronState, err := jn.GetState(common.InitStateKey)
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				return
			}
			neuronID, err := initNeuronState.GetBytes(common.CharacterIDKey)
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				return
			}
			initCharacterNeuron, err := initNeuronState.GetNeuronByID(spec.ObjectID(neuronID))
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				return
			}
			neu := initCharacterNeuron.Copy()

			defaultNeuronState.SetNeuron(neu)

			imp, _, err = imp.WalkThrough(neu)
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				return
			}

			impState, err := imp.GetState(common.DefaultStateKey)
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				return
			}
			response, err := impState.GetBytes("response")
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				return
			}
			defaultNeuronState.SetBytes("response", response)
			defaultNeuronState.SetBytes("state", []byte("finished"))
		}(imp)

		// Return the impulse ID to signal a job registration.
		impState.SetBytes("response", []byte(impState.GetObjectID()))

		return imp, nil, nil
	case "in-progress":
		// Return to keep waiting.
		return imp, nil, nil
	case "finished":
		// Return response.
		response, err := neuronState.GetBytes("response")
		if err != nil {
			return nil, nil, maskAny(err)
		}

		impState.SetBytes("response", response)

		return imp, nil, nil
	}

	return nil, nil, maskAny(invalidImpulseStateError)
}
