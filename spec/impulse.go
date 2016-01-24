package spec

type Impulse interface {
	GetObjectID() ObjectID

	GetObjectType() ObjectType

	GetState() State

	SetState(state State)

	WalkThrough(neu Neuron) (Impulse, Neuron, error)
}
