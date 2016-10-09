package spec

import (
	"golang.org/x/net/context"
)

// LogControl provides a way to control Anna's log behaviour. E.g. filters can
// be set and reset using it.
type LogControl interface {
	// ResetLevels sets the log levels allowed to be logged back to the default
	// value.
	ResetLevels(ctx context.Context) error

	// ResetObjects sets the log objects allowed to be logged back to the
	// default value.
	ResetObjects(ctx context.Context) error

	// ResetVerbosity sets the log verbosity allowed to be logged back to the
	// default value.
	ResetVerbosity(ctx context.Context) error

	// SetLevels adds the given log levels to the list of log levels allowed to
	// be logged.
	SetLevels(ctx context.Context, levels string) error

	// SetObjects adds the given log objects to the list of log objects allowed
	// to be logged.
	SetObjects(ctx context.Context, objects string) error

	// SetVerbosity sets the log verbosity allowed to be logged to the given
	// value.
	SetVerbosity(ctx context.Context, verbosity int) error
}
