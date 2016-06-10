package spec

import (
	"encoding/json"
	"reflect"
)

// Strategy represents an executable and permutable CLG tree structure.
type Strategy interface {
	// Execute executes the current strategy. In case of a static strategy, the
	// strategy's configured argument will simply be returned without any further
	// magic. In case of none static strategies, the strategy's configured Root
	// will be executed by applying any configured Nodes. Note that a strategy
	// can only be executed in case it is valid.
	Execute() ([]reflect.Value, error)

	// GetOutputs returns the strategies output interface. That is a list of
	// types the strategy's Root CLG implements.
	GetOutputs() ([]reflect.Type, error)

	// IsStatic describes whether the strategy has an Root configured or not.
	// Note a Root here is an CLG. A static strategy has no Root configured. Thus
	// it is not executable and will always return the same result when executed,
	// because of its configured argument.
	IsStatic() bool

	json.Marshaler

	Object

	json.Unmarshaler

	// RemoveNode removes the node under the index given by indizes. The index is
	// multi dimensional. The first index within indizes applies to the current
	// strategy. The second index within indizes applies to one of the strategy's
	// Nodes, and so forth. Note removing nodes does only work for none static
	// strategies.
	RemoveNode(indizes []int) error

	// SetNode sets the given node under the index given by indizes. The index is
	// multi dimensional. The first index within indizes applies to the current
	// strategy. The second index within indizes applies to one of the strategy's
	// Nodes, and so forth. Note removing nodes does only work for none static
	// strategies.
	SetNode(indizes []int, node Strategy) error

	// Validate throws an error if the strategy is not valid. Reasons might be
	// not matching interfaces between Root and Nodes, or circular Strategies.
	// Note a static strategy is always valid.
	Validate() error
}
