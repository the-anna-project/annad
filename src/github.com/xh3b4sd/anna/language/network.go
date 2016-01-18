package language

import (
	"github.com/xh3b4sd/anna/core"
)

func NewNetwork() core.Network {
	n := network{
		State: core.NewState(),
	}

	return n
}

type network struct {
	State core.State
}

func (n network) GetState() core.State {
	return n.State
}

func (n network) SetState(state core.State) {
	n.State = state
}

func (n network) Trigger(impulse core.Impulse) (core.Impulse, core.Connection, error) {
	return nil, nil, nil
}
