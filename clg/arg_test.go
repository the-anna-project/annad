package clg

import (
	"reflect"
	"testing"
)

func Test_ArgToInt(t *testing.T) {
	testCases := []struct {
		Args         []interface{}
		Index        int
		Def          []int
		Expected     int
		ErrorMatcher func(err error) bool
	}{
		{
			Args:         []interface{}{1, 2},
			Index:        0,
			Def:          []int{},
			Expected:     1,
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{1, 2},
			Index:        1,
			Def:          []int{},
			Expected:     2,
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{1, 2},
			Index:        0,
			Def:          []int{3},
			Expected:     1,
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{1, 2},
			Index:        1,
			Def:          []int{3},
			Expected:     2,
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{1, 2},
			Index:        2,
			Def:          []int{3},
			Expected:     3,
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{1, 2},
			Index:        2,
			Def:          []int{3, 4},
			Expected:     0,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := ArgToInt(testCase.Args, testCase.Index, testCase.Def...)
		if testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if output != testCase.Expected {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

func Test_ArgToString(t *testing.T) {
	testCases := []struct {
		Args         []interface{}
		Index        int
		Def          []string
		Expected     string
		ErrorMatcher func(err error) bool
	}{
		{
			Args:         []interface{}{"a", "b"},
			Index:        0,
			Def:          []string{},
			Expected:     "a",
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{"a", "b"},
			Index:        1,
			Def:          []string{},
			Expected:     "b",
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{"a", "b"},
			Index:        0,
			Def:          []string{"c"},
			Expected:     "a",
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{"a", "b"},
			Index:        1,
			Def:          []string{"c"},
			Expected:     "b",
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{"a", "b"},
			Index:        2,
			Def:          []string{"c"},
			Expected:     "c",
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{"a", "b"},
			Index:        2,
			Def:          []string{"c", "d"},
			Expected:     "",
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := ArgToString(testCase.Args, testCase.Index, testCase.Def...)
		if testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if output != testCase.Expected {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

func Test_ArgToStringSlice(t *testing.T) {
	testCases := []struct {
		Args         []interface{}
		Index        int
		Def          [][]string
		Expected     []string
		ErrorMatcher func(err error) bool
	}{
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        0,
			Def:          [][]string{},
			Expected:     []string{"a"},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        1,
			Def:          [][]string{},
			Expected:     []string{"b"},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        0,
			Def:          [][]string{[]string{"c"}},
			Expected:     []string{"a"},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        1,
			Def:          [][]string{[]string{"c"}},
			Expected:     []string{"b"},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        2,
			Def:          [][]string{[]string{"c"}},
			Expected:     []string{"c"},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        2,
			Def:          [][]string{[]string{"c"}, []string{"d"}},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := ArgToStringSlice(testCase.Args, testCase.Index, testCase.Def...)
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
