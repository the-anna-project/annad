// Package network implements spec.Network to provide a neural network based on
// dynamic and self improving CLG execution. The network provides input and
// output channels. When input is received it is injected into the neural
// communication. The following neural activity calculates output which is
// streamed through the output channel back to the requestor.
package network

import (
	"reflect"
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectType represents the object type of the network object. This is used
	// e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "network"
)

// Config represents the configuration used to create a new network object.
type Config struct {
	// Dependencies.
	Log                spec.Log
	PermutationFactory spec.PermutationFactory
	Storage            spec.Storage
	TextInput          chan spec.TextRequest
	TextOutput         chan spec.TextResponse

	// Settings.

	// Delay causes each CLG execution to be delayed. This value represents a
	// default value. The actually used value is optimized based on experience
	// and learning.
	// TODO implement
	Delay time.Duration
}

// DefaultConfig provides a default configuration to create a new network
// object by best effort.
func DefaultConfig() Config {
	newPermutationFactory, err := permutation.NewFactory(permutation.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

	newStorage, err := memory.NewStorage(memory.DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		Log:                log.NewLog(log.DefaultConfig()),
		PermutationFactory: newPermutationFactory,
		Storage:            newStorage,
		TextInput:          make(chan spec.TextRequest, 1000),
		TextOutput:         make(chan spec.TextResponse, 1000),
	}

	return newConfig
}

// New creates a new configured network object.
func New(config Config) (spec.Network, error) {
	newNetwork := &network{
		Config: config,

		BootOnce:     sync.Once{},
		CLGIDs:       map[string]spec.ObjectID{},
		CLGs:         newCLGs(),
		ID:           id.MustNew(),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectType,
	}

	if newNetwork.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newNetwork.Storage == nil {
		return nil, maskAnyf(invalidConfigError, "storage must not be empty")
	}
	if newNetwork.TextInput == nil {
		return nil, maskAnyf(invalidConfigError, "text input channel must not be empty")
	}
	if newNetwork.TextOutput == nil {
		return nil, maskAnyf(invalidConfigError, "text output channel must not be empty")
	}

	newNetwork.Log.Register(newNetwork.GetType())

	return newNetwork, nil
}

type network struct {
	Config

	BootOnce sync.Once

	// CLGIDs provides a mapping of CLG names pointing to their corresponding CLG
	// ID.
	CLGIDs map[string]spec.ObjectID

	CLGs         map[spec.ObjectID]spec.CLG // TODO there is probably no reason to index the CLGs like this
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (n *network) Activate(CLG spec.CLG, payload spec.NetworkPayload, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Activate")

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
		if permutation.IsMaxGrowthReached(err) {
			// We cannot permute the given list anymore. So far the requested CLG's
			// interface could not be satisfied.
			return nil, nil, maskAnyf(invalidInterfaceError, "types must match")
		}
	}
}

func (n *network) Boot() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Boot")

	n.BootOnce.Do(func() {
		n.CLGs = n.configureCLGs(n.CLGs)
		n.CLGIDs = n.mapCLGIDs(n.CLGs)

		go func() {
			err := n.Listen()
			if err != nil {
				n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
			}
		}()
	})
}

func (n *network) Calculate(CLG spec.CLG, payload spec.NetworkPayload) (spec.NetworkPayload, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Calculate")

	calculatedPayload, err := CLG.Calculate(payload)
	if err != nil {
		return nil, maskAny(err)
	}

	return calculatedPayload, nil
}

// Forward to other CLGs
// Split the neural connection path
// TODO
func (n *network) Forward(CLG spec.CLG, payload spec.NetworkPayload) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Forward")

	// Check if the given context provides a CLG tree ID.

	clgTreeID := ""
	if clgTreeID == "" {
		// create new
	} else {
		// lookup existing
	}

	// for _, r := range requests {
	// 	err := n.Send(r) send does not exist anymore
	// 	if err != nil {
	// 		return maskAny(err)
	// 	}
	// }

	return nil
}

func (n *network) Listen() error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Listen")

	// Listen on TextInput from the outside to receive text requests.
	CLG, err := n.clgByName("input")
	if err != nil {
		return maskAny(err)
	}
	go func() {
		clgID := CLG.GetID()
		networkID := n.GetID()
		clgChannel := CLG.GetInputChannel()

		for {
			select {
			case textRequest := <-n.TextInput:
				// TODO this is only used for testing to bypass the neural network and
				// directly respond with the received input.

				//newTextResponseConfig := api.DefaultTextResponseConfig()
				//newTextResponseConfig.Output = textRequest.GetInput()
				//newTextResponse, err := api.NewTextResponse(newTextResponseConfig)
				//if err != nil {
				//	n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
				//}
				//n.TextOutput <- newTextResponse
				//continue

				// TODO

				newPayloadConfig := api.DefaultNetworkPayloadConfig()
				newPayloadConfig.Args = []reflect.Value{reflect.ValueOf(context.Background()), reflect.ValueOf(textRequest.GetInput())}
				newPayloadConfig.Destination = clgID
				newPayloadConfig.Sources = []spec.ObjectID{networkID}
				newPayload, err := api.NewNetworkPayload(newPayloadConfig)
				if err != nil {
					n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
				}

				clgChannel <- newPayload
			}
		}
	}()

	// Make all CLGs listening in their specific input channel.
	for ID, CLG := range n.CLGs {
		go func(ID spec.ObjectID, CLG spec.CLG) {
			var queue []spec.NetworkPayload
			clgChannel := CLG.GetInputChannel()

			for {
				select {
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
							n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
						}
						queue = newQueue

						// Calculate based on the CLG's implemented business logic.
						calculatedPayload, err := n.Calculate(CLG, newPayload)
						if err != nil {
							n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
						}

						// Forward to other CLG's, if necessary.
						err = n.Forward(CLG, calculatedPayload)
						if err != nil {
							n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
						}

						// Return the calculated output to the requesting client, if the
						// current CLG is the output CLG.
						if CLG.GetName() == "output" {
							var output string
							// The first argument is always a context, which is ignored,
							// because it only serves internal purposes.
							for _, v := range calculatedPayload.GetArgs()[1:] {
								output += v.String()
							}
							newTextResponseConfig := api.DefaultTextResponseConfig()
							newTextResponseConfig.Output = output
							newTextResponse, err := api.NewTextResponse(newTextResponseConfig)
							if err != nil {
								n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
							}
							n.TextOutput <- newTextResponse
						}
					}(payload)
				}
			}
		}(ID, CLG)
	}

	return nil
}

func (n *network) Shutdown() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Shutdown")

	n.ShutdownOnce.Do(func() {
		//
	})
}
