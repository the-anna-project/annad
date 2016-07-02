package api

// CoreRequest represents a request being send to the core of the neural
// network. This is basically good for requesting calculations from the neural
// network by providing an input and an optional expectation.
type CoreRequest struct {
	// Input represents the input being fed into the neural network. There must
	// be a none empty input given when requesting calculations from the neural
	// network.
	Input string `json:"input"`

	// ExpectationRequest represents the expectation object. This is used to
	// match against the calculated output. In case there is an expectation
	// given, the neural network tries to calculate an output until it generated
	// one that matches the given expectation.
	ExpectationRequest ExpectationRequest `json:"expectation,omitempty"`
}

// IsEmpty checks whether the current core request is empty. An empty core
// request can be consider invalid.
func (cr CoreRequest) IsEmpty() bool {
	return cr.Input == "" || cr.ExpectationRequest.IsEmpty()
}
