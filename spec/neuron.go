package spec

// Neuron implementations provide some sort of specialized business logic. They
// are used to make networks dynamic and functional. Neurons here should not be
// too much aligned with the biological equivalent. Recent implementations
// turned out to be way too complex and stood in the way of getting things
// done. Anyway, neurons can be used for very specific very tiny tasks though.
// At the end the question is what makes sense and what works out. Neurons
// simply represent a pragmatic way to separate concerns where appropriate.
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
