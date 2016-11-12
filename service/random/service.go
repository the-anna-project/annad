// Package random provides a service implementation creating random numbers.
package random

import (
	"crypto/rand"
	"io"
	"math/big"
	"time"

	"github.com/cenk/backoff"

	"github.com/xh3b4sd/anna/service/spec"
)

// New creates a new random service.
func New() spec.Random {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection spec.Collection

	// Settings.

	// backoffFactory is supposed to be able to create a new spec.Backoff. Retry
	// implementations can make use of this to decide when to retry.
	backoffFactory func() spec.Backoff
	metadata       map[string]string
	// randFactory represents a service returning random values. Here e.g.
	// crypto/rand.Int can be used.
	randFactory func(rand io.Reader, max *big.Int) (n *big.Int, err error)
	// randReader represents an instance of a cryptographically strong
	// pseudo-random generator. Here e.g. crypto/rand.Reader can be used.
	randReader io.Reader
	// timeout represents the deadline being waited during random number creation
	// before returning a timeout error.
	timeout time.Duration
}

func (s *service) Configure() error {
	// Settings.

	id, err := s.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "random",
		"type": "service",
	}

	s.backoffFactory = func() spec.Backoff {
		return &backoff.StopBackOff{}
	}
	s.randFactory = rand.Int
	s.randReader = rand.Reader
	s.timeout = 1 * time.Second

	return nil
}

func (s *service) CreateMax(max int) (int, error) {
	// Define the action.
	var result int
	action := func() error {
		done := make(chan struct{}, 1)
		fail := make(chan error, 1)

		go func() {
			m := big.NewInt(int64(max))
			j, err := s.randFactory(s.randReader, m)
			if err != nil {
				fail <- maskAny(err)
				return
			}

			result = int(j.Int64())

			done <- struct{}{}
		}()

		select {
		case <-time.After(s.timeout):
			return maskAnyf(timeoutError, "after %s", s.timeout)
		case err := <-fail:
			return maskAny(err)
		case <-done:
			return nil
		}
	}

	// Execute the action wrapped with a retrier.
	err := backoff.Retry(action, s.backoffFactory())
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

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() spec.Collection {
	return s.serviceCollection
}

func (s *service) SetBackoffFactory(bf func() spec.Backoff) {
	s.backoffFactory = bf
}

func (s *service) SetServiceCollection(sc spec.Collection) {
	s.serviceCollection = sc
}

func (s *service) Validate() error {
	// Dependencies.

	if s.serviceCollection == nil {
		return maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	return nil
}
