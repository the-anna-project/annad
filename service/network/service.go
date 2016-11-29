// Package network implements spec.NetworkService to provide a neural network based on
// dynamic and self improving CLG execution. The network provides input and
// output channels. When input is received it is injected into the neural
// communication. The following neural activity calculates output which is
// streamed through the output channel back to the requestor.
package network

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/the-anna-project/annad/object/context"
	"github.com/the-anna-project/annad/object/networkpayload"
	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new network service.
func New() servicespec.NetworkService {
	return &service{
		// Dependencies.
		serviceCollection: nil,

		// Settings.
		closer:       make(chan struct{}, 1),
		metadata:     map[string]string{},
		shutdownOnce: sync.Once{},
	}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	bootOnce sync.Once
	// CLGIDs provides a mapping of CLG names pointing to their corresponding CLG.
	clgs   map[string]servicespec.CLGService
	closer chan struct{}
	// Delay causes each CLG execution to be delayed. This value represents a
	// default value. A delay can be used harden the internal synchronization of
	// the network. For instance some chaos monkey could be implemented to cause
	// unusual delays in neural communications. Analysing situations in which such
	// chaos monkey takes place might shed some light on faulty implementations
	// within the neural network.
	//
	// TODO implement the actual usage of the delay and make it dynamically
	// configurable on demand like we already do with the log control.
	delay time.Duration

	metadata     map[string]string
	shutdownOnce sync.Once
}

func (s *service) Activate(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) (objectspec.NetworkPayload, error) {
	s.Service().Log().Line("func", "Activate")

	networkPayload, err := s.Service().Activator().Activate(CLG, networkPayload)
	if err != nil {
		return nil, maskAny(err)
	}

	return networkPayload, nil
}

func (s *service) Boot() {
	s.Service().Log().Line("func", "Boot")

	s.bootOnce.Do(func() {
		id, err := s.Service().ID().New()
		if err != nil {
			panic(err)
		}
		s.metadata = map[string]string{
			"id":   id,
			"name": "log",
			"type": "service",
		}

		s.clgs = s.newCLGs()
		s.delay = 0

		go func() {
			// Create a new execute config for the worker service to execute the
			// input listener.
			executeConfig := s.Service().Worker().ExecuteConfig()
			executeConfig.SetActions([]func(canceler <-chan struct{}) error{s.InputListener})
			executeConfig.SetCanceler(s.closer)
			executeConfig.SetNumWorkers(1)
			err := s.Service().Worker().Execute(executeConfig)
			if err != nil {
				s.Service().Log().Line("msg", maskAny(err))
			}
		}()

		go func() {
			// Create a new execute config for the worker service to execute the
			// event listener.
			executeConfig := s.Service().Worker().ExecuteConfig()
			executeConfig.SetActions([]func(canceler <-chan struct{}) error{s.EventListener})
			executeConfig.SetCanceler(s.closer)
			executeConfig.SetNumWorkers(10)
			err := s.Service().Worker().Execute(executeConfig)
			if err != nil {
				s.Service().Log().Line("msg", maskAny(err))
			}
		}()
	})
}

func (s *service) Calculate(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) (objectspec.NetworkPayload, error) {
	s.Service().Log().Line("func", "Calculate")

	outputs, err := filterError(reflect.ValueOf(CLG.GetCalculate()).Call(networkPayload.GetCLGInput()))
	if err != nil {
		return nil, maskAny(err)
	}

	newNetworkPayloadConfig := networkpayload.DefaultConfig()
	newNetworkPayloadConfig.Args = outputs
	newNetworkPayloadConfig.Context = networkPayload.GetContext()
	newNetworkPayloadConfig.Destination = networkPayload.GetDestination()
	newNetworkPayloadConfig.Sources = networkPayload.GetSources()
	newNetworkPayload, err := networkpayload.New(newNetworkPayloadConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newNetworkPayload, nil
}

func (s *service) EventListener(canceler <-chan struct{}) error {
	invokeEventHandler := func() error {
		// Fetch the next network payload from the queue. This call blocks until one
		// network payload was fetched from the queue. As soon as we receive the
		// network payload, it is removed from the queue automatically, so it is not
		// handled twice.
		eventKey := fmt.Sprintf("event:network-payload")
		element, err := s.Service().Storage().General().PopFromList(eventKey)
		if err != nil {
			return maskAny(err)
		}
		networkPayload := networkpayload.MustNew()
		err = json.Unmarshal([]byte(element), &networkPayload)
		if err != nil {
			return maskAny(err)
		}

		// Lookup the CLG that is supposed to be executed. The CLG object is
		// referenced by name. When being executed it is referenced by its behaviour
		// ID. The behaviour ID represents a specific peer within a connection path.
		clgName, ok := networkPayload.GetContext().GetCLGName()
		if !ok {
			return maskAnyf(invalidCLGNameError, "must not be empty")
		}
		CLG, ok := s.clgs[clgName]
		if !ok {
			return maskAnyf(clgNotFoundError, "name: %s", clgName)
		}

		// Invoke the event handler to execute the given CLG using the given network
		// payload. Here we execute one distinct behaviour within its own scope. The
		// CLG decides if and how it is activated, how it calculates its output, if
		// any, and where to forward signals to, if any.
		err = s.EventHandler(CLG, networkPayload)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	for {
		select {
		case <-canceler:
			return maskAny(workerCanceledError)
		default:
			err := invokeEventHandler()
			if err != nil {
				s.Service().Log().Line("msg", maskAny(err))
			}
		}
	}
}

func (s *service) EventHandler(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) error {
	var err error

	// Activate if the CLG's interface is satisfied by the given
	// network payload.
	networkPayload, err = s.Activate(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	// Calculate based on the CLG's implemented business logic.
	networkPayload, err = s.Calculate(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	// Forward to other CLG's, if necessary.
	err = s.Forward(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	// Track the the given CLG and network payload to learn more about the
	// connection paths created.
	err = s.Track(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Forward(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) error {
	s.Service().Log().Line("func", "Forward")

	err := s.Service().Forwarder().Forward(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) InputListener(canceler <-chan struct{}) error {
	CLG, ok := s.clgs["input"]
	if !ok {
		return maskAnyf(clgNotFoundError, "name: %s", "input")
	}

	for {
		select {
		case <-canceler:
			return maskAny(workerCanceledError)
		case textInput := <-s.Service().Input().Text().Channel():
			err := s.InputHandler(CLG, textInput)
			if err != nil {
				s.Service().Log().Line("msg", "%#v", maskAny(err))
			}
		}
	}
}

func (s *service) InputHandler(CLG servicespec.CLGService, textInput objectspec.TextInput) error {
	// In case the text request defines the echo flag, we overwrite the given CLG
	// directly to the output CLG. This will cause the created network payload to
	// be forwarded to the output CLG without indirection. Note that this should
	// only be used for testing purposes to bypass more complex neural network
	// activities to directly respond with the received input.
	if textInput.Echo() {
		var ok bool
		CLG, ok = s.clgs["output"]
		if !ok {
			return maskAnyf(clgNotFoundError, "name: %s", "output")
		}
	}

	// Create new IDs for the new CLG tree and the input CLG.
	clgTreeID, err := s.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	behaviourID, err := s.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}

	// Create a new context and adapt it using the information of the current scope.
	ctx := context.MustNew()
	ctx.SetBehaviourID(string(behaviourID))
	ctx.SetCLGName(CLG.Metadata()["name"])
	ctx.SetCLGTreeID(string(clgTreeID))
	ctx.SetExpectation(textInput.Expectation())
	ctx.SetSessionID(textInput.SessionID())

	// We transform the received text request to a network payload to have a
	// conventional data structure within the neural network.
	newNetworkPayloadConfig := networkpayload.DefaultConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(textInput.Input())}
	newNetworkPayloadConfig.Context = ctx
	// TODO destination and sources should be metadata objects
	newNetworkPayloadConfig.Destination = behaviourID
	newNetworkPayloadConfig.Sources = []string{s.Metadata()["id"]}
	newNetworkPayload, err := networkpayload.New(newNetworkPayloadConfig)
	if err != nil {
		return maskAny(err)
	}

	// Write the new CLG tree ID to reference the input CLG ID and add the CLG
	// tree ID to the new context.
	firstBehaviourIDKey := fmt.Sprintf("clg-tree-id:%s:first-behaviour-id", clgTreeID)
	err = s.Service().Storage().General().Set(firstBehaviourIDKey, string(behaviourID))
	if err != nil {
		return maskAny(err)
	}

	// Write the transformed network payload to the queue.
	eventKey := fmt.Sprintf("event:network-payload")
	b, err := json.Marshal(newNetworkPayload)
	if err != nil {
		return maskAny(err)
	}
	err = s.Service().Storage().General().PushToList(eventKey, string(b))
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Shutdown() {
	s.Service().Log().Line("func", "Shutdown")

	s.shutdownOnce.Do(func() {
		close(s.closer)
	})
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}

func (s *service) Track(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) error {
	s.Service().Log().Line("func", "Track")

	err := s.Service().Tracker().Track(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
