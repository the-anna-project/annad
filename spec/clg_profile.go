package spec

import (
	"encoding/json"
	"reflect"
)

// CLGProfile contains information of a certain CLG.
type CLGProfile interface {
	// Equals checks whether the current CLG profile is equal to the given one.
	Equals(CLGProfile) bool

	GetHash() string

	GetInputTypes() []reflect.Kind

	GetInputExamples() []interface{}

	GetMethodName() string

	GetMethodBody() string

	GetOutputTypes() []reflect.Kind

	GetOutputExamples() []interface{}

	// TODO comment
	GetRightSideNeighbours() []string

	json.Marshaler

	json.Unmarshaler
}
