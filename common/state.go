package common

import (
	"github.com/xh3b4sd/anna/spec"
)

var (
	DefaultStateFile = "state.json"

	StateType = stateType{
		FSReader:   spec.StateType("fs"),
		FSWriter:   spec.StateType("fs"),
		NoneReader: spec.StateType("none"),
		NoneWriter: spec.StateType("none"),
	}
)

type stateType struct {
	FSReader   spec.StateType
	FSWriter   spec.StateType
	NoneReader spec.StateType
	NoneWriter spec.StateType
}
