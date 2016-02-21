package spec

type ObjectID string

type ObjectType string

// Object represents the interface for identification and persistence.
type Object interface {
	GetID() ObjectID
	GetType() ObjectType
}
