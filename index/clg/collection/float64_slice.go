package collection

import (
	"sort"
	"strconv"
	"strings"
)

func (c *collection) AppendFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	f, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	fs = append(fs, f)

	return []interface{}{fs}, nil
}

func containsFloat64(fs []float64, f float64) bool {
	for _, i := range fs {
		if i == f {
			return true
		}
	}

	return false
}

func (c *collection) ContainsFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	f, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	contains := containsFloat64(fs, f)

	return []interface{}{contains}, nil
}

func (c *collection) CountFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	count := len(fs)

	return []interface{}{count}, nil
}

func differenceFloat64(fs1, fs2 []float64) []float64 {
	var newDifference []float64

	for _, f1 := range fs1 {
		if !containsFloat64(fs2, f1) {
			newDifference = append(newDifference, f1)
		}
	}

	return newDifference
}

func (c *collection) DifferenceFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs1, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	fs2, err := ArgToFloat64Slice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(fs1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs1))
	}
	if len(fs2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs2))
	}

	newDifference := differenceFloat64(fs1, fs2)

	return []interface{}{newDifference}, nil
}

func (c *collection) EqualMatcherFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	f, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var m []float64
	var u []float64
	for _, i := range fs {
		if i == f {
			m = append(m, i)
		} else {
			u = append(u, i)
		}
	}

	return []interface{}{m, u}, nil
}

func (c *collection) GlobMatcherFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	f, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var m []float64
	var u []float64
	for _, i := range fs {
		if strings.Contains(strconv.FormatFloat(i, 'f', -1, 64), strconv.FormatFloat(f, 'f', -1, 64)) {
			m = append(m, i)
		} else {
			u = append(u, i)
		}
	}

	return []interface{}{m, u}, nil
}

func (c *collection) IndexFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	index, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(fs) <= index {
		return nil, maskAny(indexOutOfRangeError)
	}

	newFloat64 := fs[index]

	return []interface{}{newFloat64}, nil
}

func intersectionFloat64(fs1, fs2 []float64) []float64 {
	var newIntersection []float64

	for _, f1 := range fs1 {
		for _, f2 := range fs2 {
			if f2 == f1 {
				newIntersection = append(newIntersection, f2)
				continue
			}
		}
	}

	return newIntersection
}

func (c *collection) IntersectionFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs1, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	fs2, err := ArgToFloat64Slice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(fs1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs1))
	}
	if len(fs2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs2))
	}

	newIntersection := intersectionFloat64(fs1, fs2)

	return []interface{}{newIntersection}, nil
}

func (c *collection) IsUniqueFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(fs) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 2 got %d", len(args))
	}

	unique := true
	seen := map[float64]struct{}{}
	for _, f := range fs {
		if _, ok := seen[f]; ok {
			unique = false
			break
		}
		seen[f] = struct{}{}
	}

	return []interface{}{unique}, nil
}

func maxFloat64(list []float64) float64 {
	if len(list) == 0 {
		return 0
	}

	max := list[0]

	for _, i := range list {
		if i > max {
			max = i
		}
	}

	return max
}

func (c *collection) MaxFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(fs) < 1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 1 got %d", len(fs))
	}

	m := maxFloat64(fs)

	return []interface{}{m}, nil
}

func meanFloat64(list []float64) float64 {
	l := len(list)
	if l == 0 {
		return 0
	}

	var sum float64
	for _, i := range list {
		sum += i
	}

	mean := sum / float64(l)

	return mean
}

func (c *collection) MeanFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(fs) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs))
	}

	m := meanFloat64(fs)

	return []interface{}{m}, nil
}

func medianFloat64(list []float64) float64 {
	l := len(list)
	if l == 0 {
		return 0
	}

	// The median can only be calculated on a sorted list of numbers. Thus we
	// create a copy first to keep the input as it is.
	c := list
	sort.Float64s(c)

	var median float64
	if l%2 == 0 {
		// In case the amount of numbers is even, the median consists of the mean
		// (average) of the two middle numbers.
		median = (c[l/2-1] + c[l/2]) / 2
	} else {
		// In case the amount of numbers is odd, the median is the middle number.
		median = float64(c[l/2])
	}

	return median
}

func (c *collection) MedianFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(fs) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs))
	}

	m := medianFloat64(fs)

	return []interface{}{m}, nil
}

func minFloat64(list []float64) float64 {
	if len(list) == 0 {
		return 0
	}

	min := list[0]

	for _, i := range list {
		if i < min {
			min = i
		}
	}

	return min
}

func (c *collection) MinFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(fs) < 1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 1 got %d", len(fs))
	}

	m := minFloat64(fs)

	return []interface{}{m}, nil
}

func modeFloat64(list []float64) []float64 {
	if len(list) == 0 {
		return nil
	}

	// Collect the counts of all items and also find the maximum number of
	// occurrences.
	max := 1
	counts := map[float64]int{}
	for _, item := range list {
		if _, ok := counts[item]; !ok {
			counts[item] = 1
		} else {
			counts[item]++

			count := counts[item]
			if count > max {
				max = count
			}
		}
	}

	// Collect the most occurred items and sort the result.
	var mode []float64
	for item, count := range counts {
		if count == max {
			mode = append(mode, item)
		}
	}
	sort.Float64s(mode)

	return mode
}

func (c *collection) ModeFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(fs) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs))
	}

	m := modeFloat64(fs)

	return []interface{}{m}, nil
}

func (c *collection) NewFloat64Slice(args ...interface{}) ([]interface{}, error) {
	if len(args) > 0 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 0 got %d", len(args))
	}

	var fs []float64

	return []interface{}{fs}, nil
}

func percentilesFloat64(fs, ps []float64) ([]float64, error) {
	// Validate the input.
	l := float64(len(fs))
	if l == 0 || len(ps) == 0 {
		return nil, nil
	}
	for _, p := range ps {
		if !betweenFloat64(p, 0, 100) {
			return nil, maskAnyf(indexOutOfRangeError, "percentiles must be between 0 and 100")
		}
	}

	// The percentiles can only be calculated on a sorted list of numbers. Thus
	// we create a copy first to keep the input as it is.
	c := fs
	sort.Float64s(c)

	var percentiles []float64

	for _, p := range ps {
		var percentile float64

		index := (p / 100) * l
		i := int(index)

		if index == float64(i) {
			percentile = c[i-1]
		} else {
			pd := index - float64(i)
			id := c[i] - c[i-1]
			percentile = index + (pd * id / 1)
		}

		percentiles = append(percentiles, percentile)
	}

	return percentiles, nil
}

func (c *collection) PercentilesFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	ps, err := ArgToFloat64Slice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(fs) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs))
	}
	if len(ps) < 1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 1 got %d", len(ps))
	}

	newPercentiles, err := percentilesFloat64(fs, ps)
	if err != nil {
		return nil, maskAny(err)
	}

	return []interface{}{newPercentiles}, nil
}

func (c *collection) ReverseFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(fs) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs))
	}

	var newFloat64Slice []float64
	for i := len(fs) - 1; i >= 0; i-- {
		newFloat64Slice = append(newFloat64Slice, fs[i])
	}

	return []interface{}{newFloat64Slice}, nil
}

func (c *collection) SortFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(fs) < 1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 1 got %d", len(fs))
	}

	sort.Float64s(fs)

	return []interface{}{fs}, nil
}

func (c *collection) SwapLeftFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(fs) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs))
	}

	newFloat64Slice := append(fs[1:], fs[0])

	return []interface{}{newFloat64Slice}, nil
}

func (c *collection) SwapRightFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(fs) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs))
	}

	newFloat64Slice := append([]float64{fs[len(fs)-1]}, fs[:len(fs)-1]...)

	return []interface{}{newFloat64Slice}, nil
}

func symmetricDifferenceFloat64(fs1, fs2 []float64) []float64 {
	var newSymmetricDifference []float64

	for _, f1 := range fs1 {
		if !containsFloat64(fs2, f1) {
			newSymmetricDifference = append(newSymmetricDifference, f1)
		}
	}
	for _, f2 := range fs2 {
		if !containsFloat64(fs1, f2) {
			newSymmetricDifference = append(newSymmetricDifference, f2)
		}
	}

	return newSymmetricDifference
}

func (c *collection) SymmetricDifferenceFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs1, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	fs2, err := ArgToFloat64Slice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(fs1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs1))
	}
	if len(fs2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs2))
	}

	newSymmetricDifference := symmetricDifferenceFloat64(fs1, fs2)

	return []interface{}{newSymmetricDifference}, nil
}

func unionFloat64(fs1, fs2 []float64) []float64 {
	var newUnion []float64

	for _, f := range fs1 {
		newUnion = append(newUnion, f)
	}
	for _, f := range fs2 {
		newUnion = append(newUnion, f)
	}

	return newUnion
}

func (c *collection) UnionFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs1, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	fs2, err := ArgToFloat64Slice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(fs1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs1))
	}
	if len(fs2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs2))
	}

	newUnion := unionFloat64(fs1, fs2)

	return []interface{}{newUnion}, nil
}

func (c *collection) UniqueFloat64Slice(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFloat64Slice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(fs) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(fs))
	}

	seen := map[float64]struct{}{}
	var newFloat64Slice []float64
	for _, f := range fs {
		if _, ok := seen[f]; ok {
			continue
		}
		seen[f] = struct{}{}
		newFloat64Slice = append(newFloat64Slice, f)
	}

	return []interface{}{newFloat64Slice}, nil
}
