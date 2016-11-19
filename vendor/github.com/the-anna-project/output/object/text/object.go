// Package text provides a simple object for receiving text output.
package text

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// New creates a new text output object.
func New() objectspec.TextOutput {
	return &object{}
}

type object struct {
	// Settings.

	// output represents the output being calculated by the neural network.
	output string `json:"output"`
}

func (ti *object) Output() string {
	return ti.output
}

func (ti *object) SetOutput(output string) {
	ti.output = output
}
