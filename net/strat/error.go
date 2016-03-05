package stratnet

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

var combinationLimitError = errgo.New("combination limit")

// IsCombinationLimit asserts combinationLimitError.
func IsCombinationLimit(err error) bool {
	return errgo.Cause(err) == combinationLimitError
}

var invalidContextError = errgo.New("invalid context")

// IsInvalidContext asserts invalidContextError.
func IsInvalidContext(err error) bool {
	return errgo.Cause(err) == invalidContextError
}

var invalidScopeError = errgo.New("invalid scope")

// IsInvalidScope asserts invalidScopeError.
func IsInvalidScope(err error) bool {
	return errgo.Cause(err) == invalidScopeError
}

var noStrategyError = errgo.New("no strategy")

// IsInvalidScope asserts noStrategyError.
func IsNoStrategy(err error) bool {
	return errgo.Cause(err) == noStrategyError
}

var invalidStrategyError = errgo.New("invalid strategy")

// IsInvalidScope asserts invalidStrategyError.
func IsInvalidStrategy(err error) bool {
	return errgo.Cause(err) == invalidStrategyError
}
