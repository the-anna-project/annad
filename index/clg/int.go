package clg

import (
	"math"
)

func (c *clgCollection) DivideInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 / i2

	return []interface{}{s}, nil
}

func (c *clgCollection) GreaterThanInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	greater := i1 > i2

	return []interface{}{greater}, nil
}

func (c *clgCollection) LesserThanInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	lesser := i1 < i2

	return []interface{}{lesser}, nil
}

func (c *clgCollection) MultiplyInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 * i2

	return []interface{}{s}, nil
}

func (c *clgCollection) PowInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := math.Pow(float64(i1), float64(i2))

	return []interface{}{s}, nil
}

func (c *clgCollection) SqrtInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	s := math.Sqrt(float64(i1))

	return []interface{}{s}, nil
}

func (c *clgCollection) SubtractInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 - i2

	return []interface{}{s}, nil
}

func (c *clgCollection) SumInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 + i2

	return []interface{}{s}, nil
}
