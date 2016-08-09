package sum

// This file is generated by the CLG generator. Don't edit it manually. The CLG
// generator is invoked by go generate. For more information about the usage of
// the CLG generator check https://github.com/xh3b4sd/clggen or have a look at
// the network package. There is the go generate statement to invoke clggen.

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
	ObjectType spec.ObjectType = "clg"
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

	if newCLG.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newCLG.Storage == nil {
		return nil, maskAnyf(invalidConfigError, "storage must not be empty")
	}

	newCLG.Log.Register(newCLG.GetType())

	return newCLG, nil
}

// MustNew creates either a new default configured CLG object, or panics.
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

func (c *clg) Calculate(inputs []reflect.Value) ([]reflect.Value, error) {
	outputs, err := filterError(reflect.ValueOf(c.calculate).Call(inputs))
	if err != nil {
		return nil, maskAny(err)
	}

	return outputs, nil
}

func (c *clg) GetName() string {
	return "sum"
}

func (c *clg) Inputs() []reflect.Type {
	t := reflect.TypeOf(c.calculate)

	var clgInputs []reflect.Type

	for i := 0; i < t.NumIn(); i++ {
		clgInputs = append(clgInputs, t.In(i))
	}

	return clgInputs
}

func (c *clg) SetLog(log spec.Log) {
	c.Log = log
}

func (c *clg) SetStorage(storage spec.Storage) {
	c.Storage = storage
}
