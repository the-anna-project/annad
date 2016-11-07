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
	"github.com/xh3b4sd/anna/service/spec"
)

// Config represents the configuration used to create a new permutation
// service object.
type Config struct{}

// DefaultConfig provides a default configuration to create a new
// permutation service object by best effort.
func DefaultConfig() Config {
	newConfig := Config{}

	return newConfig
}

// New creates a new configured permutation service object.
func New(config Config) (spec.Permutation, error) {
	// Create new object.
	newService := &service{
		Config: config,
	}

	return newService, nil
}

// MustNew creates either a new default configured random service object,
// or panics.
func MustNew() spec.Permutation {
	newService, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newService
}

type service struct {
	Config
}

func (s *service) PermuteBy(list spec.PermutationList, delta int) error {
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
