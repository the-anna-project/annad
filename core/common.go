package core

import (
	"sync/atomic"

	"github.com/xh3b4sd/anna/spec"
)

const (
	objectTypeStrategyNetwork spec.ObjectType = "strategy-network"
)

func (c *core) bootObjectTree() {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call bootObjectTree")

	errorHandler := func(err error) {
		c.Log.WithTags(spec.Tags{L: "F", O: c, T: nil, V: 1}, "%#v", maskAny(err))
	}

	_, err := c.GetNetworkByType(objectTypeStrategyNetwork)
	if IsNetworkNotFound(err) {
		strategyNetwork, err := c.FactoryClient.NewStrategyNetwork()
		if err != nil {
			errorHandler(maskAny(err))
		}
		err = c.SetNetworkByType(objectTypeStrategyNetwork, strategyNetwork)
		if err != nil {
			errorHandler(maskAny(err))
		}
	} else if err != nil {
		errorHandler(maskAny(err))
	}
}

func (c *core) gatewayListener(newSignal spec.Signal) (spec.Signal, error) {
	ID := newSignal.GetID()
	input := newSignal.GetInput()

	newImpulse, err := c.FactoryClient.NewImpulse()
	if err != nil {
		return nil, maskAny(err)
	}
	err = newImpulse.SetID(spec.ObjectID(ID))
	if err != nil {
		return nil, maskAny(err)
	}
	err = newImpulse.SetInput(input.(string))
	if err != nil {
		return nil, maskAny(err)
	}

	// Increment the impulse count to track how many impulses are processed
	// inside the core.
	c.ImpulsesInProgress = atomic.AddInt64(&c.ImpulsesInProgress, 1)

	newImpulse, err = c.Trigger(newImpulse)

	// Decrement the impulse count once all hard work is done. Note that this
	// is important to be done before the error handling of Core.Trigger to
	// ensure the impulse count is properly decreased.
	c.ImpulsesInProgress = atomic.AddInt64(&c.ImpulsesInProgress, -1)

	if err != nil {
		return nil, maskAny(err)
	}

	output, err := newImpulse.GetOutput()
	if err != nil {
		return nil, maskAny(err)
	}

	newSignal.SetOutput(output)

	return newSignal, nil
}
