package network

import (
	"github.com/xh3b4sd/anna/service/clg/divide"
	"github.com/xh3b4sd/anna/service/clg/greater"
	"github.com/xh3b4sd/anna/service/clg/input"
	"github.com/xh3b4sd/anna/service/clg/isbetween"
	"github.com/xh3b4sd/anna/service/clg/isgreater"
	"github.com/xh3b4sd/anna/service/clg/islesser"
	"github.com/xh3b4sd/anna/service/clg/lesser"
	"github.com/xh3b4sd/anna/service/clg/multiply"
	"github.com/xh3b4sd/anna/service/clg/output"
	"github.com/xh3b4sd/anna/service/clg/pairsyntactic"
	"github.com/xh3b4sd/anna/service/clg/readinformationid"
	"github.com/xh3b4sd/anna/service/clg/readseparator"
	"github.com/xh3b4sd/anna/service/clg/round"
	"github.com/xh3b4sd/anna/service/clg/splitfeatures"
	"github.com/xh3b4sd/anna/service/clg/subtract"
	"github.com/xh3b4sd/anna/service/clg/sum"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// logNetworkError logs the given error in a specific way dependening on the
// given error. If the given error is nil, nothing will be logged.
func (s *service) logNetworkError(err error) {
	if output.IsExpectationNotMet(err) {
		s.Service().Log().Line("msg", "%#v", maskAny(err))
	} else if err != nil {
		s.Service().Log().Line("msg", "%#v", maskAny(err))
	}
}

// logWorkerErrors logs all errors that are may be queued by the provided error
// channel using the configured logger with log level E and verbosity 4.
func (s *service) logWorkerErrors(errors chan error) {
	for err := range errors {
		if IsWorkerCanceled(err) {
			continue
		}

		s.Service().Log().Line("msg", "%#v", maskAny(err))
	}
}

// newCLGs returns a list of all CLGs which are configured and ready to be used
// within the neural network.
func (s *service) newCLGs() map[string]servicespec.CLG {
	// TODO this should be initialized with the service collection
	list := []servicespec.CLG{
		divide.New(),
		input.New(),
		divide.New(),
		greater.New(),
		input.New(),
		isbetween.New(),
		isgreater.New(),
		islesser.New(),
		lesser.New(),
		multiply.New(),
		output.New(),
		pairsyntactic.New(),
		readinformationid.New(),
		readseparator.New(),
		round.New(),
		splitfeatures.New(),
		subtract.New(),
		sum.New(),
	}

	newCLGs := map[string]servicespec.CLG{}

	for _, CLG := range list {
		newCLGs[CLG.Metadata()["name"]] = CLG
	}

	for name := range newCLGs {
		newCLGs[name].SetServiceCollection(s.serviceCollection)
	}

	return newCLGs
}
