package serverspec

import (
	"golang.org/x/net/context"
)

type LogControl interface {
	ResetLevels(ctx context.Context) error
	ResetObjects(ctx context.Context) error
	ResetVerbosity(ctx context.Context) error
	SetLevels(ctx context.Context, levels string) error
	SetObjects(ctx context.Context, objects string) error
	SetVerbosity(ctx context.Context, verbosity int) error
}
