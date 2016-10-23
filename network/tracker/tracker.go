package tracker

import (
	"github.com/xh3b4sd/anna/factory"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/network/tracker/event"
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
func (t *tracker) ExecuteEvents(sources []string, destination string) error {
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
	queueConfig := event.DefaultEventQueueConfig()
	queueConfig.Destination = destination
	queueConfig.Sources = sources
	queue, err := event.NewQueue(queueConfig)
	if err != nil {
		return maskAny(err)
	}

	// TODO comment
	// We need to provide a way to control the operations against the underlying
	// storage. In case we tracked one event for each connection path we can stop
	// the walk through the key space. Further we also need to be aware of
	// external cancelation. In case the tracker is shut down, we need to stop all
	// work.
	go func() {
		// This is the list of lookup functions which is executed sequentially.
		lookups := []func(e spec.Event) error{
			t.ExecuteExtendHeadEvent,
			t.ExecuteExtendTailEvent,
			t.ExecuteNewPathEvent,
			t.ExecuteSplitPathEvent,
		}

		for {
			select {
			case <-t.Closer:
				return
			case <-done:
				return
			case <-queue.Complete():
				return
			case e := <-queue.Out():
				// Execute one lookup after another to track connection path patterns.
				go func(e spec.Event) {
					for _, l := range lookups {
						err := l(e)
						if err != nil {
							return maskAny(err)
						}
					}
				}(e)
			}
		}
	}()

	err := t.Storage().Connection().WalkKeys("*", t.Closer, queue.In())
	if err != nil {
		return maskAny(err)
	}

	close(done)

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
