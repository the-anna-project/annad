package network

import (
	"reflect"
	"sync/atomic"

	"github.com/xh3b4sd/anna/clg/find-connections"
	"github.com/xh3b4sd/anna/spec"
)

func (n *network) configureCLGs() {
	for clgName, _ := range n.CLGs {
		n.CLGs[clgName].SetStorage(n.Storage)
	}
}

func (n *network) getGatewayListener() func(newSignal spec.Signal) (spec.Signal, error) {
	newListener := func(newSignal spec.Signal) (spec.Signal, error) {
		input := newSignal.GetInput()

		newImpulse, err := n.NewImpulse()
		if err != nil {
			return nil, maskAny(err)
		}
		newImpulse.SetInputByImpulseID(newImpulse.GetID(), input.(string))

		// Increment the impulse count to track how many impulses are processed
		// inside the core network.
		atomic.AddInt64(&n.ImpulsesInProgress, 1)
		newImpulse, err = n.Trigger(newImpulse)
		// Decrement the impulse count once all hard work is done. Note that this
		// is important to be done before the error handling of Core.Trigger to
		// ensure the impulse count is properly decreased.
		atomic.AddInt64(&n.ImpulsesInProgress, -1)

		if err != nil {
			return nil, maskAny(err)
		}

		output := newImpulse.GetOutput()
		newSignal.SetOutput(output)

		return newSignal, nil
	}

	return newListener
}

func (n *network) getMethodValue(name string) (reflect.Value, error) {
	v := reflect.ValueOf(n).MethodByName(name)
	if !v.IsValid() {
		return reflect.Value{}, maskAnyf(methodNotFoundError, n)
	}

	return v, nil
}

func (n *network) getNewCLGs() map[string]spec.CLG {
	newList := []spec.CLG{
		findconnections.MustNew(),
	}

	newCLGs := map[string]spec.CLG{}

	for _, CLG := range newList {
		newCLGs[CLG.GetType()] = CLG
	}

	return newCLGs
}
