package network

import (
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
	"github.com/xh3b4sd/anna/clg/read-separator"
	"github.com/xh3b4sd/anna/clg/split-features"
	"github.com/xh3b4sd/anna/clg/subtract"
	"github.com/xh3b4sd/anna/clg/sum"
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
		CLGs[ID].SetFactoryCollection(n.FactoryCollection)
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

					// In case the current queue exeeds a certain amount of payloads, it
					// is unlikely that the queue is going to be helpful when growing any
					// further. Thus we cut the queue at some point.
					//
					// TODO the limit is hardcoded and should be configured by the neural
					// network itself.
					if len(queue) > 10 {
						queue = queue[:10]
					}

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

// helper

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
		readseparator.MustNew(),
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
