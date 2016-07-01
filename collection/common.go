package clg

import (
	"reflect"

	"github.com/xh3b4sd/anna/clg/find-connections"
	"github.com/xh3b4sd/anna/spec"
)

func (c *Collection) configureCLGs() {
	for clgName, _ := range c.CLGs {
		c.CLGs[clgName].SetStorage(c.Storage)
	}
}

func (c *Collection) getMethodValue(name spec.CLG) (reflect.Value, error) {
	n := string(name)
	v := reflect.ValueOf(c).MethodByName(n)
	if !v.IsValid() {
		return reflect.Value{}, maskAnyf(methodNotFoundError, n)
	}

	return v, nil
}

func (c *Collection) newCLGs() map[string]spec.CLG {
	newList := []spec.CLG{
		findconnections.MustNew(),
	}

	newCLGs := map[string]spec.CLG{}

	for _, CLG := range newList {
		newCLGs[CLG.GetType()] = CLG
	}

	return newCLGs
}
