package permutationlist

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
)

// Config represents the configuration used to create a new permutation
// list object.
type Config struct {
	// Settings.

	// Indizes represents an ordered list where each index represents a raw value
	// position.
	Indizes []int

	// MaxGrowth represents the maximum length PermutedValues is allowed to grow.
	MaxGrowth int

	// RawValues represents the values being used to permute PermutedValues.
	RawValues []interface{}
}

// DefaultConfig provides a default configuration to create a new
// permutation list object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Indizes:   []int{},
		MaxGrowth: 5,
		RawValues: []interface{}{},
	}

	return newConfig
}

// New creates a new configured permutation list object.
func New(config Config) (objectspec.PermutationList, error) {
	// Create new object.
	newObject := &list{
		Config: config,

		PermutedValues: []interface{}{},
	}

	// Validate new object.
	if newObject.MaxGrowth < 1 {
		return nil, maskAnyf(invalidConfigError, "max growth must be 1 or greater")
	}
	if len(newObject.RawValues) < 2 {
		return nil, maskAnyf(invalidConfigError, "raw values must be 2 or greater")
	}

	return newObject, nil
}

type list struct {
	Config

	// PermutedValues represents the permuted list of RawValues. Initially this is
	// the zero value []interface{}{}.
	PermutedValues []interface{}
}

func (l *list) GetIndizes() []int {
	return l.Indizes
}

func (l *list) GetMaxGrowth() int {
	return l.MaxGrowth
}

func (l *list) GetPermutedValues() []interface{} {
	return l.PermutedValues
}

func (l *list) GetRawValues() []interface{} {
	return l.RawValues
}

func (l *list) SetIndizes(indizes []int) {
	l.Indizes = indizes
	l.PermutedValues = permuteValues(l)
}
