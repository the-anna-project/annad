package permutation

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/service/spec"
)

func testMaybeNewList(t *testing.T, values []interface{}) spec.PermutationList {
	newConfig := DefaultListConfig()
	newConfig.MaxGrowth = 3
	newConfig.RawValues = values

	newList, err := NewList(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newList
}

func Test_Permutation_NewList_Error_MaxGrowth(t *testing.T) {
	newConfig := DefaultListConfig()
	newConfig.MaxGrowth = 0
	newConfig.RawValues = []interface{}{"a", "b"}

	_, err := NewList(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Permutation_NewList_Error_Values(t *testing.T) {
	newConfig := DefaultListConfig()
	newConfig.MaxGrowth = 3
	newConfig.RawValues = []interface{}{"a"}

	_, err := NewList(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Permutation_List_GetIndizes(t *testing.T) {
	testValues := []interface{}{"a", "b"}

	newService := MustNewService()
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

	newService := MustNewService()
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
