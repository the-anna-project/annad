package language

import (
	"github.com/xh3b4sd/anna/core"
)

func NewNeuron() core.Neuron {
	return neuron{
		State: core.NewState(),
	}
}

type neuron struct {
	State core.State `json:"state"`
}

func (n neuron) GetState() core.State {
	return n.State
}

func (n neuron) SetState(state core.State) {
	n.State = state
}

func (n neuron) Trigger(impulse core.Impulse) (core.Impulse, core.Connection, error) {
	return nil, nil, nil
}
