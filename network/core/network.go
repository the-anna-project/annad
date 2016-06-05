package core

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/log"
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
	Log         spec.Log
	Storage     spec.Storage
	TextGateway spec.Gateway
	Scheduler   spec.Scheduler
}

// DefaultNetworkConfig provides a default configuration to create a new core
// network object by best effort.
func DefaultNetworkConfig() NetworkConfig {
	newStorage, err := memory.NewStorage(memory.DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	newConfig := NetworkConfig{
		Log:         log.NewLog(log.DefaultConfig()),
		Storage:     newStorage,
		TextGateway: gateway.NewGateway(gateway.DefaultConfig()),
		Scheduler:   nil, // TODO initialize
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

	newNetwork.Log.Register(newNetwork.GetType())

	if newNetwork.Scheduler == nil {
		return nil, maskAnyf(invalidConfigError, "scheduler must not be empty")
	}

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

	return nil, nil
}
