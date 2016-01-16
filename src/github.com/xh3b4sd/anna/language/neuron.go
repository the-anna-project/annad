package language

import (
	"github.com/xh3b4sd/anna/network"
)

func NewNeuron() network.Neuron {
	return neuron{
		State: network.NewState(),
	}
}

type neuron struct {
	Connections []network.Connection
	CreatedAt   time.Time `json:"create_at"`
	Impulses    []network.Impuls
	Networks    []network.Network
	State       network.State
}

func (n neuron) Age() time.Duration {
	return time.Since(lb.CreatedAt)
}

func (n neuron) Connection(neuron network.Neuron) (network.Connection, error) {
	return nil, nil
}

func (n neuron) Connections() ([]network.Connection, error) {
	return nil, nil
}

func (n neuron) Continue() {
}

func (n neuron) Impulses() ([]network.Impuls, error) {
}

func (n neuron) Load(State) {
	return n.State.Capture(n)
}

func (n neuron) Merge(dst, src network.Connection) (network.Connection, error) {
	return nil, nil
}

func (n neuron) Networks() []network.Network {
	return n.Networks
}

func (n neuron) Pause() {
}

func (n neuron) State() network.State {
	return n.State.Capture(n)
}
