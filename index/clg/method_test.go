package clg

import (
	"reflect"
	"testing"
)

func Test_Method_CallMethodByName(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"SortStringSlice", []string{"c", "b", "d", "a"}},
			Expected:     []interface{}{[]string{"a", "b", "c", "d"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"TypeInterface", 3.8},
			Expected:     []interface{}{"float64"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"RepeatString", "abc", 3},
			Expected:     []interface{}{"abcabcabc"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"RepeatString"},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{"not found"},
			Expected:     nil,
			ErrorMatcher: IsMethodNotFound,
		},
		{
			Input:        []interface{}{3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.CallMethodByName(testCase.Input...)
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

func Test_Method_GetMethodNames_Expected(t *testing.T) {
	testCases := []struct {
		Input          []interface{}
		ExpectedSubSet []string
		ErrorMatcher   func(err error) bool
	}{
		{
			Input:          []interface{}{},
			ExpectedSubSet: []string{"CallMethodByName", "GetMethodNames", "RepeatString"},
			ErrorMatcher:   nil,
		},
		{
			Input:          []interface{}{"Method"},
			ExpectedSubSet: []string{"CallMethodByName", "GetMethodNames"},
			ErrorMatcher:   nil,
		},
		{
			Input:          []interface{}{"Method", "foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsTooManyArguments,
		},
		{
			Input:          []interface{}{3.4},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsWrongArgumentType,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.GetMethodNames(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			ss, err := ArgToStringSlice(output, 0)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
			for j, e := range testCase.ExpectedSubSet {
				var contains bool
				for _, s := range ss {
					if s == e {
						contains = true
						break
					}
				}
				if !contains {
					t.Fatal("case", j+1, "of", i+1, "expected", e, "got", "")
				}
			}
		}
	}
}

func Test_Method_GetMethodNames_Unexpected(t *testing.T) {
	testCases := []struct {
		Input            []interface{}
		UnexpectedSubSet []string
		ErrorMatcher     func(err error) bool
	}{
		{
			Input:            []interface{}{"Method"},
			UnexpectedSubSet: []string{"RepeatString"},
			ErrorMatcher:     nil,
		},
		{
			Input:        []interface{}{"Method", "foo"},
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{3.4},
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.GetMethodNames(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			ss, err := ArgToStringSlice(output, 0)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
			for j, e := range testCase.UnexpectedSubSet {
				var contains bool
				for _, s := range ss {
					if s == e {
						contains = true
						break
					}
				}
				if contains {
					t.Fatal("case", j+1, "of", i+1, "expected", e, "got", "")
				}
			}
		}
	}
}

func Test_Method_GetNumMethods(t *testing.T) {
	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	_, err = newCLGIndex.GetNumMethods("foo")
	if !IsTooManyArguments(err) {
		t.Fatal("expected", true, "got", false)
	}

	output, err := newCLGIndex.GetNumMethods()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(output) > 1 {
		t.Fatal("expected", 1, "got", len(output))
	}
	num, err := ArgToInt(output, 0)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	// There shouldn't be a test that expects an exact amount of methods, thus
	// we simply expect that there are more than a given threshold.
	if num < 130 {
		t.Fatal("expected", nil, "got", num)
	}
}
