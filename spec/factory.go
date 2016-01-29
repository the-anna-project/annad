package spec

type Factory interface {
	NewCore() (Core, error)

	NewImpulse() (Impulse, error)

	NewCharacterNeuron() (Neuron, error)

	NewFirstNeuron() (Neuron, error)

	NewJobNeuron() (Neuron, error)

	NewNetwork() (Network, error)

	NewState(objectType ObjectType) (State, error)
}
