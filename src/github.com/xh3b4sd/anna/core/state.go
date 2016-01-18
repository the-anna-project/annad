package core

import (
	"time"
)

type State interface {
	GetAge() time.Duration

	GetConnections() ([]Connection, error)

	GetImpulses() ([]Impulse, error)

	MarshalJSON() ([]byte, error)

	GetNetworks() ([]Network, error)

	GetNeurons() ([]Neuron, error)

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
	Connections []Connection `json:"connections"`
	CreatedAt   time.Time    `json:"created_at"`
	Impulses    []Impulse    `json:"impulses"`
	Networks    []Network    `json:"networks"`
	Neurons     []Neuron     `json:"neurons"`
}

func (s state) GetAge() time.Duration {
	return time.Since(s.CreatedAt)
}

func (s state) GetConnections() ([]Connection, error) {
	return nil, nil
}

func (s state) GetImpulses() ([]Impulse, error) {
	return nil, nil
}

func (s state) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (s state) GetNetworks() ([]Network, error) {
	return nil, nil
}

func (s state) GetNeurons() ([]Neuron, error) {
	return nil, nil
}

func (s state) SetConnection(connection Connection) error {
	return nil
}

func (s state) SetImpulse(impulse Impulse) error {
	return nil
}

func (s state) SetNetwork(network Network) error {
	return nil
}

func (s state) SetNeuron(neuron Neuron) error {
	return nil
}

func (s state) UnmarshalJSON([]byte) error {
	return nil
}
