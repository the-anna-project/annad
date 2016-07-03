package network

import (
	"reflect"
	"sync/atomic"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/clg/divide"
	"github.com/xh3b4sd/anna/clg/find-connections"
	"github.com/xh3b4sd/anna/spec"
)

// receiver

func (n *network) configureCLGs(CLGs map[spec.ObjectID]clgScope) map[spec.ObjectID]clgScope {
	for name := range CLGs {
		CLGs[name].CLG.SetLog(n.Log)
		CLGs[name].CLG.SetStorage(n.Storage)
	}

	return CLGs
}

func (n *network) getGatewayListener() func(newSignal spec.Signal) (spec.Signal, error) {
	newListener := func(newSignal spec.Signal) (spec.Signal, error) {
		newImpulse, err := n.NewImpulse(newSignal.GetInput().(api.CoreRequest))
		if err != nil {
			return nil, maskAny(err)
		}

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

// private

type clgScope struct {
	CLG    spec.CLG
	Input  chan []reflect.Value
	Output chan []reflect.Value
}

func newCLGs() map[spec.ObjectID]clgScope {
	newList := []spec.CLG{
		divide.MustNew(),
		findconnections.MustNew(),
	}

	newCLGs := map[spec.ObjectID]clgScope{}

	for _, CLG := range newList {
		newCLGs[CLG.GetID()] = clgScope{
			CLG:    CLG,
			Input:  make(chan []reflect.Value, 10),
			Output: make(chan []reflect.Value, 10),
		}
	}

	return newCLGs
}

func equalInputs(provided []reflect.Value, implemented []reflect.Type) bool {
	var p []string
	for _, v := range provided {
		p = append(p, v.Type().String())
	}

	var i []string
	for _, t := range implemented {
		i = append(i, t.String())
	}

	if !reflect.DeepEqual(p, i) {
		return false
	}

	return true
}

func prepareInput(imp spec.Impulse) []reflect.Value {
	values := []reflect.Value{reflect.ValueOf(imp), reflect.ValueOf(imp.GetInput())}
	return values
}

func prepareOutput(values []reflect.Value) (spec.Impulse, error) {
	if len(values) == 0 {
		return nil, maskAnyf(invalidInterfaceError, "output must not be empty")
	}

	imp, ok := values[0].Interface().(spec.Impulse)
	if !ok {
		return nil, maskAnyf(invalidInterfaceError, "impulse must be first")
	}

	var output string
	for _, v := range values[1:] {
		output += v.String()
	}
	imp.SetOutput(output)

	return imp, nil
}
