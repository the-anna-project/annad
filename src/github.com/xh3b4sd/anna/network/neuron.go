package network

import (
	"time"
)

type Neuron interface {
	Age() time.Time

	Connection(neuron Neuron) (Connection, error)

	Connections() ([]Connection, error)

	Impulses() ([]Impuls, error)

	Load(state State)

	Merge(dst, src Connection) (Connection, error)

	Networks() ([]Network, error)

	State() State
}
