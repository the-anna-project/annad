package core

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/network/knowledge"
	"github.com/xh3b4sd/anna/scheduler"
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
	KnowledgeNetwork spec.Network
	Log              spec.Log
	Scheduler        spec.Scheduler
	Storage          spec.Storage
	TextGateway      spec.Gateway
}

// DefaultNetworkConfig provides a default configuration to create a new core
// network object by best effort.
func DefaultNetworkConfig() NetworkConfig {
	newKnowledgeNetwork, err := knowledge.NewNetwork(knowledge.DefaultNetworkConfig())
	if err != nil {
		panic(err)
	}

	newScheduler, err := scheduler.NewScheduler(scheduler.DefaultConfig())
	if err != nil {
		panic(err)
	}

	newStorage, err := memory.NewStorage(memory.DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	newConfig := NetworkConfig{
		KnowledgeNetwork: newKnowledgeNetwork,
		Log:              log.NewLog(log.DefaultConfig()),
		Scheduler:        newScheduler,
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
	if newNetwork.Scheduler == nil {
		return nil, maskAnyf(invalidConfigError, "scheduler must not be empty")
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

func (n *network) NewImpulse() (spec.Impulse, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 15}, "call NewImpulse")

	newConfig := impulse.DefaultConfig()
	newConfig.Log = n.Log
	newImpulse, err := impulse.New(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newImpulse, nil
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

	// TODO

	// peers, err := n.KnowledgeNetwork.GetPeers(imp.GetInput())
	// if err != nil {
	// 	return nil, maskAny(err)
	// }

	// strategyChannel := n.getStrategyChannel(peers)

	// for {
	// 	select {
	// 	case s := <-strategyChannel:
	// 		// Execute the current strategy.
	// 		output, err := s.Excute(imp.GetInput())
	// 		if err != nil {
	// 			return nil, maskAny(err)
	// 		}

	// 		// Match output against expectation, if any.
	// 		e := imp.GetExpectation()
	// 		if e == nil {
	// 			imp.SetOutput(output)
	// 			return imp, nil
	// 		}

	// 		matched, err := output.Match(e)
	// 		if err != nil {
	// 			return nil, maskAny(err)
	// 		}

	// 		if matched {
	// 			imp.SetOutput(output)
	// 			return imp, nil
	// 		}

	// 		// Output did not match expectation. Go ahead.
	// 		continue
	// 	default:
	// 		// There is no strategy available. Thus we create a new one. Here we have
	// 		// the entry point of strategy creation. There can be very sophisticated
	// 		// mechanisms for that. For now we make use of a permutation factory.
	// 		err = n.StrategyFactory.PermuteBy(strategyList, delta)
	// 		if err != nil {
	// 			return nil, maskAny(err)
	// 		}
	// 		err = n.StrategyFactory.MapTo(strategyList)
	// 		if err != nil {
	// 			return nil, maskAny(err)
	// 		}

	// 		strategyChannel <- n.membersToStrategy(strategyList.GetMembers())
	// 	}
	// }

	return nil, nil
}
