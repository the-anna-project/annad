package spec

import (
	"reflect"
)

// CLG TODO
type CLG interface {
	// TODO
	Execute(inputs []reflect.Value) ([]reflect.Value, error)

	GetName() string

	Inputs() []reflect.Type

	Object

	SetStorage(storage Storage)
}
