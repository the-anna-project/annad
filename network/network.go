// Package network implements spec.Network. Any network implementation is a
// collection of neurons.
package network

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
	newStateConfig.ObjectType = common.ObjectType.Network

	newConfig := Config{
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		State:         state.NewState(newStateConfig),
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

	Mutex sync.Mutex `json:"-"`
}

func (n *network) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	n.Log.V(11).Debugf("call Network.Trigger")

	// Get first neuron.
	var err error
	var neu spec.Neuron
	neurons := n.GetState().GetNeurons()
	if len(neurons) == 0 {
		n.Log.V(12).Debugf("create first neuron")

		neu, err = n.FactoryClient.NewFirstNeuron()
		if err != nil {
			return nil, maskAny(err)
		}

		n.GetState().SetNeuron(neu)
		n.GetState().SetBytes(string(common.ObjectType.FirstNeuron), []byte(neu.GetObjectID()))
	} else {
		n.Log.V(12).Debugf("forwarding impulse to first neuron")

		neuronID, err := n.GetState().GetBytes(string(common.ObjectType.FirstNeuron))
		if err != nil {
			return nil, maskAny(err)
		}
		neu, err = n.GetState().GetNeuronByID(spec.ObjectID(neuronID))
		if err != nil {
			return nil, maskAny(err)
		}
	}

	imp, _, err = imp.WalkThrough(neu)
	if err != nil {
		return nil, maskAny(err)
	}

	return imp, nil
}
