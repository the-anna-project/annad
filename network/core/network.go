package core

import (
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xh3b4sd/anna/clg"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/network/knowledge"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeCoreNetwork represents the object type of the core network
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeCoreNetwork spec.ObjectType = "core-network"
)

// NetworkConfig represents the configuration used to create a new core network
// object.
type NetworkConfig struct {
	// Dependencies.
	Collection       *clg.Collection
	KnowledgeNetwork spec.Network
	Log              spec.Log
	Storage          spec.Storage
	TextGateway      spec.Gateway
}

// DefaultNetworkConfig provides a default configuration to create a new core
// network object by best effort.
func DefaultNetworkConfig() NetworkConfig {
	newCollection, err := clg.NewCollection(clg.DefaultCollectionConfig())
	if err != nil {
		panic(err)
	}

	newKnowledgeNetwork, err := knowledge.NewNetwork(knowledge.DefaultNetworkConfig())
	if err != nil {
		panic(err)
	}

	newStorage, err := memory.NewStorage(memory.DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	newConfig := NetworkConfig{
		Collection:       newCollection,
		KnowledgeNetwork: newKnowledgeNetwork,
		Log:              log.NewLog(log.DefaultConfig()),
		Storage:          newStorage,
		TextGateway:      gateway.NewGateway(gateway.DefaultConfig()),
	}

	return newConfig
}

// NewNetwork creates a new configured core network object.
func NewNetwork(config NetworkConfig) (spec.Network, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		return nil, maskAny(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		return nil, maskAny(err)
	}

	newNetwork := &network{
		NetworkConfig: config,

		BootOnce:           sync.Once{},
		ID:                 newID,
		ImpulsesInProgress: 0,
		Mutex:              sync.Mutex{},
		ShutdownOnce:       sync.Once{},
		Type:               ObjectTypeCoreNetwork,
	}

	if newNetwork.KnowledgeNetwork == nil {
		return nil, maskAnyf(invalidConfigError, "knowledge network must not be empty")
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
	NetworkConfig

	BootOnce           sync.Once
	ID                 spec.ObjectID
	ImpulsesInProgress int64
	Mutex              sync.Mutex
	ShutdownOnce       sync.Once
	Type               spec.ObjectType
}

func (n *network) Boot() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Boot")

	n.BootOnce.Do(func() {
		go n.TextGateway.Listen(n.gatewayListener, nil)
	})
}

func (n *network) Shutdown() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Shutdown")

	n.ShutdownOnce.Do(func() {
		n.TextGateway.Close()

		for {
			impulsesInProgress := atomic.LoadInt64(&n.ImpulsesInProgress)
			if impulsesInProgress == 0 {
				// As soon as all impulses are processed we can go ahead to shutdown the
				// core network.
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
