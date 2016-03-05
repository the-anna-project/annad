package spec

// Strategy implements a container for a sequence of actions to be carried
// around, e.g. by an impulse.
type Strategy interface {
	// String returns the string representation of the strategy's action
	// sequence, e.g. "one,two,three".
	String() string

	// GetActions returns the ordered list of the strategy's action items.
	GetActions() []ObjectType

	Object
}
