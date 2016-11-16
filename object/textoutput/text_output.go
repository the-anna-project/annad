package textoutput

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// Config represents the configuration used to create a new text
// response object.
type Config struct {
	// Settings.

	// Output represents the output being calculated by the neural network.
	Output string `json:"output"`
}

// DefaultConfig provides a default configuration to create a new
// text response object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Output: "",
	}

	return newConfig
}

// New creates a new configured text response object.
func New(config Config) (objectspec.TextOutput, error) {
	newObject := &textOutput{
		Config: config,
	}

	return newObject, nil
}

type textOutput struct {
	Config
}

func (to *textOutput) GetOutput() string {
	return to.Output
}
