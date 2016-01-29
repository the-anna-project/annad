package common

import (
	"github.com/xh3b4sd/anna/spec"
)

func MustObjectToNeuron(object spec.Object) spec.Neuron {
	if i, ok := object.(spec.Neuron); ok {
		return i
	}

	panic(objectNotNeuronError)
}
