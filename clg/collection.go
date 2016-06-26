package clg

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCollection represents the object type of the collection object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeCollection spec.ObjectType = "collection"
)

// CollectionConfig represents the configuration used to create a new
// collection object.
type CollectionConfig struct {
	// Dependencies.
	Log spec.Log

	// Settings.
	Expectation spec.Expectation
	Input       string
	Output      string
	SessionID   string
}

// DefaultCollectionConfig provides a default configuration to create a new
// collection object by best effort.
func DefaultCollectionConfig() CollectionConfig {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newConfig := CollectionConfig{
		// Dependencies.
		Log: log.NewLog(log.DefaultConfig()),

		// Settings.
		Expectation: nil,
		Input:       "",
		Output:      "",
		SessionID:   string(newID),
	}

	return newConfig
}

// New creates a new configured collection object.
func NewCollection(config CollectionConfig) (*Collection, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		return &Collection{}, maskAny(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		return &Collection{}, maskAny(err)
	}

	newCollection := &Collection{
		CollectionConfig: config,
		ID:               newID,
		Mutex:            sync.Mutex{},
		Type:             ObjectTypeCollection,
	}

	newCollection.Log.Register(newCollection.GetType())

	return newCollection, nil
}

// Collection represents the object holding all available CLGs.
type Collection struct {
	CollectionConfig

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}
