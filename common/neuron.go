package common

import (
	"github.com/xh3b4sd/anna/spec"
)

const (
	CharacterNeuronIDKey = "character-neuron-id"
	FirstNeuronIDKey     = "first-neuron-id"
	JobNeuronIDKey       = "job-neuron-id"
)

func MustObjectToNeuron(object spec.Object) spec.Neuron {
	if i, ok := object.(spec.Neuron); ok {
		return i
	}

	panic(objectNotNeuronError)
}

func GetInitNeuronCopy(key string, object spec.Object) (spec.Neuron, error) {
	objectState, err := object.GetState(InitStateKey)
	if err != nil {
		return nil, maskAny(err)
	}
	bytes, err := objectState.GetBytes(key)
	if err != nil {
		return nil, maskAny(err)
	}
	initNeuron, err := objectState.GetNeuronByID(spec.ObjectID(bytes))
	if err != nil {
		return nil, maskAny(err)
	}

	return initNeuron.Copy(), nil
}
