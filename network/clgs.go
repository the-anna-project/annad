package network

import (
	"reflect"

	"github.com/xh3b4sd/anna/clg/find-connections"
	"github.com/xh3b4sd/anna/spec"
)

// receiver

func (n *network) configureCLGs(CLGs []spec.CLG) []spec.CLG {
	for name, _ := range CLGs {
		CLGs[name].SetStorage(n.Storage)
	}

	return CLGs
}

// private

type clgScope struct {
	CLG    spec.CLG
	Input  chan []reflect.Value
	Output chan []reflect.Value
}

func newCLGs() map[string]spec.CLG {
	newList := []spec.CLG{
		findconnections.MustNew(),
	}

	newCLGs := map[string]clgScope{}

	for _, CLG := range newList {
		newCLGs[CLG.GetType()] = clgScope{
			CLG:    CLG,
			Input:  make(chan []reflect.Value, 10),
			Output: make(chan []reflect.Value, 10),
		}
	}

	return newCLGs
}
