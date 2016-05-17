package profile

import (
	"reflect"

	"github.com/xh3b4sd/anna/spec"
)

var (
	// NumArgs is an ordered list of numbers used to find out how many input
	// arguments a CLG expects. Usually CLGs do not expect more than 5 input
	// arguments. For special cases we try to find out how many they expect
	// beyond 5 arguments. Here we assume that a CLG might expect 10 or even 50
	// arguments. In case a CLG expects 50 or more arguments, we assume it
	// expects infinite arguments.
	NumArgs = []int{0, 1, 2, 3, 4, 5, 10, 20, 30, 40, 50}

	// ArgTypes represents a list of well known types used to identify CLG input-
	// and output types. Here we want to have a list of types only. That is why
	// we make use of the builtin function new. It takes a type and returns an
	// initialized zero value pointer to that type. That way we don't need to
	// initialized values or need to know what zero values certain types provide.
	// Further we prevent messy type assertions like int(0), int8(0), int16(0),
	// etc..
	ArgTypes = []interface{}{
		// Simple types.
		*new(string),
		*new(bool),
		*new(int),
		*new(int64),
		*new(float64),
		*new(interface{}),
		*new(string),
		*new(struct{}),

		// Slices of simple types.
		*new([]string),
		*new([]bool),
		*new([]int),
		*new([]int64),
		*new([]float64),
		*new([]interface{}),
		*new([]string),
		*new([]struct{}),

		// Slices of slices of simple types.
		*new([][]string),
		*new([][]bool),
		*new([][]int),
		*new([][]int64),
		*new([][]float64),
		*new([][]interface{}),
		*new([][]string),
		*new([][]struct{}),
	}
)

func (g *generator) CreateInputs(clgName string) ([]string, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateInputs")

	methodValue := reflect.ValueOf(g.Collection).MethodByName(clgName)
	if !g.isMethodValue(methodValue) {
		return nil, maskAnyf(invalidCLGError, clgName)
	}
	t := methodValue.Type()

	var newInputs []string

	for i := 0; i < t.NumIn(); i++ {
		newInputs = append(newInputs, t.In(i).String())
	}

	return newInputs, nil
}
