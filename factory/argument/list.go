package argument

import (
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/spec"
)

// ListConfig represents the configuration used to create a new argument list
// object.
type ListConfig struct {
	// Settings.

	// Indizes represents an ordered list where each index represents a members
	// position.
	Indizes []int

	// MaxGrowth represents the maximum length Members is allowed to grow. It
	// used to find out how many input arguments a CLG expects.
	MaxGrowth int

	// Type descibes the concrete type the argument list represents.
	Type spec.PermutationType

	// Values represents the values being used to permute Members. For the
	// argument list this is a list of permutation lists.
	Values []spec.PermutationList
}

// DefaultListConfig provides a default configuration to create a new
// argument list object by best effort.
func DefaultListConfig() ListConfig {
	newBoolConfig := permutation.DefaultListConfig()
	newBoolConfig.Values = []interface{}{true, false}
	newBoolPermutationList, err := permutation.NewList(newBoolConfig)
	panicOnError(err)

	newDistributionConfig := permutation.DefaultListConfig()
	// TODO this somehow makes no sense. What to set as values?
	newDistributionConfig.Values = []interface{}{}
	newDistributionPermutationList, err := permutation.NewList(newDistributionConfig)
	panicOnError(err)

	newFeatureSetConfig := permutation.DefaultListConfig()
	// TODO this somehow makes no sense. What to set as values?
	newFeatureSetConfig.Values = []interface{}{}
	newFeatureSetPermutationList, err := permutation.NewList(newFeatureSetConfig)
	panicOnError(err)

	newFloat64Config := permutation.DefaultListConfig()
	newFloat64Config.Values = []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9)}
	newFloat64PermutationList, err := permutation.NewList(newFloat64Config)
	panicOnError(err)

	newIntConfig := permutation.DefaultListConfig()
	newIntConfig.Values = []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	newIntPermutationList, err := permutation.NewList(newIntConfig)
	panicOnError(err)

	newStringConfig := permutation.DefaultListConfig()
	newStringConfig.Values = []interface{}{"a", "b", "c", "d", "f", "g", "h", "i", "j", "k"}
	newStringPermutationList, err := permutation.NewList(newStringConfig)
	panicOnError(err)

	newFloat64SliceConfig := permutation.DefaultListConfig()
	newFloat64SlicePermutationList, err := permutation.NewList(newFloat64SliceConfig)
	panicOnError(err)

	newIntSliceConfig := permutation.DefaultListConfig()
	newIntSlicePermutationList, err := permutation.NewList(newIntSliceConfig)
	panicOnError(err)

	newStringSliceConfig := permutation.DefaultListConfig()
	newStringSlicePermutationList, err := permutation.NewList(newStringSliceConfig)
	panicOnError(err)

	newFloat64SliceSliceConfig := permutation.DefaultListConfig()
	newFloat64SliceSlicePermutationList, err := permutation.NewList(newFloat64SliceSliceConfig)
	panicOnError(err)

	newIntSliceSliceConfig := permutation.DefaultListConfig()
	newIntSliceSlicePermutationList, err := permutation.NewList(newIntSliceSliceConfig)
	panicOnError(err)

	newStringSliceSliceConfig := permutation.DefaultListConfig()
	newStringSliceSlicePermutationList, err := permutation.NewList(newStringSliceSliceConfig)
	panicOnError(err)

	newConfig := ListConfig{
		// Settings.
		Indizes:   []int{},
		MaxGrowth: 50,
		Type:      permutation.TypeArgumentList,
		Values: []spec.PermutationList{
			newBoolPermutationList,
			newDistributionPermutationList,
			newFeatureSetPermutationList,
			newFloat64PermutationList,
			newIntPermutationList,
			newStringPermutationList,
			newFloat64SlicePermutationList,
			newIntSlicePermutationList,
			newStringSlicePermutationList,
			newFloat64SliceSlicePermutationList,
			newIntSliceSlicePermutationList,
			newStringSliceSlicePermutationList,
		},
	}

	return newConfig
}

// NewList creates a new configured argument list object.
func NewList(config ListConfig) (spec.ArgumentList, error) {
	// Create new object.
	newList := &list{
		ListConfig: config,

		Members: []interface{}{},
		Pointer: 0,
	}

	// Validate new object.
	if newList.MaxGrowth < 2 {
		return nil, maskAnyf(invalidConfigError, "max growth must be 2 or greater")
	}
	if len(newList.Values) < 2 {
		return nil, maskAnyf(invalidConfigError, "values must be 2 or greater")
	}

	return newList, nil
}

type list struct {
	ListConfig

	// Members represents the list being permuted. Initially this is the zero
	// value of []interface{}: []interface{}{}.
	Members []interface{}

	Pointer int
}

// TODO
func (l *list) CreateArguments() ([]interface{}, error) {
	var newArguments []interface{}

	return newArguments, nil
}

func (l *list) GetIndizes() []int {
	return l.Indizes
}

func (l *list) GetMaxGrowth() int {
	return l.MaxGrowth
}

func (l *list) GetMembers() []interface{} {
	return l.Members
}

func (l *list) GetPointer() int {
	return l.Pointer
}

func (l *list) GetType() spec.PermutationType {
	return l.Type
}

func (l *list) GetValues() []interface{} {
	return l.Values
}

func (l *list) SetIndizes(indizes []int) {
	l.Indizes = indizes
}

func (l *list) SetMembers(members []interface{}) {
	l.Members = members
}

func (l *list) SetPointer(pointer int) {
	l.Pointer = pointer
}
