package network

import (
	"reflect"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/clg/output"
	"github.com/xh3b4sd/anna/context"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

// receiver

func (n *network) listenCLGs() {
	// Make all CLGs listening in their specific input channel.
	for ID, CLG := range n.CLGs {
		go func(ID spec.ObjectID, CLG spec.CLG) {
			var queue []spec.NetworkPayload
			queueBuffer := len(CLG.GetInputTypes()) + 1
			inputChannel := CLG.GetInputChannel()

			for {
				select {
				case <-n.Closer:
					return
				case payload := <-inputChannel:
					// In case the current queue exeeds a certain amount of payloads, it
					// is unlikely that the queue is going to be helpful when growing any
					// further. Thus we cut the queue at some point beyond the interface
					// capabilities of the requested CLG.
					queue = append(queue, payload)
					if len(queue) > queueBuffer {
						queue = queue[1:]
					}

					// We pass the context of the received payload separately because we
					// merged the payload to the queue and have no reliable and easy way
					// to make the current context available otherwise.
					ctx := payload.GetContext()

					// Setting the CLG name to the current context helps us looking up the
					// CLG later on when needed. That way we do not need to pass it
					// through all the way even though it is not required to be available
					// in that many places.
					ctx.SetCLGName(CLG.GetName())

					go func(ctx spec.Context, queue []spec.NetworkPayload) {
						// Activate if the CLG's interface is satisfied by the given
						// network payload.
						newPayload, newQueue, err := n.Activate(ctx, queue)
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
						calculatedPayload, err := n.Calculate(ctx, newPayload)
						if output.IsExpectationNotMet(err) {
							n.Log.WithTags(spec.Tags{C: nil, L: "W", O: n, V: 7}, "%#v", maskAny(err))

							err = n.forwardInputCLG(ctx, calculatedPayload)
							if err != nil {
								n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
							}

							return
						} else if err != nil {
							n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
						}

						// Forward to other CLG's, if necessary.
						err = n.Forward(ctx, calculatedPayload)
						if err != nil {
							n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
						}
					}(ctx, queue)
				}
			}
		}(ID, CLG)
	}
}

func (n *network) listenInputCLG() {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Listen")

	// Listen on TextInput from the outside to receive text requests.
	CLG, err := n.clgByName("input")
	if err != nil {
		n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
	}

	networkID := n.GetID()
	clgChannel := CLG.GetInputChannel()

	for {
		select {
		case <-n.Closer:
			break
		case textRequest := <-n.TextInput:
			go func(textRequest spec.TextRequest) {
				// This should only be used for testing to bypass the neural network
				// and directly respond with the received input.
				if textRequest.GetEcho() {
					newTextResponseConfig := api.DefaultTextResponseConfig()
					newTextResponseConfig.Output = textRequest.GetInput()
					newTextResponse, err := api.NewTextResponse(newTextResponseConfig)
					if err != nil {
						n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
					}
					n.TextOutput <- newTextResponse
					return
				}

				// Create new IDs for the new CLG tree and the input CLG.
				clgTreeID, err := n.Factory().ID().New()
				if err != nil {
					n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
					return
				}
				behaviorID, err := n.Factory().ID().New()
				if err != nil {
					n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
					return
				}

				// Write the new CLG tree ID to reference the input CLG ID and add the CLG
				// tree ID to the new context.
				firstBehaviorIDKey := key.NewCLGKey("clg-tree-id:%s:first-behavior-id", clgTreeID)
				err = n.Storage().General().Set(firstBehaviorIDKey, string(behaviorID))
				if err != nil {
					n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
					return
				}

				// Prepare the context and a unique behavior ID for the input CLG.
				ctx := context.MustNew()
				ctx.SetBehaviorID(string(behaviorID))
				ctx.SetCLGName("input")
				ctx.SetCLGTreeID(string(clgTreeID))
				ctx.SetExpectation(textRequest.GetExpectation())
				ctx.SetSessionID(textRequest.GetSessionID())

				// We transform the received input to a network payload to have a
				// conventional data structure within the neural network. Note the
				// following details.
				//
				//     The list of arguments always contains a context as first argument.
				//
				//     Destination is always the behavior ID of the input CLG, because
				//     this one is the connecting building block to other CLGs within the
				//     neural network. This behavior ID is always a new one, because it
				//     will eventually be part of a completely new CLG tree within the
				//     connection space.
				//
				//     Sources is here only the individual network ID to have at least
				//     any reference of origin.
				//
				payloadConfig := api.DefaultNetworkPayloadConfig()
				payloadConfig.Args = []reflect.Value{reflect.ValueOf(textRequest.GetInput())}
				payloadConfig.Context = ctx
				payloadConfig.Destination = behaviorID
				payloadConfig.Sources = []spec.ObjectID{networkID}
				newPayload, err := api.NewNetworkPayload(payloadConfig)
				if err != nil {
					n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
					return
				}

				// Send the new network payload to the input CLG.
				clgChannel <- newPayload
			}(textRequest)
		}
	}
}
