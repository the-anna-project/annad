package spec

// ObjectID represents an object's ID. This is a unique identifier. Here it is
// most likely a simple hexadecimal encoded string.
type ObjectID string

// ObjectType represents an object's type. E.g. "neuron", "network", or
// "redis-storage". The object type should be defined within the scope of the
// object that is implementing it. Please prevent a global list of all possible
// object types.
type ObjectType string

// Object represents the interface for identification. Each object using a
// logger or being involved in tasks of the neural network should implement
// Object so others know what object they are dealing with.
type Object interface {
	// GetID returns the objects's object ID.
	GetID() string

	// GetType returns the object's object type.
	GetType() ObjectType
}
