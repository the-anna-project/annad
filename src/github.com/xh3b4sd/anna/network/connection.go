package network

import (
	"time"
)

type Connection interface {
	Age() time.Time

	Impulses() ([]Impuls, error)

	Load(state State)

	Merge(dst, src Impuls) (Impuls, error)

	Networks() ([]Network, error)

	Neurons() ([]Neuron, error)

	Trigger(impuls Impuls) (Impuls, error)

	State() State
}
