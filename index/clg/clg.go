// Package clg implementes fundamental actions used to create strategies that
// allow to discover new behavior for problem solving.
package clg

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCLGIndex represents the object type of the CLG index object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeCLGIndex spec.ObjectType = "clg-index"
)

// Config represents the configuration used to create a new CLG index object.
type Config struct {
	// Dependencies.
	Collection spec.CLGCollection
	Log        spec.Log
}

// DefaultConfig provides a default configuration to create a new CLG index
// object by best effort.
func DefaultConfig() Config {
	newCLGCollection, err := NewCLGCollection(DefaultCLGCollectionConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		Collection: newCLGCollection,
		Log:        log.NewLog(log.DefaultConfig()),
	}

	return newConfig
}

// NewCLGIndex creates a new configured CLG index object.
func NewCLGIndex(config Config) (spec.CLGIndex, error) {
	newCLGIndex := &clgIndex{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		Type:         ObjectTypeCLGIndex,
		ShutdownOnce: sync.Once{},
	}

	newCLGIndex.Log.Register(newCLGIndex.GetType())

	return newCLGIndex, nil
}

type clgIndex struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	Type         spec.ObjectType
	ShutdownOnce sync.Once
}

func (i *clgIndex) Boot() {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call Boot")

	i.BootOnce.Do(func() {
		go func() {
			err := i.CreateCLGProfiles(i.GetCLGCollection())
			if err != nil {
				i.Log.WithTags(spec.Tags{L: "E", O: i, T: nil, V: 4}, "%#v", maskAny(err))
			}
		}()
	})
}

func (i *clgIndex) CreateCLGProfiles(clgCollection spec.CLGCollection) error {
	return nil
}

func (i *clgIndex) GetCLGCollection() spec.CLGCollection {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call GetCLGCollection")

	return i.Collection
}

func (i *clgIndex) Shutdown() {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call Shutdown")

	i.ShutdownOnce.Do(func() {
	})
}
