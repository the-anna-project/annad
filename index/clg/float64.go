package clg

import (
	"fmt"
	"math"
	"strconv"
)

func (c *clgCollection) DivideFloat64(args ...interface{}) ([]interface{}, error) {
	f1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	f2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := f1 / f2

	return []interface{}{s}, nil
}

func (c *clgCollection) GreaterThanFloat64(args ...interface{}) ([]interface{}, error) {
	f1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	f2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	greater := f1 > f2

	return []interface{}{greater}, nil
}

func (c *clgCollection) LesserThanFloat64(args ...interface{}) ([]interface{}, error) {
	f1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	f2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	lesser := f1 < f2

	return []interface{}{lesser}, nil
}

func (c *clgCollection) MultiplyFloat64(args ...interface{}) ([]interface{}, error) {
	f1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	f2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := f1 * f2

	return []interface{}{s}, nil
}

func (c *clgCollection) PowFloat64(args ...interface{}) ([]interface{}, error) {
	f1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	f2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := math.Pow(f1, f2)

	return []interface{}{s}, nil
}

func roundFloat64(f float64, p int) (float64, error) {
	newFloat, err := strconv.ParseFloat(fmt.Sprintf(fmt.Sprintf("%%.%df", p), f), 64)
	if err != nil {
		return 0, maskAnyf(cannotParseError, "%s", err)
	}

	return newFloat, nil
}

func (c *clgCollection) RoundFloat64(args ...interface{}) ([]interface{}, error) {
	f, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	p, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	newFloat64, err := roundFloat64(f, p)
	if err != nil {
		return nil, maskAny(err)
	}

	return []interface{}{newFloat64}, nil
}

func (c *clgCollection) SqrtFloat64(args ...interface{}) ([]interface{}, error) {
	f, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	s := math.Sqrt(f)

	return []interface{}{s}, nil
}

func (c *clgCollection) SubtractFloat64(args ...interface{}) ([]interface{}, error) {
	f1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	f2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := f1 - f2

	return []interface{}{s}, nil
}

func (c *clgCollection) SumFloat64(args ...interface{}) ([]interface{}, error) {
	f1, err := ArgToFloat64(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	f2, err := ArgToFloat64(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := f1 + f2

	return []interface{}{s}, nil
}
