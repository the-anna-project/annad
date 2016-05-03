package clg

import (
	"reflect"
	"testing"
)

func Test_IntSlice_AppendIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{}, 4},
			Expected:     []interface{}{[]int{4}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, 4},
			Expected:     []interface{}{[]int{1, 2, 3, 4}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 1},
			Expected:     []interface{}{[]int{2, 0, 3, 1}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 1, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, "foo"},
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
		output, err := testMaybeNewCLGCollection(t).AppendIntSlice(testCase.Input...)
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

func Test_IntSlice_ContainsIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{2, 0, 3}, 1},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 3},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 0},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 2},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 22},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 0000}, // 0000 translates to 0
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 2222},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 0000, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 0000, 23},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 0000, false},
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
			Input:        []interface{}{[]int{2, 0, 3}},
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
			Input:        []interface{}{[]int{2, 0, 3}, "ab"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, 0.82},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]int{2, 0, 3}, []bool{false}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).ContainsIntSlice(testCase.Input...)
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

func Test_IntSlice_CountIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2, 3}},
			Expected:     []interface{}{3},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{}},
			Expected:     []interface{}{0},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{0, 0, 0, 3, 0, 4}},
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
			Input:        []interface{}{[]int{}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).CountIntSlice(testCase.Input...)
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

func Test_IntSlice_DifferenceIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2, 3}, []int{4, 7}},
			Expected:     []interface{}{[]int{1, 2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 4, 3}, []int{4, 7}},
			Expected:     []interface{}{[]int{1, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{7, 2, 3}, []int{3, 7}},
			Expected:     []interface{}{[]int{2}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{9, 2, 3}, []int{3, 7}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{9}, []int{3, 7}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{9, 2, 3}, []int{3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{9, 2, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []int{3, 7}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]int{9, 2, 3}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).DifferenceIntSlice(testCase.Input...)
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

func Test_IntSlice_EqualMatcherIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2, 3}, 1},
			Expected:     []interface{}{[]int{1}, []int{2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, 12},
			Expected:     []interface{}{[]int(nil), []int{1, 2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3, 3, 3}, 3},
			Expected:     []interface{}{[]int{3, 3, 3}, []int(nil)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3, 3, 3}, 3, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{3, 3, 3}},
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
		output, err := testMaybeNewCLGCollection(t).EqualMatcherIntSlice(testCase.Input...)
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

func Test_IntSlice_GlobMatcherIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2, 3}, 1},
			Expected:     []interface{}{[]int{1}, []int{2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, 12},
			Expected:     []interface{}{[]int{1, 2}, []int{3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3, 3, 3}, 3},
			Expected:     []interface{}{[]int{3, 3, 3}, []int(nil)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3, 3, 3}, 3, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{3, 3, 3}},
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
		output, err := testMaybeNewCLGCollection(t).GlobMatcherIntSlice(testCase.Input...)
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

func Test_IntSlice_IndexIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2, 3}, 0},
			Expected:     []interface{}{1},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, 1},
			Expected:     []interface{}{2},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, 2},
			Expected:     []interface{}{3},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{3, 2},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, "a"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, 3},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, 3, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).IndexIntSlice(testCase.Input...)
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

func Test_IntSlice_IntersectionIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2, 3}, []int{4, 7}},
			Expected:     []interface{}{[]int(nil)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 4, 3}, []int{4, 7}},
			Expected:     []interface{}{[]int{4}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{7, 2, 3}, []int{3, 7}},
			Expected:     []interface{}{[]int{7, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{9, 2, 3}, []int{3, 7}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{9}, []int{3, 7}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{9, 2, 3}, []int{3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{9, 2, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []int{3, 7}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]int{9, 2, 3}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).IntersectionIntSlice(testCase.Input...)
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

func Test_IntSlice_IsUniqueIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2, 3}},
			Expected:     []interface{}{true},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3, 1, 2, 3}},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3, 1, 2, 1, 3}},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3, 3, 3, 3, 3}},
			Expected:     []interface{}{false},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{3, 2, 1}, true},
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
		output, err := testMaybeNewCLGCollection(t).IsUniqueIntSlice(testCase.Input...)
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

func Test_IntSlice_JoinIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2}},
			Expected:     []interface{}{12},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}},
			Expected:     []interface{}{123},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{1}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{1, 2}, "foo"},
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
		output, err := testMaybeNewCLGCollection(t).JoinIntSlice(testCase.Input...)
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

func Test_maxInt(t *testing.T) {
	i := maxInt(nil)
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
	}
	i = maxInt([]int{})
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
	}
}

func Test_IntSlice_MaxIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{4, 26, 12}},
			Expected:     []interface{}{26},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{4, 26, 312}},
			Expected:     []interface{}{312},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{94, 26, 12}},
			Expected:     []interface{}{94},
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
			Input:        []interface{}{[]int{4, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{4, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{5}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{}},
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
		output, err := testMaybeNewCLGCollection(t).MaxIntSlice(testCase.Input...)
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

func Test_meanInt(t *testing.T) {
	i := meanInt(nil)
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
	}
	i = meanInt([]int{})
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
	}
}

func Test_IntSlice_MeanIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{4, 5, 6}},
			Expected:     []interface{}{float64(5)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{4, 5, 6, 7, 8}},
			Expected:     []interface{}{float64(6)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{94, 26, 12}},
			Expected:     []interface{}{float64(44)},
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
			Input:        []interface{}{[]int{4, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{4, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{5}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{}},
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
		output, err := testMaybeNewCLGCollection(t).MeanIntSlice(testCase.Input...)
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

func Test_medianInt_ZeroValue(t *testing.T) {
	i := medianInt(nil)
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
	}
	i = medianInt([]int{})
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
	}
}

// Test_medianInt_EnsureInputUnchanged verifies that the input arguments
// stay unchanged after calculating the median of it. This check is necessary
// due to the fact that the input needs to be sorted in order to being able of
// calculating the median of it. Thus the check verifies that internally a copy
// of the input is created before calculating the median.
func Test_medianInt_EnsureInputUnchanged(t *testing.T) {
	input := []int{3, 2, 9, 5}
	expected := input
	medianInt(input)
	if !reflect.DeepEqual(input, expected) {
		t.Fatal("expected", expected, "got", input)
	}

	input = []int{3, 4, 5, 6, 7, 8}
	expected = input
	medianInt(input)
	if !reflect.DeepEqual(input, expected) {
		t.Fatal("expected", expected, "got", input)
	}
}

func Test_IntSlice_MedianIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{4, 5, 6}},
			Expected:     []interface{}{float64(5)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{0, 0, 1, 1, 1, 3, 4, 4, 55, 67}},
			Expected:     []interface{}{float64(2)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{4, 5, 6, 7, 8}},
			Expected:     []interface{}{float64(6)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{4, 5, 6, 6, 7, 8}},
			Expected:     []interface{}{float64(6)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{4, 5, 5, 7, 7, 8}},
			Expected:     []interface{}{float64(6)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{94, 13, 26, 12}},
			Expected:     []interface{}{float64(19.5)},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{94, 25, 12}},
			Expected:     []interface{}{float64(25)},
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
			Input:        []interface{}{[]int{4, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{4, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{5}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{}},
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
		output, err := testMaybeNewCLGCollection(t).MedianIntSlice(testCase.Input...)
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

func Test_minInt(t *testing.T) {
	i := minInt(nil)
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
	}
	i = minInt([]int{})
	if i != 0 {
		t.Fatal("expected", 0, "got", i)
	}
}

func Test_IntSlice_MinIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{4, 26, 12}},
			Expected:     []interface{}{4},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{4, 26, 312}},
			Expected:     []interface{}{4},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{94, 26, 12}},
			Expected:     []interface{}{12},
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
			Input:        []interface{}{[]int{4, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{4, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{5}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{}},
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
		output, err := testMaybeNewCLGCollection(t).MinIntSlice(testCase.Input...)
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

func Test_modeInt(t *testing.T) {
	is := modeInt(nil)
	if is != nil {
		t.Fatal("expected", nil, "got", is)
	}
	is = modeInt([]int{})
	if is != nil {
		t.Fatal("expected", nil, "got", is)
	}
}

func Test_IntSlice_ModeIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{4, 5, 6}},
			Expected:     []interface{}{[]int{4, 5, 6}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{4, 5, 5, 6, 7, 8}},
			Expected:     []interface{}{[]int{5}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{94, 94, 94, 26, 12, 3, 6, 3, 81, 3}},
			Expected:     []interface{}{[]int{3, 94}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{94, 94, 26, 12, 3, 6, 3, 26, 6}},
			Expected:     []interface{}{[]int{3, 6, 26, 94}},
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
			Input:        []interface{}{[]int{4, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{4, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{5}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{}},
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
		output, err := testMaybeNewCLGCollection(t).ModeIntSlice(testCase.Input...)
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

func Test_IntSlice_NewIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{},
			Expected:     []interface{}{[]int(nil)},
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
		output, err := testMaybeNewCLGCollection(t).NewIntSlice(testCase.Input...)
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

func Test_percentilesInt_ZeroValue(t *testing.T) {
	testCases := []struct {
		List        []int
		Percentiles []float64
	}{
		{
			List:        nil,
			Percentiles: nil,
		},
		{
			List:        []int{},
			Percentiles: nil,
		},
		{
			List:        nil,
			Percentiles: []float64{},
		},
		{
			List:        []int{},
			Percentiles: []float64{},
		},
	}

	for i, testCase := range testCases {
		ps, err := percentilesInt(testCase.List, testCase.Percentiles)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		if ps != nil {
			t.Fatal("case", i+1, "expected", nil, "got", ps)
		}
	}
}

// Test_percentilesInt_EnsureInputUnchanged verifies that the input
// arguments stay unchanged after calculating the percentiles of it. This check
// is necessary due to the fact that the input needs to be sorted in order to
// being able of calculating the percentiles of it. Thus the check verifies
// that internally a copy of the input is created before calculating the
// percentiles.
func Test_percentilesInt_EnsureInputUnchanged(t *testing.T) {
	percentiles := []float64{50, 100}

	list := []int{3, 2, 9, 5}
	expected := list
	percentilesInt(list, percentiles)
	if !reflect.DeepEqual(list, expected) {
		t.Fatal("expected", expected, "got", list)
	}

	list = []int{3, 4, 5, 6, 7, 8}
	expected = list
	percentilesInt(list, percentiles)
	if !reflect.DeepEqual(list, expected) {
		t.Fatal("expected", expected, "got", list)
	}
}

func Test_IntSlice_PercentilesIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{50, 100}},
			Expected:     []interface{}{[]float64{5, 9}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3, 1, 2, 4, 8, 6, 9, 5, 7}, []float64{50, 100}},
			Expected:     []interface{}{[]float64{5, 9}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{20, 40, 60, 80, 100}},
			Expected:     []interface{}{[]float64{2.6, 4.2, 5.799999999999999, 7.4, 9}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{2, 1, 3, 9, 5, 8, 7, 6, 4}, []float64{20, 40, 60, 80, 100}},
			Expected:     []interface{}{[]float64{2.6, 4.2, 5.799999999999999, 7.4, 9}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{75, 90, 95, 99, 99.999}},
			Expected:     []interface{}{[]float64{7.5, 8.2, 9.099999999999998, 9.82, 9.99982}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{9, 2, 3, 4, 5, 6, 8, 7, 1}, []float64{75, 90, 95, 99, 99.999}},
			Expected:     []interface{}{[]float64{7.5, 8.2, 9.099999999999998, 9.82, 9.99982}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{-0.001, 101}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{500, 1001}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{50, 101}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{50, -100}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{-0.0001, -100}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{-1, 100}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{-0.001, 100}},
			Expected:     nil,
			ErrorMatcher: IsIndexOutOfRange,
		},
		{
			Input:        []interface{}{[]int{3, 9}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{3.2, []float64{50}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]int{3, 9}, []float64{50}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{3}, []float64{50}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{3}, []float64{}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{}, []float64{}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{}},
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
		output, err := testMaybeNewCLGCollection(t).PercentilesIntSlice(testCase.Input...)
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

func Test_IntSlice_ReverseIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2}},
			Expected:     []interface{}{[]int{2, 1}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}},
			Expected:     []interface{}{[]int{3, 2, 1}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{3, 1, 2}},
			Expected:     []interface{}{[]int{2, 1, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{4, 13, 1, 2, 3, 1}},
			Expected:     []interface{}{[]int{1, 3, 2, 1, 13, 4}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, "foo"},
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
		output, err := testMaybeNewCLGCollection(t).ReverseIntSlice(testCase.Input...)
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

func Test_IntSlice_SortIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{4, 26, 12}},
			Expected:     []interface{}{[]int{4, 12, 26}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{4, 26, 312}},
			Expected:     []interface{}{[]int{4, 26, 312}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{94, 26, 12}},
			Expected:     []interface{}{[]int{12, 26, 94}},
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
			Input:        []interface{}{[]int{4, 26, 12}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{4, 26, 12}, 3},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{}},
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
		output, err := testMaybeNewCLGCollection(t).SortIntSlice(testCase.Input...)
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

func Test_IntSlice_SwapLeftIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2}},
			Expected:     []interface{}{[]int{2, 1}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}},
			Expected:     []interface{}{[]int{2, 3, 1}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5}},
			Expected:     []interface{}{[]int{2, 3, 4, 5, 1}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, "foo"},
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
		output, err := testMaybeNewCLGCollection(t).SwapLeftIntSlice(testCase.Input...)
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

func Test_IntSlice_SwapRightIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2}},
			Expected:     []interface{}{[]int{2, 1}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}},
			Expected:     []interface{}{[]int{3, 1, 2}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3, 4, 5}},
			Expected:     []interface{}{[]int{5, 1, 2, 3, 4}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, "foo"},
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
		output, err := testMaybeNewCLGCollection(t).SwapRightIntSlice(testCase.Input...)
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

func Test_IntSlice_SymmetricDifferenceIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{11, 2, 33}, []int{4, 7}},
			Expected:     []interface{}{[]int{11, 2, 33, 4, 7}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{11, 5, 33}, []int{5, 33, 31}},
			Expected:     []interface{}{[]int{11, 31}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{5, 2, 33}, []int{33, 7}},
			Expected:     []interface{}{[]int{5, 2, 7}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{5, 2, 33}, []int{33, 7}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{5}, []int{33, 7}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{5, 2, 33}, []int{33}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{5, 2, 33}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []int{33, 7}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]int{5, 2, 33}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).SymmetricDifferenceIntSlice(testCase.Input...)
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

func Test_IntSlice_UnionIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2, 3}, []int{4, 7}},
			Expected:     []interface{}{[]int{1, 2, 3, 4, 7}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 8, 3}, []int{4, 7}},
			Expected:     []interface{}{[]int{1, 8, 3, 4, 7}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{8, 2, 3}, []int{3, 7}},
			Expected:     []interface{}{[]int{8, 2, 3, 3, 7}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{8, 2, 3}, []int{3, 7}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{8}, []int{3, 7}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{8, 2, 3}, []int{3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{[]int{8, 2, 3}},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{true, []int{3, 7}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]int{8, 2, 3}, 8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCLGCollection(t).UnionIntSlice(testCase.Input...)
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

func Test_IntSlice_UniqueIntSlice(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[]int{1, 2, 3}},
			Expected:     []interface{}{[]int{1, 2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 1, 2, 3, 1, 3}},
			Expected:     []interface{}{[]int{1, 2, 3}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[]int{1, 2, 3}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[]int{1}},
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
		output, err := testMaybeNewCLGCollection(t).UniqueIntSlice(testCase.Input...)
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
