// Package textinput provides a simple service for receiving text input
// requests.
package textinput

import (
	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new text input service.
func New() servicespec.InputService {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	channel  chan objectspec.TextInput
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
		"name": "input",
		"type": "service",
	}

	s.channel = make(chan objectspec.TextInput, 1000)
}

func (s *service) Channel() chan objectspec.TextInput {
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
