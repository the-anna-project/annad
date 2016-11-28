package tracker

import (
	"fmt"

	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new tracker service.
func New() servicespec.TrackerService {
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

	closer   chan struct{}
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

	// Prepare a queue to synchronise the workload.
	queue := make(chan string, len(sourceIDs))
	for _, sourceID := range sourceIDs {
		queue <- sourceID
	}

	// Define the action being executed by the worker service. This action is
	// supposed to be executed concurrently. Therefore the queue we just created
	// is used to synchronize the workload.
	action := func(canceler <-chan struct{}) error {
		sourceID := <-queue

		// Connect source and destination ID of the CLG in the behaviour layer of
		// the connection space.
		err := s.Service().Layer().Behaviour().CreateConnection(sourceID, destinationID)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	executeConfig := s.Service().Worker().ExecuteConfig()
	executeConfig.SetActions([]func(canceler <-chan struct{}) error{action})
	executeConfig.SetCanceler(s.closer)
	executeConfig.SetNumWorkers(len(sourceIDs))
	err := s.Service().Worker().Execute(executeConfig)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) CLGNames(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) error {
	destinationName := CLG.Metadata()["name"]
	sourceIDs := networkPayload.GetSources()

	// Prepare a queue to synchronise the workload.
	queue := make(chan string, len(sourceIDs))
	for _, sourceID := range sourceIDs {
		queue <- sourceID
	}

	// Define the action being executed by the worker service. This action is
	// supposed to be executed concurrently. Therefore the queue we just created
	// is used to synchronize the workload.
	action := func(canceler <-chan struct{}) error {
		sourceID := <-queue

		// Resolve behaviour ID to CLG name.
		//
		// TODO handle mapping of CLG ID/Name in separate service
		behaviourNameKey := fmt.Sprintf("behaviour-id:%s:behaviour-name", sourceID)
		sourceName, err := s.Service().Storage().General().Get(behaviourNameKey)
		if err != nil {
			return maskAny(err)
		}

		// Connect source and destination name of the CLG in the behaviour layer
		// of the connection space.
		err = s.Service().Layer().Behaviour().CreateConnection(sourceName, destinationName)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	executeConfig := s.Service().Worker().ExecuteConfig()
	executeConfig.SetActions([]func(canceler <-chan struct{}) error{action})
	executeConfig.SetCanceler(s.closer)
	executeConfig.SetNumWorkers(len(sourceIDs))
	err := s.Service().Worker().Execute(executeConfig)
	if err != nil {
		return maskAny(err)
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
	//
	// TODO execute concurrently
	var err error
	for _, l := range lookups {
		err = l(CLG, networkPayload)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}
