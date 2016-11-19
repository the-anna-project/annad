package text

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// New creates a new text input object.
func New() objectspec.TextInput {
	return &object{}
}

type object struct {
	// Settings.

	// echo being set to true causes the provided input simply to be echoed back.
	// The provided input goes through the whole stack and is streamed back and
	// forth, but bypasses neural network. This is useful to test the
	// client/server integration of the gRPC stream implementation.
	echo bool
	// expectation represents the expectation object. This is used to match
	// against the calculated output. In case there is an expectation given, the
	// neural network tries to calculate an output until it generated one that
	// matches the given expectation.
	expectation objectspec.Expectation
	// input represents the input being fed into the neural network. There must
	// be a none empty input given when requesting calculations from the neural
	// network.
	input string
	// sessionID represents the session the current text request is associated
	// with. This is provided to differentiate streams between different users.
	sessionID string
}

func (o *object) Echo() bool {
	return o.echo
}

func (o *object) Expectation() objectspec.Expectation {
	return o.expectation
}

func (o *object) Input() string {
	return o.input
}

func (o *object) SessionID() string {
	return o.sessionID
}

func (o *object) SetEcho(echo bool) {
	o.echo = echo
}

func (o *object) SetExpectation(expectation objectspec.Expectation) {
	o.expectation = expectation
}

func (o *object) SetInput(input string) {
	o.input = input
}

func (o *object) SetSessionID(sessionID string) {
	o.sessionID = sessionID
}
