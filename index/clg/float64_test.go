package clg

import (
	"math"
	"reflect"
	"testing"
)

func Test_Float64_DivideFloat64(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{3.5, 12.5},
			Expected:     []interface{}{0.28},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{35.5, 14.5},
			Expected:     []interface{}{2.4482758620689653},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{-3.5, 7.5},
			Expected:     []interface{}{-0.4666666666666667},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{12.5, 4.5},
			Expected:     []interface{}{2.7777777777777777},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36.5, 6.5},
			Expected:     []interface{}{5.615384615384615},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"a", 7.5},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{2.5, "7"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3.5},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{3.5, 4.5, 5.5},
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
		output, err := newCLGIndex.DivideFloat64(testCase.Input...)
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

func Test_Float64_GreaterThanFloat64(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{3.5, 3.5},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.5, 12.5},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{35.5, 14.5},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{-3.5, 7.5},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{12.5, 4.5},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36.5, 6.5},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"a", 7.5},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{2.5, "7"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3.5},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{3.5, 4.5, 5.5},
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
		output, err := newCLGIndex.GreaterThanFloat64(testCase.Input...)
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

func Test_Float64_LesserThanFloat64(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{3.5, 3.5},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.5, 12.5},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{35.5, 14.5},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{-3.5, 7.5},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{12.5, 4.5},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36.5, 6.5},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"a", 7.5},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{2.5, "7"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3.5},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{3.5, 4.5, 5.5},
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
		output, err := newCLGIndex.LesserThanFloat64(testCase.Input...)
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

func Test_Float64_MultiplyFloat64(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{3.5, 12.5},
			Expected:     []interface{}{43.75},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{35.5, 14.5},
			Expected:     []interface{}{514.75},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{-3.5, 7.5},
			Expected:     []interface{}{-26.25},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{12.5, 4.5},
			Expected:     []interface{}{56.25},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36.5, 6.5},
			Expected:     []interface{}{237.25},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"a", 7.5},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{2.5, "7"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3.5},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{3.5, 4.5, 5.5},
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
		output, err := newCLGIndex.MultiplyFloat64(testCase.Input...)
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

func Test_Float64_PowFloat64(t *testing.T) {
	testCases := []struct {
		Input           []interface{}
		Expected        []interface{}
		ExpectedMatcher func(f float64) bool
		ErrorMatcher    func(err error) bool
	}{
		{
			Input:        []interface{}{3.5, 12.5},
			Expected:     []interface{}{6.32194268775406e+06},
			ErrorMatcher: nil,
		},
		{
			Input:           []interface{}{-3.5, 7.5},
			ExpectedMatcher: math.IsNaN,
			ErrorMatcher:    nil,
		},
		{
			Input:        []interface{}{12.5, 4.5},
			Expected:     []interface{}{86316.74575031098},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36.5, 3.5},
			Expected:     []interface{}{293781.893469365},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"a", 7.5},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{2.5, "7"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3.5},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{3.5, 4.5, 5.5},
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
		output, err := newCLGIndex.PowFloat64(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if len(output) != 1 {
				t.Fatal("expected", 1, "got", len(output))
			}
			f, err := ArgToFloat64(output, 0)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			if testCase.ExpectedMatcher != nil && !testCase.ExpectedMatcher(f) {
				t.Fatal("case", i+1, "expected", true, "got", false)
			}
			if testCase.ExpectedMatcher == nil {
				if !reflect.DeepEqual(output, testCase.Expected) {
					t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
				}
			}
		}
	}
}

func Test_Float64_RoundFloat64(t *testing.T) {
	testCases := []struct {
		Input           []interface{}
		Expected        []interface{}
		ExpectedMatcher func(f float64) bool
		ErrorMatcher    func(err error) bool
	}{
		{
			Input:        []interface{}{3.5, 0},
			Expected:     []interface{}{float64(4)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.5, 1},
			Expected:     []interface{}{3.5},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.5, 2},
			Expected:     []interface{}{3.5},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.4, 0},
			Expected:     []interface{}{float64(3)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.4, 1},
			Expected:     []interface{}{3.4},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.4, 2},
			Expected:     []interface{}{3.4},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.87263, 2},
			Expected:     []interface{}{3.87},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.87263, 3},
			Expected:     []interface{}{3.873},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.87263, 4},
			Expected:     []interface{}{3.8726},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3.87263, -4},
			Expected:     nil,
			ErrorMatcher: IsCannotParse,
		},
		{
			Input:        []interface{}{3.87263, -2},
			Expected:     nil,
			ErrorMatcher: IsCannotParse,
		},
		{
			Input:        []interface{}{3.87263, 4, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{3.87263},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, 4},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3, 4},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}, 4},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3.87263, "a"},
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
		output, err := newCLGIndex.RoundFloat64(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if len(output) != 1 {
				t.Fatal("case", i+1, "expected", 1, "got", len(output))
			}
			f, err := ArgToFloat64(output, 0)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
			if testCase.ExpectedMatcher != nil && !testCase.ExpectedMatcher(f) {
				t.Fatal("case", i+1, "expected", true, "got", false)
			}
			if testCase.ExpectedMatcher == nil {
				if !reflect.DeepEqual(output, testCase.Expected) {
					t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
				}
			}
		}
	}
}

func Test_Float64_SqrtFloat64(t *testing.T) {
	testCases := []struct {
		Input           []interface{}
		Expected        []interface{}
		ExpectedMatcher func(f float64) bool
		ErrorMatcher    func(err error) bool
	}{
		{
			Input:        []interface{}{9.5},
			Expected:     []interface{}{3.082207001484488},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36.5},
			Expected:     []interface{}{6.041522986797286},
			ErrorMatcher: nil,
		},
		{
			Input:           []interface{}{-81.5},
			ExpectedMatcher: math.IsNaN,
			ErrorMatcher:    nil,
		},
		{
			Input:        []interface{}{144.5},
			Expected:     []interface{}{12.020815280171307},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{12.5},
			Expected:     []interface{}{3.5355339059327378},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"a"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{true},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{3.5, 4.5},
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
		output, err := newCLGIndex.SqrtFloat64(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if len(output) != 1 {
				t.Fatal("expected", 1, "got", len(output))
			}
			f, err := ArgToFloat64(output, 0)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			if testCase.ExpectedMatcher != nil && !testCase.ExpectedMatcher(f) {
				t.Fatal("case", i+1, "expected", true, "got", false)
			}
			if testCase.ExpectedMatcher == nil {
				if !reflect.DeepEqual(output, testCase.Expected) {
					t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
				}
			}
		}
	}
}

func Test_Float64_SubtractFloat64(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{3.5, 12.5},
			Expected:     []interface{}{float64(-9)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{35.5, 14.5},
			Expected:     []interface{}{float64(21)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{-3.5, 7.5},
			Expected:     []interface{}{float64(-11)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{12.5, 4.5},
			Expected:     []interface{}{float64(8)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36.5, 6.5},
			Expected:     []interface{}{float64(30)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"a", 7.5},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{2.5, "7"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3.5},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{3.5, 4.5, 5.5},
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
		output, err := newCLGIndex.SubtractFloat64(testCase.Input...)
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

func Test_Float64_SumFloat64(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{3.5, 12.5},
			Expected:     []interface{}{float64(16)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{35.5, 14.5},
			Expected:     []interface{}{float64(50)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{-3.5, 7.5},
			Expected:     []interface{}{float64(4)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{12.5, 4.5},
			Expected:     []interface{}{float64(17)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{36.5, 6.5},
			Expected:     []interface{}{float64(43)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"a", 7.5},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{2.5, "7"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3.5},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{3.5, 4.5, 5.5},
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
		output, err := newCLGIndex.SumFloat64(testCase.Input...)
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
