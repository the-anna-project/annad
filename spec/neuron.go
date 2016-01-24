package spec

type Neuron interface {
	GetObjectID() ObjectID

	// GetNetwork fetches the neurons related network from the internal state.
	// Since State.GetNetworks is very generic, this call will always only return
	// one network. This is due to the fact that a neuron can by design only
	// relate to one network.
	GetNetwork() (Network, error)

	GetObjectType() ObjectType

	GetState() State

	SetState(state State)

	// Trigger starts processing of the given impulse within the current neuron.
	// Magic happens here based on the implemented behaviour. There is always an
	// Impulse returned, as long as there is no error. The returned neuron might
	// be nil in case the current neuron decided to not forward the impulse to
	// any further neuron. This indicates the end of the impulses walk through.
	Trigger(impulse Impulse) (Impulse, Neuron, error)
}
