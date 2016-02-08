package common

import (
	"github.com/xh3b4sd/anna/spec"
)

type objectType struct {
	Core            spec.ObjectType
	Impulse         spec.ObjectType
	FactoryClient   spec.ObjectType
	FactoryServer   spec.ObjectType
	CharacterNeuron spec.ObjectType
	FirstNeuron     spec.ObjectType
	JobNeuron       spec.ObjectType
	Main            spec.ObjectType
	Network         spec.ObjectType
	None            spec.ObjectType
	Server          spec.ObjectType
	State           spec.ObjectType
	TextInterface   spec.ObjectType
}

var (
	ObjectType = objectType{
		Core:            spec.ObjectType("core"),
		Impulse:         spec.ObjectType("impulse"),
		FactoryClient:   spec.ObjectType("factory-client"),
		FactoryServer:   spec.ObjectType("factory-server"),
		CharacterNeuron: spec.ObjectType("character-neuron"),
		FirstNeuron:     spec.ObjectType("first-neuron"),
		JobNeuron:       spec.ObjectType("job-neuron"),
		Main:            spec.ObjectType("main"),
		Network:         spec.ObjectType("network"),
		None:            spec.ObjectType("none"),
		Server:          spec.ObjectType("server"),
		State:           spec.ObjectType("state"),
		TextInterface:   spec.ObjectType("text-interface"),
	}
)
