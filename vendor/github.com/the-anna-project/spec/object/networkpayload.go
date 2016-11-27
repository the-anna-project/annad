package object

import (
	"encoding/json"
	"reflect"
)

// NetworkPayload represents the data container carried around within the
// neural network.
type NetworkPayload interface {
	// GetArgs returns the arguments of the current network payload.
	GetArgs() []reflect.Value

	// GetCLGInput returns a list of arguments intended to be provided as input
	// for a CLG's execution. The list of arguments exists of the arguments
	// configured to the network payload and the context configured to the network
	// payload. Note that the context is always the first argument in the list.
	GetCLGInput() []reflect.Value

	// GetContext returns the context of the current network payload.
	GetContext() Context

	// GetArgs returns the destination of the current network payload, which must
	// be the ID of a CLG registered within the neural network.
	GetDestination() string

	// GetID returns the object ID of the current network payload.
	GetID() string

	// GetArgs returns the sources of the current network payload, which must be
	// the ID of a CLG registered within the neural network. One allowed exception
	// is the very first source of the very first network payload, which is
	// created within the network when user input is received to forward it to the
	// input CLG.
	GetSources() []string

	json.Marshaler
	json.Unmarshaler

	// SetArgs sets the arguments of the current network payload.
	SetArgs(args []reflect.Value)

	// String returns the concatenated string representations of the currently
	// configured arguments.
	String() string
}
