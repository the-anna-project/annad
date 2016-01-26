package core

import (
	"sync"
	"time"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/gateway"
	gatewayspec "github.com/xh3b4sd/anna/gateway/spec"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type Config struct {
	TextGateway gatewayspec.Gateway `json:"-"`

	Log spec.Log `json:"-"`

	States map[string]spec.State `json:"states,omitempty"`
}

const (
	ObjectType spec.ObjectType = "core"
)

func DefaultConfig() Config {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newConfig := Config{
		TextGateway: gateway.NewGateway(),
		Log:         log.NewLog(log.DefaultConfig()),
		States: map[string]spec.State{
			common.DefaultStateKey: state.NewState(newStateConfig),
		},
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

	Mutex sync.Mutex `json:"-"`
}

func (c *core) Boot() {
	c.Log.V(12).Debugf("call Core.Boot")

	go c.listen()
}

func (c *core) Copy() spec.Core {
	coreCopy := *c

	for key, state := range coreCopy.States {
		coreCopy.States[key] = state.Copy()
	}

	return &coreCopy
}

func (c *core) GetObjectID() spec.ObjectID {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	return c.States[common.DefaultStateKey].GetObjectID()
}

func (c *core) GetObjectType() spec.ObjectType {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	return c.States[common.DefaultStateKey].GetObjectType()
}

func (c *core) GetState(key string) (spec.State, error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if state, ok := c.States[key]; ok {
		return state, nil
	}

	return nil, maskAny(stateNotFoundError)
}

func (c *core) listen() {
	c.Log.V(12).Debugf("call Core.listen")

	for {
		newSignal, err := c.TextGateway.ReceiveSignal()
		if gateway.IsGatewayClosed(err) {
			c.Log.V(6).Warnf("gateway is closed")
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			c.Log.V(3).Errorf("%#v", maskAny(err))
			continue
		}
		c.Log.V(12).Debugf("core received new signal '%s'", newSignal.GetID())

		responder, err := newSignal.GetResponder()
		if gateway.IsSignalCanceled(err) {
			c.Log.V(6).Warnf("signal is canceled")
			continue
		} else if err != nil {
			c.Log.V(3).Errorf("%#v", maskAny(err))
			continue
		}

		go func(newSignal gatewayspec.Signal) {
			request, err := newSignal.GetBytes("request")
			if err != nil {
				c.Log.V(3).Errorf("%#v", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}

			newImpulse, err := common.GetInitImpulseCopy(common.ImpulseIDKey, c)
			if err != nil {
				c.Log.V(3).Errorf("%#v", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}

			newStateConfig := state.DefaultConfig()
			newStateConfig.Bytes["request"] = request
			newStateConfig.ObjectID = spec.ObjectID(newSignal.GetID())
			newStateConfig.ObjectType = impulse.ObjectType

			newImpulse.SetState(common.DefaultStateKey, state.NewState(newStateConfig))

			resImpulse, err := c.Trigger(newImpulse)
			if err != nil {
				c.Log.V(3).Errorf("%#v", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}

			impState, err := resImpulse.GetState(common.DefaultStateKey)
			if err != nil {
				c.Log.V(3).Errorf("%#v", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}
			response, err := impState.GetBytes("response")
			if err != nil {
				c.Log.V(3).Errorf("%#v", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				return
			}
			newSignal.SetBytes("response", response)
			responder <- newSignal
		}(newSignal)
	}
}

func (c *core) SetState(key string, state spec.State) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.States[key] = state
}

func (c *core) Shutdown() {
	c.Log.V(12).Debugf("call Core.Shutdown")

	// TODO close gateway
	// TODO stop listening
	// TODO wait for impulses to be processed
	// TODO backup state
}

func (c *core) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	c.Log.V(12).Debugf("call Core.Trigger")

	// Track state.
	impState, err := imp.GetState(common.DefaultStateKey)
	if err != nil {
		return nil, maskAny(err)
	}
	impState.SetCore(c)
	coreState, err := c.GetState(common.DefaultStateKey)
	if err != nil {
		return nil, maskAny(err)
	}
	coreState.SetImpulse(imp)

	// Initialize network within core state if not already done.
	defaultCoreState, err := c.GetState(common.DefaultStateKey)
	if err != nil {
		return nil, maskAny(err)
	}
	networks := defaultCoreState.GetNetworks()
	if len(networks) == 0 {
		initCoreState, err := c.GetState(common.InitStateKey)
		if err != nil {
			return nil, maskAny(err)
		}
		networkID, err := initCoreState.GetBytes(common.NetworkIDKey)
		if err != nil {
			return nil, maskAny(err)
		}
		network, err := initCoreState.GetNetworkByID(spec.ObjectID(networkID))
		if err != nil {
			return nil, maskAny(err)
		}
		defaultCoreState.SetNetwork(network.Copy())
	}

	// Get network. Note that there is potential for multiple networks. For now
	// we just have one.
	coreState, err = c.GetState(common.DefaultStateKey)
	if err != nil {
		return nil, maskAny(err)
	}
	for _, n := range coreState.GetNetworks() {
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
