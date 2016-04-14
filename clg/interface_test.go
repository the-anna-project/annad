package clg

import (
	"reflect"
	"testing"
)

func Test_Interface_DiscardInterface(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"ab"},
			Expected:     nil,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{23},
			Expected:     nil,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{true},
			Expected:     nil,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{23.8},
			Expected:     nil,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{map[int]bool{4: true}},
			Expected:     nil,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab", "foo"},
			Expected:     nil,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab", 23},
			Expected:     nil,
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab", false},
			Expected:     nil,
			ErrorMatcher: nil,
		},
		{
			// Note all args are missing.
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: nil,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.DiscardInterface(testCase.Input...)
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

func Test_Interface_EqualInterface(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"ab", "ab"},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab", "xy"},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36, 36},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36, 41},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{4.3, 4.3},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{4.3, 6.8},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{true, true},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{true, false},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3, false},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab", 4.3},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, []string{"c", "a", "b"}},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3, 4, 5}, []int{3, 4, 5}},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3, 4, 5}, []int{5, 3, 4}},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab"},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{"ab", "ab", "ab"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.EqualInterface(testCase.Input...)
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

func Test_Interface_InsertArgInterface(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]interface{}{}, "b", []int{0, 1}},
			Expected:     []interface{}{[]interface{}{"b", "b"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]interface{}{"a", "c"}, "b", []int{1}},
			Expected:     []interface{}{[]interface{}{"a", "b", "c"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]interface{}{"a", "c"}, "b", []int{1, 2, 4}},
			Expected:     []interface{}{[]interface{}{"a", "b", "b", "c", "b"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]interface{}{"a", "b", "c"}, "b", []int{1, 2, 5}},
			Expected:     []interface{}{[]interface{}{"a", "b", "b", "b", "c", "b"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]interface{}{"a", "c"}, "b", []int{0, 4}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]interface{}{"a", "c"}, "b", []int{0, 0}},
			Expected:     nil,
			ErrorMatcher: IsDuplicatedMember,
		},
		{
			Input:        []interface{}{[]interface{}{"a", "c"}, "b", []int{0, 1}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]interface{}{"a", "c"}, "b"},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{"a", "b", []int{0, 1}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]interface{}{"a", "c"}, "b", 3.4},
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
		output, err := newCLGIndex.InsertArgInterface(testCase.Input...)
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

func Test_Interface_ReturnInterface(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{},
			Expected:     []interface{}{},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab"},
			Expected:     []interface{}{"ab"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{true, 3, "foo"},
			Expected:     []interface{}{true, 3, "foo"},
			ErrorMatcher: nil,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.ReturnInterface(testCase.Input...)
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

func Test_Interface_SwapInterface(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"a", "b"},
			Expected:     []interface{}{"b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3, true},
			Expected:     []interface{}{true, 3},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{true, 3, "foo"},
			Expected:     []interface{}{true, 3, "foo"},
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{8.1},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
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
		output, err := newCLGIndex.SwapInterface(testCase.Input...)
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

func Test_Interface_TypeInterface(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"ab"},
			Expected:     []interface{}{"string"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{23},
			Expected:     []interface{}{"int"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{true},
			Expected:     []interface{}{"bool"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{23.8},
			Expected:     []interface{}{"float64"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{map[int]bool{4: true}},
			Expected:     []interface{}{"map[int]bool"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"ab", "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"ab", 23},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"ab", false},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			// Note all args are missing.
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
		output, err := newCLGIndex.TypeInterface(testCase.Input...)
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
