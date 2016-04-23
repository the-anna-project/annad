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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.AppendFloat64Slice(testCase.Input...)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.ContainsFloat64Slice(testCase.Input...)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.CountFloat64Slice(testCase.Input...)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.EqualMatcherFloat64Slice(testCase.Input...)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.GlobMatcherFloat64Slice(testCase.Input...)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.IndexFloat64Slice(testCase.Input...)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.IsUniqueFloat64Slice(testCase.Input...)
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
	i := maxFloat64(nil)
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
	}
	i = maxFloat64([]float64{})
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.MaxFloat64Slice(testCase.Input...)
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
	i := minFloat64(nil)
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
	}
	i = minFloat64([]float64{})
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.MinFloat64Slice(testCase.Input...)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.SortFloat64Slice(testCase.Input...)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.SwapLeftFloat64Slice(testCase.Input...)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.SwapRightFloat64Slice(testCase.Input...)
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

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.UniqueFloat64Slice(testCase.Input...)
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
