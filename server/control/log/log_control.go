package logcontrol

import (
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/server/spec"
	"github.com/xh3b4sd/anna/spec"
)

type Config struct {
	Log spec.Log
}

func DefaultConfig() Config {
	return Config{
		Log: log.NewLog(log.DefaultConfig()),
	}
}

func NewLogControl(config Config) serverspec.LogControl {
	return logControl{
		Config: config,
	}
}

type logControl struct {
	Config
}

func (lc logControl) ResetLevels(ctx context.Context) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call ResetLevels")

	err := lc.Log.ResetLevels()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc logControl) ResetObjectTypes(ctx context.Context) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call ResetObjectTypes")

	err := lc.Log.ResetObjectTypes()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc logControl) ResetVerbosity(ctx context.Context) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call ResetVerbosity")

	err := lc.Log.ResetVerbosity()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc logControl) SetLevels(ctx context.Context, levels string) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call SetLevels")

	err := lc.Log.SetLevels(levels)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc logControl) SetObjectTypes(ctx context.Context, objectTypes string) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call SetObjectTypes")

	err := lc.Log.SetObjectTypes(objectTypes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc logControl) SetVerbosity(ctx context.Context, verbosity int) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call SetVerbosity")

	err := lc.Log.SetVerbosity(verbosity)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
