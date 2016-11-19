// Package text implements spec.EndpointService and provides a way to feed
// neural networks with text input. To make Anna consume text, there is the text
// endpoint implemented through the network API.
package text

import (
	"io"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	textinputobject "github.com/the-anna-project/input/object/text"
	apispec "github.com/the-anna-project/spec/api"
	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new text endpoint service.
func New() servicespec.EndpointService {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	// address is the host:port representation based on the golang convention
	// for net.Listen to serve gRPC traffic.
	address      string
	bootOnce     sync.Once
	closer       chan struct{}
	gRPCServer   *grpc.Server
	metadata     map[string]string
	shutdownOnce sync.Once
}

func (s *service) Boot() {
	s.Service().Log().Line("func", "Boot")

	s.bootOnce.Do(func() {
		id, err := s.Service().ID().New()
		if err != nil {
			panic(err)
		}
		s.metadata = map[string]string{
			"id":   id,
			"kind": "text",
			"name": "endpoint",
			"type": "service",
		}

		s.bootOnce = sync.Once{}
		s.closer = make(chan struct{}, 1)
		s.gRPCServer = grpc.NewServer()
		s.shutdownOnce = sync.Once{}

		RegisterTextEndpointServer(s.gRPCServer, s)

		// Create the gRPC server. The Serve method below is returning listener
		// errors, if any. In case net.Listener.Accept is called and waits for
		// connections while the listener was closed, a net.OpError will be thrown.
		// For this case we only log errors from the fail channel in case the
		// server's closer was not closed yet.
		fail := make(chan error, 1)
		go func() {
			select {
			case <-s.closer:
			case err := <-fail:
				s.Service().Log().Line("msg", "%#v", maskAny(err))
			}
		}()

		s.Service().Log().Line("msg", "gRPC server listens on '%s'", s.address)
		listener, err := net.Listen("tcp", s.address)
		if err != nil {
			s.Service().Log().Line("error", maskAny(err))
		}
		err = s.gRPCServer.Serve(listener)
		if err != nil {
			fail <- maskAny(err)
		}
	})
}

func (s *service) DecodeResponse(textOutput objectspec.TextOutput) *StreamTextResponse {
	streamTextResponse := &StreamTextResponse{
		Code: apispec.CodeData,
		Data: &StreamTextResponseData{
			Output: textOutput.Output(),
		},
		Text: apispec.TextData,
	}

	return streamTextResponse
}

func (s *service) EncodeRequest(streamTextRequest *StreamTextRequest) (objectspec.TextInput, error) {
	textInputObject := textinputobject.New()
	textInputObject.SetEcho(streamTextRequest.Echo)
	//textInputObject.SetExpectation(streamTextRequest.Expectation)
	textInputObject.SetInput(streamTextRequest.Input)
	textInputObject.SetSessionID(streamTextRequest.SessionID)

	return textInputObject, nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetAddress(address string) {
	s.address = address
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}

func (s *service) Shutdown() {
	s.Service().Log().Line("func", "Shutdown")

	s.shutdownOnce.Do(func() {
		close(s.closer)

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			// Stop the gRPC server gracefully and wait some time for open
			// connections to be closed. Then force it to be stopped.
			go s.gRPCServer.GracefulStop()
			time.Sleep(3 * time.Second)
			s.gRPCServer.Stop()
			wg.Done()
		}()

		wg.Wait()
	})
}

func (s *service) StreamText(stream TextEndpoint_StreamTextServer) error {
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
			s.Service().Input().Text().Channel() <- textRequest
		}
	}()

	// Listen on the outout of the text interface and stream it back to the
	// client.
	go func() {
		for {
			select {
			case <-done:
				return
			case textOutput := <-s.Service().Output().Text().Channel():
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
