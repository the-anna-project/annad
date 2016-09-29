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
			// TODO queue calculation might belong to Activator
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
