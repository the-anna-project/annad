package feature

import (
	"reflect"
	"testing"
)

func Test_NewFeature_Error_Positions(t *testing.T) {
	newConfig := DefaultFeatureConfig()
	// Note positions configuration is missing.
	newConfig.Sequence = "test"
	_, err := NewFeature(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_NewFeature_Error_Sequence(t *testing.T) {
	newConfig := DefaultFeatureConfig()
	newConfig.Positions = [][]float64{{0, 1}, {3, 4}}
	// Note sequence configuration is missing.
	_, err := NewFeature(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_NewFeature_Error_Distribution(t *testing.T) {
	newConfig := DefaultFeatureConfig()
	// Note positions configuration is invalid.
	newConfig.Positions = [][]float64{{0}, {3, 4, 5, 6}}
	newConfig.Sequence = "test"
	_, err := NewFeature(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Feature_AddPosition(t *testing.T) {
	newConfig := DefaultFeatureConfig()
	newConfig.Positions = [][]float64{{0, 1}}
	newConfig.Sequence = "test"
	newFeature, err := NewFeature(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	output := newFeature.GetPositions()
	if !reflect.DeepEqual(output, [][]float64{{0, 1}}) {
		t.Fatal("expected", [][]float64{{0, 1}}, "got", output)
	}

	err = newFeature.AddPosition([]float64{3, 4})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	output = newFeature.GetPositions()
	if !reflect.DeepEqual(output, [][]float64{{0, 1}, {3, 4}}) {
		t.Fatal("expected", [][]float64{{0, 1}, {3, 4}}, "got", output)
	}

	err = newFeature.AddPosition([]float64{3})
	if !IsInvalidPosition(err) {
		t.Fatal("expected", true, "got", false)
	}
	err = newFeature.AddPosition([]float64{4})
	if !IsInvalidPosition(err) {
		t.Fatal("expected", true, "got", false)
	}
	err = newFeature.AddPosition([]float64{3, 4, 5})
	if !IsInvalidPosition(err) {
		t.Fatal("expected", true, "got", false)
	}
}
