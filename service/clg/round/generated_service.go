package round

// This file is generated by the CLG generator. Don't edit it manually. The CLG
// generator is invoked by go generate. For more information about the usage of
// the CLG generator check https://github.com/xh3b4sd/clggen or have a look at
// the clg package. There is the go generate statement placed to invoke clggen.

import (
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// New creates a new round CLG service.
func New() servicespec.CLG {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.Collection

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
		"kind": "round",
		"name": "clg",
		"type": "service",
	}
}

func (s *service) GetCalculate() interface{} {
	return s.calculate
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.Collection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(serviceCollection servicespec.Collection) {
	s.serviceCollection = serviceCollection
}
