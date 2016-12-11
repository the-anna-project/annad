// Package service provides a service to manage position peers within the
// connection space.
package service

import (
	"strconv"
	"strings"

	servicespec "github.com/the-anna-project/spec/service"
)

// Config represents the configuration used to create a new position service.
type Config struct {
	// Settings.
	DimensionCount int
	DimensionDepth int
}

// DefaultConfig provides a default configuration to create a new position
// service by best effort.
func DefaultConfig() Config {
	return Config{
		// Settings.
		DimensionCount: 0,
		DimensionDepth: 0,
	}
}

// New creates a new position service.
func New(config Config) (servicespec.PositionService, error) {
	// Settings.
	if config.DimensionCount == 0 {
		return nil, maskAnyf(invalidConfigError, "dimension count must not be empty")
	}
	if config.DimensionDepth == 0 {
		return nil, maskAnyf(invalidConfigError, "dimension depth must not be empty")
	}

	newService := &service{
		// Dependencies.
		serviceCollection: nil,

		// Settings.
		closer:         make(chan struct{}, 1),
		dimensionCount: config.DimensionCount,
		dimensionDepth: config.DimensionDepth,
		metadata:       map[string]string{},
	}

	return newService, nil
}

type service struct {
	// Dependencies.
	serviceCollection servicespec.ServiceCollection

	// Settings.
	// TODO add Shutdown
	closer         chan struct{}
	dimensionCount int
	dimensionDepth int
	metadata       map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "position",
		"type": "service",
	}
}

func (s *service) Default() (string, error) {
	s.Service().Log().Line("func", "Default")

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

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}
