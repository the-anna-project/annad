package clg

import (
	"reflect"
	"testing"
)

func Test_Float64Slice_AppendFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{}, 4.221},
			Expected:     []interface{}{[]float64{4.221}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, 4.221},
			Expected:     []interface{}{[]float64{1.98, 2.2, 3, 4.221}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}, 1.98},
			Expected:     []interface{}{[]float64{2.2, 0.0034, 3, 1.98}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}, 1.98, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"foo", 3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).AppendFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_ContainsFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}, 1.98},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}, float64(3)},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}, 0.0034},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}, 2.2},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}, float64(22)},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0, 3}, float64(0000)}, // 0000 translates to 0
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0, 3}, float64(2222)},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0, 3}, 0.3, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0, 3}, 0.3, 23},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0, 3}, 0.3, false},
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
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}},
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
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}, "ab"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}, 82},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{2.2, 0.0034, 3}, []bool{false}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).ContainsFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_CountFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}},
			Expected:     []interface{}{3},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{}},
			Expected:     []interface{}{0},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{0, 0, 0, 3, 0, 4.221}},
			Expected:     []interface{}{6},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).CountFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_DifferenceFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1, 2, 3}, []float64{4, 7}},
			Expected:     []interface{}{[]float64{1, 2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.44, 2, 3.578}, []float64{4, 7.1}},
			Expected:     []interface{}{[]float64{1.44, 2, 3.578}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1, 4, 3}, []float64{4, 7}},
			Expected:     []interface{}{[]float64{1, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.47, 4.47, 3}, []float64{4.47, 7}},
			Expected:     []interface{}{[]float64{1.47, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{7, 2, 3}, []float64{3, 7}},
			Expected:     []interface{}{[]float64{2}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{7, 2.833, 3.9}, []float64{3.9, 7}},
			Expected:     []interface{}{[]float64{2.833}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{9, 2, 3.9}, []float64{3.9, 7}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{9.9}, []float64{3, 7}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{9, 2.2, 3.3}, []float64{3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{9.3, 2, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []float64{3, 7}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{9.3, 2.3, 3.3}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).DifferenceFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_EqualMatcherFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, 1.98},
			Expected:     []interface{}{[]float64{1.98}, []float64{2.2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, float64(12)},
			Expected:     []interface{}{[]float64(nil), []float64{1.98, 2.2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{3, 3, 3}, float64(3)},
			Expected:     []interface{}{[]float64{3, 3, 3}, []float64(nil)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{3, 3, 3}, float64(3), "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{3, 3, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"c", "c", "c"}, float64(3)},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3, 3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).EqualMatcherFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_GlobMatcherFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, 1.98},
			Expected:     []interface{}{[]float64{1.98}, []float64{2.2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 1.933, 3}, 1.9},
			Expected:     []interface{}{[]float64{1.98, 1.933}, []float64{3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 1.933, 1.9}, 1.9},
			Expected:     []interface{}{[]float64{1.98, 1.933, 1.9}, []float64(nil)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{3, 3, 3}, float64(3), "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{3, 3, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]string{"c", "c", "c"}, 3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3, 3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).GlobMatcherFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_IndexFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, 0},
			Expected:     []interface{}{1.98},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, 1},
			Expected:     []interface{}{2.2},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, 2},
			Expected:     []interface{}{float64(3)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3, 2.2},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, "a"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, 3},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, 3, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).IndexFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_IntersectionFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1, 2, 3}, []float64{4, 7}},
			Expected:     []interface{}{[]float64(nil)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.44, 2, 3.578}, []float64{4, 7.1}},
			Expected:     []interface{}{[]float64(nil)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1, 4, 3}, []float64{4, 7}},
			Expected:     []interface{}{[]float64{4}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.47, 4.47, 3}, []float64{4.47, 7}},
			Expected:     []interface{}{[]float64{4.47}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{7, 2, 3}, []float64{3, 7}},
			Expected:     []interface{}{[]float64{7, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{7, 2.833, 3.9}, []float64{3.9, 7}},
			Expected:     []interface{}{[]float64{7, 3.9}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{9, 2, 3.9}, []float64{3.9, 7}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{9.9}, []float64{3, 7}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{9, 2.2, 3.3}, []float64{3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{9.3, 2, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []float64{3, 7}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{9.3, 2.3, 3.3}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).IntersectionFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_IsUniqueFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{3, 1.98, 2.2, 3}},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{3, 1.98, 2.2, 1.98, 3}},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{3, 3, 3, 3, 3}},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{3, 2.2, 1.98}, true},
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
		output, err := testMaybeNewCLGCollection(t).IsUniqueFloat64Slice(testCase.Input...)
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

func Test_maxFloat64(t *testing.T) {
	f := maxFloat64(nil)
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
	f = maxFloat64([]float64{})
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
}

func Test_Float64Slice_MaxFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{4.221, 26, 12}},
			Expected:     []interface{}{float64(26)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4.221, 26, 31.2}},
			Expected:     []interface{}{31.2},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{91.4, 26, 12}},
			Expected:     []interface{}{91.4},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{4.221, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{4.221, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).MaxFloat64Slice(testCase.Input...)
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

func Test_meanFloat64(t *testing.T) {
	f := meanFloat64(nil)
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
	f = meanFloat64([]float64{})
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
}

func Test_Float64Slice_MeanFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{4, 5, 6}},
			Expected:     []interface{}{float64(5)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4.3, 5.55, 6.789}},
			Expected:     []interface{}{5.546333333333333},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4, 5, 6, 7, 8}},
			Expected:     []interface{}{float64(6)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{94, 26, 12}},
			Expected:     []interface{}{float64(44)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{94.4, 25.15, 12.15}},
			Expected:     []interface{}{float64(43.900000000000006)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{4, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{4, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{5}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).MeanFloat64Slice(testCase.Input...)
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

func Test_medianFloat64_ZeroValue(t *testing.T) {
	f := medianFloat64(nil)
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
	f = medianFloat64([]float64{})
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
}

// Test_medianFloat64_EnsureInputUnchanged verifies that the input arguments
// stay unchanged after calculating the median of it. This check is necessary
// due to the fact that the input needs to be sorted in order to being able of
// calculating the median of it. Thus the check verifies that internally a copy
// of the input is created before calculating the median.
func Test_medianFloat64_EnsureInputUnchanged(t *testing.T) {
	input := []float64{3, 2.1, 9, 5.812}
	expected := input
	medianFloat64(input)
	if !reflect.DeepEqual(input, expected) {
		t.Fatal("expected", expected, "got", input)
	}

	input = []float64{3, 4, 5, 6, 7, 8}
	expected = input
	medianFloat64(input)
	if !reflect.DeepEqual(input, expected) {
		t.Fatal("expected", expected, "got", input)
	}
}

func Test_Float64Slice_MedianFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{4, 5, 6}},
			Expected:     []interface{}{float64(5)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{0.4, 0.88, 1.1, 1.2086, 1.3, 3.9, 4.0, 4.3, 55.5, 67.89}},
			Expected:     []interface{}{2.6},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4, 5, 6, 7, 8}},
			Expected:     []interface{}{float64(6)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4.8, 5, 6, 6, 7.63, 8}},
			Expected:     []interface{}{float64(6)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4.8, 5, 5, 7, 7.63, 8}},
			Expected:     []interface{}{float64(6)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{94, 13, 26, 12}},
			Expected:     []interface{}{float64(19.5)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{94.4, 25.15, 12.15}},
			Expected:     []interface{}{float64(25.15)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{4, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{4, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{5}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).MedianFloat64Slice(testCase.Input...)
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

func Test_minFloat64(t *testing.T) {
	f := minFloat64(nil)
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
	f = minFloat64([]float64{})
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
}

func Test_Float64Slice_MinFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{4.221, 26, 12}},
			Expected:     []interface{}{4.221},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4.221, 26, 312}},
			Expected:     []interface{}{4.221},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{94, 26, 12}},
			Expected:     []interface{}{float64(12)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{4.221, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{4.221, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).MinFloat64Slice(testCase.Input...)
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

func Test_modeFloat64(t *testing.T) {
	fs := modeFloat64(nil)
	if fs != nil {
		t.Fatal("expected", nil, "got", fs)
	}
	fs = modeFloat64([]float64{})
	if fs != nil {
		t.Fatal("expected", nil, "got", fs)
	}
}

func Test_Float64Slice_ModeFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{4, 5, 6}},
			Expected:     []interface{}{[]float64{4, 5, 6}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4.3, 5.55, 6.789}},
			Expected:     []interface{}{[]float64{4.3, 5.55, 6.789}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4, 5, 5, 6, 7, 8}},
			Expected:     []interface{}{[]float64{5}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4, 5.32, 5.32, 6, 7, 8}},
			Expected:     []interface{}{[]float64{5.32}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{94, 94, 94, 26.23, 12, 3.1, 6, 3.1, 81.777, 3.1}},
			Expected:     []interface{}{[]float64{3.1, 94}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{94, 94, 26.23, 12, 3.1, 6, 3.1, 26.23, 6}},
			Expected:     []interface{}{[]float64{3.1, 6, 26.23, 94}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{4, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{4, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{5}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).ModeFloat64Slice(testCase.Input...)
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

func Test_percentilesFloat64_ZeroValue(t *testing.T) {
	testCases := []struct {
		List        []float64
		Percentiles []float64
	}{
		{
			List:        nil,
			Percentiles: nil,
		},
		{
			List:        []float64{},
			Percentiles: nil,
		},
		{
			List:        nil,
			Percentiles: []float64{},
		},
		{
			List:        []float64{},
			Percentiles: []float64{},
		},
	}

	for i, testCase := range testCases {
		ps, err := percentilesFloat64(testCase.List, testCase.Percentiles)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		if ps != nil {
			t.Fatal("case", i+1, "expected", nil, "got", ps)
		}
	}
}

// Test_percentilesFloat64_EnsureInputUnchanged verifies that the input
// arguments stay unchanged after calculating the percentiles of it. This check
// is necessary due to the fact that the input needs to be sorted in order to
// being able of calculating the percentiles of it. Thus the check verifies
// that internally a copy of the input is created before calculating the
// percentiles.
func Test_percentilesFloat64_EnsureInputUnchanged(t *testing.T) {
	percentiles := []float64{50, 100}

	list := []float64{3, 2.1, 9, 5.812}
	expected := list
	percentilesFloat64(list, percentiles)
	if !reflect.DeepEqual(list, expected) {
		t.Fatal("expected", expected, "got", list)
	}

	list = []float64{3, 4, 5, 6, 7, 8}
	expected = list
	percentilesFloat64(list, percentiles)
	if !reflect.DeepEqual(list, expected) {
		t.Fatal("expected", expected, "got", list)
	}
}

func Test_Float64Slice_PercentilesFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{50, 100}},
			Expected:     []interface{}{[]float64{5, 9}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.33, 2, 3.44, 4, 5, 6.82, 7, 8.8, 9}, []float64{50, 100}},
			Expected:     []interface{}{[]float64{5, 9}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{3, 1, 2, 4, 8, 6, 9, 5, 7}, []float64{50, 100}},
			Expected:     []interface{}{[]float64{5, 9}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{20, 40, 60, 80, 100}},
			Expected:     []interface{}{[]float64{2.6, 4.2, 5.799999999999999, 7.4, 9}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1, 2.3322, 3, 4.818, 5, 6.33, 7, 8, 9.001}, []float64{20, 40, 60, 80, 100}},
			Expected:     []interface{}{[]float64{2.86576, 4.6908, 5.931999999999999, 7.4, 9.001}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{2, 1, 3, 9, 5, 8, 7, 6, 4}, []float64{20, 40, 60, 80, 100}},
			Expected:     []interface{}{[]float64{2.6, 4.2, 5.799999999999999, 7.4, 9}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{75, 90, 95, 99, 99.999}},
			Expected:     []interface{}{[]float64{7.5, 8.2, 9.099999999999998, 9.82, 9.99982}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.001, 2, 3.3, 3.002, 3.35, 3.801, 4, 5, 6, 7, 8, 9}, []float64{75, 90, 95, 99, 99.999}},
			Expected:     []interface{}{[]float64{6, 11.600000000000001, 11.799999999999997, 12.759999999999998, 12.999759999999998}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{9, 2, 3, 4, 5, 6, 8, 7, 1}, []float64{75, 90, 95, 99, 99.999}},
			Expected:     []interface{}{[]float64{7.5, 8.2, 9.099999999999998, 9.82, 9.99982}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{-0.001, 101}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{500, 1001}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{50, 101}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{50, -100}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{-0.0001, -100}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{-1, 100}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{-0.001, 100}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]float64{3.2, 9.1}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3.2, []float64{50}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{3.2, 9.1}, []float64{50}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{3.2}, []float64{50}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{3.2}, []float64{}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{}, []float64{}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).PercentilesFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_NewFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{},
			Expected:     []interface{}{[]float64(nil)},
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
		output, err := testMaybeNewCLGCollection(t).NewFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_ReverseFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1, 2}},
			Expected:     []interface{}{[]float64{2, 1}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1, 2, 3}},
			Expected:     []interface{}{[]float64{3, 2, 1}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{3.2, 1, 2.8}},
			Expected:     []interface{}{[]float64{2.8, 1, 3.2}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4, 13.444, 1, 2.8, 3, 1}},
			Expected:     []interface{}{[]float64{1, 3, 2.8, 1, 13.444, 4}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{1, 2, 3}, "foo"},
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
		output, err := testMaybeNewCLGCollection(t).ReverseFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_SortFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{4.221, 26, 12}},
			Expected:     []interface{}{[]float64{4.221, 12, 26}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{4.221, 26, 312}},
			Expected:     []interface{}{[]float64{4.221, 26, 312}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{94, 26, 12}},
			Expected:     []interface{}{[]float64{12, 26, 94}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]string{"a", "b", "c"}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{4.221, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{4.221, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).SortFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_SwapLeftFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1.98, 2.2}},
			Expected:     []interface{}{[]float64{2.2, 1.98}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}},
			Expected:     []interface{}{[]float64{2.2, 3, 1.98}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3, 4.221, 5}},
			Expected:     []interface{}{[]float64{2.2, 3, 4.221, 5, 1.98}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, "foo"},
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
		output, err := testMaybeNewCLGCollection(t).SwapLeftFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_SwapRightFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1.98, 2.2}},
			Expected:     []interface{}{[]float64{2.2, 1.98}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}},
			Expected:     []interface{}{[]float64{3, 1.98, 2.2}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3, 4.221, 5}},
			Expected:     []interface{}{[]float64{5, 1.98, 2.2, 3, 4.221}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, "foo"},
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
		output, err := testMaybeNewCLGCollection(t).SwapRightFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_SymmetricDifferenceFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{11, 2, 33}, []float64{4, 7}},
			Expected:     []interface{}{[]float64{11, 2, 33, 4, 7}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.1, 2, 33.23}, []float64{4, 7.1}},
			Expected:     []interface{}{[]float64{1.1, 2, 33.23, 4, 7.1}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{11, 5, 33}, []float64{5, 33, 31}},
			Expected:     []interface{}{[]float64{11, 31}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{11.5586, 5.733, 33.733}, []float64{5.733, 33.733, 31}},
			Expected:     []interface{}{[]float64{11.5586, 31}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{5.32, 2.32, 33.32}, []float64{33.32, 7.32}},
			Expected:     []interface{}{[]float64{5.32, 2.32, 7.32}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{5, 2.111, 33}, []float64{33.3, 7.44}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{5.2}, []float64{33, 7}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{5, 2.821, 33.21}, []float64{33}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{5.43, 2.43, 33}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []float64{33.83, 7}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{5.22, 2.83, 3.3}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).SymmetricDifferenceFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_UnionFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1, 2, 3}, []float64{4, 7}},
			Expected:     []interface{}{[]float64{1, 2, 3, 4, 7}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1, 2.8, 3}, []float64{4.784, 7}},
			Expected:     []interface{}{[]float64{1, 2.8, 3, 4.784, 7}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, []float64{4.3, 7.9}},
			Expected:     []interface{}{[]float64{1.98, 2.2, 3, 4.3, 7.9}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, []float64{4.3, 7.9}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{1.98}, []float64{4.3, 7.9}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, []float64{4.3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []float64{4.3, 7.9}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).UnionFloat64Slice(testCase.Input...)
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

func Test_Float64Slice_UniqueFloat64Slice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}},
			Expected:     []interface{}{[]float64{1.98, 2.2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 1.98, 2.2, 3, 1.98, 3}},
			Expected:     []interface{}{[]float64{1.98, 2.2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]float64{1.98, 2.2, 3}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]float64{1.98}},
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
		output, err := testMaybeNewCLGCollection(t).UniqueFloat64Slice(testCase.Input...)
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
