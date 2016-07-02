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
		CLGs:               getNewCLGs(),
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
	CLGs               map[string]spec.CLG
	ID                 spec.ObjectID
	ImpulsesInProgress int64
	Mutex              sync.Mutex
	ShutdownOnce       sync.Once
	Type               spec.ObjectType
}

// Check if the requested CLG should be activated.
// TODO
func (n *network) Activate(clgName string, inputs []reflect.Value) (bool, error) {
	// Check if the interface of the requested CLG matches the provided inputs.
	// In case the interface does not match, it is not possible to call the
	// requested CLG using the provided inputs. Then we return an error.
	{
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
		n.configureCLGs()
		go n.TextGateway.Listen(n.getGatewayListener(), nil)
	})
}

// Process the calculation of the requested CLG
// TODO
func (n *network) Calculate(clgName string, inputs []reflect.Value) ([]reflect.Value, error) {
	return nil, nil
}

// Execute CLG
// TODO
func (n *network) Execute(clgName string, inputs []reflect.Value) ([]reflect.Value, error) {
	// Check if the requested CLG should be activated.
	var activate bool
	var err error
	{
		activate, err = n.Activate(clgName, inputs)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	if !activate {
		return nil, nil
	}

	// Calculate
	var outputs []reflect.Value
	{
		outputs, err = n.Calculate(clgName, inputs)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	// Forward
	{
		err = n.Forward(clgName, inputs, outputs)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	return outputs, nil
}

// Forward to other CLGs
// TODO
func (n *network) Forward(clgName string, inputs, outputs []reflect.Value) error {
	return nil, nil
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

	stage := 0
	input := prepareInput(stage, reflect.ValueOf(imp.GetInput()))

	for {
		// Execute the current stage. Therefore contextual relevant connections are
		// looked up. The context is provided by the given input. In the first
		// stage this the input is provided by the incoming impulse. In all
		// following stages the input will be output of the preceding strategy. See
		// /doc/concept/stage.md for more information.
		connections, err := n.Collection.ExecuteCLG(spec.CLG("FindConnections"), input)
		if err != nil {
			return nil, maskAny(err)
		}
		//
		// TODO what about a CLG that learns about errors and decides if the
		// complete calculation should be aborted or can be go on?
		//
		// Having all necessary connections in place enables to create a strategy
		// based on peer relationships.
		strategy, err := n.Collection.ExecuteCLG(spec.CLG("CreateStrategy"), connections)
		if err != nil {
			return nil, maskAny(err)
		}
		// Finally the created strategy can simply be executed.
		output, err := n.Collection.ExecuteCLG(spec.CLG("ExecuteStrategy"), strategy)
		if err != nil {
			return nil, maskAny(err)
		}

		// The iteration of this stage is over. We need to increment the stage in
		// case there is another iteration. For the case of another iteration we
		// also prepare the generated output to be the input for the next stage.
		// For the case of not having another iteration, we set the calculated
		// output to the current impulse.
		stage++
		input = prepareInput(stage, output...)
		imp.SetOutput(prepareOutput(output...))

		// TODO we probably want to also track connections and strategies.

		// TODO check if there were enough stages iterated. The eventual decent
		// number of required stages could be calculated by the number of stages
		// that were required in the past.
		//
		//     // TODO this should be a CLG/strategy as well.
		//     numStages := mean(mean(len(connections.Stages)), mean(len(strategy.Stages)))
		//
		//     if stage < numStages {
		//       continue
		//     }
		//

		// Check the calculated output aganst the provided expectation, if any.
		if imp.GetExpectation().IsEmpty() {
			// There is no expectation provided. We simply go with what we
			// calculated.
			break
		}

		// There is an expectation provided. Thus we are going to check the
		// calculated output against it.
		match, err := imp.GetExpectation().Match(imp)
		if err != nil {
			return nil, maskAny(err)
		}
		if match {
			// The provided expectation did match the calculated result. We apply the
			// information as output to the current impulse and return it.
			break
		}
	}

	return imp, nil
}
