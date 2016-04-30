package clg

import (
	"reflect"
	"strings"
)

func (c *clgCollection) CallMethodByName(args ...interface{}) ([]interface{}, error) {
	methodName, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	inputValues := ArgsToValues(args[1:])
	methodValue := reflect.ValueOf(c).MethodByName(methodName)
	if !methodValue.IsValid() {
		return nil, maskAnyf(methodNotFoundError, methodName)
	}

	outputValues := methodValue.Call(inputValues)
	results, err := ValuesToArgs(outputValues)
	if err != nil {
		return nil, maskAny(err)
	}

	return results, nil
}

func (c *clgCollection) GetMethodNames(args ...interface{}) ([]interface{}, error) {
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	var pattern string
	if len(args) == 1 {
		var err error
		pattern, err = ArgToString(args, 0)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	var allMethodNames []string

	t := reflect.TypeOf(c)
	for i := 0; i < t.NumMethod(); i++ {
		methodName := t.Method(i).Name
		if pattern != "" && !strings.Contains(methodName, pattern) {
			continue
		}
		allMethodNames = append(allMethodNames, methodName)
	}

	return []interface{}{allMethodNames}, nil
}

func (c *clgCollection) GetNumMethods(args ...interface{}) ([]interface{}, error) {
	if len(args) > 0 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 0 got %d", len(args))
	}

	t := reflect.TypeOf(c)
	num := t.NumMethod()

	return []interface{}{num}, nil
}
