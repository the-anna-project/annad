package spec

type Neuron interface {
	Object

	// Trigger starts processing of the given impulse within the current neuron.
	// Magic happens here based on the implemented behavior. The received impulse
	// is returned, but maybe modified. If there is an error, Impulse is nil. The
	// returned Neuron might be nil in case the current neuron decided to not
	// forward the impulse to any further neuron. This indicates the end of the
	// impulses walk through of the neurons sub tree.
	Trigger(impulse Impulse) (Impulse, Neuron, error)
}
