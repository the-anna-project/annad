package permutation

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/object/permutationlist"
	objectspec "github.com/xh3b4sd/anna/object/spec"
)

func testMaybeNewList(t *testing.T, values []interface{}) objectspec.PermutationList {
	newConfig := permutationlist.DefaultConfig()
	newConfig.MaxGrowth = 3
	newConfig.RawValues = values

	newList, err := permutationlist.New(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newList
}

// Test_Permutation_Service_PermuteBy_AbsoluteDelta tests permutations by
// providing deltas always to a new Service. That we we need to provide
// absolute deltas.
func Test_Permutation_Service_PermuteBy_AbsoluteDelta(t *testing.T) {
	testMaybeNewList := func(t *testing.T) objectspec.PermutationList {
		newConfig := permutationlist.DefaultConfig()
		newConfig.MaxGrowth = 3
		newConfig.RawValues = []interface{}{"a", "b"}

		newList, err := permutationlist.New(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		return newList
	}

	testCases := []struct {
		Input        int
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        0,
			Expected:     []interface{}{},
			ErrorMatcher: nil,
		},
		{
			Input:        1,
			Expected:     []interface{}{"a"},
			ErrorMatcher: nil,
		},
		{
			Input:        2,
			Expected:     []interface{}{"b"},
			ErrorMatcher: nil,
		},
		{
			Input:        3,
			Expected:     []interface{}{"a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        4,
			Expected:     []interface{}{"a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        5,
			Expected:     []interface{}{"b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        6,
			Expected:     []interface{}{"b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        7,
			Expected:     []interface{}{"a", "a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        8,
			Expected:     []interface{}{"a", "a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        9,
			Expected:     []interface{}{"a", "b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        10,
			Expected:     []interface{}{"a", "b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        11,
			Expected:     []interface{}{"b", "a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        12,
			Expected:     []interface{}{"b", "a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        13,
			Expected:     []interface{}{"b", "b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        14,
			Expected:     []interface{}{"b", "b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        15,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
		{
			Input:        23,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
		{
			Input:        583,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
	}

	newService := MustNew()

	for i, testCase := range testCases {
		// Note we use list service for all test cases.
		newList := testMaybeNewList(t)

		err := newService.PermuteBy(newList, testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		output := newList.GetPermutedValues()

		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

// Test_Permutation_Service_PermuteBy_Increment tests if increments by 1 always
// work.
func Test_Permutation_Service_PermuteBy_Increment(t *testing.T) {
	testMaybeNewList := func(t *testing.T) objectspec.PermutationList {
		newConfig := permutationlist.DefaultConfig()
		newConfig.MaxGrowth = 3
		newConfig.RawValues = []interface{}{"a", "b"}

		newList, err := permutationlist.New(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		return newList
	}

	testCases := []struct {
		Input        int
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        1,
			Expected:     []interface{}{"a"},
			ErrorMatcher: nil,
		},
		{
			Input:        1,
			Expected:     []interface{}{"b"},
			ErrorMatcher: nil,
		},
		{
			Input:        1,
			Expected:     []interface{}{"a", "a"},
			ErrorMatcher: nil,
		},
	}

	// Note we use the same service for all test cases.
	newService := MustNew()
	newList := testMaybeNewList(t)

	for i, testCase := range testCases {
		err := newService.PermuteBy(newList, testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		output := newList.GetPermutedValues()

		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

// Test_Permutation_Service_PermuteBy_RelativeDelta tests permutations by
// providing deltas always to an already existing Service. That we we need to
// provide relative deltas.
func Test_Permutation_Service_PermuteBy_RelativeDelta(t *testing.T) {
	testMaybeNewList := func(t *testing.T) objectspec.PermutationList {
		newConfig := permutationlist.DefaultConfig()
		newConfig.MaxGrowth = 3
		newConfig.RawValues = []interface{}{"a", "b"}

		newList, err := permutationlist.New(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		return newList
	}

	testCases := []struct {
		Input        int
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        3,
			Expected:     []interface{}{"a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        4,
			Expected:     []interface{}{"a", "a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        5,
			Expected:     []interface{}{"b", "a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        2,
			Expected:     []interface{}{"b", "b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        1,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
		{
			Input:        32,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
		{
			Input:        772,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
	}

	// Note we use the same service for all test cases.
	newService := MustNew()
	newList := testMaybeNewList(t)

	for i, testCase := range testCases {
		err := newService.PermuteBy(newList, testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		output := newList.GetPermutedValues()

		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

func Test_Permutation_List_GetIndizes(t *testing.T) {
	testValues := []interface{}{"a", "b"}

	newService := MustNew()
	newList := testMaybeNewList(t, testValues)

	// Make sure the initial index is empty.
	newIndizes := newList.GetIndizes()
	if len(newIndizes) != 0 {
		t.Fatal("expected", 0, "got", newIndizes)
	}

	// Make sure the initial index is obtained even after some permutations. Note
	// that we have 2 values to permutate. We are going to calculate permutations
	// of the base 2 number system. This is the binary number system. The 4th
	// permutation in the binary system is 10.
	err := newService.PermuteBy(newList, 4)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newIndizes = newList.GetIndizes()
	if !reflect.DeepEqual(newIndizes, []int{0, 1}) {
		t.Fatal("expected", []int{1, 1}, "got", newIndizes)
	}

	// The 12th permutation (current index already is 10) in the binary system is
	// 110.
	err = newService.PermuteBy(newList, 8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newIndizes = newList.GetIndizes()
	if !reflect.DeepEqual(newIndizes, []int{1, 0, 1}) {
		t.Fatal("expected", []int{1, 0, 1}, "got", newIndizes)
	}
}

func Test_Permutation_List_GetRawValues(t *testing.T) {
	testValues := []interface{}{"a", "b"}

	newService := MustNew()
	newList := testMaybeNewList(t, testValues)

	// Make sure the initial values are still obtained on the fresh Service.
	newValues := newList.GetRawValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}

	// Make sure the initial values are still obtained even after some
	// permutations.
	err := newService.PermuteBy(newList, 4)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newValues = newList.GetRawValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}

	err = newService.PermuteBy(newList, 8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newValues = newList.GetRawValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}
}
