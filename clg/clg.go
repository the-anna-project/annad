// Package clg provides CLGs that implement basic behaviour.
package clg

import (
	"reflect"

	"github.com/xh3b4sd/anna/spec"
)

// ExecuteCLG is convenient to call CLGs by their method name and abstracted
// input values.
func (c *Collection) ExecuteCLG(name spec.CLG, inputs []reflect.Value) ([]reflect.Value, error) {
	v, err := c.getMethodValue(name)
	if err != nil {
		return nil, maskAny(err)
	}

	outputs := v.Call(inputs)

	return outputs, nil
}

// CLGInputs returns the input types of the CLG identified by name.
func (c *Collection) CLGInputs(name spec.CLG) ([]reflect.Type, error) {
	v, err := c.getMethodValue(name)
	if err != nil {
		return nil, maskAny(err)
	}

	var inputs []reflect.Type

	t := v.Type()
	for i := 0; i < t.NumIn(); i++ {
		inputs = append(inputs, t.In(i))
	}

	return inputs, nil
}

// CLGOutputs returns the output types of the CLG identified by name.
func (c *Collection) CLGOutputs(name spec.CLG) ([]reflect.Type, error) {
	v, err := c.getMethodValue(name)
	if err != nil {
		return nil, maskAny(err)
	}

	var outputs []reflect.Type

	t := v.Type()
	for i := 0; i < t.NumOut(); i++ {
		outputs = append(outputs, t.Out(i))
	}

	return outputs, nil
}

// CLGNames returns all available CLG method names.
func (c *Collection) CLGNames() []spec.CLG {
	t := reflect.TypeOf(c)
	var names []spec.CLG

	for i := 0; i < t.NumMethod(); i++ {
		names = append(names, spec.CLG(t.Method(i).Name))
	}

	return names
}
