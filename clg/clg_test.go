package clg

import (
	"reflect"
	"testing"
)

func Test_Index_Call(t *testing.T) {
	testCases := []struct {
		MethodName   string
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			MethodName:   "SortStringSlice",
			Input:        []interface{}{[]string{"c", "b", "d", "a"}},
			Expected:     []interface{}{[]string{"a", "b", "c", "d"}},
			ErrorMatcher: nil,
		},
		{
			MethodName:   "ArgType",
			Input:        []interface{}{3.8},
			Expected:     []interface{}{"float64"},
			ErrorMatcher: nil,
		},
		{
			MethodName:   "RepeatString",
			Input:        []interface{}{"abc", 3},
			Expected:     []interface{}{"abcabcabc"},
			ErrorMatcher: nil,
		},
	}

	newConfig := DefaultConfig()
	newIndex, err := NewIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newIndex.Call(testCase.MethodName, testCase.Input...)
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
