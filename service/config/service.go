// Package config provides configuration for all services within the service
// collection.
package config

import (
	"github.com/xh3b4sd/anna/object/config"
	"github.com/xh3b4sd/anna/object/config/endpoint"
	"github.com/xh3b4sd/anna/object/config/space"
	"github.com/xh3b4sd/anna/object/config/storage"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// New creates a new config service.
func New() servicespec.Config {
	return &service{}
}

type service struct {
	// Dependencies.

	configCollection  *config.Collection
	serviceCollection servicespec.Collection

	// Settings.

	metadata map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "feature",
		"type": "service",
	}
}

func (s *service) Endpoint() *endpoint.Collection {
	return s.config.Endpoint()
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.Collection {
	return s.serviceCollection
}

func (s *service) SetConfigCollection(configCollection *config.Collection) {
	s.configCollection = configCollection
}

func (s *service) SetServiceCollection(sc servicespec.Collection) {
	s.serviceCollection = sc
}

func (s *service) Space() *space.Collection {
	return s.config.Space()
}

func (s *service) Storage() *storage.Collection {
	return s.config.Storage()
}
