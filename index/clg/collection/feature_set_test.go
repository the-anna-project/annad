package collection

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testFeatureSet(t *testing.T, ss []string) spec.FeatureSet {
	newConfig := featureset.DefaultConfig()
	newConfig.Sequences = ss
	newFeatureSet, err := featureset.New(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newFeatureSet.Scan()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	return newFeatureSet
}

func Test_FeatureSet_GetFeaturesFeatureSet(t *testing.T) {
	testCases := []struct {
		Input          []interface{}
		ExpectedSubSet []string
		ErrorMatcher   func(err error) bool
	}{
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."})},
			ExpectedSubSet: []string{"This", "is", "test", "."},
			ErrorMatcher:   nil,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), "foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsTooManyArguments,
		},
		{
			Input:          []interface{}{"foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsWrongArgumentType,
		},
		{
			Input:          []interface{}{},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetFeaturesFeatureSet(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			fs, err := ArgToFeatures(output, 0)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			if len(output) > 1 {
				t.Fatal("expected", 1, "got", len(output))
			}
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

func Test_FeatureSet_GetFeaturesByCountFeatureSet_Expected(t *testing.T) {
	testCases := []struct {
		Input          []interface{}
		ExpectedSubSet []string
		ErrorMatcher   func(err error) bool
	}{
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 1},
			ExpectedSubSet: []string{"This", "is", "test", "."},
			ErrorMatcher:   nil,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 3},
			ExpectedSubSet: []string{" ", "is"},
			ErrorMatcher:   nil,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 11, "foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsTooManyArguments,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), "foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsWrongArgumentType,
		},
		{
			Input:          []interface{}{"foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsWrongArgumentType,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."})},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsNotEnoughArguments,
		},
		{
			Input:          []interface{}{},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetFeaturesByCountFeatureSet(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			fs, err := ArgToFeatures(output, 0)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			if len(output) > 1 {
				t.Fatal("expected", 1, "got", len(output))
			}
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

func Test_FeatureSet_GetFeaturesByCountFeatureSet_Unexpected(t *testing.T) {
	testCases := []struct {
		Input            []interface{}
		UnexpectedSubSet []string
		ErrorMatcher     func(err error) bool
	}{
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 1},
			UnexpectedSubSet: nil,
			ErrorMatcher:     nil,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 3},
			UnexpectedSubSet: []string{"This", "."},
			ErrorMatcher:     nil,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 11, "foo"},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsTooManyArguments,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), "foo"},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsWrongArgumentType,
		},
		{
			Input:            []interface{}{"foo"},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsWrongArgumentType,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."})},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsNotEnoughArguments,
		},
		{
			Input:            []interface{}{},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetFeaturesByCountFeatureSet(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			fs, err := ArgToFeatures(output, 0)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			if len(output) > 1 {
				t.Fatal("expected", 1, "got", len(output))
			}
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

func Test_FeatureSet_GetFeaturesByLengthFeatureSet_Expected(t *testing.T) {
	testCases := []struct {
		Input          []interface{}
		ExpectedSubSet []string
		ErrorMatcher   func(err error) bool
	}{
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 1, -1},
			ExpectedSubSet: []string{"This", "is", "test", "."},
			ErrorMatcher:   nil,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 2, 2},
			ExpectedSubSet: []string{" a", "is", "te"},
			ErrorMatcher:   nil,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 11, 22, "foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsTooManyArguments,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), "foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsWrongArgumentType,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 3, "foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsWrongArgumentType,
		},
		{
			Input:          []interface{}{"foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsWrongArgumentType,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."})},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsNotEnoughArguments,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 3},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsNotEnoughArguments,
		},
		{
			Input:          []interface{}{},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetFeaturesByLengthFeatureSet(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			fs, err := ArgToFeatures(output, 0)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			if len(output) > 1 {
				t.Fatal("expected", 1, "got", len(output))
			}
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

func Test_FeatureSet_GetFeaturesByLengthFeatureSet_Unexpected(t *testing.T) {
	testCases := []struct {
		Input            []interface{}
		UnexpectedSubSet []string
		ErrorMatcher     func(err error) bool
	}{
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 1, -1},
			UnexpectedSubSet: nil,
			ErrorMatcher:     nil,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 2, 2},
			UnexpectedSubSet: []string{"This", "is ", "a", "."},
			ErrorMatcher:     nil,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 11, 22, "foo"},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsTooManyArguments,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), "foo"},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsWrongArgumentType,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 3, "foo"},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsWrongArgumentType,
		},
		{
			Input:            []interface{}{"foo"},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsWrongArgumentType,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."})},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsNotEnoughArguments,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 3},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsNotEnoughArguments,
		},
		{
			Input:            []interface{}{},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetFeaturesByLengthFeatureSet(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			fs, err := ArgToFeatures(output, 0)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			if len(output) > 1 {
				t.Fatal("expected", 1, "got", len(output))
			}
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

func Test_FeatureSet_GetFeaturesBySequenceFeatureSet_Expected(t *testing.T) {
	testCases := []struct {
		Input          []interface{}
		ExpectedSubSet []string
		ErrorMatcher   func(err error) bool
	}{
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), ""},
			ExpectedSubSet: []string{"This", "is", "test", "."},
			ErrorMatcher:   nil,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), "This"},
			ExpectedSubSet: []string{"This", "This is a"},
			ErrorMatcher:   nil,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), "."},
			ExpectedSubSet: []string{"is a test.", ".", "test."},
			ErrorMatcher:   nil,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), ".", "foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsTooManyArguments,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 8.1},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsWrongArgumentType,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 43},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsWrongArgumentType,
		},
		{
			Input:          []interface{}{"foo"},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsWrongArgumentType,
		},
		{
			Input:          []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."})},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsNotEnoughArguments,
		},
		{
			Input:          []interface{}{},
			ExpectedSubSet: nil,
			ErrorMatcher:   IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetFeaturesBySequenceFeatureSet(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			fs, err := ArgToFeatures(output, 0)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			if len(output) > 1 {
				t.Fatal("expected", 1, "got", len(output))
			}
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

func Test_FeatureSet_GetFeaturesBySequenceFeatureSet_Unexpected(t *testing.T) {
	testCases := []struct {
		Input            []interface{}
		UnexpectedSubSet []string
		ErrorMatcher     func(err error) bool
	}{
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), ""},
			UnexpectedSubSet: nil,
			ErrorMatcher:     nil,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), "This"},
			UnexpectedSubSet: []string{"is ", "a", "."},
			ErrorMatcher:     nil,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), "."},
			UnexpectedSubSet: []string{"is ", "test", "a"},
			ErrorMatcher:     nil,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), ".", "foo"},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsTooManyArguments,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 8.1},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsWrongArgumentType,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."}), 43},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsWrongArgumentType,
		},
		{
			Input:            []interface{}{"foo"},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsWrongArgumentType,
		},
		{
			Input:            []interface{}{testFeatureSet(t, []string{"This is a test.", "This is another test."})},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsNotEnoughArguments,
		},
		{
			Input:            []interface{}{},
			UnexpectedSubSet: nil,
			ErrorMatcher:     IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetFeaturesBySequenceFeatureSet(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			fs, err := ArgToFeatures(output, 0)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
			if len(output) > 1 {
				t.Fatal("expected", 1, "got", len(output))
			}
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

func Test_FeatureSet_GetMaxLengthFeatureSet(t *testing.T) {
	testFeatureSet := func(maxLength int) spec.FeatureSet {
		newConfig := featureset.DefaultConfig()
		newConfig.MaxLength = maxLength
		newConfig.Sequences = []string{"a", "b"}
		newFeatureSet, err := featureset.New(newConfig)
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
		{
			Input:        []interface{}{testFeatureSet(8)},
			Expected:     []interface{}{8},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeatureSet(8), "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetMaxLengthFeatureSet(testCase.Input...)
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

func Test_FeatureSet_GetMinLengthFeatureSet(t *testing.T) {
	testFeatureSet := func(minLength int) spec.FeatureSet {
		newConfig := featureset.DefaultConfig()
		newConfig.MinLength = minLength
		newConfig.Sequences = []string{"a", "b"}
		newFeatureSet, err := featureset.New(newConfig)
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
		{
			Input:        []interface{}{testFeatureSet(8)},
			Expected:     []interface{}{8},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeatureSet(8), "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetMinLengthFeatureSet(testCase.Input...)
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

func Test_FeatureSet_GetMinCountFeatureSet(t *testing.T) {
	testFeatureSet := func(minCount int) spec.FeatureSet {
		newConfig := featureset.DefaultConfig()
		newConfig.MinCount = minCount
		newConfig.Sequences = []string{"a", "b"}
		newFeatureSet, err := featureset.New(newConfig)
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
		{
			Input:        []interface{}{testFeatureSet(8)},
			Expected:     []interface{}{8},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeatureSet(8), "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetMinCountFeatureSet(testCase.Input...)
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
			Input:        []interface{}{[]string{"foo", "bar"}, -1, 1, -1},
			ErrorMatcher: featureset.IsInvalidConfig,
		},
		{
			Input:        []interface{}{[]string{"foo", "bar"}, -1, 1, 0, "", "foo"},
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		_, err := testMaybeNewCollection(t).GetNewFeatureSet(testCase.Input...)
		if !testCase.ErrorMatcher(err) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
	}
}

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
			Input: []interface{}{[]string{"foo", "bar"}, 4, 3, 0, " "},
		},
		{
			Input: []interface{}{[]string{"f|o|o", "b|a|r", "f|o|o|d", "b|a|r|k|e|e|p|e|r", "tap"}, 2, 1, 1, "|"},
		},
	}

	for i, testCase := range testCases {
		// Test.
		output, err := testMaybeNewCollection(t).GetNewFeatureSet(testCase.Input...)
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
		defaultConfig := featureset.DefaultConfig()

		// Assert.
		if len(fs.GetFeatures()) == 0 {
			t.Fatal("case", i+1, "expected", ">0", "got", 0)
		}
		sequences, err := ArgToStringSlice(testCase.Input, 0, defaultConfig.Sequences)
		testMaybeFatalCase(t, i, err)
		if !reflect.DeepEqual(fs.GetSequences(), sequences) {
			t.Fatal("case", i+1, "expected", sequences, "got", fs.GetSequences())
		}
		maxLength, err := ArgToInt(testCase.Input, 1, defaultConfig.MaxLength)
		testMaybeFatalCase(t, i, err)
		if fs.GetMaxLength() != maxLength {
			t.Fatal("case", i+1, "expected", maxLength, "got", fs.GetMaxLength())
		}
		minLength, err := ArgToInt(testCase.Input, 2, defaultConfig.MinLength)
		testMaybeFatalCase(t, i, err)
		if fs.GetMinLength() != minLength {
			t.Fatal("case", i+1, "expected", minLength, "got", fs.GetMinLength())
		}
		minCount, err := ArgToInt(testCase.Input, 3, defaultConfig.MinCount)
		testMaybeFatalCase(t, i, err)
		if fs.GetMinCount() != minCount {
			t.Fatal("case", i+1, "expected", minCount, "got", fs.GetMinCount())
		}
		separator, err := ArgToString(testCase.Input, 4, defaultConfig.Separator)
		testMaybeFatalCase(t, i, err)
		if fs.GetSeparator() != separator {
			t.Fatal("case", i+1, "expected", separator, "got", fs.GetSeparator())
		}
	}
}

func Test_FeatureSet_GetSeparatorFeatureSet(t *testing.T) {
	testFeatureSet := func(separator string) spec.FeatureSet {
		newConfig := featureset.DefaultConfig()
		newConfig.Separator = separator
		newConfig.Sequences = []string{"a", "b"}
		newFeatureSet, err := featureset.New(newConfig)
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
			Input:        []interface{}{testFeatureSet(",")},
			Expected:     []interface{}{","},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeatureSet("|")},
			Expected:     []interface{}{"|"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeatureSet(" ")},
			Expected:     []interface{}{" "},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeatureSet(" "), "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetSeparatorFeatureSet(testCase.Input...)
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

func Test_FeatureSet_GetSequencesFeatureSet(t *testing.T) {
	testFeatureSet := func(sequences []string) spec.FeatureSet {
		newConfig := featureset.DefaultConfig()
		newConfig.Sequences = sequences
		newFeatureSet, err := featureset.New(newConfig)
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
			Input:        []interface{}{testFeatureSet([]string{"a", "b"})},
			Expected:     []interface{}{[]string{"a", "b"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeatureSet([]string{"1", "2"})},
			Expected:     []interface{}{[]string{"1", "2"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeatureSet([]string{"foo", "bar"})},
			Expected:     []interface{}{[]string{"foo", "bar"}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testFeatureSet([]string{"foo", "bar"}), "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"foo"},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetSequencesFeatureSet(testCase.Input...)
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
