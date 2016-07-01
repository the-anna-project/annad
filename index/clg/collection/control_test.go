package collection

import (
	"reflect"
	"testing"
)

func Test_Control_ForControl(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[][]interface{}{{"a"}, {"b"}, {"c"}}, "ToUpperString"},
			Expected:     []interface{}{[]interface{}{"A", "B", "C"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[][]interface{}{{1, 2}, {3, 4}, {5, 6}}, "not found"},
			Expected:     nil,
			ErrorMatcher: IsMethodNotFound,
		},
		{
			Input:        []interface{}{[][]interface{}{{"a"}, {"b"}, {"c"}}, "ToUpperString", "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[][]interface{}{{"a"}, {"b"}, {"c"}}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[][]interface{}{{"a"}}, "ToUpperString"},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, "ToUpperString"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).ForControl(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}
