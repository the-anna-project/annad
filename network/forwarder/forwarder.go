package forwarder

import (
	"encoding/json"
	"reflect"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/factory"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

const (
	// ObjectType represents the object type of the forwarder object. This is used
	// e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "forwarder"
)

// Config represents the configuration used to create a new forwarder object.
type Config struct {
	// Dependencies.
	FactoryCollection spec.FactoryCollection
	Log               spec.Log
	StorageCollection spec.StorageCollection

	// Settings.

	// MaxSignals represents the maximum number of signals being forwarded by one
	// CLG. When a requested CLG needs to decide where to forward signals to, it
	// may will forward up to MaxSignals signals to other CLGs, if any.
	MaxSignals int
}

// DefaultConfig provides a default configuration to create a new forwarder
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		FactoryCollection: factory.MustNewCollection(),
		Log:               log.New(log.DefaultConfig()),
		StorageCollection: storage.MustNewCollection(),

		// Settings.
		MaxSignals: 5,
	}

	return newConfig
}

// New creates a new configured forwarder object.
func New(config Config) (spec.Forwarder, error) {
	newForwarder := &forwarder{
		Config: config,

		ID:   id.MustNew(),
		Type: ObjectType,
	}

	// Dependencies.
	if newForwarder.FactoryCollection == nil {
		return nil, maskAnyf(invalidConfigError, "factory collection must not be empty")
	}
	if newForwarder.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newForwarder.StorageCollection == nil {
		return nil, maskAnyf(invalidConfigError, "storage collection must not be empty")
	}

	// Settings.
	if newForwarder.MaxSignals == 0 {
		return nil, maskAnyf(invalidConfigError, "maximum signals must not be empty")
	}

	newForwarder.Log.Register(newForwarder.GetType())

	return newForwarder, nil
}

// MustNew creates either a new default configured forwarder object, or panics.
func MustNew() spec.Forwarder {
	newForwarder, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newForwarder
}

type forwarder struct {
	Config

	ID   spec.ObjectID
	Type spec.ObjectType
}

func (f *forwarder) Forward(CLG spec.CLG, networkPayload spec.NetworkPayload) error {
	f.Log.WithTags(spec.Tags{C: nil, L: "D", O: f, V: 13}, "call Forward")

	// This is the list of lookup functions which is executed seuqentially.
	lookups := []func(CLG spec.CLG, networkPayload spec.NetworkPayload) ([]spec.NetworkPayload, error){
		f.GetNetworkPayloads,
		f.NewNetworkPayloads,
	}

	// Execute one lookup after another. As soon as we find some behaviour IDs, we
	// use them to forward the given network payload.
	var newNetworkPayloads []spec.NetworkPayload
	var err error
	for _, lookup := range lookups {
		newNetworkPayloads, err = lookup(CLG, networkPayload)
		if IsNetworkPayloadsNotFound(err) {
			// There could no behaviour IDs be found by this lookup. Go on and try the
			// next one.
			continue
		} else if err != nil {
			return maskAny(err)
		}

		// The current lookup was successful. We do not need to execute any further
		// lookup, but can go on with the behaviour IDs found.
		break
	}

	// Forward the found network payloads to other CLGs by adding them to the
	// queue so other processes can fetch them.
	for _, np := range newNetworkPayloads {
		networkPayloadKey := key.NewNetworkKey("events:network-payload")
		b, err := json.Marshal(np)
		if err != nil {
			return maskAny(err)
		}
		// TODO store asynchronuously
		err = f.Storage().General().PushToSet(networkPayloadKey, string(b))
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}

func (f *forwarder) GetMaxSignals() int {
	return f.MaxSignals
}

func (f *forwarder) GetNetworkPayloads(CLG spec.CLG, networkPayload spec.NetworkPayload) ([]spec.NetworkPayload, error) {
	ctx := networkPayload.GetContext()

	// Check if there are behaviour IDs known that we can use to forward the
	// current network payload to.
	behaviourID, ok := ctx.GetBehaviourID()
	if !ok {
		return nil, maskAnyf(invalidBehaviourIDError, "must not be empty")
	}
	behaviourIDsKey := key.NewNetworkKey("forward:configuration:behaviour-id:%s:behaviour-ids", behaviourID)
	newBehaviourIDs, err := f.Storage().General().GetAllFromSet(behaviourIDsKey)
	if storage.IsNotFound(err) {
		// No configuration of behaviour IDs is stored. Thus we return an error.
		// Eventually some other lookup is able to find sufficient network payloads.
		return nil, maskAny(networkPayloadsNotFoundError)
	} else if err != nil {
		return nil, maskAny(err)
	}

	// Create a list of new network payloads.
	var newNetworkPayloads []spec.NetworkPayload
	for _, behaviourID := range newBehaviourIDs {
		// Prepare a new context for the new network payload.
		newCtx := ctx.Clone()
		newCtx.SetBehaviourID(behaviourID)

		// Create a new network payload.
		newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
		newNetworkPayloadConfig.Args = networkPayload.GetArgs()
		newNetworkPayloadConfig.Context = newCtx
		newNetworkPayloadConfig.Destination = spec.ObjectID(behaviourID)
		newNetworkPayloadConfig.Sources = []spec.ObjectID{networkPayload.GetDestination()}
		newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
		if err != nil {
			return nil, maskAny(err)
		}

		newNetworkPayloads = append(newNetworkPayloads, newNetworkPayload)
	}

	return newNetworkPayloads, nil
}

func (f *forwarder) NewNetworkPayloads(CLG spec.CLG, networkPayload spec.NetworkPayload) ([]spec.NetworkPayload, error) {
	ctx := networkPayload.GetContext()

	// Decide how many new behaviour IDs should be created. This defines the
	// number of signals being forwarded to other CLGs. Here we want to make a
	// pseudo random decision. CreateMax takes a max paramater which is exclusive.
	// Therefore we increment the configuration for the maximum signals desired by
	// one, to reflect the maximum setting properly.
	maxSignals, err := f.Factory().Random().CreateMax(f.GetMaxSignals() + 1)
	if err != nil {
		return nil, maskAny(err)
	}

	// Create the desired number of behaviour IDs.
	var newBehaviourIDs []string
	for i := 0; i < maxSignals; i++ {
		newBehaviourID, err := f.Factory().ID().New()
		if err != nil {
			return nil, maskAny(err)
		}
		newBehaviourIDs = append(newBehaviourIDs, string(newBehaviourID))
	}

	// Store each new behaviour ID in the underlying storage.
	behaviourID, ok := ctx.GetBehaviourID()
	if !ok {
		return nil, maskAnyf(invalidBehaviourIDError, "must not be empty")
	}
	behaviourIDsKey := key.NewNetworkKey("forward:configuration:behaviour-id:%s:behaviour-ids", behaviourID)
	for _, behaviourID := range newBehaviourIDs {
		// TODO store asynchronuously
		err = f.Storage().General().PushToSet(behaviourIDsKey, behaviourID)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	// Create a list of new network payloads.
	var newNetworkPayloads []spec.NetworkPayload
	for _, behaviourID := range newBehaviourIDs {
		// Prepare a new context for the new network payload.
		newCtx := ctx.Clone()
		newCtx.SetBehaviourID(behaviourID)

		// Create a new network payload.
		newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
		newNetworkPayloadConfig.Args = networkPayload.GetArgs()
		newNetworkPayloadConfig.Context = newCtx
		newNetworkPayloadConfig.Destination = spec.ObjectID(behaviourID)
		newNetworkPayloadConfig.Sources = []spec.ObjectID{networkPayload.GetDestination()}
		newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
		if err != nil {
			return nil, maskAny(err)
		}

		newNetworkPayloads = append(newNetworkPayloads, newNetworkPayload)
	}

	return newNetworkPayloads, nil
}

// TODO this should probably be moved to the output CLG
func (f *forwarder) ToInputCLG(CLG spec.CLG, networkPayload spec.NetworkPayload) error {
	ctx := networkPayload.GetContext()

	// Find the original information sequence using the information ID from the
	// context.
	informationID, ok := ctx.GetInformationID()
	if !ok {
		return maskAnyf(invalidInformationIDError, "must not be empty")
	}
	informationSequenceKey := key.NewNetworkKey("information-id:%s:information-sequence", informationID)
	informationSequence, err := f.Storage().General().Get(informationSequenceKey)
	if err != nil {
		return maskAny(err)
	}

	// Find the first behaviour ID using the CLG tree ID from the context. The
	// behaviour ID we are looking up here is the ID of the initial input CLG.
	clgTreeID, ok := ctx.GetCLGTreeID()
	if !ok {
		return maskAnyf(invalidCLGTreeIDError, "must not be empty")
	}
	firstBehaviourIDKey := key.NewNetworkKey("clg-tree-id:%s:first-behaviour-id", clgTreeID)
	behaviourID, err := f.Storage().General().Get(firstBehaviourIDKey)
	if err != nil {
		return maskAny(err)
	}

	// Adapt the given context with the information of the current scope.
	ctx.SetBehaviourID(behaviourID)
	ctx.SetCLGName("input")
	ctx.SetCLGTreeID(clgTreeID)
	// We do not need to set the expectation because it never changes.
	// We do not need to set the session ID because it never changes.

	// Create a new network payload.
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(informationSequence)}
	newNetworkPayloadConfig.Context = ctx
	newNetworkPayloadConfig.Destination = spec.ObjectID(behaviourID)
	newNetworkPayloadConfig.Sources = []spec.ObjectID{networkPayload.GetDestination()}
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		return maskAny(err)
	}

	// Write the transformed network payload to the queue.
	networkPayloadKey := key.NewNetworkKey("events:network-payload")
	b, err := json.Marshal(newNetworkPayload)
	if err != nil {
		return maskAny(err)
	}
	err = f.Storage().General().PushToList(networkPayloadKey, string(b))
	if err != nil {
		return maskAny(err)
	}

	return nil
}
