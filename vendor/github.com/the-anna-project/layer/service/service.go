// Package service implements a service to manage connections inside network
// layers.
package service

import (
	"fmt"

	servicespec "github.com/the-anna-project/spec/service"
)

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

	// TODO add Shutdown
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

func (s *service) CreatePeer(peer string) (string, error) {
	s.Service().Log().Line("func", "CreatePeer")

	peer = fmt.Sprintf("layer:%s:peer:%s", s.kind, peer)

	actions := []func(canceler <-chan struct{}) error{
		func(canceler <-chan struct{}) error {
			err := s.Service().Peer().Create(peer)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
		func(canceler <-chan struct{}) error {
			position, err := s.Service().Position().Create(peer)
			if err != nil {
				return maskAny(err)
			}

			err = s.Service().Connection().Create(peer, position)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
	}

	executeConfig := s.Service().Worker().ExecuteConfig()
	executeConfig.SetActions(actions)
	executeConfig.SetCanceler(s.closer)
	executeConfig.SetNumWorkers(len(actions))
	err := s.Service().Worker().Execute(executeConfig)
	if err != nil {
		return "", maskAny(err)
	}

	return peer, nil
}

func (s *service) DeletePeer(peer string) (string, error) {
	s.Service().Log().Line("func", "DeletePeer")

	peer = fmt.Sprintf("layer:%s:peer:%s", s.kind, peer)

	actions := []func(canceler <-chan struct{}) error{
		func(canceler <-chan struct{}) error {
			err := s.Service().Peer().Delete(peer)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
		func(canceler <-chan struct{}) error {
			position, err := s.Service().Position().Delete(peer)
			if err != nil {
				return maskAny(err)
			}
			err = s.Service().Connection().Delete(peer, position)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
	}

	executeConfig := s.Service().Worker().ExecuteConfig()
	executeConfig.SetActions(actions)
	executeConfig.SetCanceler(s.closer)
	executeConfig.SetNumWorkers(len(actions))
	err := s.Service().Worker().Execute(executeConfig)
	if err != nil {
		return "", maskAny(err)
	}

	return peer, nil
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
