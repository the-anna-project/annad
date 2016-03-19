package spec

// Strategy implements a container for a sequence of actions to be carried
// around, e.g. by an impulse.
type Strategy interface {
	// ActionsToString returns the strategy's actions as comma separated string.
	// Note that this method is used to be store already created strategies using
	// descriptive keys for fast lookups.
	ActionsToString() string

	// GetActions returns the ordered list of the strategy's action items.
	GetActions() []ObjectType

	// GetHashMap returns the strategy's storable information as hash map.
	GetHashMap() map[string]string

	// GetRequestor returns the object type if the strategies requestor.
	GetRequestor() ObjectType

	Object
}
