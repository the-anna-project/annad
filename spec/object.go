package spec

type ObjectID string

type ObjectType string

type Object interface {
	GetObjectID() ObjectID
	GetObjectType() ObjectType
	GetState(key string) (State, error)
	SetState(key string, state State)
}
