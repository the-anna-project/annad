// Package text provides a simple service for receiving text output.
package text

import (
	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new text output service.
func New() servicespec.OutputService {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	channel  chan objectspec.TextOutput
	metadata map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"kind": "text",
		"name": "output",
		"type": "service",
	}

	s.channel = make(chan objectspec.TextOutput, 1000)
}

func (s *service) Channel() chan objectspec.TextOutput {
	return s.channel
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(serviceCollection servicespec.ServiceCollection) {
	s.serviceCollection = serviceCollection
}
