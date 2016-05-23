package permutation

import (
	"github.com/xh3b4sd/anna/index/clg/collection/distribution"
	"github.com/xh3b4sd/anna/index/clg/collection/feature-set"
	"github.com/xh3b4sd/anna/spec"
)

func createIndizesWithDelta(list spec.PermutationList, delta int) ([]int, error) {
	// Initialize scope variables.
	base := len(list.GetValues())
	newIndizes := list.GetIndizes()
	operation := 0

	// Check for the initial situation. This is special and the only exception
	// within the algorithm.
	if len(newIndizes) == 0 {
		newIndizes = append(newIndizes, 0)
		operation++
	}

	for {
		// Check amount of operations in the first place. That way the initial
		// situation as well as all other operations are covered.
		operation++
		if operation > delta {
			break
		}

		// Increment the least significant digit. That is, the right most index.
		// This is the only incrementation being done on the index.
		i := len(newIndizes) - 1
		lsd := newIndizes[i]
		lsd++
		newIndizes[i] = lsd

		// Cap the indizes and shift them if necessary. In case the least
		// significant digit was incremented above the base capacity, indizes need
		// to be shifted from right to left. This is like counting a number.
		var msdShifted bool
		newIndizes, msdShifted = shiftIndizes(newIndizes, base)
		if msdShifted {
			// Make sure the permutation does not growth more than allowed.
			if len(newIndizes)+1 > list.GetMaxGrowth() {
				return nil, maskAny(maxGrowthReachedError)
			}

			// In case all the indizes where shifted, we zeroed out all indizes. Then
			// we need to prepend another zero as new most significant digit of the
			// index.
			newIndizes = prepend(newIndizes, 0, 0)
		}
	}

	return newIndizes, nil
}

func createMembers(list spec.PermutationList) []interface{} {
	newMembers := make([]interface{}, len(list.GetIndizes()))

	for i, index := range list.GetIndizes() {
		newMembers[i] = list.GetValues()[index]
	}

	return newMembers
}

func prepend(s []int, i, x int) []int {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = x

	return s
}

func shiftIndizes(indizes []int, base int) ([]int, bool) {
	var msdShifted bool
	var reminder int

	for i := len(indizes) - 1; i >= 0; i-- {
		if reminder > 0 {
			current := indizes[i] + reminder
			reminder = 0
			indizes[i] = current
		}

		if indizes[i] >= base {
			indizes[i] = 0
			reminder = 1

			if i == 0 {
				msdShifted = true
			}
		}
	}

	return indizes, msdShifted
}

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
