package permutationlist

import "testing"

func Test_Permutation_NewList_Error_MaxGrowth(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.MaxGrowth = 0
	newConfig.RawValues = []interface{}{"a", "b"}

	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Permutation_NewList_Error_Values(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.MaxGrowth = 3
	newConfig.RawValues = []interface{}{"a"}

	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}
