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
	"github.com/xh3b4sd/anna/network/activator"
	"github.com/xh3b4sd/anna/network/forwarder"
	"github.com/xh3b4sd/anna/network/tracker"
	"github.com/xh3b4sd/anna/object/context"
	"github.com/xh3b4sd/anna/object/networkpayload"
	objectspec "github.com/xh3b4sd/anna/object/spec"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
	storagespec "github.com/xh3b4sd/anna/storage/spec"

	workerpool "github.com/xh3b4sd/worker-pool"
)

// Config represents the configuration used to create a new network object.
type Config struct {
	// Dependencies.
	Activator         systemspec.Activator
	ServiceCollection servicespec.Collection
	Forwarder         systemspec.Forwarder
	StorageCollection storagespec.Collection
	Tracker           systemspec.Tracker

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
	Delay time.Duration
}

// DefaultConfig provides a default configuration to create a new network
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Activator:         activator.MustNew(),
		ServiceCollection: nil,
		Forwarder:         forwarder.MustNew(),
		StorageCollection: storage.MustNewCollection(),
		Tracker:           tracker.MustNew(),

		// Settings.
		Delay: 0,
	}

	return newConfig
}

// New creates a new configured network object.
func New(config Config) (systemspec.Network, error) {
	newNetwork := &service{
		Config: config,

		BootOnce: sync.Once{},
		Closer:   make(chan struct{}, 1),
		Metadata: map[string]string{
			"id":   id.MustNewID(),
			"name": "network",
			"type": "service",
		},
		ShutdownOnce: sync.Once{},
	}

	if newNetwork.Activator == nil {
		return nil, maskAnyf(invalidConfigError, "activator must not be empty")
	}
	if newNetwork.Forwarder == nil {
		return nil, maskAnyf(invalidConfigError, "forwarder must not be empty")
	}
	if newNetwork.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "service collection must not be empty")
	}
	if newNetwork.StorageCollection == nil {
		return nil, maskAnyf(invalidConfigError, "storage collection must not be empty")
	}
	if newNetwork.Tracker == nil {
		return nil, maskAnyf(invalidConfigError, "tracker must not be empty")
	}

	newNetwork.CLGs = newNetwork.newCLGs()

	return newNetwork, nil
}

// MustNew creates either a new default configured network object, or panics.
func MustNew() systemspec.Network {
	newNetwork, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newNetwork
}

type service struct {
	Config

	BootOnce sync.Once

	// CLGIDs provides a mapping of CLG names pointing to their corresponding CLG.
	CLGs map[string]systemspec.CLG

	Closer       chan struct{}
	Metadata     map[string]string
	ShutdownOnce sync.Once
}

func (s *service) Activate(CLG systemspec.CLG, networkPayload objectspec.NetworkPayload) (objectspec.NetworkPayload, error) {
	s.Service().Log().Line("func", "Activate")

	networkPayload, err := s.Activator.Activate(CLG, networkPayload)
	if err != nil {
		return nil, maskAny(err)
	}

	return networkPayload, nil
}

func (s *service) Boot() {
	s.Service().Log().Line("func", "Boot")

	s.BootOnce.Do(func() {
		s.CLGs = s.newCLGs()

		go func() {
			// Create a new worker pool for the input listener.
			inputPoolConfig := workerpool.DefaultConfig()
			inputPoolConfig.Canceler = s.Closer
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
			eventPoolConfig.Canceler = s.Closer
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

func (s *service) Calculate(CLG systemspec.CLG, networkPayload objectspec.NetworkPayload) (objectspec.NetworkPayload, error) {
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
		CLG, ok := s.CLGs[clgName]
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

func (s *service) EventHandler(CLG systemspec.CLG, networkPayload objectspec.NetworkPayload) error {
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

func (s *service) Forward(CLG systemspec.CLG, networkPayload objectspec.NetworkPayload) error {
	s.Service().Log().Line("func", "Forward")

	err := s.Forwarder.Forward(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) InputListener(canceler <-chan struct{}) error {
	CLG, ok := s.CLGs["input"]
	if !ok {
		return maskAnyf(clgNotFoundError, "name: %s", "input")
	}

	for {
		select {
		case <-canceler:
			return maskAny(workerCanceledError)
		case textInput := <-s.Service().TextInput().GetChannel():
			err := s.InputHandler(CLG, textInput)
			if err != nil {
				s.Service().Log().Line("msg", "%#v", maskAny(err))
			}
		}
	}
}

func (s *service) InputHandler(CLG systemspec.CLG, textInput objectspec.TextInput) error {
	// In case the text request defines the echo flag, we overwrite the given CLG
	// directly to the output CLG. This will cause the created network payload to
	// be forwarded to the output CLG without indirection. Note that this should
	// only be used for testing purposes to bypass more complex neural network
	// activities to directly respond with the received input.
	if textInput.GetEcho() {
		var ok bool
		CLG, ok = s.CLGs["output"]
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
	ctx.SetCLGName(CLG.GetName())
	ctx.SetCLGTreeID(string(clgTreeID))
	ctx.SetExpectation(textInput.GetExpectation())
	ctx.SetSessionID(textInput.GetSessionID())

	// We transform the received text request to a network payload to have a
	// conventional data structure within the neural network.
	newNetworkPayloadConfig := networkpayload.DefaultConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(textInput.GetInput())}
	newNetworkPayloadConfig.Context = ctx
	newNetworkPayloadConfig.Destination = behaviourID
	newNetworkPayloadConfig.Sources = []string{s.GetID()}
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

func (s *service) Shutdown() {
	s.Service().Log().Line("func", "Shutdown")

	s.ShutdownOnce.Do(func() {
		close(s.Closer)
	})
}

func (s *service) Track(CLG systemspec.CLG, networkPayload objectspec.NetworkPayload) error {
	s.Service().Log().Line("func", "Track")

	err := s.Tracker.Track(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
