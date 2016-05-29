package log

import (
	"sync"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeLogControl represents the object type of the log control object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeLogControl spec.ObjectType = "log-control"
)

// ControlConfig represents the configuration used to create a new log control
// object.
type ControlConfig struct {
	Log spec.Log
}

// DefaultControlConfig provides a default configuration to create a new log control
// object by best effort.
func DefaultControlConfig() ControlConfig {
	return ControlConfig{
		Log: log.NewLog(log.DefaultConfig()),
	}
}

// NewControl creates a new configured log control object.
func NewControl(config ControlConfig) (spec.LogControl, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		return nil, maskAny(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		return nil, maskAny(err)
	}

	newControl := &control{
		ControlConfig: config,
		ID:            newID,
		Mutex:         sync.Mutex{},
		Type:          spec.ObjectType(ObjectTypeLogControl),
	}

	newControl.Log.Register(newControl.GetType())

	return newControl, nil
}

type control struct {
	ControlConfig

	ID spec.ObjectID

	Mutex sync.Mutex

	Type spec.ObjectType
}

func (c *control) ResetLevels(ctx context.Context) error {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call ResetLevels")

	err := c.Log.ResetLevels()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (c *control) ResetObjects(ctx context.Context) error {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call ResetObjects")

	err := c.Log.ResetObjects()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (c *control) ResetVerbosity(ctx context.Context) error {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call ResetVerbosity")

	err := c.Log.ResetVerbosity()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (c *control) SetLevels(ctx context.Context, levels string) error {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call SetLevels")

	err := c.Log.SetLevels(levels)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (c *control) SetObjects(ctx context.Context, objectTypes string) error {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call SetObjects")

	err := c.Log.SetObjects(objectTypes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (c *control) SetVerbosity(ctx context.Context, verbosity int) error {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call SetVerbosity")

	err := c.Log.SetVerbosity(verbosity)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
