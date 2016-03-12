package spec

// Strategy implements a container for a sequence of actions to be carried
// around, e.g. by an impulse.
type Strategy interface {
	// ActionsToString returns the strategy's actions as comma separated string.
	ActionsToString() string

	// GetActions returns the ordered list of the strategy's action items.
	GetActions() []ObjectType

	// GetHashMap returns the strategies storable information as hash map.
	GetHashMap() map[string]string

	// GetRequestor returns the object type if the strategies requestor.
	GetRequestor() ObjectType

	Object
}
