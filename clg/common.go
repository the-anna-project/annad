package clg

import (
	"reflect"
)

// argsToValues converts the given []interface{} to reflect values. The given
// args can be the output of a node. That is, the output of a strategy
// execution. The converted list of reflect values can be used to provide it as
// input of a CLG.
func argsToValues(args []interface{}) []reflect.Value {
	values := make([]reflect.Value, len(args))

	for i := range args {
		values[i] = reflect.ValueOf(args[i])
	}

	return values
}
