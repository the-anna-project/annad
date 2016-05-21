package argument

import (
	"strconv"
	"strings"

	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/spec"
)

// FactoryConfig represents the configuration used to create a new argument
// factory object.
type FactoryConfig struct {
	// Dependencies.
	PermutationFactory spec.PermutationFactory
}

// DefaultFactoryConfig provides a default configuration to create a new
// argument factory object by best effort.
func DefaultFactoryConfig() FactoryConfig {
	newFactory, err := permutation.NewFactory(permutation.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

	newConfig := FactoryConfig{
		PermutationFactory: newPermutationFactory,
	}

	return newConfig
}

// NewFactory creates a new configured argument factory object.
func NewFactory(config FactoryConfig) (spec.ArgumentFactory, error) {
	// Create new object.
	newFactory := &factory{
		FactoryConfig: config,
	}

	// Validate new object.
	if newFactory.PermutationFactory == nil {
		return nil, maskAnyf(invalidConfigError, "permutation factory must not be empty")
	}

	return newFactory, nil
}

type factory struct {
	FactoryConfig
}

func (f *factory) CreateIndex(list spec.PermutationList) string {
	return f.PermutationFactory.CreateIndex(list)
}

func (f *factory) PermuteBy(list spec.ArgumentList, delta int) error {
	var c int

	for {
		// Increment or reset pointer accordingly to its current position.
		newPointer, reset := shiftPointer(list.GetPointer(), len(list.GetMembers()))

		// In case the pointer was reset the given list needs to be permuted as a
		// whole. This also applies in case there are not enough members in the given
		// list. Then the permutation of the whole list will cause an additional
		// member to be created.
		if reset || list.GetPointer() >= len(list.GetMembers())-1 {
			// Permute the whole list.
			err := f.PermutationFactory.PermuteBy(list, delta)
			if err != nil {
				return maskAny(err)
			}
		} else {
			// Fetch the current member going to be permuted.
			newMembers := list.GetMembers()
			currentMember := newMembers[list.GetPointer()]
			// Permute the member the pointer is referring to.
			if cm, ok := currentMember.(spec.PermutationList); ok {
				err := f.PermutationFactory.PermuteBy(currentMember, delta)
				if permutation.IsMaxGrowthReached(err) {
					if c >= len(newMembers) {
						// All members already reached their max growth limit. Thus we
						// throw an error to notify the caller about this issue.
						return maskAny(err)
					}
					// The current member already reached the max growth limit. Thus we
					// are going ahead to try the next one. Therefore we count how many
					// times we already tried and save the state of the pointer to
					// permute the next member in the next iteration.
					c++
					list.SetPointer(newPointer)
					continue
				} else if err != nil {
					return maskAny(err)
				}
			} else {
				return maskAnyf(invalidTypeError, "expected spec.PermutationList got %T", currentMember)
			}
			// Put the permuted member back into the list.
			newMembers[list.GetPointer()] = currentMember
			list.SetMembers(newMembers)
		}

		// Finally we can set the pointer in case all operations are done
		// successfully.
		list.SetPointer(newPointer)

		// We permuted one member, so we can break out of the loop.
		break
	}

	return nil
}
