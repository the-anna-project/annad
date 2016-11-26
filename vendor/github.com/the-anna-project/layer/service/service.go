// Package service implements a service to manage connections inside network
// layers.
package service

import servicespec "github.com/the-anna-project/spec/service"

// New creates a new layer service.
func New() servicespec.LayerService {
	return &service{
		// Dependencies.
		serviceCollection: nil,

		// Settings.
		kind:     "",
		metadata: map[string]string{},
	}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	kind     string
	metadata map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"kind": s.kind,
		"name": "storage",
		"type": "service",
	}
}

// TODO
func (s *service) CreateConnection(peerA, peerB string) error {
	s.Service().Log().Line("func", "CreateConnection")
	return nil
}

// TODO
func (s *service) DeleteConnection(peerA, peerB string) error {
	s.Service().Log().Line("func", "DeleteConnection")
	return nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetKind(kind string) {
	s.kind = kind
}

func (s *service) SetServiceCollection(serviceCollection servicespec.ServiceCollection) {
	s.serviceCollection = serviceCollection
}
