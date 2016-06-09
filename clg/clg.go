// Package clg provides CLGs that implement basic behaviour.
package clg

import (
	"reflect"

	"github.com/xh3b4sd/anna/spec"
)

// Collection represents the object holding all available CLGs. This is only
// intended to be used package internally and for documentation reasons. The
// clg package implements a global and stateless collection that is used in
// Execute and Names.
type Collection struct{}

var (
	collection = Collection{}
)

// Execute is convenient to call CLGs by their method name and abstracted input
// values.
func Execute(name spec.CLG, inputs []reflect.Value) ([]reflect.Value, error) {
	n := string(name)
	v := reflect.ValueOf(collection).MethodByName(n)
	if !v.IsValid() {
		return nil, maskAnyf(methodNotFoundError, n)
	}

	outputs := v.Call(inputs)

	return outputs, nil
}

// Inputs returns the input types of the CLG identified by name.
func Inputs(name spec.CLG) ([]reflect.Type, error) {
	n := string(name)
	v := reflect.ValueOf(collection).MethodByName(n)
	if !v.IsValid() {
		return nil, maskAnyf(methodNotFoundError, n)
	}
	t := v.Type()
	var inputs []reflect.Type

	for i := 0; i < t.NumIn(); i++ {
		inputs = append(inputs, t.In(i))
	}

	return inputs, nil
}

// Outputs returns the output types of the CLG identified by name.
func Outputs(name spec.CLG) ([]reflect.Type, error) {
	n := string(name)
	v := reflect.ValueOf(collection).MethodByName(n)
	if !v.IsValid() {
		return nil, maskAnyf(methodNotFoundError, n)
	}
	t := v.Type()
	var outputs []reflect.Type

	for i := 0; i < t.NumOut(); i++ {
		outputs = append(outputs, t.Out(i))
	}

	return outputs, nil
}

// Names returns all available CLG method names.
func Names() []spec.CLG {
	t := reflect.TypeOf(collection)
	var names []spec.CLG

	for i := 0; i < t.NumMethod(); i++ {
		names = append(names, spec.CLG(t.Method(i).Name))
	}

	return names
}
