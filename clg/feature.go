package clg

import (
	"github.com/xh3b4sd/anna/feature-set"
)

func (i *clgIndex) AddPositionFeature(args ...interface{}) ([]interface{}, error) {
	f, err := ArgToFeature(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	position, err := ArgToFloat64Slice(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	err = f.AddPosition(position)
	if err != nil {
		return nil, maskAny(err)
	}

	return []interface{}{}, nil
}

func (i *clgIndex) GetCountFeature(args ...interface{}) ([]interface{}, error) {
	f, err := ArgToFeature(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	count := f.GetCount()

	return []interface{}{count}, nil
}

func (i *clgIndex) GetDistributionFeature(args ...interface{}) ([]interface{}, error) {
	f, err := ArgToFeature(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	distribution := f.GetDistribution()

	return []interface{}{distribution}, nil
}

func (i *clgIndex) GetNewFeature(args ...interface{}) ([]interface{}, error) {
	var err error
	newConfig := featureset.DefaultFeatureConfig()

	newConfig.Positions, err = ArgToFloat64SliceSlice(args, 0, newConfig.Positions)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.Sequence, err = ArgToString(args, 1, newConfig.Sequence)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	newFeature, err := featureset.NewFeature(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return []interface{}{newFeature}, nil
}

func (i *clgIndex) GetPositionsFeature(args ...interface{}) ([]interface{}, error) {
	f, err := ArgToFeature(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	positions := f.GetPositions()

	return []interface{}{positions}, nil
}

func (i *clgIndex) GetSequenceFeature(args ...interface{}) ([]interface{}, error) {
	f, err := ArgToFeature(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	sequence := f.GetSequence()

	return []interface{}{sequence}, nil
}
