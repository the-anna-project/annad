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

// receiver

// maybeReturnAndLogErrors returns the very first error that may be given by
// errors. All upcoming errors may be given by the provided error channel will
// be logged using the configured logger with log level E and verbosity 4.
func (n *network) maybeReturnAndLogErrors(errors chan error) error {
	var executeErr error

	for err := range errors {
		if IsWorkerCanceled(err) {
			continue
		}

		if executeErr == nil {
			// Only return the first error.
			executeErr = err
		} else {
			// Log all errors but the first one
			n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
		}
	}

	if executeErr != nil {
		return maskAny(executeErr)
	}

	return nil
}

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
		newCLGs[name].SetFactoryCollection(n.FactoryCollection)
		newCLGs[name].SetLog(n.Log)
		newCLGs[name].SetStorageCollection(n.StorageCollection)
	}

	return newCLGs
}
