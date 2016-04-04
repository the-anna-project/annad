package clg

import (
	"fmt"
	"reflect"
)

// Arg

func ArgToInt(args []interface{}, index int) (int, error) {
	if len(args) < index+1 {
		return 0, maskAnyf(notEnoughArgumentsError, "expected %d got %d", index+1, len(args))
	}

	if i, ok := args[index].(int); ok {
		return i, nil
	}

	return 0, maskAnyf(wrongArgumentTypeError, "expected %T got %T", "", args[index])
}

func ArgToString(args []interface{}, index int) (string, error) {
	if len(args) < index+1 {
		return "", maskAnyf(notEnoughArgumentsError, "expected %d got %d", index+1, len(args))
	}

	if s, ok := args[index].(string); ok {
		return s, nil
	}

	return "", maskAnyf(wrongArgumentTypeError, "expected %T got %T", "", args[index])
}

func ArgToStringSlice(args []interface{}, index int) ([]string, error) {
	if len(args) < index+1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d got %d", index+1, len(args))
	}

	if ss, ok := args[index].([]string); ok {
		return ss, nil
	}

	return nil, maskAnyf(wrongArgumentTypeError, "expected %T got %T", "", args[index])
}

// Args

func ArgsToValues(args []interface{}) []reflect.Value {
	values := make([]reflect.Value, len(args))

	for i := range args {
		values[i] = reflect.ValueOf(args[i])
	}

	return values
}

func ValuesToArgs(values []reflect.Value) ([]interface{}, error) {
	fmt.Printf("values: %#v\n", values)
	if len(values) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(values))
	}
	if len(values) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 2 got %d", len(values))
	}

	if !values[1].IsValid() {
		fmt.Print(1)
		return values[0].Interface().([]interface{}), nil
	} else {
		fmt.Print(2)
		return nil, maskAny(values[1].Interface().(error))
	}
}
