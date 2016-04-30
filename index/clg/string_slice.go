package clg

import (
	"sort"
	"strings"
)

func (i *clgIndex) AppendStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	s, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	ss = append(ss, s)

	return []interface{}{ss}, nil
}

func containsString(ss []string, s string) bool {
	for _, i := range ss {
		if i == s {
			return true
		}
	}

	return false
}

func (i *clgIndex) ContainsStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	s, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	contains := containsString(ss, s)

	return []interface{}{contains}, nil
}

func (i *clgIndex) CountCharacterStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newMap := map[string]int{}
	for _, s := range ss {
		newMap[s]++
	}

	return []interface{}{newMap}, nil
}

func (i *clgIndex) CountStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	count := len(ss)

	return []interface{}{count}, nil
}

func differenceString(ss1, ss2 []string) []string {
	var newDifference []string

	for _, s1 := range ss1 {
		if !containsString(ss2, s1) {
			newDifference = append(newDifference, s1)
		}
	}

	return newDifference
}

func (i *clgIndex) DifferenceStringSlice(args ...interface{}) ([]interface{}, error) {
	ss1, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	ss2, err := ArgToStringSlice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(ss1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss1))
	}
	if len(ss2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss2))
	}

	newDifference := differenceString(ss1, ss2)

	return []interface{}{newDifference}, nil
}

func (i *clgIndex) EqualMatcherStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	str, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var m []string
	var u []string
	for _, s := range ss {
		if s == str {
			m = append(m, s)
		} else {
			u = append(u, s)
		}
	}

	return []interface{}{m, u}, nil
}

func (i *clgIndex) GlobMatcherStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	str, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var m []string
	var u []string
	for _, s := range ss {
		if strings.Contains(str, s) {
			m = append(m, s)
		} else {
			u = append(u, s)
		}
	}

	return []interface{}{m, u}, nil
}

func (i *clgIndex) IndexStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
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
	if len(ss) <= index {
		return nil, maskAny(indexOutOfRangeError)
	}

	newString := ss[index]

	return []interface{}{newString}, nil
}

func intersectionString(ss1, ss2 []string) []string {
	var newIntersection []string

	for _, s1 := range ss1 {
		for _, s2 := range ss2 {
			if s2 == s1 {
				newIntersection = append(newIntersection, s2)
				continue
			}
		}
	}

	return newIntersection
}

func (i *clgIndex) IntersectionStringSlice(args ...interface{}) ([]interface{}, error) {
	ss1, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	ss2, err := ArgToStringSlice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(ss1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss1))
	}
	if len(ss2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss2))
	}

	newIntersection := intersectionString(ss1, ss2)

	return []interface{}{newIntersection}, nil
}

func (i *clgIndex) IsUniqueStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(ss) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 2 got %d", len(args))
	}

	unique := true
	seen := map[string]struct{}{}
	for _, s := range ss {
		if _, ok := seen[s]; ok {
			unique = false
			break
		}
		seen[s] = struct{}{}
	}

	return []interface{}{unique}, nil
}

func (i *clgIndex) JoinStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	sep, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(ss) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss))
	}

	newString := strings.Join(ss, sep)

	return []interface{}{newString}, nil
}

func (i *clgIndex) NewStringSlice(args ...interface{}) ([]interface{}, error) {
	if len(args) > 0 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 0 got %d", len(args))
	}

	var ss []string

	return []interface{}{ss}, nil
}

func (i *clgIndex) ReverseStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(ss) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss))
	}

	var newStringSlice []string
	for i := len(ss) - 1; i >= 0; i-- {
		newStringSlice = append(newStringSlice, ss[i])
	}

	return []interface{}{newStringSlice}, nil
}

func stem(list []string) string {
	if len(list) == 0 {
		return ""
	}

	ri := 0
	li := 0
	ll := len(list)
	ref := list[0]
	rm := ""

	for {
		if ri > len(ref) {
			break
		}
		if ri > len(list[li]) {
			break
		}

		rm = ref[:ri]
		lm := list[li][:ri]

		if rm == lm {
			li++
			if li == ll {
				li = 0
				ri++
			}

			continue
		} else {
			break
		}
	}

	rm = ref[:ri-1]
	return rm
}

func (i *clgIndex) StemStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(ss) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss))
	}

	newString := stem(ss)

	return []interface{}{newString}, nil
}

func (i *clgIndex) SortStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(ss) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss))
	}

	newStringSlice := ss
	sort.Strings(newStringSlice)

	return []interface{}{newStringSlice}, nil
}

func (i *clgIndex) SwapLeftStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(ss) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss))
	}

	newStringSlice := append(ss[1:], ss[0])

	return []interface{}{newStringSlice}, nil
}

func (i *clgIndex) SwapRightStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(ss) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss))
	}

	newStringSlice := append([]string{ss[len(ss)-1]}, ss[:len(ss)-1]...)

	return []interface{}{newStringSlice}, nil
}

func symmetricDifferenceString(ss1, ss2 []string) []string {
	var newSymmetricDifference []string

	for _, s1 := range ss1 {
		if !containsString(ss2, s1) {
			newSymmetricDifference = append(newSymmetricDifference, s1)
		}
	}
	for _, s2 := range ss2 {
		if !containsString(ss1, s2) {
			newSymmetricDifference = append(newSymmetricDifference, s2)
		}
	}

	return newSymmetricDifference
}

func (i *clgIndex) SymmetricDifferenceStringSlice(args ...interface{}) ([]interface{}, error) {
	ss1, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	ss2, err := ArgToStringSlice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(ss1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss1))
	}
	if len(ss2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss2))
	}

	newSymmetricDifference := symmetricDifferenceString(ss1, ss2)

	return []interface{}{newSymmetricDifference}, nil
}

func unionString(ss1, ss2 []string) []string {
	var newUnion []string

	for _, s := range ss1 {
		newUnion = append(newUnion, s)
	}
	for _, s := range ss2 {
		newUnion = append(newUnion, s)
	}

	return newUnion
}

func (i *clgIndex) UnionStringSlice(args ...interface{}) ([]interface{}, error) {
	ss1, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	ss2, err := ArgToStringSlice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(ss1) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss1))
	}
	if len(ss2) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss2))
	}

	newUnion := unionString(ss1, ss2)

	return []interface{}{newUnion}, nil
}

func (i *clgIndex) UniqueStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	if len(ss) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss))
	}

	seen := map[string]struct{}{}
	var newStringSlice []string
	for _, s := range ss {
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		newStringSlice = append(newStringSlice, s)
	}

	return []interface{}{newStringSlice}, nil
}
