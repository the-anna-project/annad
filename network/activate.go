package network

import (
	"reflect"
	"strings"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

// receiver

func (n *network) payloadFromConnections(ctx spec.Context, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error) {
	// Fetch the available behavior IDs which are known to be useful connections
	// during the activation of the requested CLG. The payloads sent by the CLGs
	// being fetched here are useful because, in the past, they have already been
	// helpful within the current CLG tree.
	behaviorID, ok := ctx.GetBehaviorID()
	if !ok {
		return nil, nil, maskAnyf(invalidBehaviorIDError, "must not be empty")
	}
	behaviorIDsKey := key.NewCLGKey("behavior-id:%s:activate-behavior-ids", behaviorID)
	list, err := n.Storage().General().Get(behaviorIDsKey)
	if err != nil {
		return nil, nil, maskAny(err)
	}
	behaviorIDs := strings.Split(list, ",")

	// Check if there is a network payload for each behavior ID we found in the
	// storage. Here it is important to obtain the order of the behavior IDs
	// stored as connections. They represent the input interface of the requested
	// CLG.
	//
	// TODO there needs to be some sort of variation when executing existing CLG trees
	//
	var members []interface{}
	for _, behaviorID := range behaviorIDs {
		for _, networkPayload := range queue {
			// At this point there is only one source given. That is the CLG that
			// forwarded the current network payload to here. If this is not the case,
			// we return an error.
			sources := networkPayload.GetSources()
			if len(sources) != 1 {
				return nil, nil, maskAnyf(invalidInterfaceError, "there must be one source")
			}
			if behaviorID == string(sources[0]) {
				members = append(members, networkPayload)
				continue
			}
		}
	}

	if len(behaviorIDs) != len(members) {
		// The received network payloads have not been able to satisfy the interface
		// of the connections we looked up. These represent the interface of the
		// requested CLG. There is no match. Thus we return an error.
		return nil, nil, maskAnyf(invalidInterfaceError, "connections must match")
	}

	// The received network payloads have been able to satisfy the interface
	// of the connections we looked up. We merge the matching payloads together
	// and filter them from the queue.
	newPayload, err := membersToPayload(ctx, members)
	if err != nil {
		return nil, nil, maskAny(err)
	}
	newQueue, err := filterMembersFromQueue(members, queue)
	if err != nil {
		return nil, nil, maskAny(err)
	}

	return newPayload, newQueue, nil
}

// payloadFromPermutations tries to find a combination of payloads which
// together are able to fulfill the interface of the requested CLG. The first
// combination of payloads found, which match the interface of the requested CLG
// will be returned in one merged network payload. The returned list of network
// payloads will not contain any of the network payloads merged.
func (n *network) payloadFromPermutations(ctx spec.Context, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error) {
	clgName, ok := ctx.GetCLGName()
	if !ok {
		return nil, nil, maskAnyf(invalidCLGNameError, "must not be empty")
	}
	CLG, err := n.clgByName(clgName)
	if err != nil {
		return nil, nil, maskAny(err)
	}
	inputTypes := CLG.GetInputTypes()

	// Prepare the permutation list to find out which combination of payloads
	// satisfies the requested CLG's interface.
	newConfig := permutation.DefaultListConfig()
	newConfig.MaxGrowth = len(inputTypes)
	newConfig.Values = queueToValues(queue)
	newPermutationList, err := permutation.NewList(newConfig)
	if err != nil {
		return nil, nil, maskAny(err)
	}

	for {
		err := n.Factory().Permutation().MapTo(newPermutationList)
		if err != nil {
			return nil, nil, maskAny(err)
		}

		// Check if the given payload satisfies the requested CLG's interface.
		members := newPermutationList.GetMembers()
		types, err := membersToTypes(members)
		if err != nil {
			return nil, nil, maskAny(err)
		}
		if reflect.DeepEqual(types, inputTypes) {
			newPayload, err := membersToPayload(ctx, members)
			if err != nil {
				return nil, nil, maskAny(err)
			}
			newQueue, err := filterMembersFromQueue(members, queue)
			if err != nil {
				return nil, nil, maskAny(err)
			}

			return newPayload, newQueue, nil
		}

		err = n.Factory().Permutation().PermuteBy(newPermutationList, 1)
		if err != nil {
			// Note that also an error is thrown when the maximum growth of the
			// permutation list was reached.
			return nil, nil, maskAny(err)
		}
	}
}

// helper

func containsNetworkPayload(list []spec.NetworkPayload, item spec.NetworkPayload) bool {
	for _, p := range list {
		if p.GetID() == item.GetID() {
			return true
		}
	}

	return false
}

func filterMembersFromQueue(members []interface{}, queue []spec.NetworkPayload) ([]spec.NetworkPayload, error) {
	var memberPayloads []spec.NetworkPayload
	for _, m := range members {
		payload, ok := m.(spec.NetworkPayload)
		if !ok {
			return nil, maskAnyf(invalidInterfaceError, "member must be spec.NetworkPayload")
		}
		memberPayloads = append(memberPayloads, payload)
	}

	var newQueue []spec.NetworkPayload
	for _, queuedPayload := range queue {
		if containsNetworkPayload(memberPayloads, queuedPayload) {
			continue
		}

		newQueue = append(newQueue, queuedPayload)
	}

	return newQueue, nil
}

func membersToPayload(ctx spec.Context, members []interface{}) (spec.NetworkPayload, error) {
	var args []reflect.Value
	var sources []spec.ObjectID

	behaviorID, ok := ctx.GetBehaviorID()
	if !ok {
		return nil, maskAnyf(invalidBehaviorIDError, "must not be empty")
	}

	for _, m := range members {
		payload, ok := m.(spec.NetworkPayload)
		if !ok {
			return nil, maskAnyf(invalidInterfaceError, "member must be spec.NetworkPayload")
		}

		for _, v := range payload.GetArgs() {
			args = append(args, v)
		}

		sources = append(sources, payload.GetSources()...)
	}

	newPayloadConfig := api.DefaultNetworkPayloadConfig()
	newPayloadConfig.Args = args
	newPayloadConfig.Context = ctx
	newPayloadConfig.Destination = spec.ObjectID(behaviorID)
	newPayloadConfig.Sources = sources
	newPayload, err := api.NewNetworkPayload(newPayloadConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newPayload, nil
}

func membersToTypes(members []interface{}) ([]reflect.Type, error) {
	var types []reflect.Type

	for _, m := range members {
		payload, ok := m.(spec.NetworkPayload)
		if !ok {
			return nil, maskAnyf(invalidInterfaceError, "member must be spec.NetworkPayload")
		}

		for _, v := range payload.GetArgs() {
			types = append(types, v.Type())
		}
	}

	return types, nil
}

func queueToValues(queue []spec.NetworkPayload) []interface{} {
	var values []interface{}

	for _, p := range queue {
		values = append(values, p)
	}

	return values
}
