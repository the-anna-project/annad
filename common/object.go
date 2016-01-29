package common

import (
	"github.com/xh3b4sd/anna/spec"
)

type objectType struct {
	Core            spec.ObjectType
	Impulse         spec.ObjectType
	CharacterNeuron spec.ObjectType
	FirstNeuron     spec.ObjectType
	JobNeuron       spec.ObjectType
	Network         spec.ObjectType
	None            spec.ObjectType
	State           spec.ObjectType
}

var (
	ObjectType = objectType{
		Core:            spec.ObjectType("core"),
		Impulse:         spec.ObjectType("impulse"),
		CharacterNeuron: spec.ObjectType("character-neuron"),
		FirstNeuron:     spec.ObjectType("first-neuron"),
		JobNeuron:       spec.ObjectType("job-neuron"),
		Network:         spec.ObjectType("network"),
		None:            spec.ObjectType("none"),
		State:           spec.ObjectType("state"),
	}
)
