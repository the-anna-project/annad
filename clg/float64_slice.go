package clg

import (
	"sort"
	"strconv"
	"strings"
)

func (i *clgIndex) AppendFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) ContainsFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

	var contains bool
	for _, i := range fs {
		if i == f {
			contains = true
			break
		}
	}

	return []interface{}{contains}, nil
}

func (i *clgIndex) CountFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) EqualMatcherFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) GlobMatcherFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) IndexFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) IsUniqueFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) MaxFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) MinFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) NewFloat64Slice(args ...interface{}) ([]interface{}, error) {
	if len(args) > 0 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 0 got %d", len(args))
	}

	var fs []float64

	return []interface{}{fs}, nil
}

func (i *clgIndex) ReverseFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) SortFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) SwapLeftFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) SwapRightFloat64Slice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) UniqueFloat64Slice(args ...interface{}) ([]interface{}, error) {
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
