package network

import (
	"time"
)

type Impuls interface {
	Age() time.Duration

	Connections() ([]Connection, error)

	// Continue ends the idle state of the impulse caused by a call to Pause. It
	// is likely the way the impulse behaves after continuation changes
	// unpredictably. During the timespan between pausing and continuing, the
	// network, the impuls is going through, will probably have changed and thus
	// influence the way the impuls would have been going through it originally.
	Continue()

	Load(state State)

	Networks() ([]Network, error)

	Neurons() ([]Neuron, error)

	// Pause stops the activity of an impuls to keep it in its current state.
	// This causes no further interaction to happen until Continue is called. It
	// is important to pause a impuls before capturing its state to ensure
	// reproducable snapshots.
	Pause()

	String() string

	// Track tracks location information of objects that the impuls passes. v can
	// either be Network, Neuron, or Connection.
	Track(v interface{}) error

	State() State
}
