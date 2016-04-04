package clg

import (
	"reflect"
	"testing"
)

func Test_StringSlice_ContainsStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, ""},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, "test"},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, " "},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, "string"},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, "foo"},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, "  "},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, "test", "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, "test", 23},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, "test", false},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			// Note all args are missing.
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			// Note substr is missing.
			Input:        []interface{}{[]string{"test", " ", "string"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			// Note ss is wrong.
			Input:        []interface{}{"test string", " "},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, 23},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, 0.82},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"test", " ", "string"}, []bool{false}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	newConfig := DefaultConfig()
	newIndex, err := NewIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newIndex.ContainsStringSlice(testCase.Input...)
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

func Test_StringSlice_JoinStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b"}, ""},
			Expected:     []interface{}{"ab"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, " "},
			Expected:     []interface{}{"a b c"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c", "d", "e"}, "-"},
			Expected:     []interface{}{"a-b-c-d-e"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"a"}, ""},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, " ", "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{23},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]bool{}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	newConfig := DefaultConfig()
	newIndex, err := NewIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newIndex.JoinStringSlice(testCase.Input...)
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

func Test_StringSlice_SortStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b"}},
			Expected:     []interface{}{[]string{"a", "b"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     []interface{}{[]string{"a", "b", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"b", "a"}},
			Expected:     []interface{}{[]string{"a", "b"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"c", "a", "b"}},
			Expected:     []interface{}{[]string{"a", "b", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"4", "13", "01", "b", "c", "a"}},
			Expected:     []interface{}{[]string{"01", "13", "4", "a", "b", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{23},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]bool{}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	newConfig := DefaultConfig()
	newIndex, err := NewIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newIndex.SortStringSlice(testCase.Input...)
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

func Test_StringSlice_SwapLeftStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b"}},
			Expected:     []interface{}{[]string{"b", "a"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     []interface{}{[]string{"b", "c", "a"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c", "d", "e"}},
			Expected:     []interface{}{[]string{"b", "c", "d", "e", "a"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{23},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]bool{}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	newConfig := DefaultConfig()
	newIndex, err := NewIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newIndex.SwapLeftStringSlice(testCase.Input...)
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

func Test_StringSlice_SwapRightStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b"}},
			Expected:     []interface{}{[]string{"b", "a"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     []interface{}{[]string{"c", "a", "b"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c", "d", "e"}},
			Expected:     []interface{}{[]string{"e", "a", "b", "c", "d"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{23},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]bool{}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	newConfig := DefaultConfig()
	newIndex, err := NewIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newIndex.SwapRightStringSlice(testCase.Input...)
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
