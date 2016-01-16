package core

type Impulse interface {
	SetState(state State)

	GetState() State

	// Track tracks location information of objects that the impuls passes. v can
	// either be Network, Neuron, or Connection.
	Track(v interface{}) error
}
