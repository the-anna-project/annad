package profile

import (
	"reflect"

	"github.com/xh3b4sd/anna/key"
)

func (g *generator) key(f string, v ...interface{}) string {
	return key.NewSysKey(g, f, v...)
}

func (g *generator) isMethodValue(v reflect.Value) bool {
	if !v.IsValid() {
		return false
	}

	if v.Kind() != reflect.Func {
		return false
	}

	return true
}
