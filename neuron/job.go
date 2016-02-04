package neuron

import (
	"sync"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type JobNeuronConfig struct {
	FactoryClient spec.Factory `json:"-"`

	Log spec.Log `json:"-"`

	State spec.State `json:"state,omitempty"`
}

func DefaultJobNeuronConfig() JobNeuronConfig {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = common.ObjectType.JobNeuron

	newDefaultJobNeuronConfig := JobNeuronConfig{
		FactoryClient: factoryclient.NewClient(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		State:         state.NewState(newStateConfig),
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

func (jn *jobNeuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	jn.Log.V(11).Debugf("call JobNeuron.Trigger")

	bytes, err := jn.GetState().GetBytes("state")
	if err != nil {
		return nil, nil, maskAny(err)
	}
	jn.Log.V(12).Debugf("job neuron state: '%s'", bytes)

	switch string(bytes) {
	case "":
		// Create new impulse to process it asynchronously.
		go func(imp spec.Impulse) {
			jn.GetState().SetBytes("state", []byte("in-progress"))

			neu, err := jn.FactoryClient.NewCharacterNeuron()
			if err != nil {
				jn.Log.V(3).Errorf("%#v", maskAny(err))
				return
			}
			jn.GetState().SetNeuron(neu)

			imp, _, err = imp.WalkThrough(neu)
			if err != nil {
				jn.Log.V(3).Errorf("%#v", maskAny(err))
				return
			}

			request, err := imp.GetState().GetBytes("request")
			if err != nil {
				jn.Log.V(3).Errorf("%#v", maskAny(err))
				return
			}
			jn.GetState().SetBytes("response", request)
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
