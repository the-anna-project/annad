package network

import (
	"time"
)

type Network interface {
	Age() time.Duration

	Connections() ([]Connection, error)

	Gateway() Gateway

	Load(state State)

	// Merge merges dst with src by best effort and returns the result. Result
	// can either be Neuron or Network.
	Merge(dst, src interface{}) (interface{}, error)

	Neurons() ([]Neuron, error)

	State() State
}
