package api

// ExpectationRequest represents the request payload of an spec.Expectation.
type ExpectationRequest struct {
	// TODO
}

// IsEmpty checks whether the current expectation is empty.
func (er ExpectationRequest) IsEmpty() bool {
	return false
}
