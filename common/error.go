package common

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var objectNotImpulseError = errgo.New("object not impulse")

func IsObjectNotImpulse(err error) bool {
	return errgo.Cause(err) == objectNotImpulseError
}

var objectNotNeuronError = errgo.New("object not neuron")

func IsObjectNotNeuron(err error) bool {
	return errgo.Cause(err) == objectNotNeuronError
}

var objectNotNetworkError = errgo.New("object not network")

func IsObjectNotNetwork(err error) bool {
	return errgo.Cause(err) == objectNotNetworkError
}

var objectNotCoreError = errgo.New("object not core")

func IsObjectNotCore(err error) bool {
	return errgo.Cause(err) == objectNotCoreError
}
