package tracker

import (
	"fmt"
	"sync"

	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new tracker service.
func New() servicespec.TrackerService {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

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
		"name": "tracker",
		"type": "service",
	}
}

func (s *service) CLGIDs(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) error {
	destinationID := string(networkPayload.GetDestination())
	sourceIDs := networkPayload.GetSources()

	errors := make(chan error, len(sourceIDs))
	wg := sync.WaitGroup{}

	for _, sourceID := range sourceIDs {
		wg.Add(1)
		go func(sourceID string) {
			defer wg.Done()

			// Connect source and destination ID of the CLG in the behaviour layer of
			// the connection space.
			err := s.Service().Layer().Behaviour().CreateConnection(sourceID, destinationID)
			if err != nil {
				errors <- maskAny(err)
			}
		}(sourceID)
	}

	wg.Wait()

	select {
	case err := <-errors:
		if err != nil {
			return maskAny(err)
		}
	default:
		// Nothing do here. No error occurred. All good.
	}

	return nil
}

func (s *service) CLGNames(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) error {
	destinationName := CLG.Metadata()["name"]
	sourceIDs := networkPayload.GetSources()

	errors := make(chan error, len(sourceIDs))
	wg := sync.WaitGroup{}

	for _, sourceID := range sourceIDs {
		wg.Add(1)
		go func(sourceID string) {
			defer wg.Done()

			// Resolve behaviour ID to CLG name.
			//
			// TODO handle mapping of CLG ID/Name in separate service
			behaviourNameKey := fmt.Sprintf("behaviour-id:%s:behaviour-name", sourceID)
			sourceName, err := s.Service().Storage().General().Get(behaviourNameKey)
			if err != nil {
				errors <- maskAny(err)
				return
			}

			// Connect source and destination name of the CLG in the behaviour layer
			// of the connection space.
			err = s.Service().Layer().Behaviour().CreateConnection(sourceName, destinationName)
			if err != nil {
				errors <- maskAny(err)
			}
		}(sourceID)
	}

	wg.Wait()

	select {
	case err := <-errors:
		if err != nil {
			return maskAny(err)
		}
	default:
		// Nothing do here. No error occurred. All good.
	}

	return nil
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

func (s *service) Track(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) error {
	s.Service().Log().Line("func", "Track")

	// This is the list of lookup functions which is executed seuqentially.
	lookups := []func(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) error{
		s.CLGIDs,
		s.CLGNames,
	}

	// Execute one lookup after another to track connection path patterns.
	var err error
	for _, l := range lookups {
		err = l(CLG, networkPayload)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}
