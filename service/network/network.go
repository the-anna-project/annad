// Package network implements spec.Network to provide a neural network based on
// dynamic and self improving CLG execution. The network provides input and
// output channels. When input is received it is injected into the neural
// communication. The following neural activity calculates output which is
// streamed through the output channel back to the requestor.
package network

import (
	"encoding/json"
	"reflect"
	"sync"
	"time"

	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/object/context"
	"github.com/xh3b4sd/anna/object/networkpayload"
	objectspec "github.com/xh3b4sd/anna/object/spec"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	storagespec "github.com/xh3b4sd/anna/storage/spec"

	workerpool "github.com/xh3b4sd/worker-pool"
)

// New creates a new network service.
func New() servicespec.Network {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.Collection
	storageCollection storagespec.Collection

	// Internals.

	bootOnce sync.Once
	// CLGIDs provides a mapping of CLG names pointing to their corresponding CLG.
	clgs         map[string]servicespec.CLG
	closer       chan struct{}
	metadata     map[string]string
	shutdownOnce sync.Once

	// Settings.

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
}

func (s *service) Configure() error {
	// Internals.

	id, err := s.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "log",
		"type": "service",
	}

	s.clgs = s.newCLGs()

	// Settings.
	s.delay = 0

	return nil
}

func (s *service) Activate(CLG servicespec.CLG, networkPayload objectspec.NetworkPayload) (objectspec.NetworkPayload, error) {
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
		s.clgs = s.newCLGs()

		go func() {
			// Create a new worker pool for the input listener.
			inputPoolConfig := workerpool.DefaultConfig()
			inputPoolConfig.Canceler = s.closer
			inputPoolConfig.NumWorkers = 1
			inputPoolConfig.WorkerFunc = s.InputListener
			inputPool, err := workerpool.New(inputPoolConfig)
			if err != nil {
				s.Service().Log().Line("msg", "%#v", maskAny(err))
			}
			// Execute the worker pool and block until all work is done.
			s.logWorkerErrors(inputPool.Execute())
		}()

		go func() {
			// Create a new worker pool for the event listener.
			eventPoolConfig := workerpool.DefaultConfig()
			eventPoolConfig.Canceler = s.closer
			eventPoolConfig.NumWorkers = 10
			eventPoolConfig.WorkerFunc = s.EventListener
			eventPool, err := workerpool.New(eventPoolConfig)
			if err != nil {
				s.Service().Log().Line("msg", "%#v", maskAny(err))
			}
			// Execute the worker pool and block until all work is done.
			s.logWorkerErrors(eventPool.Execute())
		}()
	})
}

func (s *service) Calculate(CLG servicespec.CLG, networkPayload objectspec.NetworkPayload) (objectspec.NetworkPayload, error) {
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
		eventKey := key.NewNetworkKey("event:network-payload")
		element, err := s.Storage().General().PopFromList(eventKey)
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
			s.logNetworkError(invokeEventHandler())
		}
	}
}

func (s *service) EventHandler(CLG servicespec.CLG, networkPayload objectspec.NetworkPayload) error {
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

func (s *service) Forward(CLG servicespec.CLG, networkPayload objectspec.NetworkPayload) error {
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
		case textInput := <-s.Service().TextInput().Channel():
			err := s.InputHandler(CLG, textInput)
			if err != nil {
				s.Service().Log().Line("msg", "%#v", maskAny(err))
			}
		}
	}
}

func (s *service) InputHandler(CLG servicespec.CLG, textInput objectspec.TextInput) error {
	// In case the text request defines the echo flag, we overwrite the given CLG
	// directly to the output CLG. This will cause the created network payload to
	// be forwarded to the output CLG without indirection. Note that this should
	// only be used for testing purposes to bypass more complex neural network
	// activities to directly respond with the received input.
	if textInput.GetEcho() {
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
	ctx.SetExpectation(textInput.GetExpectation())
	ctx.SetSessionID(textInput.GetSessionID())

	// We transform the received text request to a network payload to have a
	// conventional data structure within the neural network.
	newNetworkPayloadConfig := networkpayload.DefaultConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(textInput.GetInput())}
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
	firstBehaviourIDKey := key.NewNetworkKey("clg-tree-id:%s:first-behaviour-id", clgTreeID)
	err = s.Storage().General().Set(firstBehaviourIDKey, string(behaviourID))
	if err != nil {
		return maskAny(err)
	}

	// Write the transformed network payload to the queue.
	eventKey := key.NewNetworkKey("event:network-payload")
	b, err := json.Marshal(newNetworkPayload)
	if err != nil {
		return maskAny(err)
	}
	err = s.Storage().General().PushToList(eventKey, string(b))
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

func (s *service) Service() servicespec.Collection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.Collection) {
	s.serviceCollection = sc
}

func (s *service) SetStorageCollection(sc storagespec.Collection) {
	s.storageCollection = sc
}

func (s *service) Storage() storagespec.Collection {
	return s.storageCollection
}

func (s *service) Track(CLG servicespec.CLG, networkPayload objectspec.NetworkPayload) error {
	s.Service().Log().Line("func", "Track")

	err := s.Service().Tracker().Track(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Validate() error {
	// Dependencies.
	if s.serviceCollection == nil {
		return maskAnyf(invalidConfigError, "service collection must not be empty")
	}
	if s.storageCollection == nil {
		return maskAnyf(invalidConfigError, "storage collection must not be empty")
	}

	return nil
}
