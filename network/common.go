package network

import (
	"reflect"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/clg/divide"
	"github.com/xh3b4sd/anna/clg/greater"
	"github.com/xh3b4sd/anna/clg/input"
	"github.com/xh3b4sd/anna/clg/is-between"
	"github.com/xh3b4sd/anna/clg/is-greater"
	"github.com/xh3b4sd/anna/clg/multiply"
	//"github.com/xh3b4sd/anna/clg/output"
	"github.com/xh3b4sd/anna/clg/pair-syntactic"
	"github.com/xh3b4sd/anna/clg/read-information-id"
	"github.com/xh3b4sd/anna/clg/split-features"
	"github.com/xh3b4sd/anna/clg/subtract"
	"github.com/xh3b4sd/anna/clg/sum"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

// receiver

func (n *network) clgByName(name string) (spec.CLG, error) {
	ID, ok := n.CLGIDs[name]
	if !ok {
		return nil, maskAnyf(clgNotFoundError, "name: %s", name)
	}
	CLG, ok := n.CLGs[ID]
	if !ok {
		return nil, maskAnyf(clgNotFoundError, "ID: %s", ID)
	}

	return CLG, nil
}

func (n *network) configureCLGs(CLGs map[spec.ObjectID]spec.CLG) map[spec.ObjectID]spec.CLG {
	for ID := range CLGs {
		CLGs[ID].SetIDFactory(n.IDFactory)
		CLGs[ID].SetLog(n.Log)
		CLGs[ID].SetStorageCollection(n.StorageCollection)
	}

	return CLGs
}

func (n *network) findConnections(ctx spec.Context, payload spec.NetworkPayload) ([]string, error) {
	var behaviorIDs []string

	behaviorID := ctx.GetBehaviorID()
	if behaviorID == "" {
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

func (n *network) listenCLGs() {
	// Make all CLGs listening in their specific input channel.
	for ID, CLG := range n.CLGs {
		go func(ID spec.ObjectID, CLG spec.CLG) {
			var queue []spec.NetworkPayload
			clgChannel := CLG.GetInputChannel()

			for {
				select {
				case <-n.Closer:
					break
				case payload := <-clgChannel:

					go func(payload spec.NetworkPayload) {
						// Activate if the CLG's interface is satisfied by the given
						// network payload.
						newPayload, newQueue, err := n.Activate(CLG, payload, queue)
						if IsInvalidInterface(err) {
							// The interface of the requested CLG was not fulfilled. We
							// continue listening for the next network payload without doing
							// any work.
							return
						} else if err != nil {
							n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
						}
						queue = newQueue

						// Calculate based on the CLG's implemented business logic.
						calculatedPayload, err := n.Calculate(CLG, newPayload)
						if err != nil {
							n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
						}

						// Forward to other CLG's, if necessary.
						err = n.Forward(CLG, calculatedPayload)
						if err != nil {
							n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
						}

						// Return the calculated output to the requesting client, if the
						// current CLG is the output CLG.
						if CLG.GetName() == "output" {
							newTextResponseConfig := api.DefaultTextResponseConfig()
							newTextResponseConfig.Output = calculatedPayload.String()
							newTextResponse, err := api.NewTextResponse(newTextResponseConfig)
							if err != nil {
								n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
							}
							n.TextOutput <- newTextResponse
						}
					}(payload)
				}
			}
		}(ID, CLG)
	}
}

func (n *network) mapCLGIDs(CLGs map[spec.ObjectID]spec.CLG) map[string]spec.ObjectID {
	clgIDs := map[string]spec.ObjectID{}

	for ID, CLG := range CLGs {
		clgIDs[CLG.GetName()] = ID
	}

	return clgIDs
}

func (n *network) permutePayload(CLG spec.CLG, payload spec.NetworkPayload, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error) {
	queue = append(queue, payload)

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
		err := n.PermutationFactory.MapTo(newPermutationList)
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

			// In case the current queue exeeds the interface of the requested CLG, it is
			// trimmed to cause a more strict behaviour of the neural network.
			if len(newPermutationList.GetValues()) > len(CLG.GetInputTypes()) {
				newQueue = newQueue[1:]
			}

			return newPayload, newQueue, nil
		}

		err = n.PermutationFactory.PermuteBy(newPermutationList, 1)
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

func newCLGs() map[spec.ObjectID]spec.CLG {
	newList := []spec.CLG{
		divide.MustNew(),
		input.MustNew(),
		divide.MustNew(),
		greater.MustNew(),
		input.MustNew(),
		isbetween.MustNew(),
		isgreater.MustNew(),
		multiply.MustNew(),
		//output.MustNew(),
		pairsyntactic.MustNew(),
		readinformationid.MustNew(),
		splitfeatures.MustNew(),
		subtract.MustNew(),
		sum.MustNew(),
	}

	newCLGs := map[spec.ObjectID]spec.CLG{}

	for _, CLG := range newList {
		newCLGs[CLG.GetID()] = CLG
	}

	return newCLGs
}

func queueToValues(queue []spec.NetworkPayload) []interface{} {
	var values []interface{}

	for _, p := range queue {
		values = append(values, p)
	}

	return values
}
