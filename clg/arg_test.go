package clg

import (
	"reflect"
	"testing"
)

func Test_ValuesToArgs(t *testing.T) {
	testCases := []struct {
		Input        []reflect.Value
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []reflect.Value{reflect.ValueOf([]interface{}{[]string{"a", "b", "c"}, 3.8, true}), reflect.ValueOf(nil)},
			Expected:     []interface{}{[]string{"a", "b", "c"}, 3.8, true},
			ErrorMatcher: nil,
		},
		{
			Input:        []reflect.Value{reflect.ValueOf(nil), reflect.ValueOf(wrongArgumentTypeError)},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []reflect.Value{reflect.ValueOf([]interface{}{[]string{"a", "b", "c"}, 3.8, true})},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []reflect.Value{reflect.ValueOf([]interface{}{[]string{"a", "b", "c"}, 3.8, true}), reflect.ValueOf(nil), reflect.ValueOf(nil)},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := ValuesToArgs(testCase.Input)
		if testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}
