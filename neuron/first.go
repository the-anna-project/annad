package neuron

import (
	"sync"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type FirstNeuronConfig struct {
	FactoryClient spec.Factory `json:"-"`

	Log spec.Log `json:"-"`

	State spec.State `json:"state,omitempty"`
}

func DefaultFirstNeuronConfig() FirstNeuronConfig {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = common.ObjectType.FirstNeuron

	newDefaultJobNeuronConfig := FirstNeuronConfig{
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		State:         state.NewState(newStateConfig),
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

func (fn *firstNeuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	fn.Log.V(11).Debugf("call FirstNeuron.Trigger")

	neu, err := fn.GetState().GetNeuronByID(imp.GetObjectID())
	if state.IsNeuronNotFound(err) {
		fn.Log.V(12).Debugf("create job neuron")

		// Create new job neuron with the impulse ID.
		neu, err = fn.FactoryClient.NewJobNeuron()
		if err != nil {
			return nil, nil, maskAny(err)
		}

		newStateConfig := state.DefaultConfig()
		newStateConfig.Bytes["state"] = []byte{}
		newStateConfig.ObjectID = imp.GetObjectID()
		newStateConfig.ObjectType = common.ObjectType.JobNeuron

		neu.SetState(state.NewState(newStateConfig))

		// Track new neuron.
		fn.GetState().SetNeuron(neu)
	} else if err != nil {
		return nil, nil, maskAny(err)
	}

	return imp, neu, nil
}
