package permutation

import (
	"github.com/xh3b4sd/anna/spec"
)

// FactoryConfig represents the configuration used to create a new permutation
// factory object.
type FactoryConfig struct{}

// DefaultFactoryConfig provides a default configuration to create a new
// permutation factory object by best effort.
func DefaultFactoryConfig() FactoryConfig {
	newConfig := FactoryConfig{}

	return newConfig
}

// NewFactory creates a new configured permutation factory object.
func NewFactory(config FactoryConfig) (spec.PermutationFactory, error) {
	// Create new object.
	newFactory := &factory{
		FactoryConfig: config,
	}

	return newFactory, nil
}

// MustNewFactory creates either a new default configured random factory object,
// or panics.
func MustNewFactory() spec.PermutationFactory {
	newPermutationFactory, err := NewFactory(DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

	return newPermutationFactory
}

type factory struct {
	FactoryConfig
}

func (f *factory) PermuteBy(list spec.PermutationList, delta int) error {
	if delta < 1 {
		return nil
	}

	newIndizes, err := createIndizesWithDelta(list, delta)
	if err != nil {
		return maskAny(err)
	}

	list.SetIndizes(newIndizes)

	return nil
}
