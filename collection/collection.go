// Package collection provides the collection of CLGs that implement basic
// behaviour.
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
}

// DefaultCollectionConfig provides a default configuration to create a new
// collection object by best effort.
func DefaultCollectionConfig() CollectionConfig {
	newConfig := CollectionConfig{
		// Dependencies.
		Log: log.NewLog(log.DefaultConfig()),
	}

	return newConfig
}

// NewCollection creates a new configured collection object.
func NewCollection(config CollectionConfig) (*Collection, error) {
	newCollection := &Collection{
		CollectionConfig: config,

		BootOnce: sync.Once{},
		CLGs:     newCLGs(),
		ID:       id.MustNew(),
		Mutex:    sync.Mutex{},
		Type:     ObjectTypeCollection,
	}

	newCollection.Log.Register(newCollection.GetType())

	return newCollection, nil
}

// Collection represents the object holding all available CLGs.
type Collection struct {
	CollectionConfig

	BootOnce sync.Once
	CLGs     map[string]spec.CLG
	ID       spec.ObjectID
	Mutex    sync.Mutex
	Type     spec.ObjectType
}

// TODO
func (c *Collection) Activate(clgName string, inputs []reflect.Value) (bool, error) {
	return nil, nil
}

func (c *Collection) Boot() error {
	c.BootOnce.Do(func() {
		c.configureCLGs()
	})

	return nil
}

// TODO
func (c *Collection) Calculate(clgName string, inputs []reflect.Value) ([]reflect.Value, error) {
	return nil, nil
}

// TODO
func (c *Collection) Execute(clgName string, inputs []reflect.Value) ([]reflect.Value, error) {
	// Activate
	activate, err := c.Activate(clgName, inputs)
	if err != nil {
		return nil
	}

	// Calculate
	// Forward
	return nil, nil
}

// TODO
func (c *Collection) Forward(clgName string, inputs []reflect.Value) error {
	return nil, nil
}
