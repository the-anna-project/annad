package clg

import (
	"testing"
)

func Test_StringSlice_ContainsStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     bool
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, ""},
			Expected:     false,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, "test"},
			Expected:     true,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, " "},
			Expected:     true,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, "string"},
			Expected:     true,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, "foo"},
			Expected:     false,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, "  "},
			Expected:     false,
			ErrorMatcher: nil,
		},
		{
			// Note all args are missing.
			Input:        []interface{}{},
			Expected:     false,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			// Note substr is missing.
			Input:        []interface{}{[]string{"test", " ", "string"}},
			Expected:     false,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			// Note ss is wrong.
			Input:        []interface{}{"test string", " "},
			Expected:     false,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, 23},
			Expected:     false,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, 0.82},
			Expected:     false,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, []bool{false}},
			Expected:     false,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		args, err := ContainsStringSlice(testCase.Input...)
		if testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err) {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		if testCase.ErrorMatcher == nil {
			output, err := ArgToBool(args, 0)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
			if output != testCase.Expected {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}
