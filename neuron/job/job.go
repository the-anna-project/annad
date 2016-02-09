package jobneuron

import (
	"sync"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type Config struct {
	FactoryClient spec.Factory `json:"-"`

	Log spec.Log `json:"-"`

	State spec.State `json:"state,omitempty"`
}

func DefaultConfig() Config {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = common.ObjectType.JobNeuron

	newDefaultConfig := Config{
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		State:         state.NewState(newStateConfig),
	}

	return newDefaultConfig
}

func NewNeuron(config Config) spec.Neuron {
	newNeuron := &neuron{
		Config: config,
		Mutex:  sync.Mutex{},
	}

	return newNeuron
}

type neuron struct {
	Config

	Mutex sync.Mutex `json:"-"`
}

func (n *neuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Trigger")

	bytes, err := n.GetState().GetBytes("state")
	if err != nil {
		return nil, nil, maskAny(err)
	}
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 15}, "job neuron state: '%s'", bytes)

	switch string(bytes) {
	case "":
		// Create new impulse to process it asynchronously.
		go func(imp spec.Impulse) {
			n.GetState().SetBytes("state", []byte("in-progress"))

			neu, err := n.FactoryClient.NewCharacterNeuron()
			if err != nil {
				n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
				return
			}
			n.GetState().SetNeuron(neu)

			imp, _, err = imp.WalkThrough(neu)
			if err != nil {
				n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
				return
			}

			request, err := imp.GetState().GetBytes("request")
			if err != nil {
				n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
				return
			}
			n.GetState().SetBytes("response", request)
			n.GetState().SetBytes("state", []byte("finished"))
		}(imp)

		// Return the impulse ID to signal a job registration.
		imp.GetState().SetBytes("response", []byte(imp.GetState().GetObjectID()))

		return imp, nil, nil
	case "in-progress":
		// Return to keep waiting.
		return imp, nil, nil
	case "finished":
		// Return response.
		response, err := n.GetState().GetBytes("response")
		if err != nil {
			return nil, nil, maskAny(err)
		}

		imp.GetState().SetBytes("response", response)

		return imp, nil, nil
	}

	return nil, nil, maskAny(invalidImpulseStateError)
}
