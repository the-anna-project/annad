package spec

type ObjectID string

type ObjectType string

type Object interface {
	GetObjectID() ObjectID
	GetObjectType() ObjectType
	GetState() State
	SetState(state State)
}
