package clg

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

var notEnoughArgumentsError = errgo.New("not enough arguments")

// IsNotEnoughArguments asserts notEnoughArgumentsError.
func IsNotEnoughArguments(err error) bool {
	return errgo.Cause(err) == notEnoughArgumentsError
}

var tooManyArgumentsError = errgo.New("not enough arguments")

// IsTooManyArguments asserts tooManyArgumentsError.
func IsTooManyArguments(err error) bool {
	return errgo.Cause(err) == tooManyArgumentsError
}

var wrongArgumentTypeError = errgo.New("wrong argument type")

// IsWrongArgumentType asserts wrongArgumentTypeError.
func IsWrongArgumentType(err error) bool {
	return errgo.Cause(err) == wrongArgumentTypeError
}

var methodNotFoundError = errgo.New("method not found")

// IsMethodNotFound asserts methodNotFoundError.
func IsMethodNotFound(err error) bool {
	return errgo.Cause(err) == methodNotFoundError
}

var indexOutOfRangeError = errgo.New("index out of range")

// IsIndexOutOfRange asserts indexOutOfRangeError.
func IsIndexOutOfRange(err error) bool {
	return errgo.Cause(err) == indexOutOfRangeError
}

var negativeIntError = errgo.New("negative integer")

// IsNegativeInt asserts negativeIntError.
func IsNegativeInt(err error) bool {
	return errgo.Cause(err) == negativeIntError
}

var invalidDividerError = errgo.New("invalid divider")

// IsInvalidDivider asserts invalidDividerError.
func IsInvalidDivider(err error) bool {
	return errgo.Cause(err) == invalidDividerError
}

var tooManyResultsError = errgo.New("too many results")

// IsTooManyResults asserts tooManyResultsError.
func IsTooManyResults(err error) bool {
	return errgo.Cause(err) == tooManyResultsError
}

var duplicatedMemberError = errgo.New("duplicated member")

// IsDuplicatedMember asserts duplicatedMemberError.
func IsDuplicatedMember(err error) bool {
	return errgo.Cause(err) == duplicatedMemberError
}

var cannotParseError = errgo.New("cannot parse")

// IsCannotParse asserts cannotParseError.
func IsCannotParse(err error) bool {
	return errgo.Cause(err) == cannotParseError
}

var cannotConvertError = errgo.New("cannot convert")

// IsCannotConvert asserts cannotConvertError.
func IsCannotConvert(err error) bool {
	return errgo.Cause(err) == cannotConvertError
}

var workerCanceledError = errgo.New("worker canceled")

// IsWorkerCanceled asserts workerCanceledError.
func IsWorkerCanceled(err error) bool {
	return errgo.Cause(err) == workerCanceledError
}

var invalidCLGError = errgo.New("invalid clg")

// IsInvalidCLG asserts invalidCLGError.
func IsInvalidCLG(err error) bool {
	return errgo.Cause(err) == invalidCLGError
}

var clgProfileNotFoundError = errgo.New("clg profile not found")

// IsCLGProfileNotFound asserts clgProfileNotFoundError.
func IsCLGProfileNotFound(err error) bool {
	return errgo.Cause(err) == clgProfileNotFoundError
}

var clgBodyNotFoundError = errgo.New("clg body not found")

// IsCLGBodyNotFound asserts clgBodyNotFoundError.
func IsCLGBodyNotFound(err error) bool {
	return errgo.Cause(err) == clgBodyNotFoundError
}
