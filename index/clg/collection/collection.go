// Package collection provides a collection of all CLGs that can be used.
//
// Note that this package defines a go generate statement to embed the CLG
// collection's source code within the compiled binary. That way the method
// bodies of the implemented CLGs are available for inspection and hashing.
// Hashes of CLG methods are used to check whether they changed. A change of a
// CLG method affects its functionality, its profile and probably even its use
// case. Thus changes of the CLGs method bodies need to be detected to trigger
// profile updates.
//
//go:generate ${GOPATH}/bin/loader generate -p collection
//
package collection

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
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
	Log spec.Log
}

// DefaultConfig provides a default configuration to create a new CLG
// collection object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log: log.NewLog(log.DefaultConfig()),
	}

	return newConfig
}

// New creates a new configured CLG collection object.
func New(config Config) (spec.CLGCollection, error) {
	newCollection := &collection{
		Config: config,

		ID:    id.NewObjectID(id.Hex128),
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
