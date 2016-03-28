package clg

import (
	"sort"
	"strings"
)

// ContainsStringSlice provides functionality to check if a string slice
// contains a certain member.
func ContainsStringSlice(args ...interface{}) ([]interface{}, error) {
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

// JoinStringSlice provides functionality of strings.Join.
func JoinStringSlice(args ...interface{}) ([]interface{}, error) {
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

// SortStringSlice provides functionality of sort.Strings.
func SortStringSlice(args ...interface{}) ([]interface{}, error) {
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

// SwapLeftStringSlice provides functionality to move the first member of a
// string slice to the left, that is, the end of the string slice.
func SwapLeftStringSlice(args ...interface{}) ([]interface{}, error) {
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

	newStringSlice := append(ss[1:len(ss)], ss[0])

	return []interface{}{newStringSlice}, nil
}

// SwapRightStringSlice provides functionality to move the last member of a
// string slice to the right, that is, the beginning of the string slice.
func SwapRightStringSlice(args ...interface{}) ([]interface{}, error) {
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
