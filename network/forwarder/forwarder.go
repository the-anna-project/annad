package forwarder

import (
	"reflect"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectType represents the object type of the forwarder object. This is used
	// e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "forwarder"
)

// Config represents the configuration used to create a new forwarder object.
type Config struct {
	// Dependencies.
	Log               spec.Log
	StorageCollection spec.StorageCollection
}

// DefaultConfig provides a default configuration to create a new forwarder
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log:               log.New(log.DefaultConfig()),
		StorageCollection: storage.MustNewCollection(),
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

	if newForwarder.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newForwarder.StorageCollection == nil {
		return nil, maskAnyf(invalidConfigError, "storage collection must not be empty")
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
	lookups := []func(CLG spec.CLG, networkPayload spec.NetworkPayload) ([]string, error){
		a.GetBehaviourIDs,
		a.NewBehaviourIDs,
	}

	// Execute one lookup after another. As soon as we find some behaviour IDs, we
	// use them to forward the given network payload.
	var behaviourIDs []string
	for _, lookup := range lookups {
		behaviourIDs, err = lookup(CLG, networkPayload)
		if IsBehaviourIDsNotFound(err) {
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

	// Forward the given network payload asynchronously to all CLGs references by
	// the found behaviour IDs.
	for _, ID := range behaviourIDs {
		go func(ID string) {
			// Prepare a new context for the new connection path.
			ctx := networkPayload.GetContext().Clone()
			ctx.SetBehaviourID(ID)

			// Create a new network payload.
			newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
			newNetworkPayloadConfig.Args = networkPayload.GetArgs()
			newNetworkPayloadConfig.Context = ctx
			newNetworkPayloadConfig.Destination = spec.ObjectID(ID)
			newNetworkPayloadConfig.Sources = []spec.ObjectID{networkPayload.GetDestination()}
			newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
			if err != nil {
				return maskAny(err)
			}

			// Write the transformed network payload to the queue.
			listKey := key.NewCLGKey("events:network-payload")
			element, err := json.Marshal(newNetworkPayload)
			if err != nil {
				return maskAny(err)
			}
			err = f.Storage().General().PushToList(listKey, element)
			if err != nil {
				return maskAny(err)
			}
		}(ID)
	}

	return nil
}

// TODO
func (f *forwarder) GetBehaviourIDs(CLG spec.CLG, networkPayload spec.NetworkPayload) ([]string, error) {
	var behaviourIDs []string

	behaviourID, ok := ctx.GetBehaviourID()
	if !ok {
		return nil, maskAnyf(invalidBehaviourIDError, "must not be empty")
	}
	behaviourIDsKey := key.NewCLGKey("behaviour-id:%s:behaviour-ids", behaviourID)
	err := f.Storage().General().WalkSet(behaviourIDsKey, f.Closer, func(element string) error {
		behaviourIDs = append(behaviourIDs, element)
		return nil
	})
	if err != nil {
		return nil, maskAny(err)
	}

	// TODO

	// Create a new network payload. Note that the old context of the old
	// network payload is removed to only append actual arguments to the new
	// network payload.
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = networkPayload.GetArgs()
	newNetworkPayloadConfig.Context = rule.Ctx
	newNetworkPayloadConfig.Destination = spec.ObjectID(ID)
	newNetworkPayloadConfig.Sources = []spec.ObjectID{networkPayload.GetDestination()}
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		return maskAny(err)
	}

	for _, rule := range forwardingRules {
		// Find the actual CLG based on its behaviour ID. Therefore we lookup the
		// behaviour name by the given behaviour ID. Data we read here is written
		// within several CLGs. That way the network creates its own connections
		// based on learned experiences.
		//
		// TODO where are these connections coming from?
		// TODO if there are none, we need to find some randomly
		// TODO there needs to be some sort of variation when executing existing CLG trees
		// TODO store network payload in queue
	}

	return behaviourIDs, nil
}

func (f *forwarder) NewBehaviourIDs(CLG spec.CLG, networkPayload spec.NetworkPayload) ([]string, error) {
	return nil, nil
}

func (f *forwarder) ToInputCLG(CLG spec.CLG, networkPayload spec.NetworkPayload) error {
	ctx := networkPayload.GetContext()

	// Find the original information sequence using the information ID from the
	// context.
	informationID, ok := ctx.GetCLGTreeID()
	if !ok {
		return maskAnyf(invalidInformationIDError, "must not be empty")
	}
	informationSequenceKey := key.NewCLGKey("information-id:%s:information-sequence", informationID)
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
	firstBehaviourIDKey := key.NewCLGKey("clg-tree-id:%s:first-behaviour-id", clgTreeID)
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
	listKey := key.NewCLGKey("events:network-payload")
	element, err := json.Marshal(newNetworkPayload)
	if err != nil {
		return maskAny(err)
	}
	err = f.Storage().General().PushToList(listKey, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
