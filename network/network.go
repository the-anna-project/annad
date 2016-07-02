// TODO Package network implements spec.Network to provide dynamic and self
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

	"github.com/xh3b4sd/anna/clg"
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
	CLGs               map[string]clgScope
	ID                 spec.ObjectID
	ImpulsesInProgress int64
	Mutex              sync.Mutex
	ShutdownOnce       sync.Once
	Type               spec.ObjectType
}

// Check if the requested CLG should be activated.
// TODO
func (n *network) Activate(clgName string, inputs []reflect.Value) (bool, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Activate")

	// Check if the interface of the requested CLG matches the provided inputs.
	// In case the interface does not match, it is not possible to call the
	// requested CLG using the provided inputs. Then we return an error.
	{
		// TODO move to generated CLG code Inputs method and use it here
		v, err := n.getMethodValue(clgName)
		if err != nil {
			return false, maskAny(err)
		}

		t := v.Type()
		for i := 0; i < t.NumIn(); i++ {
			clgInputs = append(clgInputs, t.In(i))
		}

		if !reflect.DeepEqual(inputs, clgInputs) {
			return false, maskAnyf(invalidCLGExecutionError, "inputs must match interface")
		}
	}

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

// Process the calculation of the requested CLG
// TODO
func (n *network) Calculate(clgName string, inputs []reflect.Value) ([]reflect.Value, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Calculate")

	return nil, nil
}

// Execute CLG
// TODO
func (n *network) Execute(clgName string, inputs []reflect.Value) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Execute")

	// Check if the requested CLG should be activated.
	var activate bool
	var err error
	{
		activate, err = n.Activate(clgName, inputs)
		if err != nil {
			return maskAny(err)
		}
	}

	if !activate {
		return nil
	}

	// Calculate
	var outputs []reflect.Value
	{
		outputs, err = n.Calculate(clgName, inputs)
		if err != nil {
			return maskAny(err)
		}
	}

	// Forward
	{
		err = n.Forward(clgName, inputs, outputs)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}

// Forward to other CLGs
// Split the neural connection path
// TODO
func (n *network) Forward(clgName string, inputs, outputs []reflect.Value) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Forward")

	return nil, nil
}

// Listen to every CLGs input channel
// TODO comment
func (n *network) Listen() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Listen")

	for clgName, CLG := range n.CLGs {
		go func(clgName string, CLG spec.CLG) {
			for {
				select {
				case inputs := <-n.CLGs[clgName].Input:
					go func() {
						err := Execute(clgName, inputs)
						if err != nil {
							n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
						}
					}()
				}
			}
		}(clgName, CLG)
	}

	select {}
}

// TODO comment
func (n *network) Receive(clgName string) ([]reflect.Value, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Receive")

	CLG, ok := n.CLGs[clgName]
	if !ok {
		return nil, maskAnyf(clgNotFoundError, clgName)
	}

	outputs := <-CLG

	return outputs, nil
}

// TODO comment
func (n *network) Send(clgName string, inputs []reflect.Value) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Send")

	CLG, ok := n.CLGs[clgName]
	if !ok {
		return maskAnyf(clgNotFoundError, clgName)
	}

	CLG.Input <- inputs

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
		output, err := n.Receive("Output")
		if err != nil {
			return nil, maskAny(err)
		}

		// The current iteration is over. For the case of another iteration we also
		// prepare the generated output to be the input for the next iteration.
		// For the case of not having another iteration, we set the calculated
		// output to the current impulse.
		input = output
		imp, err = prepareOutput(output)
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
