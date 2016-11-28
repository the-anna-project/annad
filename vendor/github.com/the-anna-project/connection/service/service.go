// Package service provides a service able to manage connections of the
// connection space.
package service

import (
	"fmt"
	"strconv"
	"time"

	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new connection service.
func New() servicespec.ConnectionService {
	return &service{
		// Dependencies.
		serviceCollection: nil,

		// Settings.
		metadata: map[string]string{},
		weight:   0,
	}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	metadata map[string]string
	weight   int
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

	key := fmt.Sprintf("connecion:%s:%s", peerA, peerB)

	seconds := fmt.Sprintf("%d", time.Now().Unix())
	weight := strconv.Itoa(s.weight)
	val := map[string]string{
		"created": seconds,
		"updated": seconds,
		"weight":  weight,
	}

	err := s.Service().Storage().Connection().SetStringMap(key, val)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Delete(peerA, peerB string) error {
	s.Service().Log().Line("func", "Delete")

	key := fmt.Sprintf("connecion:%s:%s", peerA, peerB)

	err := s.Service().Storage().Connection().Remove(key)
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
		return nil, maskAny(connectionNotFoundError)
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

func (s *service) SetWeight(weight int) {
	s.weight = weight
}
