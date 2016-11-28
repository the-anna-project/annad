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
		closer:   make(chan struct{}, 1),
		kind:     "",
		metadata: map[string]string{},
	}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	closer   chan struct{}
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
		"name": "layer",
		"type": "service",
	}
}

func (s *service) CreateConnection(peerA, peerB string) error {
	s.Service().Log().Line("func", "CreateConnection")

	actions := []func(canceler <-chan struct{}) error{
		func(canceler <-chan struct{}) error {
			err := s.Service().Peer().Create(peerA, peerB)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
		func(canceler <-chan struct{}) error {
			err := s.Service().Connection().Create(peerA, peerB)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
	}

	executeConfig := s.Service().Worker().ExecuteConfig()
	executeConfig.SetActions(actions)
	executeConfig.SetCanceler(s.closer)
	executeConfig.SetNumWorkers(2)
	err := s.Service().Worker().Execute(executeConfig)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) DeleteConnection(peerA, peerB string) error {
	s.Service().Log().Line("func", "DeleteConnection")

	// TODO prefix keys with layer information (kind)

	actions := []func(canceler <-chan struct{}) error{
		func(canceler <-chan struct{}) error {
			err := s.Service().Peer().Delete(peerA)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
		func(canceler <-chan struct{}) error {
			err := s.Service().Connection().Delete(peerA, peerB)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
	}

	executeConfig := s.Service().Worker().ExecuteConfig()
	executeConfig.SetActions(actions)
	executeConfig.SetCanceler(s.closer)
	executeConfig.SetNumWorkers(2)
	err := s.Service().Worker().Execute(executeConfig)
	if err != nil {
		return maskAny(err)
	}

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
