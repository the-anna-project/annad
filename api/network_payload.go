package api

import (
	"reflect"

	"github.com/xh3b4sd/anna/context"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

// NetworkPayloadConfig represents the configuration used to create a new
// network payload object.
type NetworkPayloadConfig struct {
	// Settings.

	// Args represents the arguments intended to be used for the requested CLG
	// execution, or the output values being calculated during the requested CLG
	// execution.
	Args []reflect.Value

	Context spec.Context

	// Destination represents the object ID of the CLG receiving the current
	// network payload.
	Destination spec.ObjectID

	// Sources represents the object IDs of the CLGs being involved providing the
	// current network payload. In fact, a network payload can only be sent by
	// one CLG. Reason for this being a list is the merge of network payloads
	// until a CLG's interface is satisfied. That way multiple CLG's can request
	// another CLG even if they are not able to satisfy the interface of the
	// requested CLG on their own. This gives the neural network an opportunity
	// to learn to combine CLGs as desired.
	Sources []spec.ObjectID
}

// DefaultNetworkPayloadConfig provides a default configuration to create a new
// network payload object by best effort.
func DefaultNetworkPayloadConfig() NetworkPayloadConfig {
	newConfig := NetworkPayloadConfig{
		Args:        nil,
		Context:     context.MustNew(),
		Destination: "",
		Sources:     nil,
	}

	return newConfig
}

// NewNetworkPayload creates a new configured network payload object.
func NewNetworkPayload(config NetworkPayloadConfig) (spec.NetworkPayload, error) {
	newNetworkPayload := &networkPayload{
		NetworkPayloadConfig: config,

		ID: id.MustNew(),
	}

	return newNetworkPayload, nil
}

// MustNewNetworkPayload creates either a new default configured network payload
// object, or panics.
func MustNewNetworkPayload() spec.NetworkPayload {
	newNetworkPayload, err := NewNetworkPayload(DefaultNetworkPayloadConfig())
	if err != nil {
		panic(err)
	}

	return newNetworkPayload
}

type networkPayload struct {
	NetworkPayloadConfig

	ID spec.ObjectID
}

func (np *networkPayload) GetArgs() []reflect.Value {
	return append([]reflect.Value{reflect.ValueOf(np.GetContext())}, np.Args...)
}

func (np *networkPayload) GetContext() spec.Context {
	return np.Context
}

func (np *networkPayload) GetDestination() spec.ObjectID {
	return np.Destination
}

func (np *networkPayload) GetID() spec.ObjectID {
	return np.ID
}

func (np *networkPayload) GetSources() []spec.ObjectID {
	return np.Sources
}

func (np *networkPayload) SetArgs(args []reflect.Value) {
	np.Args = args
}

func (np *networkPayload) String() string {
	var s string

	for _, v := range np.GetArgs() {
		s += v.String()
	}

	return s
}

func (np *networkPayload) Validate() error {
	if np.GetContext() == nil {
		return maskAnyf(invalidConfigError, "context must not be empty")
	}
	if np.GetDestination() == "" {
		return maskAnyf(invalidConfigError, "destination must not be empty")
	}
	if len(np.GetSources()) < 1 {
		return maskAnyf(invalidConfigError, "sources must not be empty")
	}

	return nil
}
