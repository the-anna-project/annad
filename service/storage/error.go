package storage

import (
	"fmt"

	"github.com/juju/errgo"

	"github.com/xh3b4sd/anna/service/storage/redis"
)

// IsNotFound combines IsNotFound error matchers of all storage
// implementations. IsNotFound should thus be used for error handling wherever
// spec.Storage is dealt with.
func IsNotFound(err error) bool {
	return redis.IsNotFound(err)
}

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
