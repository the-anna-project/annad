package collection

import (
	"github.com/xh3b4sd/anna/index/clg/collection/distribution"
)

func (c *collection) CalculateDistribution(args ...interface{}) ([]interface{}, error) {
	d, err := ArgToDistribution(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newCalculation := d.Calculate()

	return []interface{}{newCalculation}, nil
}

func (c *collection) DifferenceDistribution(args ...interface{}) ([]interface{}, error) {
	d1, err := ArgToDistribution(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	d2, err := ArgToDistribution(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	newDifference, err := d1.Difference(d2)
	if err != nil {
		return nil, maskAny(err)
	}

	return []interface{}{newDifference}, nil
}

func (c *collection) GetDimensionsDistribution(args ...interface{}) ([]interface{}, error) {
	d, err := ArgToDistribution(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newDimensions := d.GetDimensions()

	return []interface{}{newDimensions}, nil
}

func (c *collection) GetNameDistribution(args ...interface{}) ([]interface{}, error) {
	d, err := ArgToDistribution(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newName := d.Metadata()["name"]

	return []interface{}{newName}, nil
}

func (c *collection) GetNewDistribution(args ...interface{}) ([]interface{}, error) {
	var err error
	newConfig := distribution.DefaultConfig()

	newConfig.Name, err = ArgToString(args, 0, newConfig.Name)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.StaticChannels, err = ArgToFloat64Slice(args, 1, newConfig.StaticChannels)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.Vectors, err = ArgToFloat64SliceSlice(args, 2, newConfig.Vectors)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 3 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 3 got %d", len(args))
	}

	newDistribution, err := distribution.NewDistribution(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return []interface{}{newDistribution}, nil
}

func (c *collection) GetStaticChannelsDistribution(args ...interface{}) ([]interface{}, error) {
	d, err := ArgToDistribution(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newStaticChannels := d.GetStaticChannels()

	return []interface{}{newStaticChannels}, nil
}

func (c *collection) GetVectorsDistribution(args ...interface{}) ([]interface{}, error) {
	d, err := ArgToDistribution(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newVectors := d.GetVectors()

	return []interface{}{newVectors}, nil
}
