package permutation

import (
	"github.com/xh3b4sd/anna/service/spec"
)

// ServiceConfig represents the configuration used to create a new permutation
// service object.
type ServiceConfig struct{}

// DefaultServiceConfig provides a default configuration to create a new
// permutation service object by best effort.
func DefaultServiceConfig() ServiceConfig {
	newConfig := ServiceConfig{}

	return newConfig
}

// NewService creates a new configured permutation service object.
func NewService(config ServiceConfig) (spec.Permutation, error) {
	// Create new object.
	newService := &service{
		ServiceConfig: config,
	}

	return newService, nil
}

// MustNewService creates either a new default configured random service object,
// or panics.
func MustNewService() spec.Permutation {
	newPermutationService, err := NewService(DefaultServiceConfig())
	if err != nil {
		panic(err)
	}

	return newPermutationService
}

type service struct {
	ServiceConfig
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
