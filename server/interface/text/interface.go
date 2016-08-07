package text

import (
	"sync"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeTextInterface represents the object type of the text interface
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeTextInterface spec.ObjectType = "text-interface"
)

// InterfaceConfig represents the configuration used to create a new text
// interface object.
type InterfaceConfig struct {
	Log        spec.Log
	TextInput  chan api.TextRequest
	TextOutput chan api.TextResponse
}

// DefaultInterfaceConfig provides a default configuration to create a new text
// interface object by best effort.
func DefaultInterfaceConfig() InterfaceConfig {
	newConfig := InterfaceConfig{
		Log:        log.NewLog(log.DefaultConfig()),
		TextInput:  make(chan api.TextRequest, 1000),
		TextOutput: make(chan api.TextResponse, 1000),
	}

	return newConfig
}

// NewInterface creates a new configured text interface object.
func NewInterface(config InterfaceConfig) (spec.TextInterface, error) {
	newInterface := &tinterface{
		InterfaceConfig: config,
		ID:              id.MustNew(),
		Mutex:           sync.Mutex{},
		Type:            spec.ObjectType(ObjectTypeTextInterface),
	}

	if newInterface.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}

	newInterface.Log.Register(newInterface.GetType())

	return newInterface, nil
}

// tinterface is not named interface because this is a reserved key in golang.
type tinterface struct {
	InterfaceConfig

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (i *tinterface) StreamText(ctx context.Context, in chan api.TextRequest, out chan api.TextResponse) error {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call StreamText")

	// Start processing the text request through the text input channel.
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case textRequest <- in:
				i.TextInput <- textRequest
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Error()
		case textResponse := <-i.TextOutput:
			out <- textResponse
		}
	}
}
