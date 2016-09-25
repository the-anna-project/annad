// Package network implements spec.Network to provide a neural network based on
// dynamic and self improving CLG execution. The network provides input and
// output channels. When input is received it is injected into the neural
// communication. The following neural activity calculates output which is
// streamed through the output channel back to the requestor.
package network

import (
	"sync"
	"time"

	"github.com/xh3b4sd/anna/factory"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

const (
	// ObjectType represents the object type of the network object. This is used
	// e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "network"
)

// Config represents the configuration used to create a new network object.
type Config struct {
	// Dependencies.
	FactoryCollection spec.FactoryCollection
	Log               spec.Log
	StorageCollection spec.StorageCollection
	TextInput         chan spec.TextRequest
	TextOutput        chan spec.TextResponse

	// Settings.

	// Delay causes each CLG execution to be delayed. This value represents a
	// default value. A delay can be used harden the internal synchronization of
	// the network. For instance some chaos monkey could be implemented to cause
	// unusual delays in neural communications. Analysing situations in which such
	// chaos monkey takes place might shed some light on faulty implementations
	// within the neural network.
	//
	// TODO implement the actual usage of the delay and make it dynamically
	// configurable on demand like we already do with the log control.
	Delay time.Duration
}

// DefaultConfig provides a default configuration to create a new network
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		FactoryCollection: factory.MustNewCollection(),
		Log:               log.New(log.DefaultConfig()),
		StorageCollection: storage.MustNewCollection(),
		TextInput:         make(chan spec.TextRequest, 1000),
		TextOutput:        make(chan spec.TextResponse, 1000),

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

	if newNetwork.FactoryCollection == nil {
		return nil, maskAnyf(invalidConfigError, "factory collection must not be empty")
	}
	if newNetwork.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newNetwork.StorageCollection == nil {
		return nil, maskAnyf(invalidConfigError, "storage collection must not be empty")
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

func (n *network) Activate(CLG spec.CLG, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error) {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Activate")

	behaviorID, err := behaviorIDFromQueue(queue)
	if err != nil {
		return nil, nil, maskAny(err)
	}

	// Check if we have neural connections that tell us which payloads to use.
	payload, queue, err := n.payloadFromConnections(behaviorID, queue)
	if IsInvalidInterface(err) {
		// There are no sufficient connections. We need to come up with something
		// random.
		payload, queue, err = n.payloadFromPermutations(CLG, queue)
		if permutation.IsMaxGrowthReached(err) {
			// We could not find a sufficient payload for the requsted CLG by permuting
			// the queue of network payloads.
			return nil, nil, maskAnyf(invalidInterfaceError, "types must match")
		} else if err != nil {
			return nil, nil, maskAny(err)
		}

		// Once we found a new combination, we need to make sure the neural network
		// remembers it. Thus we store the connections between the current behavior
		// and the behaviors matching the interface of the current behavior.
		var behaviorIDs string
		for _, s := range payload.GetSources() {
			behaviorIDs += "," + string(s)
		}
		behaviorIDsKey := key.NewCLGKey("behavior-id:%s:activate-behavior-ids", behaviorID)
		err := n.Storage().General().Set(behaviorIDsKey, behaviorIDs)
		if err != nil {
			return nil, nil, maskAny(err)
		}
	}

	return payload, queue, nil
}

func (n *network) Boot() {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Boot")

	n.BootOnce.Do(func() {
		n.CLGs = n.configureCLGs(n.CLGs)
		n.CLGIDs = n.mapCLGIDs(n.CLGs)

		go n.listenInputCLG()
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

func (n *network) Factory() spec.FactoryCollection {
	return n.FactoryCollection
}

func (n *network) Forward(CLG spec.CLG, payload spec.NetworkPayload) error {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Forward")

	ctx, err := payload.GetContext()
	if err != nil {
		return maskAny(err)
	}

	behaviorIDs, err := n.findConnections(ctx, payload)
	if err != nil {
		return maskAny(err)
	}

	err = n.forwardCLGs(ctx, behaviorIDs, payload)
	if err != nil {
		return maskAny(err)
	}

	if CLG.GetName() == "output" {
		err := n.forwardOutputCLG(ctx, payload)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}

func (n *network) Shutdown() {
	n.Log.WithTags(spec.Tags{C: nil, L: "D", O: n, V: 13}, "call Shutdown")

	n.ShutdownOnce.Do(func() {
		close(n.Closer)
	})
}

func (n *network) Storage() spec.StorageCollection {
	return n.StorageCollection
}
