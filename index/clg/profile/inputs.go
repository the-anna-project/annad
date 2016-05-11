package profile

import (
	"reflect"

	"github.com/xh3b4sd/anna/spec"
)

func (g *generator) CreateInputs(clgName string) ([]reflect.Kind, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateInputs")

	methodValue := reflect.ValueOf(g.Collection).MethodByName(clgName)
	if !g.isMethodValue(methodValue) {
		return nil, maskAnyf(invalidCLGError, clgName)
	}
	t := methodValue.Type()

	var newInputs []reflect.Kind

	for i := 0; i < t.NumIn(); i++ {
		newInputs = append(newInputs, t.In(i).Kind())
	}

	return newInputs, nil
}
