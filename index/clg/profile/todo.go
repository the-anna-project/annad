package profile

import (
	"github.com/xh3b4sd/anna/index/clg/collection/distribution"
	"github.com/xh3b4sd/anna/index/clg/collection/feature-set"
)

// TODO
func createArgumentValues() ([]interface{}, error) {
	newDistributionConfig := distribution.DefaultConfig()
	newDistributionConfig.Name = "generated for argument list"
	newDistributionConfig.Vectors = [][]float64{{0, 0, 0}, {3, 5, 7}, {67, 84, 31}, {11, 1, 22}, {9, 9, 9}}
	newDistribution, err := distribution.NewDistribution(newDistributionConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	newFeatureSetConfig := featureset.DefaultConfig()
	newFeatureSetConfig.Sequences = []string{"some", "sequence", "test", "foo", "whatever"}
	newFeatureSet, err := featureset.NewFeatureSet(newFeatureSetConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	newValues := []interface{}{
		// nil
		nil,

		// bool
		true,
		false,

		// float64
		float64(-33),
		float64(-5),
		float64(0),
		float64(5),
		float64(33),

		// int
		int(-33),
		int(-5),
		int(0),
		int(5),
		int(33),

		// interface{}
		interface{}(nil),
		interface{}(true),
		interface{}(false),
		interface{}(5.8),
		interface{}(33.77),
		interface{}(-33),
		interface{}(-5),
		interface{}(0),
		interface{}(""),
		interface{}("a"),
		interface{}("foo"),

		// string
		"5",
		"5.8",
		"33",
		"33.77",
		"",
		"a",
		"foo",

		// []float64
		[]float64{-5, -33},
		[]float64{0, -5},
		[]float64{-5.8, -33.77},
		[]float64{-5.8},
		[]float64{},
		[]float64{0},
		[]float64{5.8},
		[]float64{5.8, 33.77},
		[]float64{0, 5},
		[]float64{5, 33},

		// []int
		[]int{-5, -33},
		[]int{0, -5},
		[]int{},
		[]int{0},
		[]int{0, 5},
		[]int{5, 33},

		// []interface{}
		[]interface{}{},
		[]interface{}{nil},
		[]interface{}{true, "foo", 5, 5.8},
		[]interface{}{false},
		[]interface{}{5.8},
		[]interface{}{33.77},
		[]interface{}{-33},
		[]interface{}{-5},
		[]interface{}{0, 0, 0, 0},
		[]interface{}{""},
		[]interface{}{"a", true, "b", false},
		[]interface{}{"foo"},
		[]interface{}{[]int{1, 2, 3}},
		[]interface{}{[]string{"", "", ""}},
		[]interface{}{[]string{"a", "b", "c"}},
		[]interface{}{[]int{1, 2, 3}, []float64{5.8, 33.77, -5.8, -33.7}},

		// []string
		[]string{},
		[]string{""},
		[]string{"foo"},
		[]string{"5", "5.8"},
		[]string{"a", "b", "c"},

		// [][]float64
		[][]float64{{-5, -33}, {5.8, 33.77}},
		[][]float64{{0, -5}},
		[][]float64{{-5.8, -33.77}},
		[][]float64{{-5.8}},
		[][]float64{{}},
		[][]float64{{0}},
		[][]float64{{5.8}, {}},
		[][]float64{{5.8, 33.77}},
		[][]float64{{0, 5}, {33.77}},
		[][]float64{{5, 33}},

		// [][]int
		[][]int{{-5, -33}, {5, 33}},
		[][]int{{0, -5}},
		[][]int{{}},
		[][]int{{0}, {}},
		[][]int{{0, 5}},
		[][]int{{5, 33}, {33}},

		// [][]string
		[][]string{{}},
		[][]string{{""}, {}},
		[][]string{{"foo"}},
		[][]string{{"5", "5.8"}, {"b", "c"}},
		[][]string{{"a", "b", "c"}, {"a"}},

		// spec.Distribution
		newDistribution,

		// spec.FeatureSet
		newFeatureSet,
	}

	return newValues, nil
}
