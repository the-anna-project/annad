package id

import (
	"crypto/rand"
	"io"
	"math/big"
	"time"

	"github.com/cenk/backoff"

	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeIDFactory represents the object type of the ID factory object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeIDFactory spec.ObjectType = "id-factory"
)

// FactoryConfig represents the configuration used to create a new ID factory
// object.
type FactoryConfig struct {
	// Settings.

	// BackOffFactory is supposed to be able to create a new spec.BackOff. Retry
	// implementations can make use of this to decide when to retry.
	BackOffFactory func() spec.BackOff

	// HashChars represents the characters used to create hashes.
	HashChars string

	// RandFactory represents a factory returning random values. Here e.g.
	// crypto/rand.Int can be used.
	RandFactory func(rand io.Reader, max *big.Int) (n *big.Int, err error)

	// RandReader represents an instance of a cryptographically strong
	// pseudo-random generator. Here e.g. crypto/rand.Reader can be used.
	RandReader io.Reader

	// Timeout represents the deadline being waited during ID creation before
	// returning a timeout error.
	Timeout time.Duration
}

// DefaultFactoryConfig provides a default configuration to create a new ID factory
// object by best effort.
func DefaultFactoryConfig() FactoryConfig {
	newConfig := FactoryConfig{
		// Settings.
		BackOffFactory: func() spec.BackOff {
			return &backoff.StopBackOff{}
		},
		HashChars:   "abcdef0123456789",
		RandFactory: rand.Int,
		RandReader:  rand.Reader,
		Timeout:     1 * time.Second,
	}

	return newConfig
}

// NewFactory creates a new configured ID factory object.
func NewFactory(config FactoryConfig) (spec.IDFactory, error) {
	newFactory := &factory{
		FactoryConfig: config,

		// Note the ID is assigned below, because the ID factory needs an ID
		// factory to assign an ID. So this ID factory is going to assign its own
		// ID by itself.
		Type: ObjectTypeIDFactory,
	}

	if newFactory.HashChars == "" {
		return nil, maskAnyf(invalidConfigError, "hash characters must not be empty")
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

	var err error
	newFactory.ID, err = newFactory.WithType(Hex128)
	if err != nil {
		return nil, maskAny(err)
	}

	return newFactory, nil
}

type factory struct {
	FactoryConfig

	ID   spec.ObjectID
	Type spec.ObjectType
}

func (f *factory) WithType(idType spec.IDType) (spec.ObjectID, error) {
	// Define the action.
	var result []byte
	action := func() error {
		done := make(chan []byte, 1)
		fail := make(chan error, 1)

		go func() {
			b := make([]byte, int(idType))

			for i := range b {
				max := big.NewInt(int64(len(f.HashChars)))
				j, err := f.RandFactory(f.RandReader, max)
				if err != nil {
					fail <- maskAny(err)
					return
				}

				b[i] = f.HashChars[int(j.Int64())]
			}

			done <- b
		}()

		select {
		case <-time.After(f.Timeout):
			return maskAnyf(timeoutError, "after %s", f.Timeout)
		case err := <-fail:
			return maskAny(err)
		case newID := <-done:
			result = newID
			return nil
		}
	}

	// Execute the action wrapped with a retrier.
	err := backoff.Retry(action, f.BackOffFactory())
	if err != nil {
		return "", maskAny(err)
	}

	return spec.ObjectID(result), nil
}
