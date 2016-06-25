package spec

// Expectation represents a description of what output is to be expected when
// requesting calculations by providing some input.
type Expectation interface {
	// Match verifies whether the given impulse matches the current expectation.
	Match(imp Impulse) (bool, error)
}
