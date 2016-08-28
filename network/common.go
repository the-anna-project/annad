package network

import (
	"reflect"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/clg/divide"
	"github.com/xh3b4sd/anna/clg/input"
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
		CLGs[ID].SetStorage(n.Storage)
	}

	return CLGs
}

func (n *network) listenCLGs() {
	// Make all CLGs listening in their specific input channel.
	for ID, CLG := range n.CLGs {
		go func(ID spec.ObjectID, CLG spec.CLG) {
			var queue []spec.NetworkPayload
			clgChannel := CLG.GetInputChannel()

			for {
				payload := <-clgChannel

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
