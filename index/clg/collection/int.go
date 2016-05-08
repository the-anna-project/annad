package collection

import (
	"math"
)

func betweenInt(i, min, max int) bool {
	if i < min {
		return false
	}
	if i > max {
		return false
	}
	return true
}

func (c *collection) BetweenInt(args ...interface{}) ([]interface{}, error) {
	num, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	min, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	max, err := ArgToInt(args, 2)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 3 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 3 got %d", len(args))
	}

	isBetween := betweenInt(num, min, max)

	return []interface{}{isBetween}, nil
}

func (c *collection) DivideInt(args ...interface{}) ([]interface{}, error) {
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

func (c *collection) GreaterThanInt(args ...interface{}) ([]interface{}, error) {
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

func (c *collection) LesserThanInt(args ...interface{}) ([]interface{}, error) {
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

func (c *collection) MultiplyInt(args ...interface{}) ([]interface{}, error) {
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

func (c *collection) PowInt(args ...interface{}) ([]interface{}, error) {
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

func (c *collection) SqrtInt(args ...interface{}) ([]interface{}, error) {
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

func (c *collection) SubtractInt(args ...interface{}) ([]interface{}, error) {
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

func (c *collection) SumInt(args ...interface{}) ([]interface{}, error) {
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
