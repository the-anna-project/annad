package corenet

import (
	"sync/atomic"

	"github.com/xh3b4sd/anna/spec"
)

func (cn *coreNet) bootObjectTree() {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call bootObjectTree")

	cn.EvalNet.Boot()
	cn.ExecNet.Boot()
	cn.PatNet.Boot()
	cn.PredNet.Boot()
	cn.StratNet.Boot()
}

func (cn *coreNet) gatewayListener(newSignal spec.Signal) (spec.Signal, error) {
	input := newSignal.GetInput()

	newImpulse, err := cn.FactoryClient.NewImpulse()
	if err != nil {
		return nil, maskAny(err)
	}
	newImpulse.SetInputByImpulseID(newImpulse.GetID(), input.(string))

	// Increment the impulse count to track how many impulses are processed
	// inside the core network.
	atomic.AddInt64(&cn.ImpulsesInProgress, 1)
	newImpulse, err = cn.Trigger(newImpulse)
	// Decrement the impulse count once all hard work is done. Note that this
	// is important to be done before the error handling of Core.Trigger to
	// ensure the impulse count is properly decreased.
	atomic.AddInt64(&cn.ImpulsesInProgress, -1)

	if err != nil {
		return nil, maskAny(err)
	}

	output := newImpulse.GetOutput()
	newSignal.SetOutput(output)

	return newSignal, nil
}
