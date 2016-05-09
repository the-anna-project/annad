package profile

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"

	"github.com/xh3b4sd/anna/index/clg/collection"
	"github.com/xh3b4sd/anna/spec"
)

func profileNamesFromCollection(collection spec.CLGCollection) ([]string, error) {
	args, err := collection.GetNamesMethod()
	if err != nil {
		return nil, maskAny(err)
	}
	newNames, err := collection.ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	return newNames, nil
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
