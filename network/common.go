package network

import (
	"reflect"
	"sync/atomic"

	"github.com/xh3b4sd/anna/api"
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

func (n *network) extractMatchingInputRequests(queue []spec.InputRequest, desired []reflect.Type) (bool, []spec.InputRequest, []spec.InputRequest, error) {
	permutationList, err := inputRequestsToPermutationList(queue, desired)
	if err != nil {
		return false, nil, nil, maskAny(err)
	}

	var execute bool
	var matching []spec.InputRequest
	var newQueue []spec.InputRequest
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
		matching, newQueue, err = n.filterInputRequests(permutationList, queue)
		if err != nil {
			return false, nil, nil, maskAny(err)
		}
		break
	}

	return execute, matching, newQueue, nil
}

// TODO
func (n *network) filterInputRequests(permutationList spec.PermutationList, queue []spec.InputRequest) ([]spec.InputRequest, []spec.InputRequest, error) {
	err := n.PermutationFactory.MapTo(permutationList)
	if err != nil {
		return nil, nil, maskAny(err)
	}
	members := permutationList.GetMembers()

	var matching []spec.InputRequest
	for _, m := range members {
		request, ok := m.(spec.InputRequest)
		if !ok {
			return nil, nil, maskAnyf(invalidInterfaceError, "invalid type for permutation list member")
		}
		matching = append(matching, request)
	}

	// TODO test
	var newQueue []spec.InputRequest
	matchingSeen := map[*spec.InputRequest]struct{}{}
	for _, r := range queue {
		_, ok := matchingSeen[&r]
		if containsInputRequest(matching, r) && !ok {
			matchingSeen[&r] = struct{}{}
			continue
		}
		newQueue = append(newQueue, r)
	}

	return matching, newQueue, nil
}

func (n *network) getGatewayListener() func(newSignal spec.Signal) (spec.Signal, error) {
	newListener := func(newSignal spec.Signal) (spec.Signal, error) {
		newImpulse, err := n.NewImpulse(newSignal.GetInput().(api.CoreRequest))
		if err != nil {
			return nil, maskAny(err)
		}

		// Increment the impulse count to track how many impulses are processed
		// inside the core network.
		atomic.AddInt64(&n.ImpulsesInProgress, 1)
		newImpulse, err = n.Trigger(newImpulse)
		// Decrement the impulse count once all hard work is done. Note that this
		// is important to be done before the error handling of Core.Trigger to
		// ensure the impulse count is properly decreased.
		atomic.AddInt64(&n.ImpulsesInProgress, -1)

		if err != nil {
			return nil, maskAny(err)
		}

		output := newImpulse.GetOutput()
		newSignal.SetOutput(output)

		return newSignal, nil
	}

	return newListener
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
		request, ok := m.(spec.InputRequest)
		if !ok {
			return nil, maskAnyf(invalidInterfaceError, "invalid type for permutation list member")
		}

		for _, v := range request.Inputs {
			types = append(types, v.Type())
		}
	}

	return types, nil
}

// private

func containsInputRequest(list []spec.InputRequest, item spec.InputRequest) bool {
	for _, r := range list {
		if equalInputRequest(r, item) {
			return true
		}
	}

	return false
}

func equalInputRequest(a spec.InputRequest, b spec.InputRequest) bool {
	if a.Source != b.Source {
		return false
	}
	if a.Destination != b.Destination {
		return false
	}
	if !reflect.DeepEqual(a.Inputs, b.Inputs) {
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

func inputRequestsToPermutationList(queue []spec.InputRequest, desired []reflect.Type) (spec.PermutationList, error) {
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

// joinRequestInputs joins all input lists of the given input requests
// together. The order of the joined inputs equals the order of the given input
// requests.
func joinRequestInputs(inputRequests []spec.InputRequest) []reflect.Value {
	var inputs []reflect.Value

	for _, ir := range inputRequests {
		inputs = append(inputs, ir.Inputs...)
	}

	return inputs
}

type clgScope struct {
	CLG    spec.CLG
	Input  chan spec.InputRequest
	Output chan spec.OutputResponse
}

func newCLGs() map[spec.ObjectID]clgScope {
	newList := []spec.CLG{
		divide.MustNew(),
	}

	newCLGs := map[spec.ObjectID]clgScope{}

	for _, CLG := range newList {
		newCLGs[CLG.GetID()] = clgScope{
			CLG:    CLG,
			Input:  make(chan spec.InputRequest, 10),
			Output: make(chan spec.OutputResponse, 10),
		}
	}

	return newCLGs
}

func prepareInput(imp spec.Impulse, source, destination spec.ObjectID) spec.InputRequest {
	request := spec.InputRequest{
		Source:      source,
		Destination: destination,
		Inputs:      []reflect.Value{reflect.ValueOf(imp), reflect.ValueOf(imp.GetInput())},
	}

	return request
}

func prepareOutput(response spec.OutputResponse) (spec.Impulse, error) {
	if len(response.Outputs) == 0 {
		return nil, maskAnyf(invalidInterfaceError, "output must not be empty")
	}

	imp, ok := response.Outputs[0].Interface().(spec.Impulse)
	if !ok {
		return nil, maskAnyf(invalidInterfaceError, "impulse must be first")
	}

	var output string
	for _, v := range response.Outputs[1:] {
		output += v.String()
	}
	imp.SetOutput(output)

	return imp, nil
}

func valuesToTypes(values []reflect.Value) []reflect.Type {
	var types []reflect.Type

	for _, v := range values {
		types = append(types, v.Type())
	}

	return types
}
