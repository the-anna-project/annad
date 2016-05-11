package collection

import (
	"sort"
	"strconv"
	"strings"
)

func (c *collection) AppendIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	n, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	is = append(is, n)

	return []interface{}{is}, nil
}

func containsInt(is []int, n int) bool {
	for _, i := range is {
		if i == n {
			return true
		}
	}

	return false
}

func (c *collection) ContainsIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	n, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	contains := containsInt(is, n)

	return []interface{}{contains}, nil
}

func (c *collection) CountIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	count := len(is)

	return []interface{}{count}, nil
}

func differenceInt(is1, is2 []int) []int {
	var newDifference []int

	for _, i1 := range is1 {
		if !containsInt(is2, i1) {
			newDifference = append(newDifference, i1)
		}
	}

	return newDifference
}

func (c *collection) DifferenceIntSlice(args ...interface{}) ([]interface{}, error) {
	is1, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	is2, err := ArgToIntSlice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(is1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is1))
	}
	if len(is2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is2))
	}

	newDifference := differenceInt(is1, is2)

	return []interface{}{newDifference}, nil
}

func (c *collection) EqualMatcherIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	n, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var m []int
	var u []int
	for _, i := range is {
		if i == n {
			m = append(m, i)
		} else {
			u = append(u, i)
		}
	}

	return []interface{}{m, u}, nil
}

func (c *collection) GlobMatcherIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	n, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var m []int
	var u []int
	for _, i := range is {
		if strings.Contains(strconv.Itoa(n), strconv.Itoa(i)) {
			m = append(m, i)
		} else {
			u = append(u, i)
		}
	}

	return []interface{}{m, u}, nil
}

func (c *collection) IndexIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
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
	if len(is) <= index {
		return nil, maskAny(indexOutOfRangeError)
	}

	newInt := is[index]

	return []interface{}{newInt}, nil
}

func intersectionInt(is1, is2 []int) []int {
	var newIntersection []int

	for _, i1 := range is1 {
		for _, i2 := range is2 {
			if i2 == i1 {
				newIntersection = append(newIntersection, i2)
				continue
			}
		}
	}

	return newIntersection
}

func (c *collection) IntersectionIntSlice(args ...interface{}) ([]interface{}, error) {
	is1, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	is2, err := ArgToIntSlice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(is1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is1))
	}
	if len(is2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is2))
	}

	newIntersection := intersectionInt(is1, is2)

	return []interface{}{newIntersection}, nil
}

func (c *collection) IsUniqueIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 2 got %d", len(args))
	}

	unique := true
	seen := map[int]struct{}{}
	for _, n := range is {
		if _, ok := seen[n]; ok {
			unique = false
			break
		}
		seen[n] = struct{}{}
	}

	return []interface{}{unique}, nil
}

func (c *collection) JoinIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is))
	}

	var newString string
	for _, i := range is {
		newString += strconv.Itoa(i)
	}
	newInt, _ := strconv.Atoi(newString)

	return []interface{}{newInt}, nil
}

func maxInt(list []int) int {
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

func (c *collection) MaxIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is))
	}

	m := maxInt(is)

	return []interface{}{m}, nil
}

func meanInt(list []int) float64 {
	l := len(list)
	if l == 0 {
		return 0
	}

	var sum int
	for _, i := range list {
		sum += i
	}

	mean := float64(sum) / float64(l)

	return mean
}

func (c *collection) MeanIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is))
	}

	m := meanInt(is)

	return []interface{}{m}, nil
}

func medianInt(list []int) float64 {
	l := len(list)
	if l == 0 {
		return 0
	}

	// The median can only be calculated on a sorted list of numbers. Thus we
	// create a copy first to keep the input as it is.
	c := list
	sort.Ints(c)

	var median float64
	if l%2 == 0 {
		// In case the amount of numbers is even, the median consists of the mean
		// (average) of the two middle numbers.
		median = float64(c[l/2-1]+c[l/2]) / 2
	} else {
		// In case the amount of numbers is odd, the median is the middle number.
		median = float64(c[l/2])
	}

	return median
}

func (c *collection) MedianIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is))
	}

	m := medianInt(is)

	return []interface{}{m}, nil
}

func minInt(list []int) int {
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

func (c *collection) MinIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is))
	}

	m := minInt(is)

	return []interface{}{m}, nil
}

func modeInt(list []int) []int {
	if len(list) == 0 {
		return nil
	}

	// Collect the counts of all items and also find the maximum number of
	// occurrences.
	max := 1
	counts := map[int]int{}
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
	var mode []int
	for item, count := range counts {
		if count == max {
			mode = append(mode, item)
		}
	}
	sort.Ints(mode)

	return mode
}

func (c *collection) ModeIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is))
	}

	m := modeInt(is)

	return []interface{}{m}, nil
}

func (c *collection) NewIntSlice(args ...interface{}) ([]interface{}, error) {
	if len(args) > 0 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 0 got %d", len(args))
	}

	var is []int

	return []interface{}{is}, nil
}

func percentilesInt(is []int, ps []float64) ([]float64, error) {
	// Validate the input.
	l := float64(len(is))
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
	c := is
	sort.Ints(c)

	var percentiles []float64

	for _, p := range ps {
		var percentile float64

		index := (p / 100) * l
		i := int(index)

		if index == float64(i) {
			percentile = float64(c[i-1])
		} else {
			pd := index - float64(i)
			id := float64(c[i] - c[i-1])
			percentile = index + (pd * id / 1)
		}

		percentiles = append(percentiles, percentile)
	}

	return percentiles, nil
}

func (c *collection) PercentilesIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
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
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is))
	}
	if len(ps) < 1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 1 got %d", len(ps))
	}

	newPercentiles, err := percentilesInt(is, ps)
	if err != nil {
		return nil, maskAny(err)
	}

	return []interface{}{newPercentiles}, nil
}

func (c *collection) ReverseIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is))
	}

	var newIntSlice []int
	for i := len(is) - 1; i >= 0; i-- {
		newIntSlice = append(newIntSlice, is[i])
	}

	return []interface{}{newIntSlice}, nil
}

func (c *collection) SortIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 1 got %d", len(is))
	}

	sort.Ints(is)

	return []interface{}{is}, nil
}

func (c *collection) SwapLeftIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is))
	}

	newIntSlice := append(is[1:], is[0])

	return []interface{}{newIntSlice}, nil
}

func (c *collection) SwapRightIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is))
	}

	newIntSlice := append([]int{is[len(is)-1]}, is[:len(is)-1]...)

	return []interface{}{newIntSlice}, nil
}

func symmetricDifferenceInt(is1, is2 []int) []int {
	var newSymmetricDifference []int

	for _, i1 := range is1 {
		if !containsInt(is2, i1) {
			newSymmetricDifference = append(newSymmetricDifference, i1)
		}
	}
	for _, i2 := range is2 {
		if !containsInt(is1, i2) {
			newSymmetricDifference = append(newSymmetricDifference, i2)
		}
	}

	return newSymmetricDifference
}

func (c *collection) SymmetricDifferenceIntSlice(args ...interface{}) ([]interface{}, error) {
	is1, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	is2, err := ArgToIntSlice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(is1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is1))
	}
	if len(is2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is2))
	}

	newSymmetricDifference := symmetricDifferenceInt(is1, is2)

	return []interface{}{newSymmetricDifference}, nil
}

func unionInt(is1, is2 []int) []int {
	var newUnion []int

	for _, i := range is1 {
		newUnion = append(newUnion, i)
	}
	for _, i := range is2 {
		newUnion = append(newUnion, i)
	}

	return newUnion
}

func (c *collection) UnionIntSlice(args ...interface{}) ([]interface{}, error) {
	is1, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	is2, err := ArgToIntSlice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(is1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is1))
	}
	if len(is2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is2))
	}

	newUnion := unionInt(is1, is2)

	return []interface{}{newUnion}, nil
}

func (c *collection) UniqueIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(is) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(is))
	}

	seen := map[int]struct{}{}
	var newIntSlice []int
	for _, i := range is {
		if _, ok := seen[i]; ok {
			continue
		}
		seen[i] = struct{}{}
		newIntSlice = append(newIntSlice, i)
	}

	return []interface{}{newIntSlice}, nil
}
