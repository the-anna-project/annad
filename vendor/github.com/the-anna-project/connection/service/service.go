// Package service provides a service able to manage connections of the
// connection space.
package service

import (
	"fmt"
	"strconv"
	"time"

	servicespec "github.com/the-anna-project/spec/service"
)

// Config represents the configuration used to create a new connection service.
type Config struct {
	// Settings.
	Weight float64
}

// DefaultConfig provides a default configuration to create a new connection
// service by best effort.
func DefaultConfig() Config {
	return Config{
		// Settings.
		Weight: 0,
	}
}

// New creates a new connection service.
func New(config Config) (servicespec.ConnectionService, error) {
	// Settings.
	if config.Weight == 0 {
		return nil, maskAnyf(invalidConfigError, "weight must not be empty")
	}

	newService := &service{
		// Dependencies.
		serviceCollection: nil,

		// Internals.
		closer:   make(chan struct{}, 1),
		metadata: map[string]string{},

		// Settings.
		weight: config.Weight,
	}

	return newService, nil
}

type service struct {
	// Dependencies.
	serviceCollection servicespec.ServiceCollection

	// Internals.
	// TODO add Shutdown
	closer   chan struct{}
	metadata map[string]string

	// Settings.
	weight float64
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "connection",
		"type": "service",
	}
}

func (s *service) Create(peerA, peerB string) error {
	s.Service().Log().Line("func", "Create")

	actions := []func(canceler <-chan struct{}) error{
		func(canceler <-chan struct{}) error {
			key := fmt.Sprintf("peer:%s", peerA)
			value := fmt.Sprintf("peer:%s", peerB)

			err := s.Service().Storage().Connection().PushToSet(key, value)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
		func(canceler <-chan struct{}) error {
			key := fmt.Sprintf("connecion:%s:%s", peerA, peerB)

			seconds := fmt.Sprintf("%d", time.Now().Unix())
			weight := strconv.FormatFloat(s.Weight(), 'f', -1, 64)

			value := map[string]string{
				"created": seconds,
				"updated": seconds,
				"weight":  weight,
			}

			err := s.Service().Storage().Connection().SetStringMap(key, value)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
	}

	// Execute the list of actions asynchronously.
	executeConfig := s.Service().Worker().ExecuteConfig()
	executeConfig.SetActions(actions)
	executeConfig.SetCanceler(s.closer)
	executeConfig.SetNumWorkers(len(actions))
	err := s.Service().Worker().Execute(executeConfig)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Delete(peerA, peerB string) error {
	s.Service().Log().Line("func", "Delete")

	actions := []func(canceler <-chan struct{}) error{
		func(canceler <-chan struct{}) error {
			key := fmt.Sprintf("peer:%s", peerA)
			value := fmt.Sprintf("peer:%s", peerB)

			err := s.Service().Storage().Connection().RemoveFromSet(key, value)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
		func(canceler <-chan struct{}) error {
			key := fmt.Sprintf("connecion:%s:%s", peerA, peerB)

			err := s.Service().Storage().Connection().Remove(key)
			if err != nil {
				return maskAny(err)
			}

			return nil
		},
	}

	// Execute the list of actions asynchronously.
	executeConfig := s.Service().Worker().ExecuteConfig()
	executeConfig.SetActions(actions)
	executeConfig.SetCanceler(s.closer)
	executeConfig.SetNumWorkers(len(actions))
	err := s.Service().Worker().Execute(executeConfig)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Search(peerA, peerB string) (map[string]string, error) {
	s.Service().Log().Line("func", "Search")

	key := fmt.Sprintf("connecion:%s:%s", peerA, peerB)

	result, err := s.Service().Storage().Connection().GetStringMap(key)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(result) == 0 {
		return nil, maskAnyf(notFoundError, peerA, peerB)
	}

	return result, nil
}

func (s *service) SearchPeers(peer string) ([]string, error) {
	s.Service().Log().Line("func", "SearchPeers")

	key := fmt.Sprintf("peer:%s", peer)

	result, err := s.Service().Storage().Connection().GetAllFromSet(key)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(result) == 0 {
		return nil, maskAnyf(notFoundError, peer)
	}

	return result, nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}

func (s *service) Weight() float64 {
	return s.weight
}
