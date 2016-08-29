// Package collection provides a collection of all CLGs that can be used.
package collection

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCLGCollection represents the object type of the CLG collection
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeCLGCollection spec.ObjectType = "clg-collection"
)

// Config represents the configuration used to create a new CLG collection
// object.
type Config struct {
	// Dependencies.
	IDFactory spec.IDFactory
	Log       spec.Log
}

// DefaultConfig provides a default configuration to create a new CLG
// collection object by best effort.
func DefaultConfig() Config {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		IDFactory: newIDFactory,
		Log:       log.New(log.DefaultConfig()),
	}

	return newConfig
}

// New creates a new configured CLG collection object.
func New(config Config) (spec.CLGCollection, error) {
	newCollection := &collection{
		Config: config,

		ID:    id.MustNew(),
		Mutex: sync.Mutex{},
		Type:  ObjectTypeCLGCollection,
	}

	newCollection.Log.Register(newCollection.GetType())

	return newCollection, nil
}

type collection struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}
