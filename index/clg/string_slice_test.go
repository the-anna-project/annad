package clg

import (
	"reflect"
	"testing"
)

func Test_StringSlice_AppendStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{}, "d"},
			Expected:     []interface{}{[]string{"d"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"b", "g", "c"}, "d"},
			Expected:     []interface{}{[]string{"b", "g", "c", "d"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"g", "a", "c"}, "b"},
			Expected:     []interface{}{[]string{"g", "a", "c", "b"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"g", "a", "c"}, "b", "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"g", "a", "c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"g", "a", "c"}, 3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3, "c"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).AppendStringSlice(testCase.Input...)
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

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).ContainsStringSlice(testCase.Input...)
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

func Test_StringSlice_CountCharacterStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input: []interface{}{[]string{"f", "o", "o"}},
			Expected: []interface{}{map[string]int{
				"f": 1,
				"o": 2,
			}},
			ErrorMatcher: nil,
		},
		{
			Input: []interface{}{[]string{"f", "o", "o", " ", "b", "o", "o"}},
			Expected: []interface{}{map[string]int{
				"f": 1,
				"o": 4,
				" ": 1,
				"b": 1,
			}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"f", "o", "o", " ", "b", "o", "o"}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{true},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).CountCharacterStringSlice(testCase.Input...)
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

func Test_StringSlice_CountStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     []interface{}{3},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{}},
			Expected:     []interface{}{0},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"", "", "", "foo", "", "bar"}},
			Expected:     []interface{}{6},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).CountStringSlice(testCase.Input...)
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

func Test_StringSlice_DifferenceStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, []string{"d", "y"}},
			Expected:     []interface{}{[]string{"a", "b", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "d", "c"}, []string{"d", "y"}},
			Expected:     []interface{}{[]string{"a", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"y", "b", "c"}, []string{"c", "y"}},
			Expected:     []interface{}{[]string{"b"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"g", "b", "c"}, []string{"c", "y"}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"g"}, []string{"c", "y"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"g", "b", "c"}, []string{"c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"g", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []string{"c", "y"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"g", "b", "c"}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).DifferenceStringSlice(testCase.Input...)
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

func Test_StringSlice_EqualMatcherStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, "a"},
			Expected:     []interface{}{[]string{"a"}, []string{"b", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, "ab"},
			Expected:     []interface{}{[]string(nil), []string{"a", "b", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"c", "c", "c"}, "c"},
			Expected:     []interface{}{[]string{"c", "c", "c"}, []string(nil)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"c", "c", "c"}, "c", "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"c", "c", "c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"c", "c", "c"}, 3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3, "c"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).EqualMatcherStringSlice(testCase.Input...)
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

func Test_StringSlice_GlobMatcherStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, "a"},
			Expected:     []interface{}{[]string{"a"}, []string{"b", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, "ab"},
			Expected:     []interface{}{[]string{"a", "b"}, []string{"c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"c", "c", "c"}, "c"},
			Expected:     []interface{}{[]string{"c", "c", "c"}, []string(nil)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"c", "c", "c"}, "c", "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"c", "c", "c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"c", "c", "c"}, 3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3, "c"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).GlobMatcherStringSlice(testCase.Input...)
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

func Test_StringSlice_IndexStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, 0},
			Expected:     []interface{}{"a"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, 1},
			Expected:     []interface{}{"b"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, 2},
			Expected:     []interface{}{"c"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3, 2},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, "a"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, 3},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, 3, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).IndexStringSlice(testCase.Input...)
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

func Test_StringSlice_IntersectionStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, []string{"d", "y"}},
			Expected:     []interface{}{[]string(nil)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "d", "c"}, []string{"d", "y"}},
			Expected:     []interface{}{[]string{"d"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"y", "b", "c"}, []string{"c", "y"}},
			Expected:     []interface{}{[]string{"y", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}, []string{"c", "y"}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"h"}, []string{"c", "y"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}, []string{"c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []string{"c", "y"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).IntersectionStringSlice(testCase.Input...)
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

func Test_StringSlice_IsUniqueStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"c", "a", "b", "c"}},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"c", "a", "b", "a", "c"}},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"c", "c", "c", "c", "c"}},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"c", "b", "a"}, true},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).IsUniqueStringSlice(testCase.Input...)
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

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).JoinStringSlice(testCase.Input...)
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

func Test_StringSlice_NewStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{},
			Expected:     []interface{}{[]string(nil)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).NewStringSlice(testCase.Input...)
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

func Test_StringSlice_ReverseStringSlice(t *testing.T) {
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
			Expected:     []interface{}{[]string{"c", "b", "a"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"c", "a", "b"}},
			Expected:     []interface{}{[]string{"b", "a", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"4", "13", "01", "b", "c", "a"}},
			Expected:     []interface{}{[]string{"a", "c", "b", "01", "13", "4"}},
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

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).ReverseStringSlice(testCase.Input...)
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

func Test_stem(t *testing.T) {
	i := stem(nil)
	if i != "" {
		t.Fatal("expected", 0, "got", i)
	}
	i = stem([]string{})
	if i != "" {
		t.Fatal("expected", 0, "got", i)
	}
}

func Test_StringSlice_StemStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b"}},
			Expected:     []interface{}{""},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "ab"}},
			Expected:     []interface{}{"a"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "bab"}},
			Expected:     []interface{}{""},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"ab", "ac", "cd", "ce"}},
			Expected:     []interface{}{""},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"cab", "cac", "cd", "ce"}},
			Expected:     []interface{}{"c"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"abcdefg", "abcde", "abcd", "abcd"}},
			Expected:     []interface{}{"abcd"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"abcdefg", "abcde", "abcd", "abcd"}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"abcdefg"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{"foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]int{3, 5, 7}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).StemStringSlice(testCase.Input...)
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

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).SortStringSlice(testCase.Input...)
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

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).SwapLeftStringSlice(testCase.Input...)
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

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).SwapRightStringSlice(testCase.Input...)
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

func Test_StringSlice_SymmetricDifferenceStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, []string{"d", "y"}},
			Expected:     []interface{}{[]string{"a", "b", "c", "d", "y"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "h", "c"}, []string{"h", "c", "x"}},
			Expected:     []interface{}{[]string{"a", "x"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}, []string{"c", "y"}},
			Expected:     []interface{}{[]string{"h", "b", "y"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}, []string{"c", "y"}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"h"}, []string{"c", "y"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}, []string{"c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []string{"c", "y"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).SymmetricDifferenceStringSlice(testCase.Input...)
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

func Test_StringSlice_UnionStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "h", "c"}, []string{"d", "y"}},
			Expected:     []interface{}{[]string{"a", "h", "c", "d", "y"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}, []string{"c", "y"}},
			Expected:     []interface{}{[]string{"h", "b", "c", "c", "y"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}, []string{"c", "y"}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"h"}, []string{"c", "y"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}, []string{"c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []string{"c", "y"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"h", "b", "c"}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).UnionStringSlice(testCase.Input...)
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

func Test_StringSlice_UniqueStringSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     []interface{}{[]string{"a", "b", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "a", "b", "c", "a", "c"}},
			Expected:     []interface{}{[]string{"a", "b", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]string{"a"}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).UniqueStringSlice(testCase.Input...)
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
