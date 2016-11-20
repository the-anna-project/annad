package list

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// New creates a new permutation list object.
func New() objectspec.PermutationList {
	return &object{
		maxGrowth:      5,
		permutedValues: []interface{}{},
		rawValues:      []interface{}{},
	}
}

type object struct {
	// Settings.

	// indizes represents an ordered list where each index represents a raw value
	// position.
	indizes []int
	// maxGrowth represents the maximum length PermutedValues is allowed to grow.
	maxGrowth int
	// permutedValues represents the permuted list of RawValues. Initially this is
	// the zero value []interface{}{}.
	permutedValues []interface{}
	// rawValues represents the values being used to permute PermutedValues.
	rawValues []interface{}
}

func (o *object) Indizes() []int {
	return o.indizes
}

func (o *object) MaxGrowth() int {
	return o.maxGrowth
}

func (o *object) PermutedValues() []interface{} {
	return o.permutedValues
}

func (o *object) RawValues() []interface{} {
	return o.rawValues
}

func (o *object) SetIndizes(indizes []int) {
	o.indizes = indizes
	o.permutedValues = permuteValues(o)
}

func (o *object) SetMaxGrowth(maxGrowth int) {
	if maxGrowth <= 0 {
		panic(maskAnyf(invalidConfigError, "max growth must be 1 or greater"))
	}
	o.maxGrowth = maxGrowth
}

func (o *object) SetRawValues(rawValues []interface{}) {
	o.rawValues = rawValues
}
