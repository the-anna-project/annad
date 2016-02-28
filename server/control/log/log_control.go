package logcontrol

import (
	"sync"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
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

func NewLogControl(config Config) spec.LogControl {
	return &logControl{
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   spec.ObjectType(common.ObjectType.Log),
	}
}

type logControl struct {
	Config

	ID spec.ObjectID

	Mutex sync.Mutex

	Type spec.ObjectType
}

func (lc *logControl) ResetLevels(ctx context.Context) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call ResetLevels")

	err := lc.Log.ResetLevels()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc *logControl) ResetObjects(ctx context.Context) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call ResetObjects")

	err := lc.Log.ResetObjects()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc *logControl) ResetVerbosity(ctx context.Context) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call ResetVerbosity")

	err := lc.Log.ResetVerbosity()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc *logControl) SetLevels(ctx context.Context, levels string) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call SetLevels")

	err := lc.Log.SetLevels(levels)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc *logControl) SetObjects(ctx context.Context, objectTypes string) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call SetObjects")

	err := lc.Log.SetObjects(objectTypes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (lc *logControl) SetVerbosity(ctx context.Context, verbosity int) error {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 13}, "call SetVerbosity")

	err := lc.Log.SetVerbosity(verbosity)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
