package spec

// Strategy implements a container for a sequence of actions to be carried
// around, e.g. by an impulse.
type Strategy interface {
	// GetCLGNames returns the ordered list of the strategy's CLG names.
	GetCLGNames() []string

	// GetRequestor returns the object type of the strategies requestor.
	GetRequestor() ObjectType

	// GetStringMap returns the strategy's storable information as string map.
	GetStringMap() (map[string]string, error)

	Object
}
