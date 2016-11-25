// Package connection provides a service able to manage connections of the
// connection space.
package connection

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

func (s *service) Create(peerA, peerB string) error {
	// Make sure a peer is not connected with itself.
	if peerA == peerB {
		return maskAnyf(invalidPeerError, "peers must be different")
	}

	_, err := s.Find(peerA, peerB)
	if IsConnectionNotFound(err) {
		// If there is no connection yet created, we can go ahead and create the
		// requested peer and connection.
		//
		// TODO process in parallel
		err := s.CreatePeer(peerA, peerB)
		if err != nil {
			return maskAny(err)
		}
		err = s.CreateConnection(peerA, peerB)
		if err != nil {
			return maskAny(err)
		}
	} else if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) CreateConnection(peerA, peerB string) error {
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

func (s *service) CreatePeer(peerA, peerB string) error {
	key := fmt.Sprintf("peer:%s", peerA)

	err := s.Service().Storage().Connection().PushToSet(key, peerB)
	if err != nil {
		return maskAny(err)
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

func (s *service) Find(peerA, peerB string) (map[string]string, error) {
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

func (s *service) FindPeers(peer string) ([]string, error) {
	key := fmt.Sprintf("peer:%s", peer)

	result, err := s.Service().Storage().Connection().GetAllFromSet(key)
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
