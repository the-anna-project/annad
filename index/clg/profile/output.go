package profile

import (
	"reflect"

	"github.com/xh3b4sd/anna/spec"
)

func (g *generator) CreateOutputs(clgName string) ([]reflect.Kind, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateOutputs")

	methodValue := reflect.ValueOf(g.Collection).MethodByName(clgName)
	if !g.isMethodValue(methodValue) {
		return nil, maskAnyf(invalidCLGError, clgName)
	}
	t := methodValue.Type()

	var newOutputs []reflect.Kind

	for i := 0; i < t.NumOut(); i++ {
		newOutputs = append(newOutputs, t.Out(i).Kind())
	}

	return newOutputs, nil
}
