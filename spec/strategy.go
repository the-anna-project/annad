package spec

import (
	"encoding/json"
)

// Strategy implements a container for a sequence of CLG names to be carried
// around, e.g. by an impulse.
type Strategy interface {
	// GetCLGNames returns the ordered list of the strategy's CLG names.
	GetCLGNames() []string

	// GetRequestor returns the object type of the strategies requestor.
	GetRequestor() ObjectType

	json.Marshaler

	Object

	json.Unmarshaler
}
