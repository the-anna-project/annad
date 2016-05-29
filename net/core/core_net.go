// Package corenet implements spec.Network. Gateways send signals to the core
// network to ask to do some work. The core network translates a signal into an
// impulse. So the core network is the starting point for all impulses.  Once
// an impulse finished its walk through the core network, the impulse's
// response is translated back to the requesting signal and the signal is send
// back through the gateway.
package corenet

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
	// ObjectTypeCoreNet represents the object type of the core network object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeCoreNet spec.ObjectType = "core-net"
)

// Config represents the configuration used to create a new core network
// object.
type Config struct {
	Log         spec.Log
	Storage     spec.Storage
	TextGateway spec.Gateway
	Scheduler   spec.Scheduler

	EvalNet  spec.Network
	ExecNet  spec.Network
	PatNet   spec.Network
	PredNet  spec.Network
	StratNet spec.Network
}

// DefaultConfig provides a default configuration to create a new core network
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
		Scheduler:   nil,

		EvalNet:  nil,
		ExecNet:  nil,
		PatNet:   nil,
		PredNet:  nil,
		StratNet: nil,
	}

	return newConfig
}

// NewCoreNet creates a new configured core network object.
func NewCoreNet(config Config) (spec.Network, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		return nil, maskAny(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		return nil, maskAny(err)
	}

	newNet := &coreNet{
		Config:             config,
		BootOnce:           sync.Once{},
		ID:                 newID,
		ImpulsesInProgress: 0,
		Mutex:              sync.Mutex{},
		ShutdownOnce:       sync.Once{},
		Type:               ObjectTypeCoreNet,
	}

	newNet.Log.Register(newNet.GetType())

	if newNet.Scheduler == nil {
		return nil, maskAnyf(invalidConfigError, "scheduler must not be empty")
	}

	return newNet, nil
}

type coreNet struct {
	Config

	BootOnce           sync.Once
	ID                 spec.ObjectID
	ImpulsesInProgress int64
	Mutex              sync.Mutex
	ShutdownOnce       sync.Once
	Type               spec.ObjectType
}

func (cn *coreNet) Boot() {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Boot")

	cn.BootOnce.Do(func() {
		go cn.bootObjectTree()
		go cn.TextGateway.Listen(cn.gatewayListener, nil)
	})
}

func (cn *coreNet) NewImpulse() (spec.Impulse, error) {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 15}, "call NewImpulse")

	newConfig := impulse.DefaultConfig()
	newConfig.Log = cn.Log
	newImpulse, err := impulse.NewImpulse(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newImpulse, nil
}

func (cn *coreNet) Shutdown() {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Shutdown")

	cn.ShutdownOnce.Do(func() {
		cn.TextGateway.Close()

		for {
			impulsesInProgress := atomic.LoadInt64(&cn.ImpulsesInProgress)
			if impulsesInProgress == 0 {
				// As soon as all impulses are processed we can go ahead to shutdown the
				// core network.
				break
			}

			time.Sleep(100 * time.Millisecond)
		}

		cn.StratNet.Shutdown()
		cn.PredNet.Shutdown()
		cn.ExecNet.Shutdown()
		cn.EvalNet.Shutdown()
	})
}

func (cn *coreNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call Trigger")

	imp, err := cn.ExecNet.Trigger(imp)
	if err != nil {
		return nil, maskAny(err)
	}

	return imp, nil
}
