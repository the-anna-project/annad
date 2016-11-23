// Package connection provides a service able to manage connections of the
// connection space.
package connection

import (
	"fmt"
	"strconv"
	"strings"

	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new connection service.
func New() servicespec.ConnectionService {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	metadata       map[string]string
	dimensionCount int
	dimensionDepth int
	weight         int
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

func (s *service) ConnectPeers(a, b objectspec.Peer) error {
	key := fmt.Sprintf("peer:%s", a.Value())

	err := s.Service().Storage().Connection().PushToList(key, b.Value())
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Create(a, b objectspec.Peer) error {
	a, b = s.sortPeers(a, b)

	// TODO process in parallel
	err := s.CreatePeer(a)
	if err != nil {
		return maskAny(err)
	}
	err = s.CreatePeer(b)
	if err != nil {
		return maskAny(err)
	}
	err = s.ConnectPeers(a, b)
	if err != nil {
		return maskAny(err)
	}
	err = s.ConnectPeers(b, a)
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

func (s *service) CreatePosition() (string, error) {
	nums, err := s.Service().Random().CreateNMax(s.dimensionCount, s.dimensionDepth)
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

// TODO
func (s *service) GetAll(peer objectspec.Peer) ([]objectspec.Peer, error) {
	return nil, nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetDimensionCount(dimensionCount int) {
	s.dimensionCount = dimensionCount
}

func (s *service) SetDimensionDepth(dimensionDepth int) {
	s.dimensionDepth = dimensionDepth
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}

func (s *service) SetWeight(weight int) {
	s.weight = weight
}
