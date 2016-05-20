package permutation

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func Test_Permutation_Factory_GetIndex(t *testing.T) {
	testValues := []interface{}{"a", "b"}

	testMaybeNewFactory := func(t *testing.T) spec.PermutationFactory {
		newConfig := DefaultFactoryConfig()
		newConfig.MaxGrowth = 3
		newConfig.Values = testValues

		newFactory, err := NewFactory(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		return newFactory
	}

	newFactory := testMaybeNewFactory(t)

	// Make sure the initial index is empty.
	newIndex := newFactory.GetIndex()
	if newIndex != "" {
		t.Fatal("expected", "", "got", newIndex)
	}

	// Make sure the initial index is obtained even after some permutations. Note
	// that we have 2 values to permutate. We are going to calculate permutations
	// of the base 2 number system. This is the binary number system. The 4th
	// permutation in the binary system is 10.
	err := newFactory.PermuteBy(4)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newIndex = newFactory.GetIndex()
	if newIndex != "10" {
		t.Fatal("expected", "10", "got", newIndex)
	}

	// The 12th permutation (current index already is 10) in the binary system is
	// 110.
	err = newFactory.PermuteBy(8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newIndex = newFactory.GetIndex()
	if newIndex != "110" {
		t.Fatal("expected", "110", "got", newIndex)
	}
}

func Test_Permutation_Factory_GetValues(t *testing.T) {
	testValues := []interface{}{"a", "b"}

	testMaybeNewFactory := func(t *testing.T) spec.PermutationFactory {
		newConfig := DefaultFactoryConfig()
		newConfig.MaxGrowth = 3
		newConfig.Values = testValues

		newFactory, err := NewFactory(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		return newFactory
	}

	newFactory := testMaybeNewFactory(t)

	// Make sure the initial values are still obtained on the fresh factory.
	newValues := newFactory.GetValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}

	// Make sure the initial values are still obtained even after some
	// permutations.
	err := newFactory.PermuteBy(4)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newValues = newFactory.GetValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}

	err = newFactory.PermuteBy(8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newValues = newFactory.GetValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}
}

// Test_Permutation_Factory_PermuteBy_AbsoluteDelta tests permutations by
// providing deltas always to a new factory. That we we need to provide
// absolute deltas.
func Test_Permutation_Factory_PermuteBy_AbsoluteDelta(t *testing.T) {
	testMaybeNewFactory := func(t *testing.T) spec.PermutationFactory {
		newConfig := DefaultFactoryConfig()
		newConfig.MaxGrowth = 3
		newConfig.Values = []interface{}{"a", "b"}

		newFactory, err := NewFactory(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		return newFactory
	}

	testCases := []struct {
		Input        int
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        0,
			Expected:     []interface{}{"a"},
			ErrorMatcher: nil,
		},
		{
			Input:        1,
			Expected:     []interface{}{"b"},
			ErrorMatcher: nil,
		},
		{
			Input:        2,
			Expected:     []interface{}{"a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        3,
			Expected:     []interface{}{"a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        4,
			Expected:     []interface{}{"b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        5,
			Expected:     []interface{}{"b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        6,
			Expected:     []interface{}{"a", "a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        7,
			Expected:     []interface{}{"a", "a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        8,
			Expected:     []interface{}{"a", "b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        9,
			Expected:     []interface{}{"a", "b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        10,
			Expected:     []interface{}{"b", "a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        11,
			Expected:     []interface{}{"b", "a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        12,
			Expected:     []interface{}{"b", "b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        13,
			Expected:     []interface{}{"b", "b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        14,
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

	for i, testCase := range testCases {
		// Note we use a new factory for all test cases.
		newFactory := testMaybeNewFactory(t)

		err := newFactory.PermuteBy(testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		output := newFactory.GetMembers()

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
	testMaybeNewFactory := func(t *testing.T) spec.PermutationFactory {
		newConfig := DefaultFactoryConfig()
		newConfig.MaxGrowth = 3
		newConfig.Values = []interface{}{"a", "b"}

		newFactory, err := NewFactory(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		return newFactory
	}

	testCases := []struct {
		Input        int
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        3,
			Expected:     []interface{}{"a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        4,
			Expected:     []interface{}{"a", "a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        5,
			Expected:     []interface{}{"b", "b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        1,
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

	for i, testCase := range testCases {
		err := newFactory.PermuteBy(testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		output := newFactory.GetMembers()

		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}
