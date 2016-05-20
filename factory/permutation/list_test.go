package permutation

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func Test_Permutation_List_GetIndizes(t *testing.T) {
	testValues := []interface{}{"a", "b"}

	testMaybeNewList := func(t *testing.T) spec.PermutationList {
		newConfig := DefaultListConfig()
		newConfig.MaxGrowth = 3
		newConfig.Values = testValues

		newList, err := NewList(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		return newList
	}

	newFactory := testMaybeNewFactory(t)
	newList := testMaybeNewList(t)

	// Make sure the initial index is empty.
	newIndizes := newList.GetIndizes()
	if len(newIndizes) != 0 {
		t.Fatal("expected", 0, "got", newIndizes)
	}

	// Make sure the initial index is obtained even after some permutations. Note
	// that we have 2 values to permutate. We are going to calculate permutations
	// of the base 2 number system. This is the binary number system. The 4th
	// permutation in the binary system is 10.
	err := newFactory.PermuteBy(newList, 4)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newIndizes = newList.GetIndizes()
	if !reflect.DeepEqual(newIndizes, []int{1, 0}) {
		t.Fatal("expected", []int{1, 0}, "got", newIndizes)
	}

	// The 12th permutation (current index already is 10) in the binary system is
	// 110.
	err = newFactory.PermuteBy(newList, 8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newIndizes = newList.GetIndizes()
	if !reflect.DeepEqual(newIndizes, []int{1, 1, 0}) {
		t.Fatal("expected", []int{1, 1, 0}, "got", newIndizes)
	}
}

func Test_Permutation_List_GetValues(t *testing.T) {
	testValues := []interface{}{"a", "b"}

	testMaybeNewList := func(t *testing.T) spec.PermutationList {
		newConfig := DefaultListConfig()
		newConfig.MaxGrowth = 3
		newConfig.Values = testValues

		newList, err := NewList(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		return newList
	}

	newFactory := testMaybeNewFactory(t)
	newList := testMaybeNewList(t)

	// Make sure the initial values are still obtained on the fresh factory.
	newValues := newList.GetValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}

	// Make sure the initial values are still obtained even after some
	// permutations.
	err := newFactory.PermuteBy(newList, 4)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newValues = newList.GetValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}

	err = newFactory.PermuteBy(newList, 8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newValues = newList.GetValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}
}
