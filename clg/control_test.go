package clg

import (
	"reflect"
	"testing"
)

func Test_Control_ForStringControl(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, "ToUpperString"},
			Expected:     []interface{}{[]string{"A", "B", "C"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, "ToUpperString", "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"a"}, "ToUpperString"},
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.ForStringControl(testCase.Input...)
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

func Test_Control_IfControl(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{5, 3}, "SubtractInt", []interface{}{5, 3}},
			Expected:     []interface{}{2},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}},
			Expected:     []interface{}{},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt"},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{8.1, []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"GreaterThanInt", true, "SubtractInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, []int{}, []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", true},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"SplitString", []interface{}{"ab", ""}, "SubtractInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"ReturnInterface", []interface{}{true, "foo"}, "SubtractInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsTooManyResults,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.IfControl(testCase.Input...)
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

func Test_Control_IfElseControl(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{5, 3}, "SubtractInt", []interface{}{5, 3}, "SumInt", []interface{}{5, 3}},
			Expected:     []interface{}{2},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}, "SumInt", []interface{}{5, 3}},
			Expected:     []interface{}{8},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}, "SumInt", []interface{}{5, 3}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}, "SumInt"},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{8.1, []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}, "SumInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"GreaterThanInt", true, "SubtractInt", []interface{}{5, 3}, "SumInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, []int{}, []interface{}{5, 3}, "SumInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", true, "SumInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}, 8.1, []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}, "SumInt", false},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"SplitString", []interface{}{"ab", ""}, "SubtractInt", []interface{}{5, 3}, "SumInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"ReturnInterface", []interface{}{true, "foo"}, "SubtractInt", []interface{}{5, 3}, "SumInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsTooManyResults,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.IfElseControl(testCase.Input...)
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
