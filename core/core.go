// Package core implements spec.Core. Gateways sending signals to the core to
// ask to do some work. The core translates a signal into an impulse. So the
// core is the starting point for all impulses. Once an impulse finished its
// walk through the core, the impulse's response is translated back to the
// requesting signal and the signal is send back through the gateway.
package core

import (
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/gateway/spec"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type Config struct {
	FactoryClient spec.Factory `json:"-"`

	Log spec.Log `json:"-"`

	State spec.State `json:"state,omitempty"`

	TextGateway gatewayspec.Gateway `json:"-"`
}

func DefaultConfig() Config {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = common.ObjectType.Core

	newConfig := Config{
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		State:         state.NewState(newStateConfig),
		TextGateway:   gateway.NewGateway(),
	}

	return newConfig
}

func NewCore(config Config) spec.Core {
	newCore := &core{
		Config:             config,
		Mutex:              sync.Mutex{},
		ImpulsesInProgress: 0,
	}

	return newCore
}

type core struct {
	Config

	Mutex              sync.Mutex `json:"-"`
	ImpulsesInProgress int64      `json:"-"`
}

func (c *core) Boot() {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call Boot")

	err := c.GetState().Read()
	if err != nil {
		c.Log.WithTags(spec.Tags{L: "F", O: c, T: nil, V: 1}, "%#v", maskAny(err))
	}

	go c.listenToGateway()
	go c.listenToSignal()
}

func (c *core) listenToGateway() {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call listenToGateway")

	for {
		newSignal, err := c.TextGateway.ReceiveSignal()
		if gateway.IsGatewayClosed(err) {
			c.Log.WithTags(spec.Tags{L: "W", O: c, T: nil, V: 7}, "gateway is closed")
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			c.Log.WithTags(spec.Tags{L: "E", O: c, T: nil, V: 4}, "%#v", maskAny(err))
			continue
		}
		c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 14}, "core received new signal '%s'", newSignal.GetID())

		responder, err := newSignal.GetResponder()
		if gateway.IsSignalCanceled(err) {
			c.Log.WithTags(spec.Tags{L: "W", O: c, T: nil, V: 7}, "gateway is canceled")
			continue
		} else if err != nil {
			c.Log.WithTags(spec.Tags{L: "E", O: c, T: nil, V: 4}, "%#v", maskAny(err))
			continue
		}

		go func(newSignal gatewayspec.Signal) {
			request, err := newSignal.GetBytes("request")
			if err != nil {
				c.Log.WithTags(spec.Tags{L: "E", O: c, T: nil, V: 4}, "%#v", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}

			newImpulse, err := c.FactoryClient.NewImpulse()
			if err != nil {
				c.Log.WithTags(spec.Tags{L: "E", O: c, T: nil, V: 4}, "%#v", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}

			newStateConfig := state.DefaultConfig()
			newStateConfig.Bytes["request"] = request
			newStateConfig.ObjectID = spec.ObjectID(newSignal.GetID())
			newStateConfig.ObjectType = common.ObjectType.Impulse

			newImpulse.SetState(state.NewState(newStateConfig))

			// Increment the impulse count to track how many impulses are processed
			// inside the core.
			c.ImpulsesInProgress = atomic.AddInt64(&c.ImpulsesInProgress, 1)

			resImpulse, err := c.Trigger(newImpulse)

			// Decrement the impulse count once all hard work is done.
			c.ImpulsesInProgress = atomic.AddInt64(&c.ImpulsesInProgress, -1)

			if err != nil {
				c.Log.WithTags(spec.Tags{L: "E", O: c, T: nil, V: 4}, "%#v", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}

			response, err := resImpulse.GetState().GetBytes("response")
			if err != nil {
				c.Log.WithTags(spec.Tags{L: "E", O: c, T: nil, V: 4}, "%#v", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}
			newSignal.SetBytes("response", response)

			responder <- newSignal
		}(newSignal)
	}
}

func (c *core) listenToSignal() {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call listenToSignal")

	listener := make(chan os.Signal, 1)
	signal.Notify(listener, os.Interrupt, os.Kill)

	<-listener

	c.Shutdown()
}

func (c *core) Shutdown() {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call Shutdown")

	c.TextGateway.Close()

	for {
		impulsesInProgress := atomic.LoadInt64(&c.ImpulsesInProgress)
		if impulsesInProgress == 0 {
			// As soon as all impulses are processed we can go ahead to shutdown the
			// core.
			break
		}

		time.Sleep(100 * time.Millisecond)
	}

	err := c.GetState().Write()
	if err != nil {
		c.Log.WithTags(spec.Tags{L: "F", O: c, T: nil, V: 1}, "%#v", maskAny(err))
	}

	c.Log.WithTags(spec.Tags{L: "I", O: c, T: nil, V: 10}, "shutting down")
	os.Exit(0)
}

func (c *core) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call Trigger")

	// Initialize network within core state if not already done.
	networks := c.GetState().GetNetworks()
	if len(networks) == 0 {
		c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 14}, "create new network")

		newNetwork, err := c.FactoryClient.NewNetwork()
		if err != nil {
			return nil, maskAny(err)
		}
		c.GetState().SetNetwork(newNetwork)
	}

	// Get network. Note that there is potential for multiple networks. For now
	// we just have one.
	for _, n := range c.GetState().GetNetworks() {
		// Initialize the impulses walk through the core via the scheduler network.
		// The scheduler network implements the very first entry for impulses into
		// the cores networks.
		var err error
		imp, err = n.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}

		break
	}

	return imp, nil
}
