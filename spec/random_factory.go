package spec

// RandomFactory creates pseudo random numbers. The factory might implement
// retries using backoff strategies and timeouts.
type RandomFactory interface {
	// CreateNMax tries to create a list of new pseudo random numbers. n
	// represents the number of pseudo random numbers in the returned list. The
	// generated numbers are within the range [0 max).
	CreateNMax(n, max int) ([]int, error)
}
