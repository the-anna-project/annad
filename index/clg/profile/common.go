package profile

import (
	"reflect"
)

func (g *generator) isMethodValue(v reflect.Value) bool {
	if !v.IsValid() {
		return false
	}

	if v.Kind() != reflect.Func {
		return false
	}

	return true
}
