package clg

import (
	"strings"
)

func ContainsString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	substr, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}

	contains := strings.Contains(s, substr)

	return []interface{}{contains}, nil
}

func ContainsStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	str, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
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
