package text

import (
	"io"
	"sync"

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
	TextInput  chan spec.TextRequest
	TextOutput chan spec.TextResponse
}

// DefaultInterfaceConfig provides a default configuration to create a new text
// interface object by best effort.
func DefaultInterfaceConfig() InterfaceConfig {
	newConfig := InterfaceConfig{
		Log:        log.NewLog(log.DefaultConfig()),
		TextInput:  make(chan spec.TextRequest, 1000),
		TextOutput: make(chan spec.TextResponse, 1000),
	}

	return newConfig
}

// NewInterface creates a new configured text interface object.
func NewInterface(config InterfaceConfig) (api.TextInterfaceServer, error) {
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

func (i *tinterface) StreamText(stream api.TextInterface_StreamTextServer) error {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call StreamText")

	done := make(chan struct{}, 1)
	fail := make(chan error, 1)

	// Listen on the server input stream and forward it to the neural network.
	go func() {
		for {
			streamTextRequest, err := stream.Recv()
			if err == io.EOF {
				// The stream ended. We broadcast to all goroutines by closing the done
				// channel.
				close(done)
				return
			} else if err != nil {
				fail <- maskAny(err)
				return
			}

			textRequestConfig := api.DefaultTextRequestConfig()
			//newTextRequestConfig.ExpectationRequest = expectationRequest
			textRequestConfig.Input = streamTextRequest.Input
			//newTextRequestConfig.SessionID = a.SessionID
			textRequest, err := api.NewTextRequest(textRequestConfig)
			if err != nil {
				fail <- maskAny(err)
				return
			}

			i.TextInput <- textRequest
		}
	}()

	// Listen on the outout of the text interface and stream it back to the
	// client.
	go func() {
		for {
			select {
			case <-done:
				return
			case textResponse := <-i.TextOutput:
				streamTextResponse := &api.StreamTextResponse{
					Output: textResponse.GetOutput(),
				}
				err := stream.Send(streamTextResponse)
				if err != nil {
					fail <- maskAny(err)
					return
				}
			}
		}
	}()

	for {
		select {
		case <-stream.Context().Done():
			close(done)
			return maskAny(stream.Context().Err())
		case <-done:
			return nil
		case err := <-fail:
			if err != nil {
				close(done)
				return maskAny(err)
			}
		}
	}
}
