package collection

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func Test_Arg_ArgToFeatures_IndexZero_Success(t *testing.T) {
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

	one := testFeature([][]float64{{1, 1}}, ",")
	two := testFeature([][]float64{{2, 2}}, ",")
	three := testFeature([][]float64{{3, 3}}, ",")
	four := testFeature([][]float64{{4, 4}}, ",")

	args := []interface{}{
		[]spec.Feature{one, two},
		[]spec.Feature{three, four},
	}

	output, err := ArgToFeatures(args, 0)
	if err != nil {
		t.Fatal("expected", true, "got", false)
	}
	if !reflect.DeepEqual(output, []spec.Feature{one, two}) {
		t.Fatal("expected", []spec.Feature{one, two}, "got", output)
	}
}

func Test_Arg_ArgToFeatures_IndexOne_Success(t *testing.T) {
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

	one := testFeature([][]float64{{1, 1}}, ",")
	two := testFeature([][]float64{{2, 2}}, ",")
	three := testFeature([][]float64{{3, 3}}, ",")
	four := testFeature([][]float64{{4, 4}}, ",")

	args := []interface{}{
		[]spec.Feature{one, two},
		[]spec.Feature{three, four},
	}

	output, err := ArgToFeatures(args, 1)
	if err != nil {
		t.Fatal("expected", true, "got", false)
	}
	if !reflect.DeepEqual(output, []spec.Feature{three, four}) {
		t.Fatal("expected", []spec.Feature{three, four}, "got", output)
	}
}

func Test_Arg_ArgToFeatures_IndexTwo_Failure(t *testing.T) {
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

	one := testFeature([][]float64{{1, 1}}, ",")
	two := testFeature([][]float64{{2, 2}}, ",")
	three := testFeature([][]float64{{3, 3}}, ",")
	four := testFeature([][]float64{{4, 4}}, ",")

	args := []interface{}{
		[]spec.Feature{one, two},
		[]spec.Feature{three, four},
	}

	_, err := ArgToFeatures(args, 2)
	if !IsNotEnoughArguments(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Arg_ArgToFeatures_WrongType(t *testing.T) {
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

	args := []interface{}{
		testFeature([][]float64{{1, 1}}, ","),
	}
	_, err := ArgToFeatures(args, 0)
	if !IsWrongArgumentType(err) {
		t.Fatal("expected", true, "got", false)
	}

	args = []interface{}{
		true,
	}
	_, err = ArgToFeatures(args, 0)
	if !IsWrongArgumentType(err) {
		t.Fatal("expected", true, "got", false)
	}

	args = []interface{}{
		"foo",
	}
	_, err = ArgToFeatures(args, 0)
	if !IsWrongArgumentType(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Arg_ArgToFloat64Slice(t *testing.T) {
	testCases := []struct {
		Args         []interface{}
		Index        int
		Def          [][]float64
		Expected     []float64
		ErrorMatcher func(err error) bool
	}{
		{
			Args:         []interface{}{[]float64{1.1, 3.3}, []float64{2.2, 4.4}, true},
			Index:        0,
			Def:          nil,
			Expected:     []float64{1.1, 3.3},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]float64{1.1, 3.3}, []float64{2.2, 4.4}, true},
			Index:        1,
			Def:          nil,
			Expected:     []float64{2.2, 4.4},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]float64{1.1, 3.3}, []float64{2.2, 4.4}, true},
			Index:        2,
			Def:          nil,
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Args:         []interface{}{[]float64{1.1, 3.3}, []float64{2.2, 4.4}, true},
			Index:        3,
			Def:          nil,
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Args:         []interface{}{[]float64{1.1, 3.3}, []float64{2.2, 4.4}, true},
			Index:        2,
			Def:          [][]float64{{2.2, 4.4}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Args:         []interface{}{[]float64{1.1, 3.3}, []float64{2.2, 4.4}, true},
			Index:        3,
			Def:          [][]float64{{4.4, 6.6}},
			Expected:     []float64{4.4, 6.6},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]float64{1.1, 3.3}, []float64{2.2, 4.4}, true},
			Index:        3,
			Def:          [][]float64{{4.4, 6.6}, {7.7, 9.9}},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := ArgToFloat64Slice(testCase.Args, testCase.Index, testCase.Def...)
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

func Test_Arg_ArgToFloat64SliceSlice(t *testing.T) {
	testCases := []struct {
		Args         []interface{}
		Index        int
		Def          [][][]float64
		Expected     [][]float64
		ErrorMatcher func(err error) bool
	}{
		{
			Args:         []interface{}{[][]float64{{1.1, 3.3}, {12.2, 14.2}}, [][]float64{{2.2, 4.4}, {44.2, 66.2}}, true},
			Index:        0,
			Def:          nil,
			Expected:     [][]float64{{1.1, 3.3}, {12.2, 14.2}},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[][]float64{{1.1, 3.3}, {12.2, 14.2}}, [][]float64{{2.2, 4.4}, {44.2, 66.2}}, true},
			Index:        1,
			Def:          nil,
			Expected:     [][]float64{{2.2, 4.4}, {44.2, 66.2}},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[][]float64{{1.1, 3.3}, {12.2, 14.2}}, [][]float64{{2.2, 4.4}, {44.2, 66.2}}, true},
			Index:        2,
			Def:          nil,
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Args:         []interface{}{[][]float64{{1.1, 3.3}, {12.2, 14.2}}, [][]float64{{2.2, 4.4}, {44.2, 66.2}}, true},
			Index:        3,
			Def:          nil,
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Args:         []interface{}{[][]float64{{1.1, 3.3}, {12.2, 14.2}}, [][]float64{{2.2, 4.4}, {44.2, 66.2}}, true},
			Index:        2,
			Def:          [][][]float64{{{5.5, 7.7}, {47.5, 69.5}}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Args:         []interface{}{[][]float64{{1.1, 3.3}, {12.2, 14.2}}, [][]float64{{2.2, 4.4}, {44.2, 66.2}}, true},
			Index:        3,
			Def:          [][][]float64{{{5.5, 7.7}, {47.5, 69.5}}},
			Expected:     [][]float64{{5.5, 7.7}, {47.5, 69.5}},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[][]float64{{1.1, 3.3}, {12.2, 14.2}}, [][]float64{{2.2, 4.4}, {44.2, 66.2}}, true},
			Index:        3,
			Def:          [][][]float64{{{5.5, 7.7}, {47.5, 69.5}}, {{17.3, 2.4}, {48.5, 29.4}}},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := ArgToFloat64SliceSlice(testCase.Args, testCase.Index, testCase.Def...)
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

func Test_Arg_ArgToInt(t *testing.T) {
	testCases := []struct {
		Args         []interface{}
		Index        int
		Def          []int
		Expected     int
		ErrorMatcher func(err error) bool
	}{
		{
			Args:         []interface{}{1, 2},
			Index:        0,
			Def:          []int{},
			Expected:     1,
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{1, 2},
			Index:        1,
			Def:          []int{},
			Expected:     2,
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{1, 2},
			Index:        0,
			Def:          []int{3},
			Expected:     1,
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{1, 2},
			Index:        1,
			Def:          []int{3},
			Expected:     2,
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{1, 2},
			Index:        2,
			Def:          []int{3},
			Expected:     3,
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{1, 2},
			Index:        2,
			Def:          []int{3, 4},
			Expected:     0,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := ArgToInt(testCase.Args, testCase.Index, testCase.Def...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if output != testCase.Expected {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

func Test_Arg_ArgToString(t *testing.T) {
	testCases := []struct {
		Args         []interface{}
		Index        int
		Def          []string
		Expected     string
		ErrorMatcher func(err error) bool
	}{
		{
			Args:         []interface{}{"a", "b"},
			Index:        0,
			Def:          []string{},
			Expected:     "a",
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{"a", "b"},
			Index:        1,
			Def:          []string{},
			Expected:     "b",
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{"a", "b"},
			Index:        0,
			Def:          []string{"c"},
			Expected:     "a",
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{"a", "b"},
			Index:        1,
			Def:          []string{"c"},
			Expected:     "b",
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{"a", "b"},
			Index:        2,
			Def:          []string{"c"},
			Expected:     "c",
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{"a", "b"},
			Index:        2,
			Def:          []string{"c", "d"},
			Expected:     "",
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := ArgToString(testCase.Args, testCase.Index, testCase.Def...)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if output != testCase.Expected {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

func Test_Arg_ArgToStringSlice(t *testing.T) {
	testCases := []struct {
		Args         []interface{}
		Index        int
		Def          [][]string
		Expected     []string
		ErrorMatcher func(err error) bool
	}{
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        0,
			Def:          [][]string{},
			Expected:     []string{"a"},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        1,
			Def:          [][]string{},
			Expected:     []string{"b"},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        0,
			Def:          [][]string{{"c"}},
			Expected:     []string{"a"},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{DefaultArg{}, []string{"b"}},
			Index:        0,
			Def:          [][]string{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Args:         []interface{}{DefaultArg{}, []string{"b"}},
			Index:        0,
			Def:          [][]string{{"c"}},
			Expected:     []string{"c"},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        1,
			Def:          [][]string{{"c"}},
			Expected:     []string{"b"},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        2,
			Def:          [][]string{},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        2,
			Def:          [][]string{{"c"}},
			Expected:     []string{"c"},
			ErrorMatcher: nil,
		},
		{
			Args:         []interface{}{[]string{"a"}, []string{"b"}},
			Index:        2,
			Def:          [][]string{{"c"}, {"d"}},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := ArgToStringSlice(testCase.Args, testCase.Index, testCase.Def...)
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

func Test_Arg_ValuesToArgs(t *testing.T) {
	testCases := []struct {
		Input        []reflect.Value
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []reflect.Value{reflect.ValueOf([]interface{}{[]string{"a", "b", "c"}, 3.8, true}), reflect.ValueOf(nil)},
			Expected:     []interface{}{[]string{"a", "b", "c"}, 3.8, true},
			ErrorMatcher: nil,
		},
		{
			Input:        []reflect.Value{reflect.ValueOf(nil), reflect.ValueOf(wrongArgumentTypeError)},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []reflect.Value{reflect.ValueOf([]interface{}{[]string{"a", "b", "c"}, 3.8, true})},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []reflect.Value{reflect.ValueOf([]interface{}{[]string{"a", "b", "c"}, 3.8, true}), reflect.ValueOf(nil), reflect.ValueOf(nil)},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
	}

	for i, testCase := range testCases {
		output, err := ValuesToArgs(testCase.Input)
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
