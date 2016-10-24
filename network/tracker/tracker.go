package tracker

import (
	"strings"
	"sync"

	"github.com/xh3b4sd/anna/factory"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

const (
	// ObjectType represents the object type of the tracker object. This is used
	// e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "tracker"
)

// Config represents the configuration used to create a new tracker object.
type Config struct {
	// Dependencies.
	FactoryCollection spec.FactoryCollection
	Log               spec.Log
	StorageCollection spec.StorageCollection
}

// DefaultConfig provides a default configuration to create a new tracker object
// by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		FactoryCollection: factory.MustNewCollection(),
		Log:               log.New(log.DefaultConfig()),
		StorageCollection: storage.MustNewCollection(),
	}

	return newConfig
}

// New creates a new configured tracker object.
func New(config Config) (spec.Tracker, error) {
	newTracker := &tracker{
		Config: config,

		ID:   id.MustNew(),
		Type: ObjectType,
	}

	if newTracker.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newTracker.FactoryCollection == nil {
		return nil, maskAnyf(invalidConfigError, "factory collection must not be empty")
	}
	if newTracker.StorageCollection == nil {
		return nil, maskAnyf(invalidConfigError, "storage collection must not be empty")
	}

	newTracker.Log.Register(newTracker.GetType())

	return newTracker, nil
}

// MustNew creates either a new default configured tracker object, or panics.
func MustNew() spec.Tracker {
	newTracker, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newTracker
}

type tracker struct {
	Config

	ID   spec.ObjectID
	Type spec.ObjectType
}

func (t *tracker) CLGIDs(CLG spec.CLG, networkPayload spec.NetworkPayload) error {
	var sources []string
	for _, s := range networkPayload.GetSources() {
		sources = append(sources, string(s))
	}

	destination := string(networkPayload.GetDestination())

	err := t.ExecuteEvents(sources, destination)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

// TODO
func (t *tracker) CLGNames(CLG spec.CLG, networkPayload spec.NetworkPayload) error {
	// destination name is CLG.GetName()
	// lookup source names in storage using source IDs from networkPayload
}

// TODO
func (t *tracker) ProcessExtendHeadEvent(source, destination string) error {
	err := t.Storage().Connection().WalkKeys("*", t.Closer, queue.GetInput())
	if err != nil {
		return maskAny(err)
	}
}

type connectionDetail struct {
	Connection  string
	Destination string
	Source      string
}

func (t *tracker) ProcessEvents(sources []string, destination string) error {
	// TODO comment
	// Create connection paths using the sources and destination tracked in the
	// given network payload. There might be multiple connection paths because
	// there might be multiple sources requesting one CLG together. The created
	// connection paths always contain the same destination, because there is only
	// one destination tracked in the given network payload. Note that the
	// destination is the behaviour ID of the CLG currently being executed.
	// connection paths look similar to the following structure.
	//
	//     behaviourID,behaviourID
	//
	var connectionDetails []connectionDetail
	for _, s := range sources {
		cd := ConnectionDetail{
			Connection:  s + "," + destination,
			Destination: destination,
			Source:      s,
		}
		connectionDetails = append(connectionDetails, cd)
	}

	// This is the list of lookup functions which is executed sequentially.
	//
	//     extend head
	//     extend tail
	//     new path
	//     split path

	var wg sync.WaitGroup

	go func() {
		for _, cd := range connectionDetails {
			wg.Add(1)
			go func(cd connectionDetail) {
				ok, err := t.Storage().Connection().Exists(cd.Connection)
				if err != nil {
					return maskAny(err)
				}
				if !ok {
					// TODO what value to use?
					err := t.Storage().Connection().Set(cd.Connection, "")
					if err != nil {
						return maskAny(err)
					}
				}
				wg.Done()
			}(cd)
		}
	}()

	go func() {
		wg.Add(1)
		err := t.Storage().Connection().WalkKeys("*", t.Closer, func(key string) error {
			for _, cd := range connectionDetails {
				if strings.HasPrefix(cp, cd.Destination) {
					// extend head
				}
				if strings.HasSuffix(cp, cd.Source) {
					// extend tail
				}
			}
		})
		if err != nil {
			return maskAny(err)
		}
		wg.Done()
	}()

	wg.Wait()

	return nil
}

func (t *tracker) Track(CLG spec.CLG, networkPayload spec.NetworkPayload) error {
	t.Log.WithTags(spec.Tags{C: nil, L: "D", O: t, V: 13}, "call Track")

	// This is the list of lookup functions which is executed seuqentially.
	lookups := []func(CLG spec.CLG, networkPayload spec.NetworkPayload) error{
		t.CLGIDs,
		t.CLGNames,
	}

	// Execute one lookup after another to track connection path patterns.
	var err error
	for _, l := range lookups {
		err = l(CLG, networkPayload)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}
