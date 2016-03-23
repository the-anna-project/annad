package patnet

import (
	"reflect"
	"testing"
)

func Test_NewFeatureSet_Error_Sequences(t *testing.T) {
	newConfig := DefaultFeatureSetConfig()
	// Note sequences configuration is missing.
	_, err := NewFeatureSet(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_FeatureSet_Separator(t *testing.T) {
	newConfig := DefaultFeatureSetConfig()
	newConfig.Separator = " "
	newConfig.Sequences = []string{
		"This is, a test.",
		"This is, another test.",
	}
	newFeatureSet, err := NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	err = newFeatureSet.Scan()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	fs := newFeatureSet.GetFeatures()
	for _, f := range fs {
		if f.GetSequence() != "another" {
			continue
		}

		if f.GetCount() != 1 {
			t.Fatal("expected", 1, "got", f.GetCount())
		}
		calculate := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0.14285714285714285, 0.14285714285714285, 0.14285714285714285, 0.14285714285714285, 0.14285714285714285, 0.14285714285714285, 0.14285714285714285, 0, 0, 0, 0, 0}
		if !reflect.DeepEqual(f.GetDistribution().Calculate(), calculate) {
			t.Fatal("expected", calculate, "got", f.GetDistribution().Calculate())
		}
		if f.GetSequence() != "another" {
			t.Fatal("expected", "another", "got", f.GetSequence())
		}
	}
}

func Test_FeatureSet_GetFeatures(t *testing.T) {
	newConfig := DefaultFeatureSetConfig()
	newConfig.MinCount = 2
	newConfig.Sequences = []string{
		"This is, a test.",
		"This is, another test.",
	}
	newFeatureSet, err := NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	err = newFeatureSet.Scan()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	fs := newFeatureSet.GetFeatures()
	for _, f := range fs {
		if f.GetSequence() != "." {
			continue
		}

		if f.GetCount() != 2 {
			t.Fatal("expected", 2, "got", f.GetCount())
		}
		calculate := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.5, 1.5}
		if !reflect.DeepEqual(f.GetDistribution().Calculate(), calculate) {
			t.Fatal("expected", calculate, "got", f.GetDistribution().Calculate())
		}
		if f.GetSequence() != "." {
			t.Fatal("expected", ".", "got", f.GetSequence())
		}
	}
}

func Test_FeatureSet_GetSequences(t *testing.T) {
	newSequences := []string{
		"This is, a test.",
		"This is, another test.",
	}

	newConfig := DefaultFeatureSetConfig()
	newConfig.MinCount = 2
	newConfig.Sequences = newSequences
	newFeatureSet, err := NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	output := newFeatureSet.GetSequences()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !reflect.DeepEqual(newSequences, output) {
		t.Fatal("expected", newSequences, "got", output)
	}
}

func Test_FeatureSet_GetFeaturesByCount(t *testing.T) {
	newConfig := DefaultFeatureSetConfig()
	newConfig.MinCount = 2
	newConfig.Sequences = []string{
		"This is, a test.",
		"This is, another test.",
	}
	newFeatureSet, err := NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	err = newFeatureSet.Scan()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	fs := newFeatureSet.GetFeaturesByCount(2)
	for _, f := range fs {
		if f.GetSequence() != "." {
			continue
		}

		if f.GetCount() != 2 {
			t.Fatal("expected", 2, "got", f.GetCount())
		}
		calculate := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.5, 1.5}
		if !reflect.DeepEqual(f.GetDistribution().Calculate(), calculate) {
			t.Fatal("expected", calculate, "got", f.GetDistribution().Calculate())
		}
		if f.GetSequence() != "." {
			t.Fatal("expected", ".", "got", f.GetSequence())
		}
	}
}

func Test_FeatureSet_GetFeaturesByLength(t *testing.T) {
	newConfig := DefaultFeatureSetConfig()
	newConfig.MinCount = 2
	newConfig.Sequences = []string{
		"This is, a test.",
		"This is, another test.",
	}
	newFeatureSet, err := NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	err = newFeatureSet.Scan()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	fs := newFeatureSet.GetFeaturesByLength(1)
	for _, f := range fs {
		if f.GetSequence() != "." {
			continue
		}

		if f.GetCount() != 2 {
			t.Fatal("expected", 2, "got", f.GetCount())
		}
		calculate := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.5, 1.5}
		if !reflect.DeepEqual(f.GetDistribution().Calculate(), calculate) {
			t.Fatal("expected", calculate, "got", f.GetDistribution().Calculate())
		}
		if f.GetSequence() != "." {
			t.Fatal("expected", ".", "got", f.GetSequence())
		}
	}
}

func Test_FeatureSet_GetFeaturesBySequence(t *testing.T) {
	newConfig := DefaultFeatureSetConfig()
	newConfig.MinCount = 2
	newConfig.Sequences = []string{
		"This is, a test.",
		"This is, another test.",
	}
	newFeatureSet, err := NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	err = newFeatureSet.Scan()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	fs := newFeatureSet.GetFeaturesBySequence(".")
	if len(fs) != 1 {
		t.Fatal("expected", 1, "got", len(fs))
	}
	f := fs[0]
	if f.GetCount() != 2 {
		t.Fatal("expected", 2, "got", f.GetCount())
	}
	calculate := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.5, 1.5}
	if !reflect.DeepEqual(f.GetDistribution().Calculate(), calculate) {
		t.Fatal("expected", calculate, "got", f.GetDistribution().Calculate())
	}
	if f.GetSequence() != "." {
		t.Fatal("expected", ".", "got", f.GetSequence())
	}
}
