// Package textoutput provides a simple service for sending text output
// responses.
package textoutput

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// ServiceConfig represents the configuration used to create a new text output
// service object.
type ServiceConfig struct {
	// Settings.
	Channel chan objectspec.TextOutput
}

// DefaultServiceConfig provides a default configuration to create a new text
// output service object by best effort.
func DefaultServiceConfig() ServiceConfig {
	newConfig := ServiceConfig{
		// Settings.
		Channel: make(chan objectspec.TextOutput, 1000),
	}

	return newConfig
}

// NewService creates a new configured text output service object.
func NewService(config ServiceConfig) (servicespec.TextOutput, error) {
	newService := &service{
		ServiceConfig: config,
	}

	if newService.Channel == nil {
		return nil, maskAnyf(invalidConfigError, "channel must not be empty")
	}

	return newService, nil
}

// MustNew creates either a new default configured id service object, or
// panics.
func MustNew() servicespec.TextOutput {
	newService, err := NewService(DefaultServiceConfig())
	if err != nil {
		panic(err)
	}

	return newService
}

type service struct {
	ServiceConfig
}

func (s *service) GetChannel() chan objectspec.TextOutput {
	return s.Channel
}
