package textinput

import (
	"github.com/the-anna-project/spec/object"
)

// Config represents the configuration used to create a new text
// response object.
type Config struct {
	// Settings.

	// Echo being set to true causes the provided input simply to be echoed back.
	// The provided input goes through the whole stack and is streamed back and
	// forth, but bypasses neural network. This is useful to test the
	// client/server integration of the gRPC stream implementation.
	Echo bool

	// Expectation represents the expectation object. This is used to match
	// against the calculated output. In case there is an expectation given, the
	// neural network tries to calculate an output until it generated one that
	// matches the given expectation.
	Expectation spec.Expectation

	// Input represents the input being fed into the neural network. There must
	// be a none empty input given when requesting calculations from the neural
	// network.
	Input string

	// SessionID represents the session the current text request is associated
	// with. This is provided to differentiate streams between different users.
	SessionID string
}

// DefaultConfig provides a default configuration to create a new
// text request object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Echo:        false,
		Expectation: nil,
		Input:       "",
		SessionID:   "",
	}

	return newConfig
}

// New creates a new configured text request object.
func New(config Config) (spec.TextInput, error) {
	newObject := &textInput{
		Config: config,
	}

	if newObject.Input == "" {
		return nil, maskAnyf(invalidConfigError, "input must not be empty")
	}
	if newObject.SessionID == "" {
		return nil, maskAnyf(invalidConfigError, "session ID must not be empty")
	}

	return newObject, nil
}

// MustNew creates either a new default configured text request
// object, or panics.
func MustNew() spec.TextInput {
	newObject, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newObject
}

type textInput struct {
	Config
}

func (ti *textInput) GetEcho() bool {
	return ti.Echo
}

func (ti *textInput) GetExpectation() spec.Expectation {
	return ti.Expectation
}

func (ti *textInput) GetInput() string {
	return ti.Input
}

func (ti *textInput) GetSessionID() string {
	return ti.SessionID
}
