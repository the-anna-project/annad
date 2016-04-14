package clg

import (
	"reflect"
	"testing"
)

func Test_Method_CallCLGByName(t *testing.T) {
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
		output, err := newCLGIndex.CallCLGByName(testCase.Input...)
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

func Test_Method_GetCLGNames(t *testing.T) {
	testCases := []struct {
		Input            []interface{}
		ExpectedSubSet   []string
		UnexpectedSubSet []string
		ErrorMatcher     func(err error) bool
	}{
		{
			Input:          []interface{}{},
			ExpectedSubSet: []string{"CallCLGByName", "GetCLGNames", "RepeatString"},
			ErrorMatcher:   nil,
		},
		{
			Input:            []interface{}{"CLG"},
			ExpectedSubSet:   []string{"CallCLGByName", "GetCLGNames"},
			UnexpectedSubSet: []string{"RepeatString"},
			ErrorMatcher:     nil,
		},
		{
			Input:          []interface{}{"CLG", "foo"},
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
		output, err := newCLGIndex.GetCLGNames(testCase.Input...)
		if testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err) {
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
