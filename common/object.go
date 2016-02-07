package common

import (
	"fmt"

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
		Core:            spec.ObjectType(fmt.Sprintf("%-16s", "core")),
		Impulse:         spec.ObjectType(fmt.Sprintf("%-16s", "impulse")),
		FactoryClient:   spec.ObjectType(fmt.Sprintf("%-16s", "factory-client")),
		FactoryServer:   spec.ObjectType(fmt.Sprintf("%-16s", "factory-server")),
		CharacterNeuron: spec.ObjectType(fmt.Sprintf("%-16s", "character-neuron")),
		FirstNeuron:     spec.ObjectType(fmt.Sprintf("%-16s", "first-neuron")),
		JobNeuron:       spec.ObjectType(fmt.Sprintf("%-16s", "job-neuron")),
		Main:            spec.ObjectType(fmt.Sprintf("%-16s", "main")),
		Network:         spec.ObjectType(fmt.Sprintf("%-16s", "network")),
		None:            spec.ObjectType(fmt.Sprintf("%-16s", "none")),
		Server:          spec.ObjectType(fmt.Sprintf("%-16s", "server")),
		State:           spec.ObjectType(fmt.Sprintf("%-16s", "state")),
		TextInterface:   spec.ObjectType(fmt.Sprintf("%-16s", "text-interface")),
	}
)
