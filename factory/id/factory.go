package id

import (
	"github.com/xh3b4sd/anna/factory/random"
	"github.com/xh3b4sd/anna/spec"
)

// FactoryConfig represents the configuration used to create a new ID factory
// object.
type FactoryConfig struct {
	// Settings.

	// HashChars represents the characters used to create hashes.
	HashChars string

	// RandomFactory represents a factory returning random numbers.
	RandomFactory spec.RandomFactory
}

// DefaultFactoryConfig provides a default configuration to create a new ID factory
// object by best effort.
func DefaultFactoryConfig() FactoryConfig {
	newRandomFactory, err := random.NewFactory(random.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

	newConfig := FactoryConfig{
		// Settings.
		HashChars:     "abcdef0123456789", // hex character set
		RandomFactory: newRandomFactory,
	}

	return newConfig
}

// NewFactory creates a new configured ID factory object.
func NewFactory(config FactoryConfig) (spec.IDFactory, error) {
	newFactory := &factory{
		FactoryConfig: config,
	}

	if newFactory.HashChars == "" {
		return nil, maskAnyf(invalidConfigError, "hash characters must not be empty")
	}
	if newFactory.RandomFactory == nil {
		return nil, maskAnyf(invalidConfigError, "random factory must not be empty")
	}

	return newFactory, nil
}

// MustNewFactory creates either a new default configured id factory object, or
// panics.
func MustNewFactory() spec.IDFactory {
	newIDFactory, err := NewFactory(DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

	return newIDFactory
}

type factory struct {
	FactoryConfig
}

func (f *factory) WithType(idType spec.IDType) (spec.ObjectID, error) {
	n := int(idType)
	max := len(f.HashChars)

	newRandomNumbers, err := f.RandomFactory.CreateNMax(n, max)
	if err != nil {
		return "", maskAny(err)
	}

	b := make([]byte, n)

	for i, r := range newRandomNumbers {
		b[i] = f.HashChars[r]
	}

	return spec.ObjectID(b), nil
}
