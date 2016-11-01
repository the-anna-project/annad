package tracker

import (
	"sync"

	"github.com/xh3b4sd/anna/factory"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/key"
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
	destinationID := string(networkPayload.GetDestination())
	sourceIDs := networkPayload.GetSources()

	errors := make(chan error, len(sourceIDs))
	wg := sync.WaitGroup{}

	for _, s := range sourceIDs {
		wg.Add(1)
		go func(s string) {
			// Persist the single CLG ID connections.
			behaviourIDKey := key.NewNetworkKey("behaviour-id:%s:o:tracker:behaviour-ids", s)
			err := t.Storage().General().PushToSet(behaviourIDKey, destinationID)
			if err != nil {
				errors <- maskAny(err)
			}
			wg.Done()
		}(string(s))
	}

	wg.Wait()

	select {
	case err := <-errors:
		if err != nil {
			return maskAny(err)
		}
	default:
		// Nothing do here. No error occurred. All good.
	}

	return nil
}

func (t *tracker) CLGNames(CLG spec.CLG, networkPayload spec.NetworkPayload) error {
	destinationName := CLG.GetName()
	sourceIDs := networkPayload.GetSources()

	errors := make(chan error, len(sourceIDs))
	wg := sync.WaitGroup{}

	for _, s := range sourceIDs {
		wg.Add(1)
		go func(s string) {
			behaviourNameKey := key.NewNetworkKey("behaviour-id:%s:behaviour-name", s)
			name, err := t.Storage().General().Get(behaviourNameKey)
			if err != nil {
				errors <- maskAny(err)
			} else {
				// The errors channel is capable of buffering one error for each source
				// ID. The else clause is necessary to queue only one possible error for
				// each source ID. So in case the name lookup was successful, we are
				// able to actually persist the single CLG name connection.
				behaviourNameKey := key.NewNetworkKey("behaviour-name:%s:o:tracker:behaviour-names", name)
				err := t.Storage().General().PushToSet(behaviourNameKey, destinationName)
				if err != nil {
					errors <- maskAny(err)
				}
			}

			wg.Done()
		}(string(s))
	}

	wg.Wait()

	select {
	case err := <-errors:
		if err != nil {
			return maskAny(err)
		}
	default:
		// Nothing do here. No error occurred. All good.
	}

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
