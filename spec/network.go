package spec

type Network interface {
	Copy() Network

	GetObjectID() ObjectID

	GetObjectType() ObjectType

	GetState(key string) (State, error)

	SetState(key string, state State)

	Trigger(imp Impulse) (Impulse, error)
}
