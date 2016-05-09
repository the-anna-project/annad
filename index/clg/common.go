package clg

import (
	"github.com/xh3b4sd/anna/spec"
)

// maybeReturnAndLogErrors returns the very first error that may be given by
// errors. All upcoming errors may be given by the provided error channel will
// be logged using the configured logger with log level E and verbosity 4.
func (i *index) maybeReturnAndLogErrors(errors chan error) error {
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
			i.Log.WithTags(spec.Tags{L: "E", O: i, T: nil, V: 4}, "%#v", maskAny(err))
		}
	}

	if executeErr != nil {
		return maskAny(executeErr)
	}

	return nil
}
