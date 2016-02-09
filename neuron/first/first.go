package firstneuron

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
	newStateConfig.ObjectType = common.ObjectType.FirstNeuron

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

	neu, err := n.GetState().GetNeuronByID(imp.GetObjectID())
	if state.IsNeuronNotFound(err) {
		n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 14}, "create job neuron")

		// Create new job neuron with the impulse ID.
		neu, err = n.FactoryClient.NewJobNeuron()
		if err != nil {
			return nil, nil, maskAny(err)
		}

		newStateConfig := state.DefaultConfig()
		newStateConfig.Bytes["state"] = []byte{}
		newStateConfig.ObjectID = imp.GetObjectID()
		newStateConfig.ObjectType = common.ObjectType.JobNeuron

		neu.SetState(state.NewState(newStateConfig))

		// Track new neuron.
		n.GetState().SetNeuron(neu)
	} else if err != nil {
		return nil, nil, maskAny(err)
	}

	return imp, neu, nil
}
