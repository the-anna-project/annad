package featureset

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/index/clg/distribution"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeFeature represents the object type of the feature object. This
	// is used e.g. to register itself to the logger.
	ObjectTypeFeature spec.ObjectType = "feature"
)

// FeatureConfig represents the configuration used to create a new feature
// object.
type FeatureConfig struct {
	// Positions represents the index locations of a detected feature.
	Positions [][]float64

	// Sequence represents the input sequence being detected as feature. That
	// means, the sequence of a feature object is the feature itself.
	Sequence string
}

// DefaultFeatureConfig provides a default configuration to create a new
// feature object by best effort.
func DefaultFeatureConfig() FeatureConfig {
	newConfig := FeatureConfig{
		Positions: [][]float64{},
		Sequence:  "",
	}

	return newConfig
}

// NewFeature creates a new configured feature object. A feature represents a
// differentiable part of a given sequence.
func NewFeature(config FeatureConfig) (spec.Feature, error) {
	newFeature := &feature{
		Distribution:  nil,
		FeatureConfig: config,
		ID:            id.NewObjectID(id.Hex128),
		Mutex:         sync.Mutex{},
		Type:          ObjectTypeFeature,
	}

	if len(newFeature.Positions) == 0 {
		return nil, maskAnyf(invalidConfigError, "positions must not be empty")
	}
	if newFeature.Sequence == "" {
		return nil, maskAnyf(invalidConfigError, "sequence must not be empty")
	}

	newConfig := distribution.DefaultConfig()
	newConfig.Name = newFeature.Sequence
	newConfig.Vectors = newFeature.Positions
	newDistribution, err := distribution.NewDistribution(newConfig)
	if err != nil {
		return nil, maskAnyf(invalidConfigError, err.Error())
	}
	newFeature.Distribution = newDistribution

	return newFeature, nil
}

type feature struct {
	FeatureConfig

	Distribution spec.Distribution
	ID           spec.ObjectID
	Mutex        sync.Mutex
	Type         spec.ObjectType
}

func (f *feature) AddPosition(position []float64) error {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()

	if len(f.Positions) > 0 && len(f.Positions[0]) != len(position) {
		return maskAnyf(invalidPositionError, "must have length of %d", len(f.Positions))
	}

	f.Positions = append(f.Positions, position)

	return nil
}

func (f *feature) GetCount() int {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()

	return len(f.Positions)
}

func (f *feature) GetDistribution() spec.Distribution {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()

	return f.Distribution
}

func (f *feature) GetPositions() [][]float64 {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()

	return f.Positions
}

func (f *feature) GetSequence() string {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()

	return f.Sequence
}
