package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
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

var notFoundError = errgo.New("not found")

// IsNotFound checks whether a redis response was empty. Therefore it checks for
// redigo.ErrNil and notFoundError.
//
//     ErrNil indicates that a reply value is nil.
//
func IsNotFound(err error) bool {
	c := errgo.Cause(err)
	return c == notFoundError || c == redis.ErrNil
}

var invalidConfigError = errgo.New("invalid config")

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return errgo.Cause(err) == invalidConfigError
}

var queryExecutionFailedError = errgo.New("query execution failed")

// IsQueryExecutionFailed asserts queryExecutionFailedError.
func IsQueryExecutionFailed(err error) bool {
	return errgo.Cause(err) == queryExecutionFailedError
}
