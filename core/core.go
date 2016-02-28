// Package core implements spec.Core. Gateways sending signals to the core to
// ask to do some work. The core translates a signal into an impulse. So the
// core is the starting point for all impulses. Once an impulse finished its
// walk through the core, the impulse's response is translated back to the
// requesting signal and the signal is send back through the gateway.
package core

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

type Config struct {
	FactoryClient spec.Factory `json:"-"`

	Log spec.Log `json:"-"`

	Networks map[spec.ObjectType]spec.Network `json:"networks"`

	TextGateway spec.Gateway `json:"-"`
}

func DefaultConfig() Config {
	newConfig := Config{
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		Networks:      map[spec.ObjectType]spec.Network{},
		TextGateway:   gateway.NewGateway(gateway.DefaultConfig()),
	}

	return newConfig
}

func NewCore(config Config) spec.Core {
	newCore := &core{
		Closer:             make(chan struct{}, 1),
		Config:             config,
		ID:                 id.NewObjectID(id.Hex128),
		ImpulsesInProgress: 0,
		Mutex:              sync.Mutex{},
		Type:               common.ObjectType.Core,
	}

	return newCore
}

type core struct {
	Closer chan struct{}

	Config

	ID spec.ObjectID `json:"id"`

	Mutex sync.Mutex `json:"-"`

	ImpulsesInProgress int64 `json:"-"`

	Type spec.ObjectType `json:"type"`
}

func (c *core) Boot() {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call Boot")

	go c.bootObjectTree()
	go c.TextGateway.Listen(c.gatewayListener, c.Closer)
}

func (c *core) GetNetworks() (map[spec.ObjectType]spec.Network, error) {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call GetNetworkByType")

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	return c.Networks, nil
}

func (c *core) GetNetworkByType(objectType spec.ObjectType) (spec.Network, error) {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call GetNetworkByType")

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if network, ok := c.Networks[objectType]; ok {
		return network, nil
	}

	return nil, maskAny(networkNotFoundError)
}

func (c *core) SetNetworkByType(objectType spec.ObjectType, network spec.Network) error {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call SetNetworkByType")

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Networks[objectType] = network

	return nil
}

func (c *core) Shutdown() {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call Shutdown")

	errorHandler := func(err error) {
		c.Log.WithTags(spec.Tags{L: "W", O: c, T: nil, V: 7}, "%#v", maskAny(err))
	}

	c.TextGateway.Close()
	c.FactoryClient.Shutdown()

	c.Closer <- struct{}{}

	networks, err := c.GetNetworks()
	if err != nil {
		errorHandler(err)
	}
	for _, network := range networks {
		network.Shutdown()
	}

	for {
		impulsesInProgress := atomic.LoadInt64(&c.ImpulsesInProgress)
		if impulsesInProgress == 0 {
			// As soon as all impulses are processed we can go ahead to shutdown the
			// core.
			break
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// Trigger walks the impulse through the deeps of the neural networks. The
// structure of the networks basically look like this.
//
//   1.    CoreNet
//           StratNet
//           RiskNet
//           EvalNet
//           ExecNet
//
//   2.        JobNet                                 -|
//               StratNet    -|                        |
//               RiskNet      |                        |
//               EvalNet      |- job net scope         |
//               ExecNet     -|                        |
//                                                     |
//   3.        CharNet                                -|
//               StratNet    -|                        |
//               RiskNet      |                        |
//               EvalNet      |- char net scope        |
//               ExecNet     -|                        |
//                                                     |
//   4.        CtxNet                                 -|- core net scope
//               StratNet    -|                        |
//               RiskNet      |                        |
//               EvalNet      |- ctx net scope         |
//               ExecNet     -|                        |
//                                                     |
//   5.        IdeaNet                                -|
//               StratNet    -|                        |
//               RiskNet      |                        |
//               EvalNet      |- idea net scope        |
//               ExecNet     -|                        |
//                                                     |
//   6.        RespNet                                -|
//               StratNet    -|
//               RiskNet      |
//               EvalNet      |- resp net scope
//               ExecNet     -|
func (c *core) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call Trigger")

	// Dynamically walk impulse through the other networks.
	for {
		networkType, err := imp.GetObjectType()
		if err != nil {
			return nil, maskAny(err)
		}

		network, err := c.GetNetworkByType(networkType)
		if err != nil {
			return nil, maskAny(err)
		}
		imp, err = network.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	// Note that the impulse returned here is not actually the same as received
	// at the beginning of the call, but was manipulated during its walk through
	// the networks.
	return imp, nil
}
