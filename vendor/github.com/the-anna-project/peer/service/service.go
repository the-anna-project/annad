// Package service provides a service to manage peers within the connection
// space.
package service

import (
	"strconv"
	"strings"

	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new peer service.
func New() servicespec.PeerService {
	return &service{
		// Dependencies.
		serviceCollection: nil,

		// Settings.
		closer:   make(chan struct{}, 1),
		metadata: map[string]string{},
	}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	closer         chan struct{}
	metadata       map[string]string
	dimensionCount int
	dimensionDepth int
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "peer",
		"type": "service",
	}
}

func (s *service) Create(peerA, peerB string) error {
	s.Service().Log().Line("func", "Create")

	// TODO create metadata connections
	//
	//     layer:information:peer:sum     layer:information:$id1
	//     layer:information:peer:$id1    layer:information:sum,layer:information:$id2,layer:information:$id4

	// This is the backreference peer
	err := s.Service().Storage().Connection().PushToSet(peerA, peerB)
	if err != nil {
		return maskAny(err)
	}

	//
	// id, err := s.Service().ID().New()
	// if err != nil {
	// 	return maskAny(err)
	// }

	return nil
}

func (s *service) Delete(peerA string) error {
	s.Service().Log().Line("func", "Delete")

	// TODO delete metadata connections

	err := s.Service().Storage().Connection().Remove(peerA)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) CreatePosition() (string, error) {
	s.Service().Log().Line("func", "CreatePosition")

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

func (s *service) Search(peer string) ([]string, error) {
	s.Service().Log().Line("func", "Search")

	result, err := s.Service().Storage().Connection().GetAllFromSet(peer)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(result) == 0 {
		return nil, maskAny(peerNotFoundError)
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
