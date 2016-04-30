package clg

import (
	"github.com/xh3b4sd/anna/index/clg/distribution"
)

func (c *clgCollection) CalculateDistribution(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) DifferenceDistribution(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) GetDimensionsDistribution(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) GetHashMapDistribution(args ...interface{}) ([]interface{}, error) {
	d, err := ArgToDistribution(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newHashMap := d.GetHashMap()

	return []interface{}{newHashMap}, nil
}

func (c *clgCollection) GetNameDistribution(args ...interface{}) ([]interface{}, error) {
	d, err := ArgToDistribution(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newName := d.GetName()

	return []interface{}{newName}, nil
}

func (c *clgCollection) GetNewDistribution(args ...interface{}) ([]interface{}, error) {
	var err error
	newConfig := distribution.DefaultConfig()

	newConfig.HashMap, err = ArgToStringStringMap(args, 0, newConfig.HashMap)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.Name, err = ArgToString(args, 1, newConfig.Name)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.StaticChannels, err = ArgToFloat64Slice(args, 2, newConfig.StaticChannels)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.Vectors, err = ArgToFloat64SliceSlice(args, 3, newConfig.Vectors)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 4 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 4 got %d", len(args))
	}

	newDistribution, err := distribution.NewDistribution(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return []interface{}{newDistribution}, nil
}

func (c *clgCollection) GetStaticChannelsDistribution(args ...interface{}) ([]interface{}, error) {
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

func (c *clgCollection) GetVectorsDistribution(args ...interface{}) ([]interface{}, error) {
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
