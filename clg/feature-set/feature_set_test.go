package featureset

import (
	"reflect"
	"sort"
	"testing"
)

func Test_NewFeatureSet_Error_Sequences(t *testing.T) {
	newConfig := DefaultConfig()
	// Note sequences configuration is missing.
	_, err := NewFeatureSet(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_FeatureSet_GetFeatures(t *testing.T) {
	newConfig := DefaultConfig()
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

func Test_FeatureSet_GetFeaturesByCount(t *testing.T) {
	newConfig := DefaultConfig()
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
	testCases := []struct {
		Sequences []string
		Min       int
		Max       int
		Expected  string
	}{
		{
			Sequences: []string{
				"This is, a test.",
				"This is, another test.",
			},
			Min:      1,
			Max:      1,
			Expected: ".",
		},
		{
			Sequences: []string{
				"This is, a test.",
				"This is, another test.",
			},
			Min:      7,
			Max:      7,
			Expected: "This is",
		},
		{
			Sequences: []string{
				"This is, a test.",
				"This is, another test.",
			},
			Min:      1,
			Max:      -1,
			Expected: ".",
		},
		{
			Sequences: []string{
				"This is, a test.",
				"This is, another test.",
			},
			Min:      7,
			Max:      -1,
			Expected: "This is",
		},
	}

	for i, testCase := range testCases {
		newConfig := DefaultConfig()
		newConfig.MinCount = 2
		newConfig.Sequences = testCase.Sequences
		newFeatureSet, err := NewFeatureSet(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		err = newFeatureSet.Scan()
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		var found bool
		fs := newFeatureSet.GetFeaturesByLength(testCase.Min, testCase.Max)
		for _, f := range fs {
			if f.GetSequence() == testCase.Expected {
				found = true
				break
			}
		}
		if !found {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
	}
}

func Test_FeatureSet_GetFeaturesBySequence_Expected(t *testing.T) {
	testCases := []struct {
		Input          string
		ExpectedSubSet []string
		ErrorMatcher   func(err error) bool
	}{
		{
			Input:          "ab",
			ExpectedSubSet: nil,
			ErrorMatcher:   nil,
		},
		{
			Input:          "This",
			ExpectedSubSet: []string{"This", "This is"},
			ErrorMatcher:   nil,
		},
		{
			Input:          "",
			ExpectedSubSet: []string{"This", "This is", "another", "test.", "."},
			ErrorMatcher:   nil,
		},
	}

	for i, testCase := range testCases {
		newConfig := DefaultConfig()
		newConfig.Sequences = []string{"This is a test.", "This is another test."}
		newFeatureSet, err := NewFeatureSet(newConfig)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			err = newFeatureSet.Scan()
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
			fs := newFeatureSet.GetFeaturesBySequence(testCase.Input)
			for j, e := range testCase.ExpectedSubSet {
				var contains bool
				for _, f := range fs {
					if f.GetSequence() == e {
						contains = true
						break
					}
				}
				if !contains {
					t.Fatal("case", j+1, "of", i+1, "expected", e, "got", "empty string")
				}
			}
		}
	}
}

func Test_FeatureSet_GetFeaturesBySequence_Unexpected(t *testing.T) {
	testCases := []struct {
		Input            string
		UnexpectedSubSet []string
		ErrorMatcher     func(err error) bool
	}{
		{
			Input:            "ab",
			UnexpectedSubSet: nil,
			ErrorMatcher:     nil,
		},
		{
			Input:            "This",
			UnexpectedSubSet: []string{"another", "test."},
			ErrorMatcher:     nil,
		},
		{
			Input:            "",
			UnexpectedSubSet: nil,
			ErrorMatcher:     nil,
		},
	}

	for i, testCase := range testCases {
		newConfig := DefaultConfig()
		newConfig.Sequences = []string{"This is a test.", "This is another test."}
		newFeatureSet, err := NewFeatureSet(newConfig)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			err = newFeatureSet.Scan()
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
			fs := newFeatureSet.GetFeaturesBySequence(testCase.Input)
			for j, e := range testCase.UnexpectedSubSet {
				var contains bool
				for _, f := range fs {
					if f.GetSequence() == e {
						contains = true
						break
					}
				}
				if contains {
					t.Fatal("case", j+1, "of", i+1, "expected", e, "got", "empty string")
				}
			}
		}
	}
}

func Test_FeatureSet_GetMaxLength(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.MaxLength = 3
	newConfig.Sequences = []string{"a", "b"}
	newFeatureSet, err := NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if newFeatureSet.GetMaxLength() != 3 {
		t.Fatal("expected", 3, "got", newFeatureSet.GetMaxLength())
	}

	newConfig = DefaultConfig()
	newConfig.MaxLength = 28
	newConfig.Sequences = []string{"a", "b"}
	newFeatureSet, err = NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if newFeatureSet.GetMaxLength() != 28 {
		t.Fatal("expected", 28, "got", newFeatureSet.GetMaxLength())
	}
}

func Test_FeatureSet_GetMinLength(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.MinLength = 3
	newConfig.Sequences = []string{"a", "b"}
	newFeatureSet, err := NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if newFeatureSet.GetMinLength() != 3 {
		t.Fatal("expected", 3, "got", newFeatureSet.GetMinLength())
	}

	newConfig = DefaultConfig()
	newConfig.MinLength = 28
	newConfig.Sequences = []string{"a", "b"}
	newFeatureSet, err = NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if newFeatureSet.GetMinLength() != 28 {
		t.Fatal("expected", 28, "got", newFeatureSet.GetMinLength())
	}
}

func Test_FeatureSet_GetMinCount(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.MinCount = -1
	newConfig.Sequences = []string{"a", "b"}
	newFeatureSet, err := NewFeatureSet(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", false, "got", true)
	}

	newConfig = DefaultConfig()
	newConfig.MinCount = 3
	newConfig.Sequences = []string{"a", "b"}
	newFeatureSet, err = NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if newFeatureSet.GetMinCount() != 3 {
		t.Fatal("expected", 3, "got", newFeatureSet.GetMinCount())
	}

	newConfig = DefaultConfig()
	newConfig.MinCount = 28
	newConfig.Sequences = []string{"a", "b"}
	newFeatureSet, err = NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if newFeatureSet.GetMinCount() != 28 {
		t.Fatal("expected", 28, "got", newFeatureSet.GetMinCount())
	}
}

func Test_FeatureSet_GetSeparator(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Separator = ","
	newConfig.Sequences = []string{"a", "b"}
	newFeatureSet, err := NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if newFeatureSet.GetSeparator() != "," {
		t.Fatal("expected", ",", "got", newFeatureSet.GetSeparator())
	}

	newConfig = DefaultConfig()
	newConfig.Separator = "foo"
	newConfig.Sequences = []string{"a", "b"}
	newFeatureSet, err = NewFeatureSet(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if newFeatureSet.GetSeparator() != "foo" {
		t.Fatal("expected", "foo", "got", newFeatureSet.GetSeparator())
	}
}

func Test_FeatureSet_MinLengthMaxLength(t *testing.T) {
	testCases := []struct {
		MinLength    int
		MaxLength    int
		Sequences    []string
		Expected     []string
		ErrorMatcher func(err error) bool
	}{
		{
			MinLength:    1,
			MaxLength:    -1,
			Sequences:    []string{"ab"},
			Expected:     nil,
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab"},
			Expected:     []string{"a", "b", "ab"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab", "ab"},
			Expected:     []string{"a", "b", "ab"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab", "abc"},
			Expected:     []string{"a", "b", "ab"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"a", "b", "c", "ab", "bc", "abc"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    1,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"a", "b", "c"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    2,
			MaxLength:    2,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"ab", "bc"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    3,
			MaxLength:    3,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"abc"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    2,
			MaxLength:    3,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"ab", "bc", "abc"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    2,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"ab", "bc", "abc"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    -2,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     nil,
			ErrorMatcher: IsInvalidConfig,
		},
		{
			MinLength:    2,
			MaxLength:    1,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     nil,
			ErrorMatcher: IsInvalidConfig,
		},
		{
			MinLength:    0,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     nil,
			ErrorMatcher: IsInvalidConfig,
		},
	}

	for i, testCase := range testCases {
		newConfig := DefaultConfig()
		newConfig.MinCount = 2
		newConfig.MaxLength = testCase.MaxLength
		newConfig.MinLength = testCase.MinLength
		newConfig.Sequences = testCase.Sequences
		newFeatureSet, err := NewFeatureSet(newConfig)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			err = newFeatureSet.Scan()
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
			var sequences []string
			for _, f := range newFeatureSet.GetFeatures() {
				sequences = append(sequences, f.GetSequence())
			}

			sort.Strings(sequences)
			sort.Strings(testCase.Expected)
			if !reflect.DeepEqual(sequences, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", sequences)
			}
		}
	}
}

func Test_FeatureSet_GetSequences(t *testing.T) {
	newSequences := []string{
		"This is, a test.",
		"This is, another test.",
	}

	newConfig := DefaultConfig()
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

func Test_FeatureSet_Separator(t *testing.T) {
	newConfig := DefaultConfig()
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
