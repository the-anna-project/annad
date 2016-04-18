package clg

import (
	"sort"
	"strconv"
	"strings"
)

func (i *clgIndex) ContainsIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	num, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var contains bool
	for _, i := range is {
		if i == num {
			contains = true
			break
		}
	}

	return []interface{}{contains}, nil
}

func (i *clgIndex) CountIntSlice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) EqualMatcherIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	num, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var m []int
	var u []int
	for _, i := range is {
		if i == num {
			m = append(m, i)
		} else {
			u = append(u, i)
		}
	}

	return []interface{}{m, u}, nil
}

func (i *clgIndex) GlobMatcherIntSlice(args ...interface{}) ([]interface{}, error) {
	is, err := ArgToIntSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	num, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var m []int
	var u []int
	for _, i := range is {
		if strings.Contains(strconv.Itoa(num), strconv.Itoa(i)) {
			m = append(m, i)
		} else {
			u = append(u, i)
		}
	}

	return []interface{}{m, u}, nil
}

func (i *clgIndex) IndexIntSlice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) IsUniqueIntSlice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) JoinIntSlice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) MaxIntSlice(args ...interface{}) ([]interface{}, error) {
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

	m := maxInt(is)

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

func (i *clgIndex) MinIntSlice(args ...interface{}) ([]interface{}, error) {
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

	m := minInt(is)

	return []interface{}{m}, nil
}

func (i *clgIndex) SortIntSlice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) SwapLeftIntSlice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) SwapRightIntSlice(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) UniqueIntSlice(args ...interface{}) ([]interface{}, error) {
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
