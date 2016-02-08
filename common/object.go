package common

import (
	"github.com/xh3b4sd/anna/spec"
)

// Note that when adding a new object type to this struct, the two definitions
// below need to be kept in sync as well.
type objectType struct {
	Core            spec.ObjectType
	Impulse         spec.ObjectType
	FactoryClient   spec.ObjectType
	FactoryServer   spec.ObjectType
	CharacterNeuron spec.ObjectType
	FirstNeuron     spec.ObjectType
	JobNeuron       spec.ObjectType
	Log             spec.ObjectType
	Main            spec.ObjectType
	Network         spec.ObjectType
	None            spec.ObjectType
	Server          spec.ObjectType
	State           spec.ObjectType
	TextInterface   spec.ObjectType
}

var (
	// Note that this struct needs to be in sync with the below list. E.g. the
	// logger checks for valid object types using this list.
	ObjectType = objectType{
		Core:            spec.ObjectType("core"),
		Impulse:         spec.ObjectType("impulse"),
		FactoryClient:   spec.ObjectType("factory-client"),
		FactoryServer:   spec.ObjectType("factory-server"),
		CharacterNeuron: spec.ObjectType("character-neuron"),
		FirstNeuron:     spec.ObjectType("first-neuron"),
		JobNeuron:       spec.ObjectType("job-neuron"),
		Log:             spec.ObjectType("log"),
		Main:            spec.ObjectType("main"),
		Network:         spec.ObjectType("network"),
		None:            spec.ObjectType("none"),
		Server:          spec.ObjectType("server"),
		State:           spec.ObjectType("state"),
		TextInterface:   spec.ObjectType("text-interface"),
	}

	// Note that this list needs to be in sync with the above struct. E.g. the
	// logger checks for valid object types using this list.
	ObjectTypes = []spec.ObjectType{
		ObjectType.Core,
		ObjectType.Impulse,
		ObjectType.FactoryClient,
		ObjectType.FactoryServer,
		ObjectType.CharacterNeuron,
		ObjectType.FirstNeuron,
		ObjectType.JobNeuron,
		ObjectType.Log,
		ObjectType.Main,
		ObjectType.Network,
		ObjectType.None,
		ObjectType.Server,
		ObjectType.State,
		ObjectType.TextInterface,
	}
)
