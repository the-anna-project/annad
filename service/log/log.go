// Package log implements spec.Log. This logger interface is to simply log
// output to gather runtime information.
package log

import (
	"os"

	kitlog "github.com/go-kit/kit/log"

	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// Config represents the configuration used to create a new log object.
type Config struct {
	// Dependencies.

	// RootLogger is the underlying logger actually logging messages.
	RootLogger        servicespec.RootLogger
	ServiceCollection servicespec.Collection

	// Settings.

}

// DefaultConfig provides a default configuration to create a new log object by
// best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		RootLogger:        kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)),
		ServiceCollection: nil,

		// Settings.
	}

	return newConfig
}

// New creates a new configured log object.
func New(config Config) (servicespec.Log, error) {
	newService := &service{
		Config: config,
	}

	id, err := newService.Service().ID().New()
	if err != nil {
		return nil, maskAny(err)
	}
	newService.Metadata["id"] = id
	newService.Metadata["name"] = "log"
	newService.Metadata["type"] = "service"

	return newService, nil
}

// MustNew creates either a new default configured log service, or panics.
func MustNew() servicespec.Log {
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

func (s *service) Line(v ...interface{}) {
	s.RootLogger.Log(v...)
}
