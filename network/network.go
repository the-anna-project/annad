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
	Log         spec.Log
	Storage     spec.Storage
	TextGateway spec.Gateway

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
	newStorage, err := memory.NewStorage(memory.DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		Log:         log.NewLog(log.DefaultConfig()),
		Storage:     newStorage,
		TextGateway: gateway.NewGateway(gateway.DefaultConfig()),
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
	ID                 spec.ObjectID
	ImpulsesInProgress int64
	Mutex              sync.Mutex
	ShutdownOnce       sync.Once
	Type               spec.ObjectType
}

// Check if the requested CLG should be activated.
// TODO
func (n *network) Activate(clgID spec.ObjectID, inputs []reflect.Value) (bool, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Activate")

	// TODO check connections if the requested CLG should be emitted
	// TODO check connections if an error should be returned in case the CLG should not be activated (learn if error should be returned or not)

	return true, nil
}

func (n *network) Boot() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Boot")

	n.BootOnce.Do(func() {
		n.CLGs = n.configureCLGs(n.CLGs)

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

func (n *network) Execute(clgID spec.ObjectID, inputs []reflect.Value) error {
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

	return nil
}

func (n *network) Listen() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Listen")

	for clgID, clgScope := range n.CLGs {
		go func(clgID spec.ObjectID, CLG spec.CLG) {
			// This is the list of input requests for a CLG execution. We wait a
			// certain amount of time to collect inputs until the requested CLG's
			// interface is fullfilled somehow.
			var inputRequests []inputRequest

			for {
				select {
				case request := <-n.CLGs[clgID].Input:
					// Check if the interface of the requested CLG matches the provided
					// inputs. In case the interface does not match, it is not possible
					// to call the requested CLG using the provided inputs. Then we
					// collect this specific request and wait some time for other input
					// requests to arrive.
					if !equalInputs(request.Inputs, clgScope.CLG.Inputs()) {
						inputRequests = append(inputRequests, request)
					}

					// We should only hold a rather small number of pending requests. A
					// request is considered pending when it is not able to fullfill the
					// requested CLG's interface on its own. Especially the untrained
					// network tends to throw around signals without that much sense.
					// Then the request queue for each CLG can grow and increase the
					// process memory. That is why we cap the queue. TODO The cap must be
					// learned. For ow we hard code it to 10.
					var maxPending = 10
					if len(inputRequests) > maxPending {
						inputRequests = inputRequests[1:]
					}

					// TODO permute the requests until some combination causes the interface to be fullfilled.
					// TODO create a properly ordered list of input requests for the CLG execution below.

					// In case the interface is fullfilled we can finally execute the CLG.
					go func(inputRequests []inputRequest) {
						err := n.Execute(clgID, joinRequestInputs(inputRequests))
						if err != nil {
							n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
							return
						}

						// TODO we need to reward the CLG connections that forwarded signals together correctly
					}(inputRequests)
				}
			}
		}(clgID, clgScope.CLG)
	}

	select {}
}

func (n *network) Receive(clgID spec.ObjectID) (outputResponse, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Receive")

	clgScope, ok := n.CLGs[clgID]
	if !ok {
		return nil, maskAnyf(clgNotFoundError, "%s", clgID)
	}

	outputs := <-clgScope.Output

	return outputs, nil
}

func (n *network) Send(clgID spec.ObjectID, request inputRequest) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Send")

	clgScope, ok := n.CLGs[clgID]
	if !ok {
		return maskAnyf(clgNotFoundError, "%s", clgID)
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

func (n *network) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Trigger")

	input := prepareInput(imp)

	for {
		// Trigger the very first CLG for neural processing. This is the Input CLG.
		// It is provided with the impulse for contextual relevant information
		// related to the current task. Here we send the impulse through the CLG in
		// a fire and forget style. We send something along the neural network
		// without receiving anything back at this point. In the first iteration
		// the input is provided by the incoming impulse. In all following
		// iterations, if any, the input will be output of the preceding iteration.
		err := n.Send("Input", input)
		if err != nil {
			return nil, maskAny(err)
		}

		// Wait for output of the neural network. In the call before we sent some
		// input to trigger the neural connections between the CLGs. Here we wait
		// until the Output CLG was triggered. Once this happens, it means the
		// formerly sent input triggered neural connections up to a point where a
		// connection path was drawn. This connection path started with the Input
		// CLG and ended now here with the Output CLG.
		response, err := n.Receive("Output")
		if err != nil {
			return nil, maskAny(err)
		}

		// The current iteration is over. For the case of another iteration we also
		// prepare the generated output to be the input for the next iteration.
		// For the case of not having another iteration, we set the calculated
		// output to the current impulse.
		input = response.Outputs
		imp, err = prepareOutput(response.Outputs)
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
	}

	return imp, nil
}
