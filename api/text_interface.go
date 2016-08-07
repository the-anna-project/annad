package api

// stream text request

// TextRequest represents a streamed request being send to the neural network.
// This is basically good for requesting calculations from the neural network
// by providing text input and an optional expectation object.
type TextRequest struct {
	// ExpectationRequest represents the expectation object. This is used to
	// match against the calculated output. In case there is an expectation
	// given, the neural network tries to calculate an output until it generated
	// one that matches the given expectation.
	ExpectationRequest ExpectationRequest `json:"expectation,omitempty"`

	// Input represents the input being fed into the neural network. There must
	// be a none empty input given when requesting calculations from the neural
	// network.
	Input string `json:"input"`

	// SessionID represents the session the current text request is associated
	// with. This is provided to differentiate streams between different users.
	SessionID string `json:"session_id,omitempty"`
}

// IsEmpty checks whether the current text request is empty. An empty text
// request can be considered invalid.
func (tr TextRequest) IsEmpty() bool {
	return tr.Input == "" || tr.SessionID == ""
}

// TextResponse represents a streamed response being replied from the neural
// text.
type TextResponse struct {
	// Output represents the output being calculated by the neural network.
	Output string `json:"output"`
}

// StreamTextRequest represents the request payload of the route used to stream
// text.
type StreamTextRequest struct {
	TextRequest TextRequest `json:"text_request"`
}

// StreamTextResponse represents the response payload of the route used to
// stream text. This payload by convention follows the same schema as all other
// API responses.
type StreamTextResponse Response
