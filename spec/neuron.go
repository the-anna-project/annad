package spec

import (
	"encoding/json"
)

type Neuron interface {
	json.Unmarshaler

	Object

	// Trigger starts processing of the given impulse within the current neuron.
	// Magic happens here based on the implemented behaviour. There is always an
	// Impulse returned, as long as there is no error. The returned neuron might
	// be nil in case the current neuron decided to not forward the impulse to
	// any further neuron. This indicates the end of the impulses walk through.
	Trigger(impulse Impulse) (Impulse, Neuron, error)
}
