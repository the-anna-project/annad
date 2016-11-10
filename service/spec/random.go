package spec

// Random creates pseudo random numbers. The service might implement
// retries using backoff strategies and timeouts.
type Random interface {
	Configure() error

	// CreateMax tries to create one new pseudo random number. The generated
	// number is within the range [0 max), which means that max is exclusive.
	CreateMax(max int) (int, error)

	// CreateNMax tries to create a list of new pseudo random numbers. n
	// represents the number of pseudo random numbers in the returned list. The
	// generated numbers are within the range [0 max), which means that max is
	// exclusive.
	CreateNMax(n, max int) ([]int, error)

	Metadata() map[string]string

	Service() Collection

	SetServiceCollection(sc Collection)

	Validate() error
}
