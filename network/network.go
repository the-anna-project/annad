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
	// ObjectTypeNetwork represents the object type of the network object. This
	// is used e.g. to register itself to the logger.
	ObjectTypeNetwork spec.ObjectType = "network"
)

// Config represents the configuration used to create a new network object.
type Config struct {
	// Dependencies.
	Log                spec.Log
	PermutationFactory spec.PermutationFactory
	Storage            spec.Storage
	TextInput          chan api.TextRequest
	TextOutput         chan api.TextResponse

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
		TextInput:          make(chan api.TextRequest, 1000),
		TextOutput:         make(chan api.TextResponse, 1000),
	}

	return newConfig
}

// New creates a new configured network object.
func New(config Config) (spec.Network, error) {
	newNetwork := &network{
		Config: config,

		BootOnce:     sync.Once{},
		CLGs:         newCLGs(),
		ID:           id.MustNew(),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeNetwork,
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
	CLGs     map[spec.ObjectID]spec.CLG

	// CLGIDs provides a mapping of CLG names pointing to their corresponding CLG
	// ID.
	CLGIDs       map[string]spec.ObjectID
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (n *network) Activate(clgID spec.ObjectID, payload spec.NetworkPayload) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Activate")

	// Check if the given inputs can be used against the requested CLG's
	// interface.
	CLG, ok := n.CLGs[clgID]
	if !ok {
		return maskAnyf(clgNotFoundError, "ID: %s", clgID)
	}
	if !equalTypes(valuesToTypes(payload.Args), CLG.GetInputTypes()) {
		return maskAnyf(invalidInterfaceError, "types must match")
	}

	return nil
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

func (n *network) Calculate(clgID spec.ObjectID, payload spec.NetworkPayload) (spec.NetworkPayload, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Calculate")

	CLG, ok := n.CLGs[clgID]
	if !ok {
		return spec.NetworkPayload{}, maskAnyf(clgNotFoundError, "ID: %s", clgID)
	}

	calculatedPayload, err := CLG.Calculate(payload)
	if err != nil {
		return spec.NetworkPayload{}, maskAny(err)
	}

	return calculatedPayload, nil
}

func (n *network) Execute(clgID spec.ObjectID, payload spec.NetworkPayload) (spec.NetworkPayload, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Execute")

	// Each CLG that is executed needs to decide if it wants to be activated.
	// This happens using the Activate method.
	err := n.Activate(clgID, payload)
	if err != nil {
		return spec.NetworkPayload{}, maskAny(err)
	}

	// Once activated, a CLG executes its actual implemented behaviour using
	// Calculate. This behaviour can be anything. It is up to the CLG.
	calculatedPayload, err := n.Calculate(clgID, payload)
	if err != nil {
		return spec.NetworkPayload{}, maskAny(err)
	}

	// After the CLGs calculation it can decide what to do next. Like Activate,
	// it is up to the CLG if it forwards signals to further CLGs. E.g. a CLG
	// might or might not forward its calculated results to one or more CLGs.
	// This depends on the calculated payload, the CLG's connections and its
	// behaviour properties.
	err = n.Forward(clgID, calculatedPayload)
	if err != nil {
		return spec.NetworkPayload{}, maskAny(err)
	}

	return calculatedPayload, nil
}

// Forward to other CLGs
// Split the neural connection path
// TODO
func (n *network) Forward(clgID spec.ObjectID, payload spec.NetworkPayload) error {
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

// TODO
func (n *network) Listen() error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Listen")

	// Listen on TextInput from the outside to receive text requests.
	CLG, err := n.clgByName("input")
	if err != nil {
		return maskAny(err)
	}
	go func() {
		inputChannel := CLG.GetInputChannel()

		for {
			select {
			case textRequest := <-n.TextInput:
				// TODO
				ctx := context.Background()
				ctx = context.WithValue(ctx, "key", textRequest)

				payload := spec.NetworkPayload{
					Args:        []reflect.Value{reflect.ValueOf(ctx)},
					Destination: "",
					Sources:     []spec.ObjectID{""},
				}

				inputChannel <- payload
			}
		}
	}()

	// Make all CLGs listening in their specific input channel.
	for ID, CLG := range n.CLGs {
		go func(ID spec.ObjectID, CLG spec.CLG) {
			// This is the queue of input payloads. We collect inputs until the
			// requested CLG's interface is fulfilled somehow. Then the CLG is
			// executed and the inputs used to execute the CLG are removed from the
			// queue.
			var queue []spec.NetworkPayload
			//desired := CLG.GetInputTypes()
			inputChannel := CLG.GetInputChannel()

			for {
				select {
				case payload := <-inputChannel:
					// We should only hold a rather small number of pending payloads. A
					// payload is considered pending when it is not able to fulfill the
					// requested CLG's interface on its own. Especially the untrained
					// network tends to throw around signals without that much sense.
					// Then the payload queue for each CLG can grow and increase the
					// process memory. That is why we cap the queue. TODO The cap must be
					// learned. For now we hard code it to 10.
					var maxPending = 10
					queue = append(queue, payload)
					if len(queue) > maxPending {
						queue = queue[1:]
					}

					// // Check if the interface of the requested CLG matches the provided
					// // inputs. In case the interface does not match, it is not possible
					// // to call the requested CLG using the provided inputs. Then we do
					// // nothing, but wait some time for other input payloads to arrive.
					// execute, matching, newQueue, err := n.extractMatchingNetworkPayloads(queue, desired)
					// if err != nil {
					// 	n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
					// 	continue
					// }
					// if !execute {
					// 	continue
					// }
					// queue = newQueue

					// In case the interface is fulfilled we can finally execute the CLG.
					go func(payload spec.NetworkPayload) {
						calculatedPayload, err := n.Execute(ID, payload)
						if err != nil {
							n.Log.WithTags(spec.Tags{L: "E", O: n, T: nil, V: 4}, "%#v", maskAny(err))
						}

						// The output CLG is the only other special CLG, next to the input CLG. Only
						// these both are treated specially. Here we forward the ouputs of the output
						// CLG to the output channel. This will cause the output returned here to be
						// streamed back to the client waiting for calculations of the neural
						// network.
						if CLG.GetName() == "output" {
							var output string
							for _, v := range calculatedPayload.Args[1:] {
								output += v.String()
							}
							n.TextOutput <- api.TextResponse{Output: output}
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
