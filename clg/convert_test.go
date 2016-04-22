package clg

import (
	"reflect"
	"testing"
)

func Test_Convert_BoolStringConvert(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{true},
			Expected:     []interface{}{"true"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{false},
			Expected:     []interface{}{"false"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{-36},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{true, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.BoolStringConvert(testCase.Input...)
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

func Test_Convert_Float64StringConvert(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{3.3},
			Expected:     []interface{}{"3.3"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{9.433},
			Expected:     []interface{}{"9.433"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{-3.6},
			Expected:     []interface{}{"-3.6"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{-36},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3.6, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.Float64StringConvert(testCase.Input...)
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

func Test_Convert_StringBoolConvert(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"true"},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"false"},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"TRUE"},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"FALSE"},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"1"},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"0"},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"t"},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"f"},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"T"},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"F"},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"A"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"B"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"a"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"b"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"foo"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"bar"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"13"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"-82"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"1.3"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"-8.2"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"1,3"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"-8,2"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"true", "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{true},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{false},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{0},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.StringBoolConvert(testCase.Input...)
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

func Test_Convert_IntStringConvert(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{33},
			Expected:     []interface{}{"33"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{9433},
			Expected:     []interface{}{"9433"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{-36},
			Expected:     []interface{}{"-36"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.6},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{-3.6},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{36, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.IntStringConvert(testCase.Input...)
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

func Test_Convert_StringFloat64Convert(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"3.3"},
			Expected:     []interface{}{3.3},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"9.433"},
			Expected:     []interface{}{9.433},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"-3.6"},
			Expected:     []interface{}{-3.6},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"36"},
			Expected:     []interface{}{float64(36)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"-36"},
			Expected:     []interface{}{float64(-36)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"-x36"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"abc"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"3,6"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{"true"},
			Expected:     nil,
			ErrorMatcher: IsCannotConvert,
		},
		{
			Input:        []interface{}{36},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{-36},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{true},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"3.6", "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.StringFloat64Convert(testCase.Input...)
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
