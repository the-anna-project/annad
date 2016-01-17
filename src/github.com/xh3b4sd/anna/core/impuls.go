package core

import (
	"time"
)

type Impulse interface {
	SetState(state State)

	State() State

	// Track tracks location information of objects that the impuls passes. v can
	// either be Network, Neuron, or Connection.
	Track(v interface{}) error
}
