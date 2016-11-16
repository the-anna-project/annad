package spec

import (
	"io"
	"math/big"
	"time"

	objectspec "github.com/the-anna-project/spec/object"
)

// RandomService creates pseudo random numbers. The service might implement
// retries using backoff strategies and timeouts.
type RandomService interface {
	Boot()
	// CreateMax tries to create one new pseudo random number. The generated
	// number is within the range [0 max), which means that max is exclusive.
	CreateMax(max int) (int, error)
	// CreateNMax tries to create a list of new pseudo random numbers. n
	// represents the number of pseudo random numbers in the returned list. The
	// generated numbers are within the range [0 max), which means that max is
	// exclusive.
	CreateNMax(n, max int) ([]int, error)
	Metadata() map[string]string
	Service() ServiceCollection
	SetBackoffFactory(backoffFactory func() objectspec.Backoff)
	SetRandFactory(randFactory func(randReader io.Reader, max *big.Int) (n *big.Int, err error))
	SetServiceCollection(serviceCollection ServiceCollection)
	SetTimeout(timeout time.Duration)
}
