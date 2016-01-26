package neuron

import (
	"sync"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type FirstNeuronConfig struct {
	Log spec.Log `json:"-"`

	States map[string]spec.State `json:"states,omitempty"`
}

func DefaultFirstNeuronConfig() FirstNeuronConfig {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newDefaultJobNeuronConfig := FirstNeuronConfig{
		Log: log.NewLog(log.DefaultConfig()),
		States: map[string]spec.State{
			common.DefaultStateKey: state.NewState(newStateConfig),
		},
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

	Mutex sync.Mutex `json:"-"`
}

func (fn *firstNeuron) Copy() spec.Neuron {
	firstNeuronCopy := *fn

	for key, state := range firstNeuronCopy.States {
		firstNeuronCopy.States[key] = state.Copy()
	}

	return &firstNeuronCopy
}

func (fn *firstNeuron) GetObjectID() spec.ObjectID {
	fn.Mutex.Lock()
	defer fn.Mutex.Unlock()

	return fn.States[common.DefaultStateKey].GetObjectID()
}

func (fn *firstNeuron) GetObjectType() spec.ObjectType {
	fn.Mutex.Lock()
	defer fn.Mutex.Unlock()

	return fn.States[common.DefaultStateKey].GetObjectType()
}

func (fn *firstNeuron) GetState(key string) (spec.State, error) {
	fn.Mutex.Lock()
	defer fn.Mutex.Unlock()

	if state, ok := fn.States[key]; ok {
		return state, nil
	}

	return nil, maskAny(stateNotFoundError)
}

func (fn *firstNeuron) SetState(key string, state spec.State) {
	fn.Mutex.Lock()
	defer fn.Mutex.Unlock()
	fn.States[key] = state
}

func (fn *firstNeuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	fn.Log.V(12).Debugf("call FirstNetwork.Trigger")

	// Track state.
	impState, err := imp.GetState(common.DefaultStateKey)
	if err != nil {
		return nil, nil, maskAny(err)
	}
	impState.SetNeuron(fn)
	neuronState, err := fn.GetState(common.DefaultStateKey)
	if err != nil {
		return nil, nil, maskAny(err)
	}
	neuronState.SetImpulse(imp)

	neu, err := neuronState.GetNeuronByID(imp.GetObjectID())
	if state.IsNeuronNotFound(err) {
		// Create new job neuron with the impulse ID.
		neu, err := common.GetInitNeuronCopy(common.JobNeuronIDKey, fn)
		if err != nil {
			return nil, nil, maskAny(err)
		}

		newStateConfig := state.DefaultConfig()
		newStateConfig.Bytes["state"] = []byte{}
		newStateConfig.ObjectID = imp.GetObjectID()
		newStateConfig.ObjectType = ObjectType

		neu.SetState(common.DefaultStateKey, state.NewState(newStateConfig))

		// Track new neuron.
		defaultNeuronState, err := fn.GetState(common.DefaultStateKey)
		if err != nil {
			return nil, nil, maskAny(err)
		}
		defaultNeuronState.SetNeuron(neu)
	} else if err != nil {
		return nil, nil, maskAny(err)
	}

	return imp, neu, nil
}
