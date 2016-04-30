package clg

import (
	"math"
)

func (i *clgIndex) DivideInt(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) GreaterThanInt(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) LesserThanInt(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) MultiplyInt(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) PowInt(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) SqrtInt(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) SubtractInt(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) SumInt(args ...interface{}) ([]interface{}, error) {
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
