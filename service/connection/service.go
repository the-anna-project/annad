// Package connection provides a service able to manage connections of the
// connection space.
package connection

import (
	"fmt"
	"strconv"
	"strings"

	objectspec "github.com/xh3b4sd/anna/object/spec"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

const (
	// Depth is the default size of each directional coordinate within the
	// connection space. E.g. using a Depth of 3, the resulting volume being taken
	// by a 3 dimensional space would be 9.
	Depth int = 1000000
	// Dimensions is the default number of directional coordinates within the
	// connection space. E.g. a dice has 3 dimensions.
	Dimensions int = 3
	// Weight is the default score applied to a connection expressing its
	// importance.
	Weight int = 0
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

func (s *service) Create(a, b objectspec.Peer) error {
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

func (s *service) CreateConnection(a, b objectspec.Peer) error {
	key := fmt.Sprintf("peer:%s:peer:%s", a.Value(), b.Value())

	res, err := s.Service().Storage().Connection().GetStringMap(key)
	if err != nil {
		return maskAny(err)
	}
	if len(res) == 0 {
		// The connection does not exist. Therefore we create a new one.
		seconds := s.newUnixSeconds()
		weight := strconv.Itoa(Weight)
		val := map[string]string{
			"created": seconds,
			"updated": seconds,
			"weight":  weight,
		}

		err := s.Service().Storage().Connection().SetStringMap(key, val)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}

func (s *service) CreatePeer(p objectspec.Peer) error {
	key := fmt.Sprintf("peer:%s", p.Value())

	res, err := s.Service().Storage().Connection().GetStringMap(key)
	if err != nil {
		return maskAny(err)
	}
	if len(res) == 0 {
		// The peer does not exist. Therefore we create a new one.
		kind := p.Kind()
		position, err := s.CreatePosition()
		if err != nil {
			return maskAny(err)
		}
		seconds := s.newUnixSeconds()
		val := map[string]string{
			"created":  seconds,
			"kind":     kind,
			"position": position,
			"updated":  seconds,
		}

		err = s.Service().Storage().Connection().SetStringMap(key, val)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) CreatePosition() (string, error) {
	nums, err := s.Service().Random().CreateNMax(Dimensions, Depth)
	if err != nil {
		return "", maskAny(err)
	}

	coordinates := []string{}
	for _, n := range nums {
		coordinates = append(coordinates, strconv.Itoa(n))
	}
	position := strings.Join(coordinates, ",")

	return position, nil
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
