// Package textinput provides a simple service for receiving text input
// requests.
package textinput

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// Config represents the configuration used to create a new text input
// service object.
type Config struct {
	// Dependencies.
	ServiceCollection servicespec.Collection

	// Settings.
	Channel chan objectspec.TextInput
}

// DefaultConfig provides a default configuration to create a new text
// input service object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		ServiceCollection: nil,

		// Settings.
		Channel: make(chan objectspec.TextInput, 1000),
	}

	return newConfig
}

// New creates a new configured text input service object.
func New(config Config) (servicespec.TextInput, error) {
	newService := &service{
		Config: config,
	}

	// Dependencies.
	if newService.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	// Settings.
	if newService.Channel == nil {
		return nil, maskAnyf(invalidConfigError, "channel must not be empty")
	}

	id, err := newService.Service().ID().New()
	if err != nil {
		return nil, maskAny(err)
	}
	newService.Metadata["id"] = id
	newService.Metadata["name"] = "text-input"
	newService.Metadata["type"] = "service"

	return newService, nil
}

// MustNew creates either a new default configured text input service, or
// panics.
func MustNew() servicespec.TextInput {
	newService, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newService
}

type service struct {
	Config

	Metadata map[string]string
}

func (s *service) GetChannel() chan objectspec.TextInput {
	return s.Channel
}
