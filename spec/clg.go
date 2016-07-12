package spec

import (
	"reflect"
)

// CLG TODO
type CLG interface {
	// TODO
	Calculate(inputs []reflect.Value) ([]reflect.Value, error)

	GetName() string

	Inputs() []reflect.Type

	Object

	SetLog(log Log)

	SetStorage(storage Storage)
}
