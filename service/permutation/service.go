// Package permutation provides a simple permutation service implementation in
// which the order of the members of a permutation list is permuted. Advantages
// of the permutation service is memory effiency and reproducability. It is
// memory efficient because all possible combinations are not stored in memory,
// but created on demand. Depending on the provided delta this can be quiet
// fast in case it is not too big. The service is reproducible because of the
// index used to represent a permutation.
//
//     Imagine the following example.
//
//         []interface{"a", 7, []float64{2.88}}
//
//     This is how the initial service permutation looks like. In fact, there
//     is no permutation.
//
//         []interface{}
//
//     This is how the first service permutation looks like.
//
//         []interface{"a"}
//
//     This is how the second service permutation looks like.
//
//         []interface{7}
//
//     This is how the third service permutation looks like.
//
//         []interface{[]float64{2.88}}
//
//     This is how the Nth service permutation looks like.
//
//         []interface{[]float64{2.88}, "a"}
//
package permutation

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// New creates a new permutation service.
func New() servicespec.Permutation {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.Collection

	// Settings.

	metadata map[string]string
}

func (s *service) Configure() error {
	// Settings.

	id, err := s.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "permutation",
		"type": "service",
	}

	return nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) PermuteBy(list objectspec.PermutationList, delta int) error {
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

func (s *service) Service() servicespec.Collection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.Collection) {
	s.serviceCollection = sc
}

func (s *service) Validate() error {
	// Dependencies.
	if s.serviceCollection == nil {
		return maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	return nil
}
