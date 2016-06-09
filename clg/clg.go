// Package clg provides CLGs that implement basic behaviour.
package clg

import (
	"reflect"
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
func Execute(name string, inputs []reflect.Value) ([]reflect.Value, error) {
	v := reflect.ValueOf(collection).MethodByName(name)
	if !v.IsValid() {
		return nil, maskAnyf(methodNotFoundError, name)
	}

	outputs := v.Call(inputValues)

	return outputs, nil
}

// Inputs returns the input types of the CLG identified by name.
func Inputs(name string) ([]reflect.Type, error) {
	v := reflect.ValueOf(collection).MethodByName(name)
	if !v.IsValid() {
		return nil, maskAnyf(methodNotFoundError, name)
	}
	t := v.Type()
	var inputs []reflect.Type

	for i := 0; i < t.NumIn(); i++ {
		inputs = append(inputs, t.In(i))
	}

	return inputs
}

// Outputs returns the output types of the CLG identified by name.
func Outputs(name string) ([]reflect.Type, error) {
	v := reflect.ValueOf(collection).MethodByName(name)
	if !v.IsValid() {
		return nil, maskAnyf(methodNotFoundError, name)
	}
	t := v.Type()
	var outputs []reflect.Type

	for i := 0; i < t.NumOut(); i++ {
		outputs = append(outputs, t.Out(i))
	}

	return outputs
}

// Names returns all available CLG method names.
func Names() []string {
	t := reflect.TypeOf(collection)
	var names []string

	for i := 0; i < t.NumMethod(); i++ {
		names = append(names, t.Method(i).Name)
	}

	return names
}
