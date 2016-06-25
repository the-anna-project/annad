package core

import (
	"reflect"
	"sync/atomic"

	"github.com/xh3b4sd/anna/spec"
)

func (n *network) gatewayListener(newSignal spec.Signal) (spec.Signal, error) {
	input := newSignal.GetInput()

	newImpulse, err := n.NewImpulse(input.(spec.CoreRequest))
	if err != nil {
		return nil, maskAny(err)
	}

	// TODO ImpulsesInProgress should be an exposed metric.
	//
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

func prepareInput(stage int, input string) []reflect.Value {
	values := []reflect.Value{
		reflect.ValueOf(stage),    // stage
		reflect.ValueOf("string"), // input type
		reflect.ValueOf(input),    // input value
	}

	return values
}

func prepareOutput(values []reflect.Value) string {
	var output string

	for _, v := range values {
		output += v.String()
	}

	return output
}
