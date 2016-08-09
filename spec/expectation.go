package spec

// Expectation represents a description of what output is to be expected when
// requesting calculations by providing some input.
type Expectation interface {
	// IsEmpty checks whether the current expectation is empty or not. In case
	// there is no expectation request given, this should return true.
	IsEmpty() bool

	// Match verifies whether the current expectation was met.
	// TODO
	Match() (bool, error)
}
