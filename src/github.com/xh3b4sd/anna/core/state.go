package core

import (
	"time"
)

type State interface {
	Age() time.Duration

	Connections() ([]Connection, error)

	Impulses() ([]Impuls, error)

	MarshalJSON() ([]byte, error)

	Networks() ([]Network, error)

	Neurons() ([]Neuron, error)

	SetConnection(connection Connection) error

	SetImpulse(impulse Impulse) error

	SetNetwork(network Network) error

	SetNeuron(neuron Neuron) error

	UnmarshalJSON([]byte) error
}

func NewState() State {
	return state{
		CreatedAt: time.Now(),
	}
}

type state struct {
	CreatedAt time.Time `json:"created_at"`
}

func (s State) Age() time.Duration {
	return time.Since(n.CreatedAt)
}

func (s state) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (s state) UnmarshalJSON([]byte) error {
	return nil
}
