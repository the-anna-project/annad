// Package textinput provides a simple service for receiving text input
// requests.
package textinput

import (
	"github.com/xh3b4sd/anna/service/spec"
)

// Config represents the configuration used to create a new text input
// service object.
type Config struct {
	// Settings.
	Channel chan spec.TextRequest
}

// DefaultConfig provides a default configuration to create a new text
// input service object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Channel: make(chan spec.TextRequest, 1000),
	}

	return newConfig
}

// New creates a new configured text input service object.
func New(config Config) (spec.TextInput, error) {
	newService := &service{
		Config: config,
	}

	if newService.Channel == nil {
		return nil, maskAnyf(invalidConfigError, "channel must not be empty")
	}

	return newService, nil
}

// MustNew creates either a new default configured id service object, or
// panics.
func MustNew() spec.TextInput {
	newService, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newService
}

type service struct {
	Config
}

func (s *service) GetChannel() chan spec.TextRequest {
	return s.Channel
}
