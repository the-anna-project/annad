// Package log implements spec.Log. This logger interface is to simply log
// output to gather runtime information.
package log

import servicespec "github.com/the-anna-project/spec/service"

// New creates a new log service.
func New() servicespec.LogService {
	return &service{}
}

type service struct {
	// Dependencies.

	// rootLogger is the underlying logger actually logging messages.
	rootLogger        servicespec.RootLogger
	serviceCollection servicespec.ServiceCollection

	// Settings.

	metadata map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "log",
		"type": "service",
	}
}

func (s *service) Line(v ...interface{}) {
	s.rootLogger.Log(v...)
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetRootLogger(rl servicespec.RootLogger) {
	s.rootLogger = rl
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}
