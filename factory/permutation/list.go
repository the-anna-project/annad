package permutation

import (
	"github.com/xh3b4sd/anna/factory/argument"
	"github.com/xh3b4sd/anna/spec"
)

// ListConfig represents the configuration used to create a new permutation
// list object.
type ListConfig struct {
	// Settings.

	// Indizes represents an ordered list where each index represents a members
	// position.
	Indizes []int

	// MaxGrowth represents the maximum length Members is allowed to grow.
	MaxGrowth int

	// Type descibes the concrete type the permutation list represents.
	Type spec.PermutationType

	// Values represents the values being used to permute Members.
	Values []interface{}
}

// DefaultListConfig provides a default configuration to create a new
// permutation list object by best effort.
func DefaultListConfig() ListConfig {
	newConfig := ListConfig{
		// Settings.
		Indizes:   []int{},
		MaxGrowth: 10,
		Type:      argument.TypeNone,
		Values:    []interface{}{},
	}

	return newConfig
}

// NewList creates a new configured permutation list object.
func NewList(config ListConfig) (spec.PermutationList, error) {
	// Create new object.
	newList := &list{
		ListConfig: config,

		Members: []interface{}{},
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
