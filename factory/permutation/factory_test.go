package permutation

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewFactory(t *testing.T) spec.PermutationFactory {
	newConfig := DefaultFactoryConfig()
	newFactory, err := NewFactory(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newFactory
}

// Test_Permutation_Factory_PermuteBy_AbsoluteDelta tests permutations by
// providing deltas always to a new factory. That we we need to provide
// absolute deltas.
func Test_Permutation_Factory_PermuteBy_AbsoluteDelta(t *testing.T) {
	testMaybeNewList := func(t *testing.T) spec.PermutationList {
		newConfig := DefaultListConfig()
		newConfig.MaxGrowth = 3
		newConfig.Values = []interface{}{"a", "b"}

		newList, err := NewList(newConfig)
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

	newFactory := testMaybeNewFactory(t)

	for i, testCase := range testCases {
		// Note we use list factory for all test cases.
		newList := testMaybeNewList(t)

		err := newFactory.PermuteBy(newList, testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		err = newFactory.MapTo(newList)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		output := newList.GetMembers()

		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

// Test_Permutation_Factory_PermuteBy_Increment tests if increments by 1 always
// work.
func Test_Permutation_Factory_PermuteBy_Increment(t *testing.T) {
	testMaybeNewList := func(t *testing.T) spec.PermutationList {
		newConfig := DefaultListConfig()
		newConfig.MaxGrowth = 3
		newConfig.Values = []interface{}{"a", "b"}

		newList, err := NewList(newConfig)
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

	// Note we use the same factory for all test cases.
	newFactory := testMaybeNewFactory(t)
	newList := testMaybeNewList(t)

	for i, testCase := range testCases {
		err := newFactory.PermuteBy(newList, testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		err = newFactory.MapTo(newList)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		output := newList.GetMembers()

		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

// Test_Permutation_Factory_PermuteBy_RelativeDelta tests permutations by
// providing deltas always to an already existing factory. That we we need to
// provide relative deltas.
func Test_Permutation_Factory_PermuteBy_RelativeDelta(t *testing.T) {
	testMaybeNewList := func(t *testing.T) spec.PermutationList {
		newConfig := DefaultListConfig()
		newConfig.MaxGrowth = 3
		newConfig.Values = []interface{}{"a", "b"}

		newList, err := NewList(newConfig)
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

	// Note we use the same factory for all test cases.
	newFactory := testMaybeNewFactory(t)
	newList := testMaybeNewList(t)

	for i, testCase := range testCases {
		err := newFactory.PermuteBy(newList, testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		err = newFactory.MapTo(newList)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		output := newList.GetMembers()

		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}
