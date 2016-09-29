package network

import (
	"reflect"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

// receiver

func (n *network) forwardCLGs(ctx spec.Context, behaviorIDs []string, payload spec.NetworkPayload) error {
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
		clgName, err := n.Storage().General().Get(key.NewCLGKey("behavior-id:%s:behavior-name", ID))
		if err != nil {
			return maskAny(err)
		}
		CLG, err := n.clgByName(clgName)
		if err != nil {
			return maskAny(err)
		}
		CLG.GetInputChannel() <- newPayload
	}

	return nil
}

func (n *network) forwardInputCLG(networkPayload spec.NetworkPayload) error {
	// Find the original information sequence using the information ID from the
	// context.
	informationID, ok := networkPayload.GetContext().GetCLGTreeID()
	if !ok {
		return maskAnyf(invalidInformationIDError, "must not be empty")
	}
	informationSequenceKey := key.NewCLGKey("information-id:%s:information-sequence", informationID)
	informationSequence, err := n.Storage().General().Get(informationSequenceKey)
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
	behaviorID, err := n.Storage().General().Get(firstBehaviorIDKey)
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
	element, err := n.Storage().General().PopFromList(listKey, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

// helper

func (n *network) findConnections(ctx spec.Context, payload spec.NetworkPayload) ([]string, error) {
	var behaviorIDs []string

	behaviorID, ok := ctx.GetBehaviorID()
	if !ok {
		return nil, maskAnyf(invalidBehaviorIDError, "must not be empty")
	}
	behaviorIDsKey := key.NewCLGKey("behavior-id:%s:behavior-ids", behaviorID)

	err := n.Storage().General().WalkSet(behaviorIDsKey, n.Closer, func(element string) error {
		behaviorIDs = append(behaviorIDs, element)
		return nil
	})
	if err != nil {
		return nil, maskAny(err)
	}

	return behaviorIDs, nil
}
