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
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	contains := strings.Contains(s, substr)

	return []interface{}{contains}, nil
}

func LongerString(args ...interface{}) ([]interface{}, error) {
	s1, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	s2, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var longer bool
	if len(s1) > len(s2) {
		longer = true
	}

	return []interface{}{longer}, nil
}

func RepeatString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	count, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	repeated := strings.Repeat(s, count)

	return []interface{}{repeated}, nil
}

func ShorterString(args ...interface{}) ([]interface{}, error) {
	s1, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	s2, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var longer bool
	if len(s1) < len(s2) {
		longer = true
	}

	return []interface{}{longer}, nil
}

// TODO SplitString
