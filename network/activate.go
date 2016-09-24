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

func (n *network) payloadFromConnections(CLG spec.CLG, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error) {
	// Check the queue is valid for our operations.
	if len(queue) < 1 {
		return nil, nil, maskAnyf(invalidNetworkPayloadError, "must not be empty")
	}

	// Lookup the behavior ID from the context of any payload of the queue. We can
	// simply take the first payload because the behavior ID should always be the
	// same. It is set within Forward, where signals are sent to further CLGs,
	// until the payload is passed to here.
	ctx, err := queue[0].GetContext()
	if err != nil {
		return nil, nil, maskAny(err)
	}
	behaviorID := ctx.GetBehaviorID()
	if behaviorID == "" {
		return nil, nil, maskAnyf(invalidBehaviorIDError, "must not be empty")
	}

	// Fetch the available behavior IDs which are known to be useful connections
	// during the activation of the requested CLG. The payloads sent by the CLGs
	// being fetched here are useful because, in the past, they have already been
	// helpful within the current CLG tree.
	//
	// TODO who writes this data?
	behaviorIDsKey := key.NewCLGKey("behavior-id:%s:activate-behavior-ids", behaviorID)
	list, err := n.Storage().General().Get(behaviorIDsKey)
	if err != nil {
		return nil, nil, maskAny(err)
	}
	behaviorIDs := strings.Split(list, ",")

	// Check if there is a network payload for each behavior ID we found in the
	// storage.
	var members []interface{}
	for _, behaviorID := range behaviorIDs {
		for _, networkPayload := range queue {
			// At this point there is only one source given, the CLG that forwarded
			// the current network payload to here.
			sourceID := string(networkPayload.GetSources()[0])
			if behaviorID == string(sourceID) {
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
	newPayload, err := membersToPayload(members)
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
func (n *network) payloadFromPermutations(CLG spec.CLG, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error) {
	// Prepare the permutation list to find out which combination of payloads
	// satisfies the requested CLG's interface.
	newConfig := permutation.DefaultListConfig()
	newConfig.MaxGrowth = len(CLG.GetInputTypes())
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
		if reflect.DeepEqual(types, CLG.GetInputTypes()) {
			newPayload, err := membersToPayload(members)
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

func membersToPayload(members []interface{}) (spec.NetworkPayload, error) {
	var ctxAdded bool
	var args []reflect.Value
	var destination spec.ObjectID
	var sources []spec.ObjectID

	for _, m := range members {
		payload, ok := m.(spec.NetworkPayload)
		if !ok {
			return nil, maskAnyf(invalidInterfaceError, "member must be spec.NetworkPayload")
		}

		if !ctxAdded {
			ctx, err := payload.GetContext()
			if !ok {
				return nil, maskAny(err)
			}
			args = append(args, reflect.ValueOf(ctx))
			destination = payload.GetDestination()
			ctxAdded = true
		}

		for _, v := range payload.GetArgs()[1:] {
			args = append(args, v)
			sources = append(sources, payload.GetSources()...)
		}
	}

	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = args
	newNetworkPayloadConfig.Destination = destination
	newNetworkPayloadConfig.Sources = sources
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newNetworkPayload, nil
}

func membersToTypes(members []interface{}) ([]reflect.Type, error) {
	var types []reflect.Type
	var ctxAdded bool

	for _, m := range members {
		payload, ok := m.(spec.NetworkPayload)
		if !ok {
			return nil, maskAnyf(invalidInterfaceError, "member must be spec.NetworkPayload")
		}

		if !ctxAdded {
			ctx, err := payload.GetContext()
			if !ok {
				return nil, maskAny(err)
			}
			types = append(types, reflect.TypeOf(ctx))
			ctxAdded = true
		}

		for _, v := range payload.GetArgs()[1:] {
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
