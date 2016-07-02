package findconnections

import (
	"reflect"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectType represents the object type of the CLG object. This is used
	// e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "FindConnections"
)

// Config represents the configuration used to create a new CLG object.
type Config struct {
	// Dependencies.
	Log     spec.Log
	Storage spec.Storage
}

// DefaultConfig provides a default configuration to create a new CLG object by
// best effort.
func DefaultConfig() Config {
	newStorage, err := memory.NewStorage(memory.DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: newStorage,
	}

	return newConfig
}

// New creates a new configured CLG object.
func New(config Config) (spec.CLG, error) {
	newCLG := &clg{
		Config: config,
		ID:     id.MustNew(),
		Type:   ObjectType,
	}

	newCLG.Log.Register(newCLG.GetType())

	return newCLG, nil
}

func MustNew() spec.CLG {
	newCLG, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newCLG
}

type clg struct {
	Config

	ID   spec.ObjectID
	Type spec.ObjectType
}

func (c *clg) Execute(inputs []reflect.Value) ([]reflect.Value, error) {
	outputs, err := filterError(reflect.ValueOf(c.Calculate).Call(inputs))
	if err != nil {
		return nil, maskAny(err)
	}

	return outputs, nil
}
