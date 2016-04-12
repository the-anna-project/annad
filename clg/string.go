package clg

import (
	"strings"

	"github.com/xh3b4sd/anna/id"
)

func (i *clgIndex) ContainsString(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) LongerString(args ...interface{}) ([]interface{}, error) {
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

	longer := len(s1) > len(s2)

	return []interface{}{longer}, nil
}

func (i *clgIndex) NewIDString(args ...interface{}) ([]interface{}, error) {
	if len(args) > 0 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 0 got %d", len(args))
	}

	newID := string(id.NewObjectID(id.Hex128))

	return []interface{}{newID}, nil
}

func (i *clgIndex) RepeatString(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) ShorterString(args ...interface{}) ([]interface{}, error) {
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

	shorter := len(s1) < len(s2)

	return []interface{}{shorter}, nil
}

func (i *clgIndex) SplitString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
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

	newStringSlice := strings.Split(s, sep)

	return []interface{}{newStringSlice}, nil
}

func (i *clgIndex) ToLowerString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newString := strings.ToLower(s)

	return []interface{}{newString}, nil
}

func (i *clgIndex) ToUpperString(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newString := strings.ToUpper(s)

	return []interface{}{newString}, nil
}
