package id

import (
	"github.com/xh3b4sd/anna/factory/random"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// Hex128 creates a new hexa decimal encoded, pseudo random, 128 bit hash.
	Hex128 spec.IDType = 16

	// Hex512 creates a new hexa decimal encoded, pseudo random, 512 bit hash.
	Hex512 spec.IDType = 64

	// Hex1024 creates a new hexa decimal encoded, pseudo random, 1024 bit hash.
	Hex1024 spec.IDType = 128

	// Hex2048 creates a new hexa decimal encoded, pseudo random, 2048 bit hash.
	Hex2048 spec.IDType = 256

	// Hex4096 creates a new hexa decimal encoded, pseudo random, 4096 bit hash.
	Hex4096 spec.IDType = 512
)

// FactoryConfig represents the configuration used to create a new ID factory
// object.
type FactoryConfig struct {
	// Settings.

	// HashChars represents the characters used to create hashes.
	HashChars string

	// RandomFactory represents a factory returning random numbers.
	RandomFactory spec.RandomFactory

	Type spec.IDType
}

// DefaultFactoryConfig provides a default configuration to create a new ID factory
// object by best effort.
func DefaultFactoryConfig() FactoryConfig {
	newConfig := FactoryConfig{
		// Settings.
		HashChars:     "abcdef0123456789", // hex character set
		RandomFactory: random.MustNewFactory(),
		Type:          Hex128,
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
	if newFactory.Type == 0 {
		return nil, maskAnyf(invalidConfigError, "ID type must not be empty")
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

func (f *factory) New() (spec.ObjectID, error) {
	ID, err := f.WithType(f.Type)
	if err != nil {
		return "", maskAny(err)
	}

	return ID, nil
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
