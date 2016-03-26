package clg

import (
	"testing"
)

func Test_String_ContainsString(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     bool
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"test string", ""},
			Expected:     true,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"test string", "test"},
			Expected:     true,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"test string", "string"},
			Expected:     true,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"test string", "st str"},
			Expected:     true,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"test string", "foo"},
			Expected:     false,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"test string", "  "},
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
			Input:        []interface{}{"test string"},
			Expected:     false,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{23, "test string"},
			Expected:     false,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"test string", 23},
			Expected:     false,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"test string", 0.82},
			Expected:     false,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"test string", []bool{false}},
			Expected:     false,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		args, err := ContainsString(testCase.Input...)
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

func Test_String_RepeatString(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     string
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"ab", 0},
			Expected:     "",
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab", 1},
			Expected:     "ab",
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab", 2},
			Expected:     "abab",
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab", 4},
			Expected:     "abababab",
			ErrorMatcher: nil,
		},
		{
			// Note all args are missing.
			Input:        []interface{}{},
			Expected:     "",
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			// Note substr is missing.
			Input:        []interface{}{"ab"},
			Expected:     "",
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{23, "ab"},
			Expected:     "",
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{false, 23},
			Expected:     "",
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"ab", 0.82},
			Expected:     "",
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"ab", []bool{false}},
			Expected:     "",
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		args, err := RepeatString(testCase.Input...)
		if testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err) {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		if testCase.ErrorMatcher == nil {
			output, err := ArgToString(args, 0)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
			if output != testCase.Expected {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}
