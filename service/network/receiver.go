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
func (n *network) logNetworkError(err error) {
	if output.IsExpectationNotMet(err) {
		n.Service().Log().Line("msg", "%#v", maskAny(err))
	} else if err != nil {
		n.Service().Log().Line("msg", "%#v", maskAny(err))
	}
}

// logWorkerErrors logs all errors that are may be queued by the provided error
// channel using the configured logger with log level E and verbosity 4.
func (n *network) logWorkerErrors(errors chan error) {
	for err := range errors {
		if IsWorkerCanceled(err) {
			continue
		}

		n.Service().Log().Line("msg", "%#v", maskAny(err))
	}
}

// newCLGs returns a list of all CLGs which are configured and ready to be used
// within the neural network.
func (n *network) newCLGs() map[string]servicespec.CLG {
	// TODO this should be initialized with the service collection
	list := []servicespec.CLG{
		divide.MustNew(),
		input.MustNew(),
		divide.MustNew(),
		greater.MustNew(),
		input.MustNew(),
		isbetween.MustNew(),
		isgreater.MustNew(),
		islesser.MustNew(),
		lesser.MustNew(),
		multiply.MustNew(),
		output.MustNew(),
		pairsyntactic.MustNew(),
		readinformationid.MustNew(),
		readseparator.MustNew(),
		round.MustNew(),
		splitfeatures.MustNew(),
		subtract.MustNew(),
		sum.MustNew(),
	}

	newCLGs := map[string]servicespec.CLG{}

	for _, CLG := range list {
		newCLGs[CLG.GetName()] = CLG
	}

	for name := range newCLGs {
		newCLGs[name].SetServiceCollection(n.ServiceCollection)
		newCLGs[name].SetLog(n.Log)
		newCLGs[name].SetStorageCollection(n.StorageCollection)
	}

	return newCLGs
}
