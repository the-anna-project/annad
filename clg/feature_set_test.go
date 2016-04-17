package clg

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/feature-set"
	"github.com/xh3b4sd/anna/spec"
)

func Test_FeatureSet_NewFeatureSet_Success(t *testing.T) {
	testCases := []struct {
		// Note that this input is used to validate the output. We only want to
		// ensure that the feature set that is dynmaically created within the CLG
		// is properly configured. Thus the input configuration should equal the
		// output configuration.
		Input []interface{}
	}{
		{
			Input: []interface{}{[]string{"foo"}},
		},
		{
			Input: []interface{}{[]string{"foo", "bar"}, -1, 1},
		},
		{
			Input: []interface{}{[]string{"baz", "zap", "peng"}, 4, 3},
		},
		{
			Input: []interface{}{[]string{"foo", "bar"}, 4, 3, 5, " "},
		},
		{
			Input: []interface{}{[]string{"o", "r", "t"}, 11, 8, 1, "|"},
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		// Test.
		output, err := newCLGIndex.GetNewFeatureSet(testCase.Input...)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		// Convert.
		if len(output) > 1 {
			t.Fatal("case", i+1, "expected", 1, "got", len(output))
		}
		fs, err := ArgToFeatureSet(output, 0)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		defaultConfig := featureset.DefaultFeatureSetConfig()

		// Assert.
		sequences, err := ArgToStringSlice(testCase.Input, 0, defaultConfig.Sequences)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		if !reflect.DeepEqual(fs.GetSequences(), sequences) {
			t.Fatal("case", i+1, "expected", sequences, "got", fs.GetSequences())
		}
		maxLength, err := ArgToInt(testCase.Input, 1, defaultConfig.MaxLength)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		if fs.GetMaxLength() != maxLength {
			t.Fatal("case", i+1, "expected", maxLength, "got", fs.GetMaxLength())
		}
		minLength, err := ArgToInt(testCase.Input, 2, defaultConfig.MinLength)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		if fs.GetMinLength() != minLength {
			t.Fatal("case", i+1, "expected", minLength, "got", fs.GetMinLength())
		}
		minCount, err := ArgToInt(testCase.Input, 3, defaultConfig.MinCount)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		if fs.GetMinCount() != minCount {
			t.Fatal("case", i+1, "expected", minCount, "got", fs.GetMinCount())
		}
		separator, err := ArgToString(testCase.Input, 4, defaultConfig.Separator)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		if fs.GetSeparator() != separator {
			t.Fatal("case", i+1, "expected", separator, "got", fs.GetSeparator())
		}
	}
}

func Test_FeatureSet_NewFeatureSet_Error(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{},
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{"foo"},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"foo", "bar"}, true},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"foo", "bar"}, -1, []int{1, 2, 3}},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"foo", "bar"}, -1, 1, "invalid"},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"foo", "bar"}, -1, 1, 3, 8.1},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]string{"foo", "bar"}, -2},
			ErrorMatcher: featureset.IsInvalidConfig,
		},
		{
			Input:        []interface{}{[]string{"foo", "bar"}, -1, 0},
			ErrorMatcher: featureset.IsInvalidConfig,
		},
		{
			Input:        []interface{}{[]string{"foo", "bar"}, -1, 0},
			ErrorMatcher: featureset.IsInvalidConfig,
		},
		{
			Input:        []interface{}{[]string{"foo", "bar"}, -1, 1, 0},
			ErrorMatcher: featureset.IsInvalidConfig,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		_, err := newCLGIndex.GetNewFeatureSet(testCase.Input...)
		if !testCase.ErrorMatcher(err) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
	}
}

func Test_FeatureSet_GetMaxLengthFeatureSet(t *testing.T) {
	testFeatureSet := func(maxLength int) spec.FeatureSet {
		newConfig := featureset.DefaultFeatureSetConfig()
		newConfig.MaxLength = maxLength
		newConfig.Sequences = []string{"a", "b"}
		newFeatureSet, err := featureset.NewFeatureSet(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newFeatureSet
	}

	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{testFeatureSet(3)},
			Expected:     []interface{}{3},
			ErrorMatcher: nil,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.GetMaxLengthFeatureSet(testCase.Input...)
		if testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}
