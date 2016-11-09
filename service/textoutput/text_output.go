// Package textoutput provides a simple service for sending text output
// responses.
package textoutput

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// Config represents the configuration used to create a new text output
// service object.
type Config struct {
	// Dependencies.
	ServiceCollection servicespec.Collection

	// Settings.
	Channel chan objectspec.TextOutput
}

// DefaultServiceConfig provides a default configuration to create a new text
// output service object by best effort.
func DefaultServiceConfig() Config {
	newConfig := Config{
		// Dependencies.
		ServiceCollection: nil,

		// Settings.
		Channel: make(chan objectspec.TextOutput, 1000),
	}

	return newConfig
}

// NewService creates a new configured text output service object.
func NewService(config Config) (servicespec.TextOutput, error) {
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

// MustNew creates either a new default configured text output service, or
// panics.
func MustNew() servicespec.TextOutput {
	newService, err := NewService(DefaultServiceConfig())
	if err != nil {
		panic(err)
	}

	return newService
}

type service struct {
	Config

	Metadata map[string]string
}

func (s *service) GetChannel() chan objectspec.TextOutput {
	return s.Channel
}
