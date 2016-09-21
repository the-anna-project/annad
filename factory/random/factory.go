package random

import (
	"crypto/rand"
	"io"
	"math/big"
	"time"

	"github.com/cenk/backoff"

	"github.com/xh3b4sd/anna/spec"
)

// FactoryConfig represents the configuration used to create a new random
// factory object.
type FactoryConfig struct {
	// Settings.

	// BackOffFactory is supposed to be able to create a new spec.BackOff. Retry
	// implementations can make use of this to decide when to retry.
	BackOffFactory func() spec.BackOff

	// RandFactory represents a factory returning random values. Here e.g.
	// crypto/rand.Int can be used.
	RandFactory func(rand io.Reader, max *big.Int) (n *big.Int, err error)

	// RandReader represents an instance of a cryptographically strong
	// pseudo-random generator. Here e.g. crypto/rand.Reader can be used.
	RandReader io.Reader

	// Timeout represents the deadline being waited during random number creation
	// before returning a timeout error.
	Timeout time.Duration
}

// DefaultFactoryConfig provides a default configuration to create a new random
// factory object by best effort.
func DefaultFactoryConfig() FactoryConfig {
	newConfig := FactoryConfig{
		// Settings.
		BackOffFactory: func() spec.BackOff {
			return &backoff.StopBackOff{}
		},
		RandFactory: rand.Int,
		RandReader:  rand.Reader,
		Timeout:     1 * time.Second,
	}

	return newConfig
}

// NewFactory creates a new configured random factory object.
func NewFactory(config FactoryConfig) (spec.RandomFactory, error) {
	newFactory := &factory{
		FactoryConfig: config,
	}

	if newFactory.RandFactory == nil {
		return nil, maskAnyf(invalidConfigError, "random factory must not be empty")
	}
	if newFactory.RandReader == nil {
		return nil, maskAnyf(invalidConfigError, "random reader must not be empty")
	}
	if newFactory.Timeout == 0 {
		return nil, maskAnyf(invalidConfigError, "creation timeout must not be empty")
	}

	return newFactory, nil
}

// MustNewFactory creates either a new default configured random factory object,
// or panics.
func MustNewFactory() spec.RandomFactory {
	newRandomFactory, err := NewFactory(DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

	return newRandomFactory
}

type factory struct {
	FactoryConfig
}

func (f *factory) CreateNMax(n, max int) ([]int, error) {
	// Define the action.
	var result []int
	action := func() error {
		done := make(chan struct{}, 1)
		fail := make(chan error, 1)

		go func() {
			for i := 0; i < n; i++ {
				m := big.NewInt(int64(max))
				j, err := f.RandFactory(f.RandReader, m)
				if err != nil {
					fail <- maskAny(err)
					return
				}

				result = append(result, int(j.Int64()))
			}

			done <- struct{}{}
		}()

		select {
		case <-time.After(f.Timeout):
			return maskAnyf(timeoutError, "after %s", f.Timeout)
		case err := <-fail:
			return maskAny(err)
		case <-done:
			return nil
		}
	}

	// Execute the action wrapped with a retrier.
	err := backoff.Retry(action, f.BackOffFactory())
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}
