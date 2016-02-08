package serverspec

import (
	"golang.org/x/net/context"
)

type LogControl interface {
	ResetLevels(ctx context.Context) error
	ResetObjectTypes(ctx context.Context) error
	ResetVerbosity(ctx context.Context) error
	SetLevels(ctx context.Context, levels string) error
	SetObjectTypes(ctx context.Context, objectTypes string) error
	SetVerbosity(ctx context.Context, verbosity int) error
}
