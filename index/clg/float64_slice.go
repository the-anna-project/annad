package clg

import (
	"sort"
	"strconv"
	"strings"
)

func (c *clgCollection) AppendFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) ContainsFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) CountFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) DifferenceFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) EqualMatcherFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) GlobMatcherFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) IndexFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) IntersectionFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) IsUniqueFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) MaxFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) MinFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) NewFloat64Slice(args ...interface{}) ([]interface{}, error) {
	if len(args) > 0 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 0 got %d", len(args))
	}

	var fs []float64

	return []interface{}{fs}, nil
}

func (c *clgCollection) ReverseFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) SortFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) SwapLeftFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) SwapRightFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) SymmetricDifferenceFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) UnionFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) UniqueFloat64Slice(args ...interface{}) ([]interface{}, error) {
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
