package scheduler

import (
	"fmt"

	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func maskAnyf(err error, f string, v ...interface{}) error {
	if err == nil {
		return nil
	}

	f = fmt.Sprintf("%s: %s", err.Error(), f)
	newErr := errgo.WithCausef(nil, errgo.Cause(err), f, v...)
	newErr.(*errgo.Err).SetLocation(1)

	return newErr
}

var invalidConfigError = errgo.New("invalid config")

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return errgo.Cause(err) == invalidConfigError
}

var jobNotFoundError = errgo.New("job not found")

// IsJobNotFound asserts jobNotFoundError.
func IsJobNotFound(err error) bool {
	return errgo.Cause(err) == jobNotFoundError
}

var actionNotFoundError = errgo.New("action not found")

// IsActionNotFound asserts actionNotFoundError.
func IsActionNotFound(err error) bool {
	return errgo.Cause(err) == actionNotFoundError
}
