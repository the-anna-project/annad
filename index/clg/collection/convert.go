package collection

import (
	"strconv"
)

func (c *collection) BoolStringConvert(args ...interface{}) ([]interface{}, error) {
	t, err := ArgToBool(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newString := strconv.FormatBool(t)

	return []interface{}{newString}, nil
}

func (c *collection) Float64StringConvert(args ...interface{}) ([]interface{}, error) {
	f, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newString := strconv.FormatFloat(f, 'f', -1, 64)

	return []interface{}{newString}, nil
}

func (c *collection) IntStringConvert(args ...interface{}) ([]interface{}, error) {
	n, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newString := strconv.Itoa(n)

	return []interface{}{newString}, nil
}

func (c *collection) StringBoolConvert(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newBool, err := strconv.ParseBool(s)
	if err != nil {
		return nil, maskAnyf(cannotConvertError, err.Error())
	}

	return []interface{}{newBool}, nil
}

func (c *collection) StringFloat64Convert(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newFloat64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, maskAnyf(cannotConvertError, err.Error())
	}

	return []interface{}{newFloat64}, nil
}

func (c *collection) StringIntConvert(args ...interface{}) ([]interface{}, error) {
	s, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newInt, err := strconv.Atoi(s)
	if err != nil {
		return nil, maskAnyf(cannotConvertError, err.Error())
	}

	return []interface{}{newInt}, nil
}
