package spec

type Impulse interface {
	Copy() Impulse

	GetObjectID() ObjectID

	GetObjectType() ObjectType

	GetState(key string) (State, error)

	SetState(key string, state State)

	WalkThrough(neu Neuron) (Impulse, Neuron, error)
}
