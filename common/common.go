// TODO
package common

import (
	"github.com/xh3b4sd/anna/spec"
)

func MustObjectToImpulse(object spec.Object) spec.Impulse {
	if i, ok := object.(spec.Impulse); ok {
		return i
	}

	panic(objectNotImpulseError)
}

func MustObjectToNeuron(object spec.Object) spec.Neuron {
	if i, ok := object.(spec.Neuron); ok {
		return i
	}

	panic(objectNotNeuronError)
}

func MustObjectToNetwork(object spec.Object) spec.Network {
	if i, ok := object.(spec.Network); ok {
		return i
	}

	panic(objectNotNetworkError)
}

func MustObjectToCore(object spec.Object) spec.Core {
	if i, ok := object.(spec.Core); ok {
		return i
	}

	panic(objectNotCoreError)
}
