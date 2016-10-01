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
	"github.com/xh3b4sd/anna/factory"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/network/activator"
	"github.com/xh3b4sd/anna/network/forwarder"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"

	"github.com/xh3b4sd/worker-pool"
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
	FactoryCollection spec.FactoryCollection
	Forwarder         spec.Forwarder
	Log               spec.Log
	StorageCollection spec.StorageCollection
	TextInput         chan spec.TextRequest
	TextOutput        chan spec.TextResponse

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
		FactoryCollection: factory.MustNewCollection(),
		Forwarder:         forwarder.MustNew(),
		Log:               log.New(log.DefaultConfig()),
		StorageCollection: storage.MustNewCollection(),
		TextInput:         make(chan spec.TextRequest, 1000),
		TextOutput:        make(chan spec.TextResponse, 1000),

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
		CLGIDs:       map[string]spec.ObjectID{},
		CLGs:         newCLGs(),
		Closer:       make(chan struct{}, 1),
		ID:           id.MustNew(),
		ShutdownOnce: sync.Once{},
		Type:         ObjectType,
	}

	if newNetwork.Activator == nil {
		return nil, maskAnyf(invalidConfigError, "activator must not be empty")
	}
	if newNetwork.FactoryCollection == nil {
		return nil, maskAnyf(invalidConfigError, "factory collection must not be empty")
	}
	if newNetwork.Forwarder == nil {
		return nil, maskAnyf(invalidConfigError, "forwarder must not be empty")
	}
	if newNetwork.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newNetwork.StorageCollection == nil {
		return nil, maskAnyf(invalidConfigError, "storage collection must not be empty")
	}
	if newNetwork.TextInput == nil {
		return nil, maskAnyf(invalidConfigError, "text input channel must not be empty")
	}
	if newNetwork.TextOutput == nil {
		return nil, maskAnyf(invalidConfigError, "text output channel must not be empty")
	}

	newNetwork.Log.Register(newNetwork.GetType())

	return newNetwork, nil
}

type network struct {
	Config

	BootOnce sync.Once

	// CLGIDs provides a mapping of CLG names pointing to their corresponding CLG.
	CLGs map[string]spec.CLG

	Closer       chan struct{}
	ID           spec.ObjectID
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
			err = n.returnAndLogErrors(inputPool.Execute())
			if err != nil {
				n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
			}
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
			err = n.returnAndLogErrors(eventPool.Execute())
			if err != nil {
				n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
			}
		}()
	})
}

func (n *network) Calculate(CLG spec.CLG, networkPayload spec.NetworkPayload) (spec.NetworkPayload, error) {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Calculate")

	clgName, ok := ctx.GetCLGName()
	if !ok {
		return nil, maskAnyf(invalidCLGNameError, "must not be empty")
	}
	CLG, ok := n.CLGs[clgName]
	if !ok {
		return nil, maskAnyf(clgNotFoundError, "name: %s", clgName)
	}

	outputs, err := filterError(reflect.ValueOf(CLG.GetCalculate()).Call(networkPayload.GetArgs()))
	if err != nil {
		return nil, maskAny(err)
	}

	networkPayload.SetArgs(outputs)

	return networkPayload, nil
}

func (n *network) EventListener(canceler <-chan struct{}) error {
	invokeEventHandler := func() error {
		// Fetch the next network payload from the queue. This call blocks until one
		// network payload was fetched from the queue. As soon as we receive the
		// network payload, it is removed from the queue automatically, so it is not
		// handled twice.
		listKey := key.NewCLGKey("events:network-payload")
		element, err := Storage.PopFromList(listKey)
		if err != nil {
			return maskAny(err)
		}
		networkPayload := api.MustNewNetworkPayload()
		err = json.Unmarshal([]byte(element), &networkPayload)
		if err != nil {
			return maskAny(err)
		}

		// Lookup the CLG that is supposed to be executed. The CLG object is
		// referenced by name. When being executed it is referenced by its behavior
		// ID. The behavior ID represents a specific peer within a connection path.
		clgName, ok := networkPayload.GetContext().GetCLGName()
		if !ok {
			return maskAnyf(invalidCLGNameError, "must not be empty")
		}
		CLG, ok := n.CLGs[clgName]
		if !ok {
			return maskAnyf(clgNotFoundError, "name: %s", clgName)
		}

		// Invoke the event handler to execute the given CLG using the given network
		// payload. Here we execute one distinct behavior within its own scope. The
		// CLG decides if and how it is activated, how it calculates its output, if
		// any, and where to forward signals to, if any.
		err := n.EventHandler(CLG, networkPayload)
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
				n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
			}
		}
	}
}

func (n *network) EventHandler(CLG spec.CLG, networkPayload spec.NetworkPayload) error {
	// Activate if the CLG's interface is satisfied by the given
	// network payload.
	networkPayload, err := n.Activate(CLG, networkPayload)
	if IsInvalidInterface(err) {
		// The interface of the requested CLG was not fulfilled. We
		// continue listening for the next network payload without doing
		// any work.
		return nil
	} else if err != nil {
		return maskAny(err)
	}

	// Calculate based on the CLG's implemented business logic.
	newNetworkPayload, err := n.Calculate(CLG, networkPayload)
	if output.IsExpectationNotMet(err) {
		n.Log.WithTags(spec.Tags{C: nil, L: "W", O: n, V: 7}, "%#v", maskAny(err))

		err = n.forwardInputCLG(networkPayload)
		if err != nil {
			return maskAny(err)
		}

		return nil
	} else if err != nil {
		return maskAny(err)
	}

	// Forward to other CLG's, if necessary.
	err = n.Forward(CLG, newNetworkPayload)
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
			ctx := context.MustNew()
			err := n.InputHandler(ctx, CLG, textRequest)
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
	clgTreeID, err := n.Factory().ID().New()
	if err != nil {
		return maskAny(err)
	}
	behaviorID, err := n.Factory().ID().New()
	if err != nil {
		return maskAny(err)
	}

	// Adapt the given context with the information of the current scope.
	ctx.SetBehaviorID(string(behaviorID))
	ctx.SetCLGName(CLG.GetName())
	ctx.SetCLGTreeID(string(clgTreeID))
	ctx.SetExpectation(textRequest.GetExpectation())
	ctx.SetSessionID(textRequest.GetSessionID())

	// We transform the received text request to a network payload to have a
	// conventional data structure within the neural network.
	payloadConfig := api.DefaultNetworkPayloadConfig()
	payloadConfig.Args = []reflect.Value{reflect.ValueOf(textRequest.GetInput())}
	payloadConfig.Context = ctx
	payloadConfig.Destination = behaviorID
	payloadConfig.Sources = []spec.ObjectID{n.GetID()}
	newPayload, err := api.NewNetworkPayload(payloadConfig)
	if err != nil {
		return maskAny(err)
	}

	// Write the new CLG tree ID to reference the input CLG ID and add the CLG
	// tree ID to the new context.
	firstBehaviorIDKey := key.NewCLGKey("clg-tree-id:%s:first-behavior-id", clgTreeID)
	err = n.Storage().General().Set(firstBehaviorIDKey, string(behaviorID))
	if err != nil {
		return maskAny(err)
	}

	// Write the transformed network payload to the queue.
	listKey := key.NewCLGKey("events:network-payload")
	element, err := json.Marshal(networkPayload)
	if err != nil {
		return maskAny(err)
	}
	element, err := n.Storage().General().PopFromList(listKey, element)
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
