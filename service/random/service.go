package random

import (
	"crypto/rand"
	"io"
	"math/big"
	"time"

	"github.com/cenk/backoff"

	servicespec "github.com/xh3b4sd/anna/service/spec"
	"github.com/xh3b4sd/anna/spec"
)

// ServiceConfig represents the configuration used to create a new random
// service object.
type ServiceConfig struct {
	// Settings.

	// BackoffFactory is supposed to be able to create a new spec.Backoff. Retry
	// implementations can make use of this to decide when to retry.
	BackoffFactory func() spec.Backoff

	// RandFactory represents a service returning random values. Here e.g.
	// crypto/rand.Int can be used.
	RandFactory func(rand io.Reader, max *big.Int) (n *big.Int, err error)

	// RandReader represents an instance of a cryptographically strong
	// pseudo-random generator. Here e.g. crypto/rand.Reader can be used.
	RandReader io.Reader

	// Timeout represents the deadline being waited during random number creation
	// before returning a timeout error.
	Timeout time.Duration
}

// DefaultServiceConfig provides a default configuration to create a new random
// service object by best effort.
func DefaultServiceConfig() ServiceConfig {
	newConfig := ServiceConfig{
		// Settings.
		BackoffFactory: func() spec.Backoff {
			return &backoff.StopBackOff{}
		},
		RandFactory: rand.Int,
		RandReader:  rand.Reader,
		Timeout:     1 * time.Second,
	}

	return newConfig
}

// NewService creates a new configured random service object.
func NewService(config ServiceConfig) (servicespec.Random, error) {
	newService := &service{
		ServiceConfig: config,
	}

	if newService.BackoffFactory == nil {
		return nil, maskAnyf(invalidConfigError, "backoff factory must not be empty")
	}
	if newService.RandFactory == nil {
		return nil, maskAnyf(invalidConfigError, "random factory must not be empty")
	}
	if newService.RandReader == nil {
		return nil, maskAnyf(invalidConfigError, "random reader must not be empty")
	}
	if newService.Timeout == 0 {
		return nil, maskAnyf(invalidConfigError, "creation timeout must not be empty")
	}

	return newService, nil
}

// MustNewService creates either a new default configured random service object,
// or panics.
func MustNewService() servicespec.Random {
	newRandomService, err := NewService(DefaultServiceConfig())
	if err != nil {
		panic(err)
	}

	return newRandomService
}

type service struct {
	ServiceConfig
}

func (s *service) CreateMax(max int) (int, error) {
	// Define the action.
	var result int
	action := func() error {
		done := make(chan struct{}, 1)
		fail := make(chan error, 1)

		go func() {
			m := big.NewInt(int64(max))
			j, err := s.RandFactory(s.RandReader, m)
			if err != nil {
				fail <- maskAny(err)
				return
			}

			result = int(j.Int64())

			done <- struct{}{}
		}()

		select {
		case <-time.After(s.Timeout):
			return maskAnyf(timeoutError, "after %s", s.Timeout)
		case err := <-fail:
			return maskAny(err)
		case <-done:
			return nil
		}
	}

	// Execute the action wrapped with a retrier.
	err := backoff.Retry(action, s.BackoffFactory())
	if err != nil {
		return 0, maskAny(err)
	}

	return result, nil
}

func (s *service) CreateNMax(n, max int) ([]int, error) {
	var result []int

	for i := 0; i < n; i++ {
		j, err := s.CreateMax(max)
		if err != nil {
			return nil, maskAny(err)
		}

		result = append(result, j)
	}

	return result, nil
}
