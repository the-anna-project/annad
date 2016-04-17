package clg

import (
	"github.com/xh3b4sd/anna/feature-set"
)

func (i *clgIndex) GetNewFeatureSet(args ...interface{}) ([]interface{}, error) {
	var err error
	newConfig := featureset.DefaultFeatureSetConfig()

	newConfig.Sequences, err = ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.MaxLength, err = ArgToInt(args, 1, newConfig.MaxLength)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.MinLength, err = ArgToInt(args, 2, newConfig.MinLength)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.MinCount, err = ArgToInt(args, 3, newConfig.MinCount)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.Separator, err = ArgToString(args, 4, newConfig.Separator)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 5 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 5 got %d", len(args))
	}

	newFeatureSet, err := featureset.NewFeatureSet(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return []interface{}{newFeatureSet}, nil
}

func (i *clgIndex) GetMaxLengthFeatureSet(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFeatureSet(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	maxLength := fs.GetMaxLength()

	return []interface{}{maxLength}, nil
}
