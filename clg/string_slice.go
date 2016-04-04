package clg

import (
	"sort"
	"strings"
)

func (i *index) ContainsStringSlice(args ...interface{}) ([]interface{}, error) {
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

	var contains bool
	for _, s := range ss {
		if s == str {
			contains = true
			break
		}
	}

	return []interface{}{contains}, nil
}

func (i *index) JoinStringSlice(args ...interface{}) ([]interface{}, error) {
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

func (i *index) SortStringSlice(args ...interface{}) ([]interface{}, error) {
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

func (i *index) SwapLeftStringSlice(args ...interface{}) ([]interface{}, error) {
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

func (i *index) SwapRightStringSlice(args ...interface{}) ([]interface{}, error) {
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
