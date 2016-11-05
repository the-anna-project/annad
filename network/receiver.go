package network

import (
	"github.com/xh3b4sd/anna/clg/divide"
	"github.com/xh3b4sd/anna/clg/greater"
	"github.com/xh3b4sd/anna/clg/input"
	"github.com/xh3b4sd/anna/clg/is-between"
	"github.com/xh3b4sd/anna/clg/is-greater"
	"github.com/xh3b4sd/anna/clg/is-lesser"
	"github.com/xh3b4sd/anna/clg/lesser"
	"github.com/xh3b4sd/anna/clg/multiply"
	"github.com/xh3b4sd/anna/clg/output"
	"github.com/xh3b4sd/anna/clg/pair-syntactic"
	"github.com/xh3b4sd/anna/clg/read-information-id"
	"github.com/xh3b4sd/anna/clg/read-separator"
	"github.com/xh3b4sd/anna/clg/round"
	"github.com/xh3b4sd/anna/clg/split-features"
	"github.com/xh3b4sd/anna/clg/subtract"
	"github.com/xh3b4sd/anna/clg/sum"
	"github.com/xh3b4sd/anna/spec"
)

// logNetworkError logs the given error in a specific way dependening on the
// given error. If the given error is nil, nothing will be logged.
func (n *network) logNetworkError(err error) {
	if output.IsExpectationNotMet(err) {
		n.Log.WithTags(spec.Tags{C: nil, L: "W", O: n, V: 7}, "%#v", maskAny(err))
	} else if err != nil {
		n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
	}
}

// logWorkerErrors logs all errors that are may be queued by the provided error
// channel using the configured logger with log level E and verbosity 4.
func (n *network) logWorkerErrors(errors chan error) {
	for err := range errors {
		if IsWorkerCanceled(err) {
			continue
		}

		n.Log.WithTags(spec.Tags{L: "E", O: n, C: nil, V: 4}, "%#v", maskAny(err))
	}
}

// newCLGs returns a list of all CLGs which are configured and ready to be used
// within the neural network.
func (n *network) newCLGs() map[string]spec.CLG {
	list := []spec.CLG{
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

	newCLGs := map[string]spec.CLG{}

	for _, CLG := range list {
		newCLGs[CLG.GetName()] = CLG
	}

	for name := range newCLGs {
		newCLGs[name].SetServiceCollection(n.ServiceCollection)
		newCLGs[name].SetGatewayCollection(n.GatewayCollection)
		newCLGs[name].SetLog(n.Log)
		newCLGs[name].SetStorageCollection(n.StorageCollection)
	}

	return newCLGs
}
