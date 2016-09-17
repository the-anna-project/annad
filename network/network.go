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

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/context"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/key"
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
	FeatureStorage     spec.Storage
	GeneralStorage     spec.Storage
	IDFactory          spec.IDFactory
	Log                spec.Log
	PermutationFactory spec.PermutationFactory
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

	newConfig := Config{
		// Dependencies.
		FeatureStorage:     memory.MustNew(),
		GeneralStorage:     memory.MustNew(),
		IDFactory:          id.MustNewFactory(),
		Log:                log.New(log.DefaultConfig()),
		PermutationFactory: newPermutationFactory,
		TextInput:          make(chan spec.TextRequest, 1000),
		TextOutput:         make(chan spec.TextResponse, 1000),

		// Settings.
		Delay: 0,
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
		Closer:       make(chan struct{}, 1),
		ID:           id.MustNew(),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectType,
	}

	if newNetwork.FeatureStorage == nil {
		return nil, maskAnyf(invalidConfigError, "feature storage must not be empty")
	}
	if newNetwork.GeneralStorage == nil {
		return nil, maskAnyf(invalidConfigError, "general storage must not be empty")
	}
	if newNetwork.IDFactory == nil {
		return nil, maskAnyf(invalidConfigError, "ID factory must not be empty")
	}
	if newNetwork.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newNetwork.PermutationFactory == nil {
		return nil, maskAnyf(invalidConfigError, "permutation factory must not be empty")
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
	Closer       chan struct{}
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (n *network) Activate(CLG spec.CLG, payload spec.NetworkPayload, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error) {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Activate")

	payload, queue, err := n.permutePayload(CLG, payload, queue)
	if permutation.IsMaxGrowthReached(err) {
		// We could not find a sufficient payload for the requsted CLG by permuting
		// the queue of network payloads.
		return nil, nil, maskAnyf(invalidInterfaceError, "types must match")
	} else if err != nil {
		return nil, nil, maskAny(err)
	}

	return payload, queue, nil
}

func (n *network) Boot() {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Boot")

	n.BootOnce.Do(func() {
		n.CLGs = n.configureCLGs(n.CLGs)
		n.CLGIDs = n.mapCLGIDs(n.CLGs)

		go n.Listen()
		go n.listenCLGs()
	})
}

func (n *network) Calculate(CLG spec.CLG, payload spec.NetworkPayload) (spec.NetworkPayload, error) {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Calculate")

	calculatedPayload, err := CLG.Calculate(payload)
	if err != nil {
		return nil, maskAny(err)
	}

	return calculatedPayload, nil
}

func (n *network) Forward(CLG spec.CLG, payload spec.NetworkPayload) error {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Forward")

	// Try to find the best connections.
	oldCtx, err := payload.GetContext()
	if err != nil {
		return maskAny(err)
	}
	behaviorIDs, err := n.findConnections(oldCtx, payload)
	if err != nil {
		return maskAny(err)
	}

	for _, ID := range behaviorIDs {
		// Prepare a new context for the new connection path.
		newCtx := oldCtx.Clone()
		newCtx.SetBehaviorID(ID)

		// Create a new network payload
		newPayloadConfig := api.DefaultNetworkPayloadConfig()
		newPayloadConfig.Args = append([]reflect.Value{reflect.ValueOf(newCtx)}, payload.GetArgs()[1:]...)
		newPayloadConfig.Destination = spec.ObjectID(ID)
		newPayloadConfig.Sources = []spec.ObjectID{payload.GetDestination()}
		newPayload, err := api.NewNetworkPayload(newPayloadConfig)
		if err != nil {
			return maskAny(err)
		}

		// Find the actual CLG based on its behavior ID. Therefore we lookup the
		// behavior name by the given behavior ID. Data we read here is written
		// within several CLGs. That way the network creates its own connections
		// based on learned experiences.
		clgName, err := n.GeneralStorage.Get(key.NewCLGKey("behavior-id:%s:behavior-name", ID))
		if err != nil {
			return maskAny(err)
		}
		CLG, err := n.clgByName(clgName)
		if err != nil {
			return maskAny(err)
		}
		CLG.GetInputChannel() <- newPayload
	}

	return nil
}

func (n *network) Listen() {
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
				continue
			}

			// Prepare the context and a unique behaviour ID for the input CLG.
			ctxConfig := context.DefaultConfig()
			ctxConfig.SessionID = textRequest.GetSessionID()
			ctx, err := context.New(ctxConfig)
			if err != nil {
				n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
				continue
			}
			behaviorID, err := n.IDFactory.New()
			if err != nil {
				n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
				continue
			}

			// We transform the received input to a network payload to have a
			// conventional data structure within the neural network. Note the
			// following details.
			//
			//     The list of arguments always contains a context as first argument.
			//
			//     Destination is always the behavior ID of the input CLG, since this
			//     one is the connecting building block to other CLGs within the
			//     neural network. This behavior ID is always a new one, because it
			//     will eventually be part of a completely new CLG tree within the
			//     connection space.
			//
			//     Sources is here only the individual network ID to have at least
			//     any reference of origin.
			//
			payloadConfig := api.DefaultNetworkPayloadConfig()
			payloadConfig.Args = []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(textRequest.GetInput())}
			payloadConfig.Destination = behaviorID
			payloadConfig.Sources = []spec.ObjectID{networkID}
			newPayload, err := api.NewNetworkPayload(payloadConfig)
			if err != nil {
				n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
				continue
			}

			// Send the new network payload to the input CLG.
			clgChannel <- newPayload
		}
	}
}

func (n *network) Shutdown() {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Shutdown")

	n.ShutdownOnce.Do(func() {
		close(n.Closer)
	})
}
