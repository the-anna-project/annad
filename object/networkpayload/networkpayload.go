package networkpayload

import (
	"reflect"

	"github.com/the-anna-project/annad/object/context"
	objectspec "github.com/the-anna-project/spec/object"
)

// Config represents the configuration used to create a new
// network payload object.
type Config struct {
	// Settings.

	// Args represents the arguments intended to be used for the requested CLG
	// execution, or the output values being calculated during the requested CLG
	// execution.
	Args []reflect.Value

	Context objectspec.Context

	// Destination represents the object ID of the CLG receiving the current
	// network payload.
	Destination string

	// Sources represents the object IDs of the CLGs being involved providing the
	// current network payload. In fact, a network payload can only be sent by
	// one CLG. Reason for this being a list is the merge of network payloads
	// until a CLG's interface is satisfied. That way multiple CLG's can request
	// another CLG even if they are not able to satisfy the interface of the
	// requested CLG on their own. This gives the neural network an opportunity
	// to learn to combine CLGs as desired.
	Sources []string
}

// DefaultConfig provides a default configuration to create a new
// network payload object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Args:        nil,
		Context:     context.MustNew(),
		Destination: "",
		Sources:     nil,
	}

	return newConfig
}

// New creates a new configured network payload object.
func New(config Config) (objectspec.NetworkPayload, error) {
	newObject := &networkPayload{
		Config: config,

		ID: "",
	}

	return newObject, nil
}

// MustNew creates either a new default configured network payload
// object, or panics.
func MustNew() objectspec.NetworkPayload {
	newObject, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newObject
}

type networkPayload struct {
	Config

	ID string
}

func (np *networkPayload) GetArgs() []reflect.Value {
	return np.Args
}

func (np *networkPayload) GetContext() objectspec.Context {
	return np.Context
}

func (np *networkPayload) GetCLGInput() []reflect.Value {
	return append([]reflect.Value{reflect.ValueOf(np.GetContext())}, np.GetArgs()...)
}

func (np *networkPayload) GetDestination() string {
	return np.Destination
}

func (np *networkPayload) GetID() string {
	return np.ID
}

func (np *networkPayload) GetSources() []string {
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
