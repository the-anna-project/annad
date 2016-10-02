package permutation

import (
	"github.com/xh3b4sd/anna/spec"
)

// ListConfig represents the configuration used to create a new permutation
// list object.
type ListConfig struct {
	// Settings.

	// Indizes represents an ordered list where each index represents a raw value
	// position.
	Indizes []int

	// MaxGrowth represents the maximum length PermutedValues is allowed to grow.
	MaxGrowth int

	// RawValues represents the values being used to permute PermutedValues.
	RawValues []interface{}
}

// DefaultListConfig provides a default configuration to create a new
// permutation list object by best effort.
func DefaultListConfig() ListConfig {
	newConfig := ListConfig{
		// Settings.
		Indizes:   []int{},
		MaxGrowth: 5,
		RawValues: []interface{}{},
	}

	return newConfig
}

// NewList creates a new configured permutation list object.
func NewList(config ListConfig) (spec.PermutationList, error) {
	// Create new object.
	newList := &list{
		ListConfig: config,

		PermutedValues: []interface{}{},
	}

	// Validate new object.
	if newList.MaxGrowth < 1 {
		return nil, maskAnyf(invalidConfigError, "max growth must be 1 or greater")
	}
	if len(newList.RawValues) < 2 {
		return nil, maskAnyf(invalidConfigError, "raw values must be 2 or greater")
	}

	return newList, nil
}

type list struct {
	ListConfig

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
