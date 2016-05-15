package redisstorage

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

// IsNil is a specialized error matcher that checks whether the given error is
// nil. Simply checking if err == nil is not good enough for the storage
// implementation because of a nasty redigo detail. The redigo library used as
// redis client makes use of ErrNil. This error is returned in case redis
// returns nil. In fact this is no error at all, but we need to deal with this
// weird fact. Thus the IsNil helper.
//
//     ErrNil indicates that a reply value is nil.
//
func IsNil(err error) bool {
	return err == nil || err == redis.ErrNil
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
