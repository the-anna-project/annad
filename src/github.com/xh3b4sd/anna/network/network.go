package network

import (
	"time"
)

type Network interface {
	Add(neuron Neuron) error

	Age() time.Time

	Connections() ([]Connection, error)

	// Merge merges dst with src by best effort and returns the result. Result
	// can either be Neuron or Network.
	Merge(dst, src interface{}) (interface{}, error)

	Neurons() ([]Neuron, error)

	Remove(neuron Neuron) error
}
