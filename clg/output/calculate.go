// Package output implements spec.CLG and provides one of the two very special
// CLGs. That is the output CLG. Its purpose is to check if the calculated
// output matches the provided expectation, if any expectation given. The
// output CLG is handled in a special way because it determines the end of all
// requested calculations within the neural network. After the output CLG has
// been executed, the calculated output is returned back to the requesting
// client.
package output

import (
	"encoding/json"
	"reflect"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

// TODO there is no CLG to read from the certenty pyramid

func (c *clg) forwardNetworkPayload(ctx spec.Context, informationSequence string) error {
	// Find the original information sequence using the information ID from the
	// context.
	informationID, ok := ctx.GetInformationID()
	if !ok {
		return maskAnyf(invalidInformationIDError, "must not be empty")
	}
	informationSequenceKey := key.NewNetworkKey("information-id:%s:information-sequence", informationID)
	informationSequence, err := c.Storage().General().Get(informationSequenceKey)
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
	inputBehaviourID, err := c.Storage().General().Get(firstBehaviourIDKey)
	if err != nil {
		return maskAny(err)
	}

	// Lookup the behaviour ID of the current output CLG. Below we are using this
	// to set the source of the new network payload accordingly.
	outputBehaviourID, ok := ctx.GetBehaviourID()
	if !ok {
		return maskAnyf(invalidBehaviourIDError, "must not be empty")
	}

	// Create a new contect using the given context and adapt the new context with
	// the information of the current scope.
	newCtx := ctx.Clone()
	newCtx.SetBehaviourID(inputBehaviourID)
	newCtx.SetCLGName("input")
	newCtx.SetCLGTreeID(clgTreeID)
	// We do not need to set the expectation because it never changes.
	// We do not need to set the session ID because it never changes.

	// Create a new network payload.
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(informationSequence)}
	newNetworkPayloadConfig.Context = newCtx
	newNetworkPayloadConfig.Destination = spec.ObjectID(inputBehaviourID)
	newNetworkPayloadConfig.Sources = []spec.ObjectID{spec.ObjectID(outputBehaviourID)}
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
	err = c.Storage().General().PushToList(networkPayloadKey, string(b))
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (c *clg) calculate(ctx spec.Context, informationSequence string) error {
	// Check the calculated output against the provided expectation, if any. In
	// case there is no expectation provided, we simply go with what we
	// calculated. This then means we are probably not in a training situation.
	expectation, ok := ctx.GetExpectation()
	if !ok {
		err := c.sendTextResponse(ctx, informationSequence)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	// There is an expectation provided. Thus we are going to check the calculated
	// output against it. In case the provided expectation does match the
	// calculated result, we simply return it.
	calculatedOutput := expectation.GetOutput()
	if informationSequence == calculatedOutput {
		err := c.sendTextResponse(ctx, informationSequence)
		if err != nil {
			return maskAny(err)
		}
	}

	// The calculated output did not match the given expectation. That means we
	// need to calculate some new output to match the given expectation. To do so
	// we create a new network payload and assign the input CLG of the current CLG
	// tree to it by queueing the new network payload in the underlying storage.
	err := c.forwardNetworkPayload(ctx, informationSequence)
	if err != nil {
		return maskAny(err)
	}

	// The calculated output did not match the given expectation. We return an
	// error to let the neural network know about it.
	return maskAnyf(expectationNotMetError, "'%s' != '%s'", informationSequence, calculatedOutput)
}

func (c *clg) sendTextResponse(ctx spec.Context, informationSequence string) error {
	// Return the calculated output to the requesting client, if the
	// current CLG is the output CLG.
	newTextResponseConfig := api.DefaultTextResponseConfig()
	newTextResponseConfig.Output = informationSequence
	newTextResponse, err := api.NewTextResponse(newTextResponseConfig)
	if err != nil {
		return maskAny(err)
	}

	c.Gateway().TextOutput().GetChannel() <- newTextResponse

	return nil
}
