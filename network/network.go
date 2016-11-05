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

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/context"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/network/activator"
	"github.com/xh3b4sd/anna/network/forwarder"
	"github.com/xh3b4sd/anna/network/tracker"
	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/id"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"

	workerpool "github.com/xh3b4sd/worker-pool"
)

const (
	// ObjectType represents the object type of the network object. This is used
	// e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "network"
)

// Config represents the configuration used to create a new network object.
type Config struct {
	// Dependencies.
	Activator         spec.Activator
	ServiceCollection spec.ServiceCollection
	Forwarder         spec.Forwarder
	Log               spec.Log
	StorageCollection spec.StorageCollection
	Tracker           spec.Tracker
	TextInput         chan spec.TextRequest

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
		ServiceCollection: service.MustNewCollection(),
		Forwarder:         forwarder.MustNew(),
		Log:               log.New(log.DefaultConfig()),
		StorageCollection: storage.MustNewCollection(),
		Tracker:           tracker.MustNew(),
		TextInput:         make(chan spec.TextRequest, 1000),

		// Settings.
		Delay: 0,
	}

	return newConfig
}

// New creates a new configured network object.
func New(config Config) (spec.Network, error) {
	newNetwork := &network{
		Config: config,

		BootOnce:     sync.Once{},
		Closer:       make(chan struct{}, 1),
		ID:           id.MustNew(),
		ShutdownOnce: sync.Once{},
		Type:         ObjectType,
	}

	if newNetwork.Activator == nil {
		return nil, maskAnyf(invalidConfigError, "activator must not be empty")
	}
	if newNetwork.Forwarder == nil {
		return nil, maskAnyf(invalidConfigError, "forwarder must not be empty")
	}
	if newNetwork.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
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
	if newNetwork.TextInput == nil {
		return nil, maskAnyf(invalidConfigError, "text input channel must not be empty")
	}

	newNetwork.CLGs = newNetwork.newCLGs()
	newNetwork.Log.Register(newNetwork.GetType())

	return newNetwork, nil
}

// MustNew creates either a new default configured network object, or panics.
func MustNew() spec.Network {
	newNetwork, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newNetwork
}

type network struct {
	Config

	BootOnce sync.Once

	// CLGIDs provides a mapping of CLG names pointing to their corresponding CLG.
	CLGs map[string]spec.CLG

	Closer       chan struct{}
	ID           string
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (n *network) Activate(CLG spec.CLG, networkPayload spec.NetworkPayload) (spec.NetworkPayload, error) {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Activate")

	networkPayload, err := n.Activator.Activate(CLG, networkPayload)
	if err != nil {
		return nil, maskAny(err)
	}

	return networkPayload, nil
}

func (n *network) Boot() {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Boot")

	n.BootOnce.Do(func() {
		n.CLGs = n.newCLGs()

		go func() {
			// Create a new worker pool for the input listener.
			inputPoolConfig := workerpool.DefaultConfig()
			inputPoolConfig.Canceler = n.Closer
			inputPoolConfig.NumWorkers = 1
			inputPoolConfig.WorkerFunc = n.InputListener
			inputPool, err := workerpool.New(inputPoolConfig)
			if err != nil {
				n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
			}
			// Execute the worker pool and block until all work is done.
			n.logWorkerErrors(inputPool.Execute())
		}()

		go func() {
			// Create a new worker pool for the event listener.
			eventPoolConfig := workerpool.DefaultConfig()
			eventPoolConfig.Canceler = n.Closer
			eventPoolConfig.NumWorkers = 10
			eventPoolConfig.WorkerFunc = n.EventListener
			eventPool, err := workerpool.New(eventPoolConfig)
			if err != nil {
				n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
			}
			// Execute the worker pool and block until all work is done.
			n.logWorkerErrors(eventPool.Execute())
		}()
	})
}

func (n *network) Calculate(CLG spec.CLG, networkPayload spec.NetworkPayload) (spec.NetworkPayload, error) {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Calculate")

	outputs, err := filterError(reflect.ValueOf(CLG.GetCalculate()).Call(networkPayload.GetCLGInput()))
	if err != nil {
		return nil, maskAny(err)
	}

	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = outputs
	newNetworkPayloadConfig.Context = networkPayload.GetContext()
	newNetworkPayloadConfig.Destination = networkPayload.GetDestination()
	newNetworkPayloadConfig.Sources = networkPayload.GetSources()
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newNetworkPayload, nil
}

func (n *network) EventListener(canceler <-chan struct{}) error {
	invokeEventHandler := func() error {
		// Fetch the next network payload from the queue. This call blocks until one
		// network payload was fetched from the queue. As soon as we receive the
		// network payload, it is removed from the queue automatically, so it is not
		// handled twice.
		eventKey := key.NewNetworkKey("event:network-payload")
		element, err := n.Storage().General().PopFromList(eventKey)
		if err != nil {
			return maskAny(err)
		}
		networkPayload := api.MustNewNetworkPayload()
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
		CLG, ok := n.CLGs[clgName]
		if !ok {
			return maskAnyf(clgNotFoundError, "name: %s", clgName)
		}

		// Invoke the event handler to execute the given CLG using the given network
		// payload. Here we execute one distinct behaviour within its own scope. The
		// CLG decides if and how it is activated, how it calculates its output, if
		// any, and where to forward signals to, if any.
		err = n.EventHandler(CLG, networkPayload)
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
			n.logNetworkError(invokeEventHandler())
		}
	}
}

func (n *network) EventHandler(CLG spec.CLG, networkPayload spec.NetworkPayload) error {
	var err error

	// Activate if the CLG's interface is satisfied by the given
	// network payload.
	networkPayload, err = n.Activate(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	// Calculate based on the CLG's implemented business logic.
	networkPayload, err = n.Calculate(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	// Forward to other CLG's, if necessary.
	err = n.Forward(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	// Track the the given CLG and network payload to learn more about the
	// connection paths created.
	err = n.Track(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (n *network) Forward(CLG spec.CLG, networkPayload spec.NetworkPayload) error {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Forward")

	err := n.Forwarder.Forward(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (n *network) InputListener(canceler <-chan struct{}) error {
	CLG, ok := n.CLGs["input"]
	if !ok {
		return maskAnyf(clgNotFoundError, "name: %s", "input")
	}

	for {
		select {
		case <-canceler:
			return maskAny(workerCanceledError)
		case textRequest := <-n.TextInput:
			err := n.InputHandler(CLG, textRequest)
			if err != nil {
				n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
			}
		}
	}
}

func (n *network) InputHandler(CLG spec.CLG, textRequest spec.TextRequest) error {
	// In case the text request defines the echo flag, we overwrite the given CLG
	// directly to the output CLG. This will cause the created network payload to
	// be forwarded to the output CLG without indirection. Note that this should
	// only be used for testing purposes to bypass more complex neural network
	// activities to directly respond with the received input.
	if textRequest.GetEcho() {
		var ok bool
		CLG, ok = n.CLGs["output"]
		if !ok {
			return maskAnyf(clgNotFoundError, "name: %s", "output")
		}
	}

	// Create new IDs for the new CLG tree and the input CLG.
	clgTreeID, err := n.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	behaviourID, err := n.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}

	// Create a new context and adapt it using the information of the current scope.
	ctx := context.MustNew()
	ctx.SetBehaviourID(string(behaviourID))
	ctx.SetCLGName(CLG.GetName())
	ctx.SetCLGTreeID(string(clgTreeID))
	ctx.SetExpectation(textRequest.GetExpectation())
	ctx.SetSessionID(textRequest.GetSessionID())

	// We transform the received text request to a network payload to have a
	// conventional data structure within the neural network.
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(textRequest.GetInput())}
	newNetworkPayloadConfig.Context = ctx
	newNetworkPayloadConfig.Destination = spec.ObjectID(behaviourID)
	newNetworkPayloadConfig.Sources = []spec.ObjectID{spec.ObjectID(n.GetID())}
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		return maskAny(err)
	}

	// Write the new CLG tree ID to reference the input CLG ID and add the CLG
	// tree ID to the new context.
	firstBehaviourIDKey := key.NewNetworkKey("clg-tree-id:%s:first-behaviour-id", clgTreeID)
	err = n.Storage().General().Set(firstBehaviourIDKey, string(behaviourID))
	if err != nil {
		return maskAny(err)
	}

	// Write the transformed network payload to the queue.
	eventKey := key.NewNetworkKey("event:network-payload")
	b, err := json.Marshal(newNetworkPayload)
	if err != nil {
		return maskAny(err)
	}
	err = n.Storage().General().PushToList(eventKey, string(b))
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (n *network) Shutdown() {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Shutdown")

	n.ShutdownOnce.Do(func() {
		close(n.Closer)
	})
}

func (n *network) Track(CLG spec.CLG, networkPayload spec.NetworkPayload) error {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Track")

	err := n.Tracker.Track(CLG, networkPayload)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
