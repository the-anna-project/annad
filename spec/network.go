package spec

type Network interface {
	GetObjectID() ObjectID

	GetObjectType() ObjectType

	GetState() State

	SetState(state State)

	Trigger(imp Impulse) (Impulse, error)
}
