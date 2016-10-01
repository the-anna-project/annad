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

// TODO
func (f *forwarder) Activate(CLG spec.CLG, networkPayload spec.NetworkPayload) (spec.NetworkPayload, error) {
	f.Log.WithTags(spec.Tags{C: nil, L: "D", O: f, V: 13}, "call Activate")

	// TODO

	// get stored network payloads
	// create list of network payloads
	// get CLG combinations known to be useful from storage (payloadFromConnections)
	// if no sufficient network payload, create permutation (payloadFromPermutations)
	// if no sufficient network payload, return error
	//

	behaviorID, ok := ctx.GetBehaviorID()
	if !ok {
		return nil, nil, maskAnyf(invalidBehaviorIDError, "must not be empty")
	}

	// Check if we have neural connections that tell us which payloads to use.
	payload, queue, err := f.payloadFromConnections(ctx, queue)
	if IsInvalidInterface(err) {
		// There are no sufficient connections. We need to come up with something
		// random.
		payload, queue, err = f.payloadFromPermutations(ctx, queue)
		if permutation.IsMaxGrowthReached(err) {
			// We could not find a sufficient payload for the requsted CLG by permuting
			// the queue of network payloads.
			return nil, nil, maskAnyf(invalidInterfaceError, "types must match")
		} else if err != nil {
			return nil, nil, maskAny(err)
		}

		// Once we found a new combination, we need to make sure the neural network
		// remembers it. Thus we store the connections between the current behavior
		// and the behaviors matching the interface of the current behavior.
		var behaviorIDs string
		for _, s := range payload.GetSources() {
			behaviorIDs += "," + string(s)
		}
		behaviorIDsKey := key.NewCLGKey("behavior-id:%s:activate-behavior-ids", behaviorID)
		err := f.Storage().General().Set(behaviorIDsKey, behaviorIDs)
		if err != nil {
			return nil, nil, maskAny(err)
		}
	}

	return payload, queue, nil
}

// receiver

func (f *forwarder) findForwardingRules(ctx spec.Context, payload spec.NetworkPayload) ([]string, error) {
	var behaviorIDs []string

	behaviorID, ok := ctx.GetBehaviorID()
	if !ok {
		return nil, maskAnyf(invalidBehaviorIDError, "must not be empty")
	}
	behaviorIDsKey := key.NewCLGKey("behavior-id:%s:behavior-ids", behaviorID)

	err := f.Storage().General().WalkSet(behaviorIDsKey, f.Closer, func(element string) error {
		behaviorIDs = append(behaviorIDs, element)
		return nil
	})
	if err != nil {
		return nil, maskAny(err)
	}

	// TODO

	// Create a new network payload. Note that the old context of the old
	// network payload is removed to only append actual arguments to the new
	// network payload.
	newPayloadConfig := api.DefaultNetworkPayloadConfig()
	newPayloadConfig.Args = payload.GetArgs()
	newPayloadConfig.Context = rule.Ctx
	newPayloadConfig.Destination = spec.ObjectID(ID)
	newPayloadConfig.Sources = []spec.ObjectID{payload.GetDestination()}
	newPayload, err := api.NewNetworkPayload(newPayloadConfig)
	if err != nil {
		return maskAny(err)
	}

	for _, rule := range forwardingRules {
		// Find the actual CLG based on its behavior ID. Therefore we lookup the
		// behavior name by the given behavior ID. Data we read here is written
		// within several CLGs. That way the network creates its own connections
		// based on learned experiences.
		//
		// TODO where are these connections coming from?
		// TODO if there are none, we need to find some randomly
		// TODO there needs to be some sort of variation when executing existing CLG trees
		// TODO store network payload in queue
	}

	return behaviorIDs, nil
}

func (f *forwarder) forwardCLGs(ctx spec.Context, behaviorIDs []string, payload spec.NetworkPayload) error {
	for _, ID := range behaviorIDs {
		// Prepare a new context for the new connection path.
		newCtx := ctx.Clone()
		newCtx.SetBehaviorID(ID)

		// Create a new network payload. Note that the old context of the old
		// network payload is removed to only append actual arguments to the new
		// network payload.
		newPayloadConfig := api.DefaultNetworkPayloadConfig()
		newPayloadConfig.Args = payload.GetArgs()
		newPayloadConfig.Context = newCtx
		newPayloadConfig.Destination = spec.ObjectID(ID)
		newPayloadConfig.Sources = []spec.ObjectID{payload.GetDestination()}
		newPayload, err := api.NewNetworkPayload(newPayloadConfig)
		if err != nil {
			return maskAny(err)
		}

		// Find the actual CLG based on its behavior ID. Therefore we lookup the
		// behavior name by the given behavior ID. Data we read here is written
		// within several CLGs. That way the network creates its own connections
		// based on learned experiences.
		//
		// TODO where are these connections coming from?
		// TODO if there are none, we need to find some randomly
		// TODO there needs to be some sort of variation when executing existing CLG trees
		//
		clgName, err := f.Storage().General().Get(key.NewCLGKey("behavior-id:%s:behavior-name", ID))
		if err != nil {
			return maskAny(err)
		}
		CLG, err := f.clgByName(clgName)
		if err != nil {
			return maskAny(err)
		}
		CLG.GetInputChannel() <- newPayload
	}

	return nil
}

func (f *forwarder) forwardInputCLG(networkPayload spec.NetworkPayload) error {
	// Find the original information sequence using the information ID from the
	// context.
	informationID, ok := networkPayload.GetContext().GetCLGTreeID()
	if !ok {
		return maskAnyf(invalidInformationIDError, "must not be empty")
	}
	informationSequenceKey := key.NewCLGKey("information-id:%s:information-sequence", informationID)
	informationSequence, err := f.Storage().General().Get(informationSequenceKey)
	if err != nil {
		return maskAny(err)
	}

	// Find the first behavior ID using the CLG tree ID from the context. The
	// behavior ID we are looking up here is the ID of the initial input CLG.
	clgTreeID, ok := networkPayload.GetContext().GetCLGTreeID()
	if !ok {
		return maskAnyf(invalidCLGTreeIDError, "must not be empty")
	}
	firstBehaviorIDKey := key.NewCLGKey("clg-tree-id:%s:first-behavior-id", clgTreeID)
	behaviorID, err := f.Storage().General().Get(firstBehaviorIDKey)
	if err != nil {
		return maskAny(err)
	}

	// Adapt the given context with the information of the current scope.
	networkPayload.GetContext().SetBehaviorID(behaviorID)
	networkPayload.GetContext().SetCLGName("input")
	networkPayload.GetContext().SetCLGTreeID(clgTreeID)
	// We do not need to set the expectation because it never changes.
	// We do not need to set the session ID because it never changes.

	// Create a new network payload.
	networkPayloadConfig := api.DefaultNetworkPayloadConfig()
	networkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(informationSequence)}
	networkPayloadConfig.Context = networkPayload.GetContext()
	networkPayloadConfig.Destination = spec.ObjectID(behaviorID)
	networkPayloadConfig.Sources = []spec.ObjectID{networkPayload.GetDestination()}
	newNetworkPayload, err := api.NewNetworkPayload(networkPayloadConfig)
	if err != nil {
		return maskAny(err)
	}

	// Write the transformed network payload to the queue.
	listKey := key.NewCLGKey("events:network-payload")
	element, err := json.Marshal(newNetworkPayload)
	if err != nil {
		return maskAny(err)
	}
	element, err := f.Storage().General().PopFromList(listKey, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
