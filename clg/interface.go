package clg

import (
	"reflect"
)

func (i *clgIndex) DiscardInterface(args ...interface{}) ([]interface{}, error) {
	return nil, nil
}

func (i *clgIndex) EqualInterface(args ...interface{}) ([]interface{}, error) {
	if len(args) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 2 got %d", len(args))
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	t := reflect.DeepEqual(args[0], args[1])

	return []interface{}{t}, nil
}

func (i *clgIndex) TypeInterface(args ...interface{}) ([]interface{}, error) {
	if len(args) < 1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 1 got %d", len(args))
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	t := reflect.TypeOf(args[0]).String()

	return []interface{}{t}, nil
}
