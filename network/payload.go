package network

import (
	"reflect"
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

// PayloadConfig represents the configuration used to create a new network
// payload object.
type PayloadConfig struct {
	// Settings.

	// Args represents the arguments intended to be used for the requested CLG
	// execution, or the output values being calculated during the requested CLG
	// execution. By convention this list must always contain a context as first
	// argument, otherwise the network payload is considered invalid. In this
	// case all calls to Validate would throw errors which can be asserted using
	// IsInvalidInterface. For more information about the context being used here
	// see https://godoc.org/golang.org/x/net/context.
	Args []reflect.Value

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

// DefaultPayloadConfig provides a default configuration to create a new
// network payload object by best effort.
func DefaultPayloadConfig() PayloadConfig {
	newConfig := PayloadConfig{
		Args:        nil,
		Destination: "",
		Sources:     nil,
	}

	return newConfig
}

// NewPayload creates a new configured network payload object.
func NewPayload(config PayloadConfig) (spec.NetworkPayload, error) {
	newPayload := &payload{
		PayloadConfig: config,

		ID: id.MustNew(),
	}

	return newPayload, nil
}

type payload struct {
	PayloadConfig

	ID spec.ObjectID
}

func (p payload) GetArgs() []reflect.Value {
	return p.Args
}

func (p payload) GetContext() (conext.Context, error) {
	if len(p.Args) < 1 {
		return nil, maskAnyf(invalidConfigError, "arguments must have context")
	}
	ctx, ok := p.Args[0].Interface().(context.Context)
	if !ok {
		return nil, maskAnyf(invalidInterfaceError, "arguments must have context")
	}

	return ctx, nil
}

func (p payload) GetDestination() spec.ObjectID {
	return p.Destination
}

func (p payload) GetID() spec.ObjectID {
	return p.ID
}

func (p payload) GetSources() []spec.ObjectID {
	return p.Sources
}

func (p payload) Validate() error {
	// Check if the payload has invalid properties.
	if p.Args == nil {
		return maskAnyf(invalidConfigError, "arguments must not be empty")
	}
	if p.Destination == "" {
		return maskAnyf(invalidConfigError, "destination must not be empty")
	}
	if len(p.Sources) < 1 {
		return maskAnyf(invalidConfigError, "sources must not be empty")
	}

	// Check if the payload has an context as first argument.
	_, err := p.GetContext()
	if err != nil {
		return maskAny(err)
	}

	return nil
}
