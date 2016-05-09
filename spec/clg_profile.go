package spec

import (
	"encoding/json"
	"reflect"
)

// CLGProfile contains information of a certain CLG.
type CLGProfile interface {
	// Equals checks whether the current CLG profile is equal to the given one.
	Equals(CLGProfile) bool

	// GetBody returns the profile's implemented CLG method body as string
	// representation.
	GetBody() string

	// GetHash returns the checksum of the profile's body.
	GetHash() string

	// GetInputs returns the profile's implemented CLG method input parameter
	// types.
	GetInputs() []reflect.Kind

	// GetName returns the name of the CLG this profile is associated with.
	GetName() string

	// GetOutputs returns the profile's implemented CLG method output parameter
	// types.
	GetOutputs() []reflect.Kind

	json.Marshaler

	json.Unmarshaler
}
