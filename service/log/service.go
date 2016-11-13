// Package log implements spec.Log. This logger interface is to simply log
// output to gather runtime information.
package log

import servicespec "github.com/xh3b4sd/anna/service/spec"

// New creates a new log service.
func New() servicespec.Log {
	return &service{}
}

type service struct {
	// Dependencies.

	// rootLogger is the underlying logger actually logging messages.
	rootLogger        servicespec.RootLogger
	serviceCollection servicespec.Collection

	// Settings.

	metadata map[string]string
}

func (s *service) Configure() error {
	// Settings.

	id, err := s.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "log",
		"type": "service",
	}

	return nil
}

func (s *service) Line(v ...interface{}) {
	s.rootLogger.Log(v...)
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.Collection {
	return s.serviceCollection
}

func (s *service) SetRootLogger(rl servicespec.RootLogger) {
	s.rootLogger = rl
}

func (s *service) SetServiceCollection(sc servicespec.Collection) {
	s.serviceCollection = sc
}

func (s *service) Validate() error {
	// Dependencies.

	if s.rootLogger == nil {
		return maskAnyf(invalidConfigError, "root logger must not be empty")
	}
	if s.serviceCollection == nil {
		return maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	return nil
}
