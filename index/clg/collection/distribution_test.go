package collection

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/index/clg/collection/distribution"
	"github.com/xh3b4sd/anna/spec"
)

func Test_Distribution_CalculateDistribution(t *testing.T) {
	testDistribution := func(staticChannels []float64, vectors [][]float64) spec.Distribution {
		newConfig := distribution.DefaultConfig()
		newConfig.Name = "name"
		newConfig.StaticChannels = staticChannels
		newConfig.Vectors = vectors
		newDistribution, err := distribution.NewDistribution(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newDistribution
	}

	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{testDistribution([]float64{50, 100}, [][]float64{{11, 22}, {33, 44}})},
			Expected:     []interface{}{[]float64{2, 0}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([]float64{50, 100}, [][]float64{{11, 22}, {88, 99}})},
			Expected:     []interface{}{[]float64{1, 1}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([]float64{50, 100}, [][]float64{{54, 68}, {71, 84}})},
			Expected:     []interface{}{[]float64{0, 2}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([]float64{33, 66, 99}, [][]float64{{54, 68}, {71, 84}})},
			Expected:     []interface{}{[]float64{0, 0.5, 1.5}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([]float64{50, 100}, [][]float64{{11, 22}, {33, 44}}), "foo"},
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

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).CalculateDistribution(testCase.Input...)
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

func Test_Distribution_DifferenceDistribution(t *testing.T) {
	testDistribution := func(staticChannels []float64, vectors [][]float64) spec.Distribution {
		newConfig := distribution.DefaultConfig()
		newConfig.Name = "name"
		newConfig.StaticChannels = staticChannels
		newConfig.Vectors = vectors
		newDistribution, err := distribution.NewDistribution(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newDistribution
	}

	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input: []interface{}{
				testDistribution([]float64{50, 100}, [][]float64{{11, 22}, {33, 44}}),
				testDistribution([]float64{50, 100}, [][]float64{{35, 48}, {13, 22}}),
			},
			Expected:     []interface{}{[]float64{0, 0}},
			ErrorMatcher: nil,
		},
		{
			Input: []interface{}{
				testDistribution([]float64{50, 100}, [][]float64{{11, 22}, {33, 44}}),
				testDistribution([]float64{50, 100}, [][]float64{{57, 84}, {69, 87}}),
			},
			Expected:     []interface{}{[]float64{-2, 2}},
			ErrorMatcher: nil,
		},
		{
			Input: []interface{}{
				testDistribution([]float64{50, 100}, [][]float64{{11, 22}, {33, 44}}),
				testDistribution([]float64{50, 100}, [][]float64{{57, 84}, {69, 87}}),
				"foo",
			},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input: []interface{}{
				81,
				testDistribution([]float64{50, 100}, [][]float64{{57, 84}, {69, 87}}),
			},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input: []interface{}{
				testDistribution([]float64{50, 100}, [][]float64{{57, 84}, {69, 87}}),
				81,
			},
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
		output, err := testMaybeNewCollection(t).DifferenceDistribution(testCase.Input...)
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

func Test_Distribution_GetDimensionsDistribution(t *testing.T) {
	testDistribution := func(vectors [][]float64) spec.Distribution {
		newConfig := distribution.DefaultConfig()
		newConfig.Name = "test-name"
		newConfig.Vectors = vectors
		newDistribution, err := distribution.NewDistribution(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newDistribution
	}

	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{testDistribution([][]float64{{1, 1}, {2, 7}})},
			Expected:     []interface{}{2},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([][]float64{{1, 1, 1, 1, 1}, {2, 2, 2, 2, 2}})},
			Expected:     []interface{}{5},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([][]float64{{31, 12, 9}, {6, 2, -4}})},
			Expected:     []interface{}{3},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([][]float64{{1, 1}, {2, 7}}), "foo"},
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

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetDimensionsDistribution(testCase.Input...)
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

func Test_Distribution_GetNameDistribution(t *testing.T) {
	testDistribution := func(name string) spec.Distribution {
		newConfig := distribution.DefaultConfig()
		newConfig.Name = name
		newConfig.Vectors = [][]float64{{1, 2}, {3, 4}}
		newDistribution, err := distribution.NewDistribution(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newDistribution
	}

	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{testDistribution("foo")},
			Expected:     []interface{}{"foo"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution("name")},
			Expected:     []interface{}{"name"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution("test")},
			Expected:     []interface{}{"test"},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution("test"), "foo"},
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

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetNameDistribution(testCase.Input...)
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

func Test_Distribution_GetNewDistribution(t *testing.T) {
	testCases := []struct {
		// Note that this input is used to validate the output. We only want to
		// ensure that the distribution that is dynmaically created within the CLG
		// is properly configured. Thus the input configuration should equal the
		// output configuration.
		Input        []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"name", DefaultArg{}, [][]float64{{1, 2}, {3, 4}, {5, 6}}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"foo", DefaultArg{}, [][]float64{{4, 5}, {6, 7}, {8, 9}}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{true, DefaultArg{}, [][]float64{{4, 5}, {6, 7}, {8, 9}}},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"foo", 3.4, [][]float64{{4, 5}, {6, 7}, {8, 9}}},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"foo", DefaultArg{}, []int{1, 2, 3}},
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"foo", DefaultArg{}, [][]float64{{4, 5}, {6, 7}, {8, 9}}, "foo"},
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"foo", DefaultArg{}},
			ErrorMatcher: distribution.IsInvalidConfig,
		},
		{
			Input:        []interface{}{},
			ErrorMatcher: distribution.IsInvalidConfig,
		},
		{
			Input:        []interface{}{"foo", 3.4, [][]float64{}},
			ErrorMatcher: distribution.IsInvalidConfig,
		},
	}

	for i, testCase := range testCases {
		if i != 7 {
			continue
		}
		// Test.
		output, err := testMaybeNewCollection(t).GetNewDistribution(testCase.Input...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		// Convert.
		if testCase.ErrorMatcher == nil {
			if len(output) > 1 {
				t.Fatal("case", i+1, "expected", 1, "got", len(output))
			}
			d, err := ArgToDistribution(output, 0)
			testMaybeFatalCase(t, i, err)
			defaultConfig := distribution.DefaultConfig()

			// Assert. Note that we don't test the hash map here, since this is not a
			// configuration getter.
			name, err := ArgToString(testCase.Input, 1, defaultConfig.Name)
			testMaybeFatalCase(t, i, err)
			if d.Metadata()["name"] != name {
				t.Fatal("case", i+1, "expected", name, "got", d.Metadata()["name"])
			}
			staticChannels, err := ArgToFloat64Slice(testCase.Input, 2, defaultConfig.StaticChannels)
			testMaybeFatalCase(t, i, err)
			if !reflect.DeepEqual(d.GetStaticChannels(), staticChannels) {
				t.Fatal("case", i+1, "expected", staticChannels, "got", d.GetStaticChannels())
			}
			vectors, err := ArgToFloat64SliceSlice(testCase.Input, 3, defaultConfig.Vectors)
			testMaybeFatalCase(t, i, err)
			if !reflect.DeepEqual(d.GetVectors(), vectors) {
				t.Fatal("case", i+1, "expected", vectors, "got", d.GetVectors())
			}
		}
	}
}

func Test_Distribution_GetStaticChannelsDistribution(t *testing.T) {
	testDistribution := func(staticChannels []float64) spec.Distribution {
		newConfig := distribution.DefaultConfig()
		newConfig.Name = "test-name"
		newConfig.StaticChannels = staticChannels
		newConfig.Vectors = [][]float64{{1, 2}, {3, 4}}
		newDistribution, err := distribution.NewDistribution(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newDistribution
	}

	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{testDistribution([]float64{1, 2, 3, 4})},
			Expected:     []interface{}{[]float64{1, 2, 3, 4}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([]float64{4, 8, 12, 16})},
			Expected:     []interface{}{[]float64{4, 8, 12, 16}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([]float64{20, 40, 60, 80, 100})},
			Expected:     []interface{}{[]float64{20, 40, 60, 80, 100}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([]float64{1, 2, 3, 4}), "foo"},
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

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetStaticChannelsDistribution(testCase.Input...)
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

func Test_Distribution_GetVectorsDistribution(t *testing.T) {
	testDistribution := func(vectors [][]float64) spec.Distribution {
		newConfig := distribution.DefaultConfig()
		newConfig.Name = "test-name"
		newConfig.Vectors = vectors
		newDistribution, err := distribution.NewDistribution(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		return newDistribution
	}

	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{testDistribution([][]float64{{1, 1}, {2, 7}})},
			Expected:     []interface{}{[][]float64{{1, 1}, {2, 7}}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([][]float64{{3, 2}, {10, 3}})},
			Expected:     []interface{}{[][]float64{{3, 2}, {10, 3}}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([][]float64{{11, 22}, {33, 44}})},
			Expected:     []interface{}{[][]float64{{11, 22}, {33, 44}}},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{testDistribution([][]float64{{11, 22}, {33, 44}}), "foo"},
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

	for i, testCase := range testCases {
		output, err := testMaybeNewCollection(t).GetVectorsDistribution(testCase.Input...)
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
