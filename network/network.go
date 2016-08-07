// Package network implements spec.Network to provide dynamic and self
// improving CLG execution. Gateways send signals to the core network to
// request calculations. The core network translates a signal into an impulse.
// So the core network is the starting point for all impulses. Once an impulse
// finished its walk through the core network, the impulse's output is
// translated back to the requesting signal and the signal is send back through
// the gateway to its requestor.
package network

import (
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeNetwork represents the object type of the network object. This
	// is used e.g. to register itself to the logger.
	ObjectTypeNetwork spec.ObjectType = "network"
)

// Config represents the configuration used to create a new network object.
type Config struct {
	// Dependencies.
	Log                spec.Log
	PermutationFactory spec.PermutationFactory
	Storage            spec.Storage
	TextGateway        spec.Gateway

	// Settings.

	// Delay causes each CLG execution to be delayed. This value represents a
	// default value. The actually used value is optimized based on experience
	// and learning.
	// TODO implement
	Delay time.Duration
}

// DefaultConfig provides a default configuration to create a new network
// object by best effort.
func DefaultConfig() Config {
	newPermutationFactory, err := permutation.NewFactory(permutation.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

	newStorage, err := memory.NewStorage(memory.DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		Log:                log.NewLog(log.DefaultConfig()),
		PermutationFactory: newPermutationFactory,
		Storage:            newStorage,
		TextGateway:        gateway.NewGateway(gateway.DefaultConfig()),
	}

	return newConfig
}

// New creates a new configured network object.
func New(config Config) (spec.Network, error) {
	newNetwork := &network{
		Config: config,

		BootOnce:           sync.Once{},
		CLGs:               newCLGs(),
		ID:                 id.MustNew(),
		ImpulsesInProgress: 0,
		Mutex:              sync.Mutex{},
		ShutdownOnce:       sync.Once{},
		Type:               ObjectTypeNetwork,
	}

	if newNetwork.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newNetwork.Storage == nil {
		return nil, maskAnyf(invalidConfigError, "storage must not be empty")
	}
	if newNetwork.TextGateway == nil {
		return nil, maskAnyf(invalidConfigError, "text gateway must not be empty")
	}

	newNetwork.Log.Register(newNetwork.GetType())

	return newNetwork, nil
}

type network struct {
	Config

	BootOnce           sync.Once
	CLGs               map[spec.ObjectID]clgScope
	CLGIDs             map[string]spec.ObjectID
	ID                 spec.ObjectID
	ImpulsesInProgress int64
	Mutex              sync.Mutex
	ShutdownOnce       sync.Once
	Type               spec.ObjectType
}

func (n *network) Activate(clgID spec.ObjectID, inputs []reflect.Value) (bool, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Activate")

	// Check if the given inputs can be used against the requested CLG's
	// interface.
	clgScope, ok := n.CLGs[clgID]
	if !ok {
		return false, maskAnyf(clgNotFoundError, "%s", clgID)
	}
	if !equalTypes(valuesToTypes(inputs), clgScope.CLG.Inputs()) {
		return false, maskAnyf(invalidInterfaceError, "types must match")
	}

	return true, nil
}

func (n *network) Boot() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Boot")

	n.BootOnce.Do(func() {
		n.CLGs = n.configureCLGs(n.CLGs)
		n.CLGIDs = n.mapCLGIDs(n.CLGs)

		go n.Listen()
		go n.TextGateway.Listen(n.getGatewayListener(), nil)
	})
}

func (n *network) Calculate(clgID spec.ObjectID, inputs []reflect.Value) ([]reflect.Value, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Calculate")

	clgScope, ok := n.CLGs[clgID]
	if !ok {
		return nil, maskAnyf(clgNotFoundError, "%s", clgID)
	}

	outputs, err := clgScope.CLG.Calculate(inputs)
	if err != nil {
		return nil, maskAny(err)
	}

	return outputs, nil
}

func (n *network) Execute(clgID spec.ObjectID, requests []spec.InputRequest) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Execute")

	// Each CLG that is executed needs to decide if it wants to be activated.
	// This happens using the Activate method. To make this decision the given
	// input, the CLGs connections and behaviour properties are considered.
	activate, err := n.Activate(clgID, inputs)
	if err != nil {
		return maskAny(err)
	}
	if !activate {
		return nil
	}

	// Once activated, a CLG executes its actual implemented behaviour using
	// Calculate. This behaviour can be anything. It is up to the CLG.
	outputs, err := n.Calculate(clgID, inputs)
	if err != nil {
		return maskAny(err)
	}

	// After the CLGs calculation it can decide what to do next. Like Activate,
	// it is up to the CLG if it forwards signals to further CLGs. E.g. a CLG
	// might or might not forward its calculated results to one or more CLGs.
	// All this depends on its inputs, calculated outputs, CLG connections and
	// behaviour properties.
	err = n.Forward(clgID, inputs, outputs)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

// Forward to other CLGs
// Split the neural connection path
// TODO
func (n *network) Forward(clgID spec.ObjectID, inputs, outputs []reflect.Value) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Forward")

	// Check if the impulse provides a CLG tree ID.
	imp, err := argsToImpulse(inputs)
	if err != nil {
		return nil, maskAny(err)
	}

	clgTreeID := imp.GetCLGTreeID()
	var requests []spec.InputRequest
	if clgTreeID == "" {
		// create new
	} else {
		// lookup existing
	}

	for _, r := range requests {
		err := n.Send(r)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}

// TODO
func (n *network) Listen() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Listen")

	for clgID, clgScope := range n.CLGs {
		// The Output CLG is the only other special CLG, next to the Input CLG.
		// Only these both are treated specially. Here we exclude the Output CLG
		// from the listener, because Network.Trigger is already listening for
		// responses from it.
		if clgScope.CLG.GetName() == "output" {
			continue
		}

		go func(clgID spec.ObjectID, CLG spec.CLG) {
			// This is the queue of input requests. We collect inputs until the
			// requested CLG's interface is fulfilled somehow. Then the CLG is
			// executed and the inputs used to execute the CLG are removed from the
			// queue.
			var queue []spec.InputRequest
			desired := CLG.Inputs()

			for {
				select {
				case request := <-n.CLGs[clgID].Input:
					// We should only hold a rather small number of pending requests. A
					// request is considered pending when it is not able to fulfill the
					// requested CLG's interface on its own. Especially the untrained
					// network tends to throw around signals without that much sense.
					// Then the request queue for each CLG can grow and increase the
					// process memory. That is why we cap the queue. TODO The cap must be
					// learned. For now we hard code it to 10.
					var maxPending = 10
					queue = append(queue, request)
					if len(queue) > maxPending {
						queue = queue[1:]
					}

					// Check if the interface of the requested CLG matches the provided
					// inputs. In case the interface does not match, it is not possible
					// to call the requested CLG using the provided inputs. Then we do
					// nothing, but wait some time for other input requests to arrive.
					execute, matching, newQueue, err := n.extractMatchingInputRequests(queue, desired)
					if err != nil {
						n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
						continue
					}
					if !execute {
						continue
					}
					queue = newQueue

					// In case the interface is fulfilled we can finally execute the CLG.
					go func(request spec.InputRequest) {
						err := n.Execute(clgID, request)
						if err != nil {
							n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
						}
					}(request)
				}
			}
		}(clgID, clgScope.CLG)
	}
}

func (n *network) Receive() (spec.NetworkPayload, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Receive")

	clgScope, ok := n.CLGs[clgID]
	if !ok {
		return spec.OutputResponse{}, maskAnyf(clgNotFoundError, "%s", clgID)
	}

	outputs := <-clgScope.Output

	return outputs, nil
}

func (n *network) Send(request spec.InputRequest) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Send")

	clgScope, ok := n.CLGs[request.Destination]
	if !ok {
		return maskAnyf(clgNotFoundError, "%s", request.Destination)
	}

	clgScope.Input <- request

	return nil
}

func (n *network) Shutdown() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Shutdown")

	n.ShutdownOnce.Do(func() {
		n.TextGateway.Close()

		for {
			impulsesInProgress := atomic.LoadInt64(&n.ImpulsesInProgress)
			if impulsesInProgress == 0 {
				// As soon as all impulses are processed we can go ahead to shutdown the
				// network.
				break
			}

			time.Sleep(100 * time.Millisecond)
		}
	})
}

// TODO
func (n *network) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Trigger")

	request := prepareInput(imp, n.GetID(), n.CLGIDs["input"])

	for {
		// Trigger the very first CLG for neural processing. This is the Input CLG.
		// It is provided with the impulse for contextual relevant information
		// related to the current task. Here we send the impulse through the CLG in
		// a fire and forget style. We send something along the neural network
		// without receiving anything back at this point. In the first iteration
		// the input is provided by the incoming impulse. In all following
		// iterations, if any, the input will be output of the preceding iteration.
		err := n.Send(request)
		if err != nil {
			return nil, maskAny(err)
		}

		// Wait for output of the neural network. In the call before we sent some
		// input to trigger the neural connections between the CLGs. Here we wait
		// until the Output CLG was triggered. Once this happens, it means the
		// formerly sent input triggered neural connections up to a point where a
		// connection path was drawn. This connection path started with the Input
		// CLG and ended now here with the Output CLG.
		//
		// TODO we have a concurrency issue here. The network's CLGs are connected
		// through channels. The output received here is not necessarily related to
		// the input we send above. In case we want to deal with concurrent
		// requests at a later point of time, we need to solve this issue. Current
		// idea would be to maintain some sort of output queue that contains all
		// generated outputs. When waiting on the correct output related to our
		// send input, we would need to go through the output queue until we find
		// the right output. Irrelevent outputs would be requeued. The right output
		// would be recognized by the ID of the impulse being responded with the
		// output together.
		//
		// TODO there can be multiple outputs on one input. We need to handle a
		// stream of outputs that can be streamed as responses over network.
		//
		response, err := n.Receive(n.CLGIDs["output"])
		if err != nil {
			return nil, maskAny(err)
		}

		// The current iteration is over. For the case of not having another
		// iteration, we set the calculated output to the current impulse.
		imp, err = prepareOutput(response)
		if err != nil {
			return nil, maskAny(err)
		}

		// Check the calculated output aganst the provided expectation, if any. In
		// case there is no expectation provided, we simply go with what we
		// calculated. This then means we are probably not in a training situation.
		if imp.GetExpectation().IsEmpty() {
			break
		}

		// There is an expectation provided. Thus we are going to check the
		// calculated output against it. In case the provided expectation did match
		// the calculated result, we simply return it and stop the iteration.
		match, err := imp.GetExpectation().Match(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		if match {
			break
		}

		// TODO move reward/punish to output CLG?
		// TODO expectation met == reward
	}

	return imp, nil
}
