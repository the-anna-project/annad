package core

import (
	"fmt"
	"sync"
	"time"

	"github.com/xh3b4sd/anna/gateway"
	gatewayspec "github.com/xh3b4sd/anna/gateway/spec"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/network"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type Config struct {
	TextGateway gatewayspec.Gateway `json:"-"`

	Network spec.Network `json:"network,omitempty"`

	State spec.State `json:"state,omitempty"`
}

const (
	ObjectType spec.ObjectType = "core"
)

func DefaultConfig() Config {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newConfig := Config{
		TextGateway: nil,
		Network:     network.NewNetwork(network.DefaultConfig()),
		State:       state.NewState(newStateConfig),
	}

	return newConfig
}

func NewCore(config Config) spec.Core {
	newCore := &core{
		Config: config,
		Mutex:  sync.Mutex{},
	}

	return newCore
}

type core struct {
	Config

	Mutex sync.Mutex `json:"mutex,omitempty"`
}

func (c *core) Boot() {
	go c.listen()
}

func (c *core) GetObjectID() spec.ObjectID {
	return c.GetState().GetObjectID()
}

func (c *core) GetObjectType() spec.ObjectType {
	return c.GetState().GetObjectType()
}

func (c *core) GetState() spec.State {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.State
}

func (c *core) listen() {
	for {
		newSignal, err := c.TextGateway.ReceiveSignal()
		if gateway.IsGatewayClosed(err) {
			fmt.Printf("gateway is closed\n")
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			fmt.Printf("%#v\n", maskAny(err))
			continue
		}

		responder, err := newSignal.GetResponder()
		if gateway.IsSignalCanceled(err) {
			fmt.Printf("signal is canceled\n")
			continue
		} else if err != nil {
			fmt.Printf("%#v\n", maskAny(err))
			continue
		}

		go func(newSignal gatewayspec.Signal) {
			request, err := newSignal.GetBytes("request")
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}

			newStateConfig := state.DefaultConfig()
			newStateConfig.Bytes["request"] = request
			newStateConfig.ObjectID = spec.ObjectID(newSignal.GetID())
			newStateConfig.ObjectType = impulse.ObjectType
			newState := state.NewState(newStateConfig)

			newImpulseConfig := impulse.DefaultConfig()
			newImpulseConfig.State = newState

			newImpulse := impulse.NewImpulse(newImpulseConfig)

			resImpulse, err := c.Trigger(newImpulse)
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}

			response, err := resImpulse.GetState().GetBytes("response")
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}
			newSignal.SetBytes("response", response)
			responder <- newSignal
		}(newSignal)
	}
}

func (c *core) SetState(state spec.State) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.State = state
}

func (c *core) Shutdown() {
	// TODO close gateway
	// TODO stop listening
	// TODO wait for impulses to be processed
	// TODO backup state
}

func (c *core) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	// Track state.
	imp.GetState().SetCore(c)
	c.GetState().SetImpulse(imp)

	// Initialize network within core state if not already done.
	networks := c.GetState().GetNetworks()
	if len(networks) == 0 {
		c.GetState().SetNetwork(c.Network)
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
