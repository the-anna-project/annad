// Package connection provides a service able to manage connections of the
// connection space.
package connection

import (
	"fmt"

	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// New creates a new connection service.
func New() servicespec.Connection {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.Collection

	// Settings.

	metadata map[string]string
}

func (s *service) Configure() error {
	// Settings.

	id, err := s.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "connection",
		"type": "service",
	}

	return nil
}

func (s *service) Create(a, b string) error {
	a, b = s.sortPeers(a, b)

	err := s.CreatePeer(a)
	if err != nil {
		return maskAny(err)
	}
	err = s.CreatePeer(b)
	if err != nil {
		return maskAny(err)
	}
	err = s.CreateConnection(a, b)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) CreateConnection(a, b string) error {
	// TODO ensure connection
	// TODO ensure coordinate peer
	// TODO ensure created peer
	// TODO ensure updated peer
	// TODO ensure type peer

	coordinate := s.newCoordinate()
	seconds := s.newUnixSeconds()

	m := map[string]string{}
	m["coordinate"] = coordinate
	m["created"] = seconds
	m["updated"] = seconds

	connectionKey := fmt.Sprintf("peer:%s:peer:%s", a, b)
	err := s.Service().Storage().Connection().SetStringMap(connectionKey, m)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) CreatePeer(p string) error {
	// TODO ensure peer
	// TODO ensure coordinate peer
	// TODO ensure created peer
	// TODO ensure updated peer
	// TODO ensure type peer

	connectionKey := fmt.Sprintf("service:connection:peer:%s:peer:%s", a, b)
	err := s.Service().Storage().Connection().Set(connectionKey, "{}")
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

const (
	// Depth is the default size of each directional coordinate within the
	// connection space. E.g. using a Depth of 3, the resulting volume being taken
	// by a 3 dimensional space would be 9.
	Depth int = 1000000
	// Dimensions is the default number of directional coordinates within the
	// connection space. E.g. a dice has 3 dimensions.
	Dimensions int = 3
)

func (s *service) NewCoordinate() string {
	nums := s.Service().Random().CreateNMax(Dimensions, Depth)
}

func (s *service) Service() servicespec.Collection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.Collection) {
	s.serviceCollection = sc
}

func (s *service) Validate() error {
	// Dependencies.

	if s.serviceCollection == nil {
		return maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	return nil
}
