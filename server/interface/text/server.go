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
	// ObjectType represents the object type of the text interface server object.
	// This is used e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "text-interface-server"
)

// ServerConfig represents the configuration used to create a new text
// interface object.
type ServerConfig struct {
	Log        spec.Log
	TextInput  chan spec.TextRequest
	TextOutput chan spec.TextResponse
}

// DefaultServerConfig provides a default configuration to create a new text
// interface object by best effort.
func DefaultServerConfig() ServerConfig {
	newConfig := ServerConfig{
		Log:        log.NewLog(log.DefaultConfig()),
		TextInput:  make(chan spec.TextRequest, 1000),
		TextOutput: make(chan spec.TextResponse, 1000),
	}

	return newConfig
}

// NewServer creates a new configured text interface object.
func NewServer(config ServerConfig) (api.TextInterfaceServer, error) {
	newServer := &server{
		ServerConfig: config,
		ID:           id.MustNew(),
		Mutex:        sync.Mutex{},
		Type:         ObjectType,
	}

	if newServer.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}

	newServer.Log.Register(newServer.GetType())

	return newServer, nil
}

type server struct {
	ServerConfig

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (s *server) DecodeResponse(textResponse spec.TextResponse) *api.StreamTextResponse {
	streamTextResponse := &api.StreamTextResponse{
		Code: api.CodeData,
		Data: &api.StreamTextResponseData{
			Output: textResponse.GetOutput(),
		},
		Text: api.TextData,
	}

	return streamTextResponse
}

func (s *server) EncodeRequest(streamTextRequest *api.StreamTextRequest) (spec.TextRequest, error) {
	textRequestConfig := api.DefaultTextRequestConfig()
	textRequestConfig.Echo = streamTextRequest.Echo
	//newTextRequestConfig.ExpectationRequest = expectationRequest
	textRequestConfig.Input = streamTextRequest.Input
	textRequestConfig.SessionID = streamTextRequest.SessionID
	textRequest, err := api.NewTextRequest(textRequestConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return textRequest, nil
}

func (s *server) StreamText(stream api.TextInterface_StreamTextServer) error {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call StreamText")

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

			textRequest, err := s.EncodeRequest(streamTextRequest)
			if err != nil {
				fail <- maskAny(err)
				return
			}
			s.TextInput <- textRequest
		}
	}()

	// Listen on the outout of the text interface and stream it back to the
	// client.
	go func() {
		for {
			textResponse := <-s.TextOutput
			streamTextResponse := s.DecodeResponse(textResponse)
			err := stream.Send(streamTextResponse)
			if err != nil {
				fail <- maskAny(err)
				return
			}
		}
	}()

	for {
		select {
		case <-stream.Context().Done():
			return maskAny(stream.Context().Err())
		case <-done:
			return nil
		case err := <-fail:
			return maskAny(err)
		}
	}
}
