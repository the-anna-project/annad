package clg

import (
	"reflect"
)

func (i *index) ArgType(args ...interface{}) ([]interface{}, error) {
	if len(args) < 1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 1 got %d", len(args))
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	t := reflect.TypeOf(args[0]).String()

	return []interface{}{t}, nil
}
