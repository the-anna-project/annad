package clg

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/index/clg/feature-set"
	"github.com/xh3b4sd/anna/spec"
)

func Test_Feature_AddPositionFeature(t *testing.T) {
	testFeature := func(positions [][]float64, sequence string) spec.Feature {
		newConfig := featureset.DefaultFeatureConfig()
		newConfig.Positions = positions
		newConfig.Sequence = sequence
		newFeature, err := featureset.NewFeature(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newFeature
	}

	testCases := []struct {
		Input             []interface{}
		ExpectedPositions [][]float64
		ErrorMatcher      func(err error) bool
	}{
		{
			Input:             []interface{}{testFeature([][]float64{{1, 1}}, ","), []float64{2, 2}},
			ExpectedPositions: [][]float64{{1, 1}, {2, 2}},
			ErrorMatcher:      nil,
		},
		{
			Input:             []interface{}{testFeature([][]float64{{5, 1}, {1, 1}}, "seq"), []float64{-32, 4}},
			ExpectedPositions: [][]float64{{5, 1}, {1, 1}, {-32, 4}},
			ErrorMatcher:      nil,
		},
		{
			Input:             []interface{}{testFeature([][]float64{{2, 2}, {1, 1}, {2, 2}}, "test sequence"), []float64{2, 2}},
			ExpectedPositions: [][]float64{{2, 2}, {1, 1}, {2, 2}, {2, 2}},
			ErrorMatcher:      nil,
		},
		{
			Input:             []interface{}{testFeature([][]float64{{1, 1}}, ","), []float64{1, 2, 3}},
			ExpectedPositions: nil,
			ErrorMatcher:      featureset.IsInvalidPosition,
		},
		{
			Input:             []interface{}{testFeature([][]float64{{1, 1}}, "seq"), []float64{2, 2}, "foo"},
			ExpectedPositions: nil,
			ErrorMatcher:      IsTooManyArguments,
		},
		{
			Input:             []interface{}{testFeature([][]float64{{1, 1}}, "seq"), 2.2},
			ExpectedPositions: nil,
			ErrorMatcher:      IsWrongArgumentType,
		},
		{
			Input:             []interface{}{true, []float64{2, 2}},
			ExpectedPositions: nil,
			ErrorMatcher:      IsWrongArgumentType,
		},
		{
			Input:             []interface{}{testFeature([][]float64{{1, 1}}, "seq")},
			ExpectedPositions: nil,
			ErrorMatcher:      IsNotEnoughArguments,
		},
		{
			Input:             []interface{}{},
			ExpectedPositions: nil,
			ErrorMatcher:      IsNotEnoughArguments,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.AddPositionFeature(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			f, err := ArgToFeature(testCase.Input, 0)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			if len(output) != 0 {
				t.Fatal("expected", 0, "got", len(output))
			}
			if !reflect.DeepEqual(f.GetPositions(), testCase.ExpectedPositions) {
				t.Fatal("case", i+1, "expected", testCase.ExpectedPositions, "got", f.GetPositions())
			}
		}
	}
}

func Test_Feature_GetCountFeature(t *testing.T) {
	testFeature := func(positions [][]float64, sequence string) spec.Feature {
		newConfig := featureset.DefaultFeatureConfig()
		newConfig.Positions = positions
		newConfig.Sequence = sequence
		newFeature, err := featureset.NewFeature(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newFeature
	}

	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{testFeature([][]float64{{1, 1}}, ",")},
			Expected:     []interface{}{1},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeature([][]float64{{33, 65}, {2, -1}, {-72, 337}}, ",")},
			Expected:     []interface{}{3},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeature([][]float64{{3, 3}, {7, 7}}, ",")},
			Expected:     []interface{}{2},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeature([][]float64{{1, 1}}, ","), "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.GetCountFeature(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

func Test_Feature_GetDistributionFeature(t *testing.T) {
	testFeature := func(positions [][]float64, sequence string) spec.Feature {
		newConfig := featureset.DefaultFeatureConfig()
		newConfig.Positions = positions
		newConfig.Sequence = sequence
		newFeature, err := featureset.NewFeature(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newFeature
	}

	testCases := []struct {
		Input           []interface{}
		ExpectedName    string
		ExpectedVectors [][]float64
		ErrorMatcher    func(err error) bool
	}{
		{
			Input:           []interface{}{testFeature([][]float64{{1, 1}}, ",")},
			ExpectedVectors: [][]float64{{1, 1}},
			ExpectedName:    ",",
			ErrorMatcher:    nil,
		},
		{
			Input:           []interface{}{testFeature([][]float64{{23, -1}, {2, 2}}, "seq")},
			ExpectedVectors: [][]float64{{23, -1}, {2, 2}},
			ExpectedName:    "seq",
			ErrorMatcher:    nil,
		},
		{
			Input:           []interface{}{testFeature([][]float64{{-42, -3}, {13, -6}, {-4, 8}}, "testseq")},
			ExpectedVectors: [][]float64{{-42, -3}, {13, -6}, {-4, 8}},
			ExpectedName:    "testseq",
			ErrorMatcher:    nil,
		},
		{
			Input:           []interface{}{testFeature([][]float64{{1, 1}}, ","), "foo"},
			ExpectedName:    "",
			ExpectedVectors: nil,
			ErrorMatcher:    IsTooManyArguments,
		},
		{
			Input:           []interface{}{8.1},
			ExpectedName:    "",
			ExpectedVectors: nil,
			ErrorMatcher:    IsWrongArgumentType,
		},
		{
			Input:           []interface{}{},
			ExpectedName:    "",
			ExpectedVectors: nil,
			ErrorMatcher:    IsNotEnoughArguments,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.GetDistributionFeature(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			d, err := ArgToDistribution(output, 0)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			if len(output) > 1 {
				t.Fatal("expected", 1, "got", len(output))
			}
			if d.GetName() != testCase.ExpectedName {
				t.Fatal("case", i+1, "expected", testCase.ExpectedName, "got", d.GetName())
			}
			if !reflect.DeepEqual(d.GetVectors(), testCase.ExpectedVectors) {
				t.Fatal("case", i+1, "expected", testCase.ExpectedVectors, "got", d.GetVectors())
			}
		}
	}
}

func Test_Feature_NewFeatureSet(t *testing.T) {
	testCases := []struct {
		// Note that this input is used to validate the output. We only want to
		// ensure that the feature that is dynmaically created within the CLG is
		// properly configured. Thus the input configuration should equal the
		// output configuration.
		Input        []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{[][]float64{{2, 3}, {16, 23}}, "test-sequence"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[][]float64{{0, -4}, {-28, 14}, {11, 11}}, "seq"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[][]float64{{1, 1}}, "s"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[][]float64{{1, 2}, {1, 2}}, "s"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{[][]float64{{1, 2}, {1, 2}}, "s", "foo"},
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{[][]float64{{1, 2}, {1, 2, 3}}, "s"},
			ErrorMatcher: featureset.IsInvalidConfig,
		},
		{
			Input:        []interface{}{[][]float64{}, "s"},
			ErrorMatcher: featureset.IsInvalidConfig,
		},
		{
			Input:        []interface{}{[][]float64{{1, 1}}, ""},
			ErrorMatcher: featureset.IsInvalidConfig,
		},
		{
			Input:        []interface{}{[][]float64{{1, 1}}},
			ErrorMatcher: featureset.IsInvalidConfig,
		},
		{
			Input:        []interface{}{},
			ErrorMatcher: featureset.IsInvalidConfig,
		},
		{
			Input:        []interface{}{nil, ""},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{nil, "s"},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[]float64{1, 1}, "seq"},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{1.1, "seq"},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{[][]float64{{1, 1}}, false},
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		// Test.
		output, err := newCLGIndex.GetNewFeature(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		// Convert.
		if testCase.ErrorMatcher == nil {
			if len(output) > 1 {
				t.Fatal("case", i+1, "expected", 1, "got", len(output))
			}
			f, err := ArgToFeature(output, 0)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
			defaultConfig := featureset.DefaultFeatureConfig()

			// Assert.
			positions, err := ArgToFloat64SliceSlice(testCase.Input, 0, defaultConfig.Positions)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
			if !reflect.DeepEqual(f.GetPositions(), positions) {
				t.Fatal("case", i+1, "expected", positions, "got", f.GetPositions())
			}
			sequence, err := ArgToString(testCase.Input, 1, defaultConfig.Sequence)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
			if f.GetSequence() != sequence {
				t.Fatal("case", i+1, "expected", sequence, "got", f.GetSequence())
			}
		}
	}
}

func Test_Feature_GetPositionsFeature(t *testing.T) {
	testFeature := func(positions [][]float64, sequence string) spec.Feature {
		newConfig := featureset.DefaultFeatureConfig()
		newConfig.Positions = positions
		newConfig.Sequence = sequence
		newFeature, err := featureset.NewFeature(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newFeature
	}

	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{testFeature([][]float64{{1, 1}, {2, 7}}, ",")},
			Expected:     []interface{}{[][]float64{{1, 1}, {2, 7}}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeature([][]float64{{3, 2}, {10, 3}}, ",")},
			Expected:     []interface{}{[][]float64{{3, 2}, {10, 3}}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeature([][]float64{{11, 22}, {33, 44}}, ",")},
			Expected:     []interface{}{[][]float64{{11, 22}, {33, 44}}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeature([][]float64{{11, 22}, {33, 44}}, ","), "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.GetPositionsFeature(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

func Test_Feature_GetSequenceFeature(t *testing.T) {
	testFeature := func(positions [][]float64, sequence string) spec.Feature {
		newConfig := featureset.DefaultFeatureConfig()
		newConfig.Positions = positions
		newConfig.Sequence = sequence
		newFeature, err := featureset.NewFeature(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newFeature
	}

	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{testFeature([][]float64{{1, 1}}, ",")},
			Expected:     []interface{}{","},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeature([][]float64{{1, 1}}, "|")},
			Expected:     []interface{}{"|"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeature([][]float64{{1, 1}}, "foo")},
			Expected:     []interface{}{"foo"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeature([][]float64{{1, 1}}, ","), "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{8.1},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.GetSequenceFeature(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}
