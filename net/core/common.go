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
	ID := newSignal.GetID()
	input := newSignal.GetInput()

	newImpulse, err := cn.FactoryClient.NewImpulse()
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
	// inside the core network. Note that the laborious assignment makes `go vet`
	// happy.
	v := atomic.AddInt64(&cn.ImpulsesInProgress, 1)
	cn.ImpulsesInProgress = v

	newImpulse, err = cn.Trigger(newImpulse)

	// Decrement the impulse count once all hard work is done. Note that this
	// is important to be done before the error handling of Core.Trigger to
	// ensure the impulse count is properly decreased.
	v = atomic.AddInt64(&cn.ImpulsesInProgress, -1)
	cn.ImpulsesInProgress = v

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
