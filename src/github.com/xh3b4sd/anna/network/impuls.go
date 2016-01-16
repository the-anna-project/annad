package network

import (
	"time"
)

type Impuls interface {
	Age() time.Time

	Connections() ([]Connection, error)

	Networks() ([]Network, error)

	Neurons() ([]Neuron, error)

	String() string

	// Track tracks location information of objects that the impuls passes. v can
	// either be Network, Neuron, or Connection.
	Track(v interface{}) error
}
