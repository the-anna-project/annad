// Package permutation provides a simple permutation factory in which the
// order of the members of a list is permuted. Advantages of the permutation
// factories are memory effiency and reproducability. It is memory efficient
// because all possible combinations are not stored in memory, but created on
// demand. It is reproducible because of the index used to represent a
// permutation.
//
//     This is how the initial factory permutation looks like.
//
//         []interface{"a"}
//
//     This is how the second factory permutation looks like.
//
//         []interface{"a", 7}
//
//     This is how the third factory permutation looks like.
//
//         []interface{"a", 7, []float64{2.88}}
//
//     This is how the Nth factory permutation looks like.
//
//         []interface{7, []float64{2.88}, "a"}
//
// Note that this implementation is not made for concurrent execution. A
// permutation factory is there for sequential incrementation. Concurrent
// accesses to it make no sense at all. Thus the implementation here does not
// make use of synchronization techniques as e.g. provided by the sync package.
//
package permutation

import (
	"strconv"
	"strings"

	"github.com/xh3b4sd/anna/spec"
)

// FactoryConfig represents the configuration used to create a new permutation
// factory object.
type FactoryConfig struct {
	// Settings.

	// Indizes represents the Index's translation. Each rank of the Index is
	// represented separately within Indizes.
	//
	//    345 translates to []int{3, 4, 5}
	//
	Indizes []int

	// MaxGrowth represents the maximum length Members is allowed to grow.
	MaxGrowth int

	// Values represents the values being used to permute Members.
	Values []interface{}
}

// DefaultFactoryConfig provides a default configuration to create a new
// permutation factory object by best effort.
func DefaultFactoryConfig() FactoryConfig {
	newConfig := FactoryConfig{
		// Settings.
		Indizes:   []int{},
		MaxGrowth: 50,
		Values:    []interface{}{},
	}

	return newConfig
}

// NewFactory creates a new configured factory object.
func NewFactory(config FactoryConfig) (spec.PermutationFactory, error) {
	// Create new object.
	newFactory := &factory{
		FactoryConfig: config,

		Index:   "",
		Members: []interface{}{},
	}

	// Validate new object.
	if newFactory.MaxGrowth < 2 {
		return nil, maskAnyf(invalidConfigError, "max growth must be 2 or greater")
	}
	if len(newFactory.Values) < 2 {
		return nil, maskAnyf(invalidConfigError, "values must be 2 or greater")
	}

	return newFactory, nil
}

type factory struct {
	FactoryConfig

	// Index represents the index being incremented to identify the string shape
	// permutation.
	Index string

	// Members represents the list being permuted. Initially this is the zero
	// value of []interface{}: []interface{}{}.
	Members []interface{}
}

func (f *factory) CreateIndexWithIndizes(indizes []int) string {
	var converted []string

	for _, i := range indizes {
		converted = append(converted, strconv.Itoa(i))
	}

	return strings.Join(converted, "")
}

func (f *factory) CreateMembersWithIndizes(indizes []int) []interface{} {
	// Map values for easy access.
	mapped := map[int]interface{}{}
	for i, v := range f.Values {
		mapped[i] = v
	}

	// Create the new permutation based on the updated index.
	newMembers := make([]interface{}, len(indizes))
	for i, index := range indizes {
		newMembers[i] = mapped[index]
	}

	return newMembers
}

func (f *factory) GetIndex() string {
	return f.Index
}

func (f *factory) GetIndizes() []int {
	return f.Indizes
}

func (f *factory) GetMaxGrowth() int {
	return f.MaxGrowth
}

func (f *factory) GetMembers() []interface{} {
	return f.Members
}

func (f *factory) GetValues() []interface{} {
	return f.Values
}

func (f *factory) PermuteBy(delta int) error {
	err := f.UpdateIndizesWithDelta(delta)
	if err != nil {
		return maskAny(err)
	}

	f.Members = f.CreateMembersWithIndizes(f.GetIndizes())
	f.Index = f.CreateIndexWithIndizes(f.GetIndizes())

	return nil
}

func (f *factory) UpdateIndizesWithDelta(delta int) error {
	// Initialize scope variables.
	base := len(f.Values)
	newIndizes := f.GetIndizes()
	operation := 0

	// Check for the initial situation. This is special and the only exception
	// within the algorithm.
	if len(newIndizes) == 0 {
		newIndizes = append(newIndizes, 0)
	}

	for {
		// Check amount of operations in the first place. That way the initial
		// situation as well as all other operations are covered.
		operation++
		if operation > delta {
			break
		}

		// Increment the least significant digit. That is, the right most index.
		// This is the only incrementation being done on the index.
		i := len(newIndizes) - 1
		lsd := newIndizes[i]
		lsd++
		newIndizes[i] = lsd

		// Cap the indizes and shift them if necessary. In case the least
		// significant digit was incremented above the base capacity, indizes need
		// to be shifted from right to left. This is like counting a number.
		var msdShifted bool
		newIndizes, msdShifted = shiftIndizes(newIndizes, base)
		if msdShifted {
			// Make sure the permutation does not growth more than allowed.
			if len(newIndizes)+1 > f.GetMaxGrowth() {
				return maskAny(maxGrowthReachedError)
			}

			// In case all the indizes where shifted, we zeroed out all indizes. Then
			// we need to prepend another zero as new most significant digit of the
			// index.
			newIndizes = prepend(newIndizes, 0, 0)
		}
	}

	f.Indizes = newIndizes

	return nil
}
