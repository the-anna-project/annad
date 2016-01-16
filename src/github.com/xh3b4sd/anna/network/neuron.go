package network

import (
	"time"
)

type Neuron interface {
	Age() time.Duration

	Connection(neuron Neuron) (Connection, error)

	Connections() ([]Connection, error)

	// Continue ends the idle state of the neuron caused by a call to Pause. It
	// is likely the way the neuron behaves after continuation changes
	// unpredictably. During the timespan between pausing and continuing, the
	// network, the neuron is interacting with, will probably have changed and
	// thus influence the way the neuron would have been behaving originally.
	Continue()

	Impulses() ([]Impuls, error)

	Load(state State)

	Merge(dst, src Connection) (Connection, error)

	Networks() ([]Network, error)

	// Pause stops the activity of an neuron to keep it in its current state.
	// This causes no further interaction to happen until Continue is called. It
	// is important to pause a neuron before capturing its state to ensure
	// reproducable snapshots.
	Pause()

	State() State
}
