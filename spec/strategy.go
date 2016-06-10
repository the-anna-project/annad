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

	// Get Argument returns the strategy's configured argument. This can be any
	// arbitrary value. Note when the returned argument is not nil, the current
	// strategy is considered static. Thus IsStatic returns true.
	GetArgument() interface{}

	// GetRoot returnes the strategies configured Root. This in turn is a CLG. In
	// case the returned Root is not empty, the current strategy is not
	// considered static. Thus IsStatic returns false.
	GetRoot() CLG

	// GetNodes returnes the strategies configured Nodes. These are in turn
	// strategies themself. Note that Nodes can only be configured in case the
	// strategy is not static.
	GetNodes() []Strategy

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
