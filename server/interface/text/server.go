package text

import (
	"io"

	"github.com/xh3b4sd/anna/object/networkresponse"
	objectspec "github.com/xh3b4sd/anna/object/spec"
	"github.com/xh3b4sd/anna/object/textinput"
	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/id"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// ServerConfig represents the configuration used to create a new text
// interface object.
type ServerConfig struct {
	ServiceCollection servicespec.Collection
}

// DefaultServerConfig provides a default configuration to create a new text
// interface object by best effort.
func DefaultServerConfig() ServerConfig {
	newConfig := ServerConfig{
		ServiceCollection: service.MustNewCollection(),
	}

	return newConfig
}

// NewServer creates a new configured text interface object.
func NewServer(config ServerConfig) (TextInterfaceServer, error) {
	newServer := &server{
		ServerConfig: config,

		Metadata: map[string]string{
			"id":   id.MustNewID(),
			"name": "text",
			"type": "endpoint",
		},
	}

	if newServer.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	return newServer, nil
}

type server struct {
	ServerConfig

	Metadata map[string]string
}

func (s *server) DecodeResponse(textOutput objectspec.TextOutput) *StreamTextResponse {
	streamTextResponse := &StreamTextResponse{
		Code: networkresponse.CodeData,
		Data: &StreamTextResponseData{
			Output: textOutput.GetOutput(),
		},
		Text: networkresponse.TextData,
	}

	return streamTextResponse
}

func (s *server) EncodeRequest(streamTextRequest *StreamTextRequest) (objectspec.TextInput, error) {
	textInputConfig := textinput.DefaultConfig()
	textInputConfig.Echo = streamTextRequest.Echo
	//textInputConfig.Expectation = streamTextRequest.Expectation
	textInputConfig.Input = streamTextRequest.Input
	textInputConfig.SessionID = streamTextRequest.SessionID
	textInput, err := textinput.New(textInputConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return textInput, nil
}

func (s *server) StreamText(stream TextInterface_StreamTextServer) error {
	s.Service().Log().Line("func", "StreamText")

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
			s.Service().TextInput().Channel() <- textRequest
		}
	}()

	// Listen on the outout of the text interface and stream it back to the
	// client.
	go func() {
		for {
			select {
			case <-done:
				return
			case textOutput := <-s.Service().TextOutput().Channel():
				streamTextResponse := s.DecodeResponse(textOutput)
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
			return maskAny(err)
		}
	}
}
