package network

import (
	"github.com/the-anna-project/annad/service/clg/divide"
	"github.com/the-anna-project/annad/service/clg/greater"
	"github.com/the-anna-project/annad/service/clg/input"
	"github.com/the-anna-project/annad/service/clg/isbetween"
	"github.com/the-anna-project/annad/service/clg/isgreater"
	"github.com/the-anna-project/annad/service/clg/islesser"
	"github.com/the-anna-project/annad/service/clg/lesser"
	"github.com/the-anna-project/annad/service/clg/multiply"
	"github.com/the-anna-project/annad/service/clg/output"
	"github.com/the-anna-project/annad/service/clg/pairsyntactic"
	"github.com/the-anna-project/annad/service/clg/readinformationid"
	"github.com/the-anna-project/annad/service/clg/readseparator"
	"github.com/the-anna-project/annad/service/clg/round"
	"github.com/the-anna-project/annad/service/clg/splitfeatures"
	"github.com/the-anna-project/annad/service/clg/subtract"
	"github.com/the-anna-project/annad/service/clg/sum"
	servicespec "github.com/the-anna-project/spec/service"
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
func (s *service) newCLGs() map[string]servicespec.CLGService {
	// TODO this should be initialized with the service collection
	list := []servicespec.CLGService{
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

	newCLGs := map[string]servicespec.CLGService{}

	for _, CLG := range list {
		newCLGs[CLG.Metadata()["name"]] = CLG
	}

	for name := range newCLGs {
		newCLGs[name].SetServiceCollection(s.serviceCollection)
	}

	return newCLGs
}
