package clg

import (
	"github.com/xh3b4sd/anna/clg/distribution"
)

func (i *clgIndex) GetNewDistribution(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) GetStaticChannelsDistribution(args ...interface{}) ([]interface{}, error) {
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

func (i *clgIndex) GetVectorsDistribution(args ...interface{}) ([]interface{}, error) {
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
