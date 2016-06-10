package clg

import (
	"reflect"

	"github.com/xh3b4sd/anna/spec"
)

func getMethodValue(name spec.CLG) (reflect.Value, error) {
	n := string(name)
	v := reflect.ValueOf(collection).MethodByName(n)
	if !v.IsValid() {
		return reflect.Value{}, maskAnyf(methodNotFoundError, n)
	}

	return v, nil
}
