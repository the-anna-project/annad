package clg

import (
	"reflect"
	"testing"
)

func Test_Type_ArgType(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"ab"},
			Expected:     []interface{}{"string"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{23},
			Expected:     []interface{}{"int"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{true},
			Expected:     []interface{}{"bool"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{23.8},
			Expected:     []interface{}{"float64"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{map[int]bool{4: true}},
			Expected:     []interface{}{"map[int]bool"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab", "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"ab", 23},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"ab", false},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			// Note all args are missing.
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	newConfig := DefaultConfig()
	newIndex, err := NewIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newIndex.ArgType(testCase.Input...)
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
