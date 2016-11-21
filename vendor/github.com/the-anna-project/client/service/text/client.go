package text

import (
	"sync"

	"golang.org/x/net/context"

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

	// address is the host:port representation based on the golang convention for
	// net.Listen to serve gRPC traffic.
	address      string
	cancel       func()
	context      context.Context
	metadata     map[string]string
	shutdownOnce sync.Once
}

func (s *service) Boot() {
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

	s.context = context.Background()
	s.streamText()
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetAddress(address string) {
	s.address = address
}

func (s *service) SetServiceCollection(serviceCollection servicespec.ServiceCollection) {
	s.serviceCollection = serviceCollection
}

func (s *service) Shutdown() {
	s.shutdownOnce.Do(func() {
		s.cancel()
	})
}
