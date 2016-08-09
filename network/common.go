package network

import (
	"reflect"

	"github.com/xh3b4sd/anna/clg/divide"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/spec"
)

// receiver

func (n *network) configureCLGs(CLGs map[spec.ObjectID]clgScope) map[spec.ObjectID]clgScope {
	for ID := range CLGs {
		CLGs[ID].CLG.SetLog(n.Log)
		CLGs[ID].CLG.SetStorage(n.Storage)
	}

	return CLGs
}

func (n *network) extractMatchingNetworkPayload(queue []spec.NetworkPayload, desired []reflect.Type) (bool, []spec.NetworkPayload, []spec.NetworkPayload, error) {
	permutationList, err := inputRequestsToPermutationList(queue, desired)
	if err != nil {
		return false, nil, nil, maskAny(err)
	}

	var execute bool
	var matching []spec.NetworkPayload
	var newQueue []spec.NetworkPayload
	for {
		types, err := n.permutationListToTypes(permutationList)
		if err != nil {
			return false, nil, nil, maskAny(err)
		}
		if !equalTypes(types, desired) { // to []reflect.Types{...}
			err := n.PermutationFactory.PermuteBy(permutationList, 1)
			if permutation.IsMaxGrowthReached(err) {
				// We cannot permute the given list anymore. There is nothing useful
				// for a CLG. So we only return signaling not to execute this CLG with
				// the provided input.
				return false, nil, nil, nil
			} else if err != nil {
				return false, nil, nil, maskAny(err)
			}
			continue
		}

		execute = true
		matching, newQueue, err = n.filterNetworkPayloads(permutationList, queue)
		if err != nil {
			return false, nil, nil, maskAny(err)
		}
		break
	}

	return execute, matching, newQueue, nil
}

// TODO
func (n *network) filterNetworkPayloads(permutationList spec.PermutationList, queue []spec.NetworkPayload) ([]spec.NetworkPayload, []spec.NetworkPayload, error) {
	err := n.PermutationFactory.MapTo(permutationList)
	if err != nil {
		return nil, nil, maskAny(err)
	}
	members := permutationList.GetMembers()

	var matching []spec.NetworkPayload
	for _, m := range members {
		payload, ok := m.(spec.NetworkPayload)
		if !ok {
			return nil, nil, maskAnyf(invalidInterfaceError, "invalid type for permutation list member")
		}
		matching = append(matching, payload)
	}

	// TODO test
	var newQueue []spec.NetworkPayload
	matchingSeen := map[*spec.NetworkPayload]struct{}{}
	for _, r := range queue {
		_, ok := matchingSeen[&r]
		if containsNetworkPayload(matching, r) && !ok {
			matchingSeen[&r] = struct{}{}
			continue
		}
		newQueue = append(newQueue, r)
	}

	return matching, newQueue, nil
}

func (n *network) mapCLGIDs(CLGs map[spec.ObjectID]clgScope) map[string]spec.ObjectID {
	var clgIDs map[string]spec.ObjectID

	for ID, clgScope := range CLGs {
		clgIDs[clgScope.CLG.GetName()] = ID
	}

	return clgIDs
}

func (n *network) permutationListToTypes(permutationList spec.PermutationList) ([]reflect.Type, error) {
	err := n.PermutationFactory.MapTo(permutationList)
	if err != nil {
		return nil, maskAny(err)
	}
	members := permutationList.GetMembers()

	var types []reflect.Type

	for _, m := range members {
		payload, ok := m.(spec.NetworkPayload)
		if !ok {
			return nil, maskAnyf(invalidInterfaceError, "invalid type for permutation list member")
		}

		for _, v := range payload.Args {
			types = append(types, v.Type())
		}
	}

	return types, nil
}

// private

func containsNetworkPayload(list []spec.NetworkPayload, item spec.NetworkPayload) bool {
	for _, r := range list {
		if equalNetworkPayload(r, item) {
			return true
		}
	}

	return false
}

func equalNetworkPayload(a spec.NetworkPayload, b spec.NetworkPayload) bool {
	if !reflect.DeepEqual(a.Args, b.Args) {
		return false
	}
	if a.Destination != b.Destination {
		return false
	}
	if !reflect.DeepEqual(a.Sources, b.Sources) {
		return false
	}

	return true
}

// equalTypes checks if the given two lists are equal in their ordered values.
func equalTypes(a, b []reflect.Type) bool {
	for _, at := range a {
		for _, bt := range b {
			if at.String() != bt.String() {
				return false
			}
		}
	}

	return true
}

func inputRequestsToPermutationList(queue []spec.NetworkPayload, desired []reflect.Type) (spec.PermutationList, error) {
	var values []interface{}
	for _, ir := range queue {
		values = append(values, ir)
	}

	newConfig := permutation.DefaultListConfig()
	newConfig.MaxGrowth = len(desired)
	newConfig.Values = values

	permutationList, err := permutation.NewList(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return permutationList, nil
}

// mergeNetworkPayloads joins the given network payloads to one new network
// payload. The order of the joined inputs equals the order of the given
// network payload. In the inputs of the network payloads the first element is
// always a context. When merging the network payloads the first context of the
// first inputs list is used to be applied to the new unified network payload.
// All contexts of the given network payloads should be equal anyway anyway.
// TODO
// func mergeNetworkPayloads(payloads []spec.NetworkPayload) (spec.NetworkPayload, error) {
// 	var payload spec.NetworkPayload
// 	var imp spec.Impulse
// 	var err error
//
// 	for _, ir := range payloads {
// 		if imp == nil {
// 			imp, err = argsToImpulse(ir.Args)
// 			if err != nil {
// 				return nil, maskAny(err)
// 			}
//
// 			payload.Destination = ir.Destination
// 		}
// 		// Note we ignore the first argument which should only ever be the same
// 		// impulse across all network payloads, which we already tracked above.
// 		payload.Args = append(payload.Args, ir.Args[1:]...)
// 		payload.Sources = append(payload.Sources, ir.Sources...)
// 	}
//
// 	return inputs, nil
// }

type clgScope struct {
	CLG   spec.CLG
	Input chan spec.NetworkPayload
}

func newCLGs() map[spec.ObjectID]clgScope {
	newList := []spec.CLG{
		divide.MustNew(),
	}

	newCLGs := map[spec.ObjectID]clgScope{}

	for _, CLG := range newList {
		newCLGs[CLG.GetID()] = clgScope{
			CLG:   CLG,
			Input: make(chan spec.NetworkPayload, 10),
		}
	}

	return newCLGs
}

func valuesToTypes(values []reflect.Value) []reflect.Type {
	var types []reflect.Type

	for _, v := range values {
		types = append(types, v.Type())
	}

	return types
}
