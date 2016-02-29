// Package strategynetwork provides functionality to store and optimize given
// strategies.
package strategynetwork

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

const (
	ObjectTypeStrategyNetwork spec.ObjectType = "strategy-network"
)

type NetworkConfig struct {
	// Dependencies.
	Log spec.Log

	Storage spec.Storage

	// Settings.

	// Actions represents a list of action items, that are object types. These
	// are used to find the best performing combination by ordering them in a
	// certain way. Such a ordered list of actions is called a strategy
	// maintained by a neuron. The best strategy is represented by the highest
	// score.
	Actions []spec.ObjectType

	// TODO the context should be per strategy, not per network. Means it needs
	// to be passed around during strategy creation.
	//
	// Context representes an identifier to store contextual information. E.g.
	// "math" could be a context that refers to strategies supposed to solve
	// methematical problems.
	Context string

	// MaxElements representes the maximum number of elements to fetch from a
	// list within one call. This prevents fetching too much data at once.
	MaxElements int

	// Scope representes an identifier to store information related to a scope.
	// e.g. "core" could be a scope that refers to the object using the current
	// strategy network.
	Scope string
}

func DefaultNetworkConfig() NetworkConfig {
	newConfig := NetworkConfig{
		// Dependencies.
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: storage.NewMemoryStorage(storage.DefaultMemoryStorageConfig()),

		// Settings.
		Actions:     []spec.ObjectType{},
		Context:     "default",
		MaxElements: 10,
		Scope:       "global",
	}

	return newConfig
}

// Network represents a strategy network. It manages strategies for a given set
// of action items.
type Network interface {
	// CtxKey returns a key used to scope data within the storage based on
	// contextual information.
	CtxKey(f string, v ...interface{}) string

	// GetBestNeuron returns the highest scored strategy the network is aware
	// of.
	GetBestNeuron() (Neuron, error)

	// GetHighScore returns the highest score the strategy network is aware of.
	GetHighScore() (float32, error)

	// GetNeuronByScore returns a strategy for the given score.
	GetNeuronByScore(score float32) (Neuron, error)

	spec.Network

	// StoreNeuron persists the given neuron's state.
	StoreNeuron(neu Neuron) error
}

// NewNetwork returns a new configured strategy network.
//
// The strategy network creates and maintains the following key spaces.
//
//   scope:<scope>:network:strategy:context:<context>:<key>
//
// Key can be anything of the following.
//
//   neuron:scores    Holds the weighted list of neurons that performed the best.
//   neuron:<ID>      Holds the structured data of a neuron.
//
func NewNetwork(config NetworkConfig) (Network, error) {
	newNetwork := &network{
		NetworkConfig: config,
		ID:            id.NewObjectID(id.Hex128),
		Mutex:         sync.Mutex{},
		Type:          ObjectTypeStrategyNetwork,
	}

	if newNetwork.Context == "" {
		return nil, maskAnyf(invalidContextError, "empty context")
	}

	if newNetwork.Scope == "" {
		return nil, maskAnyf(invalidScopeError, "empty scope")
	}

	newNetwork.Log.Register(newNetwork.GetType())

	return newNetwork, nil
}

type network struct {
	ID spec.ObjectID `json:"id"`

	Mutex sync.Mutex `json:"-"`

	NetworkConfig

	Type spec.ObjectType `json:"type"`
}

func (n *network) GetBestNeuron() (Neuron, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call GetBestNeuron")

	score, err := n.GetHighScore()
	if err != nil {
		return nil, maskAny(err)
	}

	strategy, err := n.GetNeuronByScore(score)
	if err != nil {
		return nil, maskAny(err)
	}

	return strategy, nil
}

func (n *network) CtxKey(f string, v ...interface{}) string {
	// keyPair is supposed to hold a structured key representation delimited by
	// colons, e.g. "neuron:score".
	keyPair := fmt.Sprintf(f, v...)
	return fmt.Sprintf("scope:%s:network:strategy:context:%s:%s", n.Scope, n.Context, keyPair)
}

func (n *network) GetHighScore() (float32, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call GetHighScore")

	_, score, err := n.Storage.GetHighestElementScore(n.CtxKey("neuron:scores"))
	if err != nil {
		return 0, maskAny(err)
	}

	return score, nil
}

func (n *network) GetNeuronByID(ID string) (Neuron, error) {
	value, err := n.Storage.Get(n.CtxKey("neuron:%s", ID))
	if err != nil {
		return nil, maskAny(err)
	}

	var neu *neuron
	err = json.Unmarshal([]byte(value), neu)
	if err != nil {
		return nil, maskAny(err)
	}

	return neu, nil
}

func (n *network) GetNeuronByScore(score float32) (Neuron, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call GetNeuronByScore")

	elements, err := n.Storage.GetElementsByScore(n.CtxKey("neuron:scores"), score, n.MaxElements)
	if err != nil {
		return nil, maskAny(err)
	}

	i := rand.Intn(len(elements))
	ID := elements[i]

	newNeuron, err := n.GetNeuronByID(ID)
	if err != nil {
		return nil, maskAny(err)
	}

	return newNeuron, nil
}

func (n *network) NewNeuron() (Neuron, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call NewNeuron")

	// Create new neuron.
	newStrategy, err := n.NewStrategy()
	if err != nil {
		return nil, maskAny(err)
	}
	newNeuronConfig := DefaultNeuronConfig()
	newNeuronConfig.Strategy = newStrategy
	neu := NewNeuron(newNeuronConfig)

	// Store neuron.
	err = n.StoreNeuron(neu)
	if err != nil {
		return nil, maskAny(err)
	}

	return neu, nil
}

func (n *network) NewStrategy() (spec.Strategy, error) {
	newConfig := StrategyConfig{
		Actions: n.Actions,
	}

	var newStrategy spec.Strategy
	for i := 0; i < 3; i++ {
		newStrategy = NewStrategy(newConfig)

		// Check if strategy already exists.
		//
		// TODO this needs to be improved. There are already ideas. See
		// https://github.com/xh3b4sd/anna/issues/79.
		_, err := n.Storage.Get(n.CtxKey("strategy:%s", newStrategy.String()))
		if err != nil {
			return nil, maskAny(err)
		}
	}

	if newStrategy == nil {
		// After 3 tries there is still no valid strategy. We already created all
		// combinations of the available actions. We cannot find any new strategy.
		// The given actions are not sufficient and lead to absolutely nothing.
		// Return an error for this serious problem.
		//
		// TODO that there is no combination actually might not be true. The more
		// combinations we already have, the more increases the probability to
		// create the same strategy three times in a row. We need to improve
		// strategy creation. A random order is a good first step, but is not
		// sufficient when we are going to reach the other end of the propability.
		// For some ideas see https://github.com/xh3b4sd/anna/issues/80.
		return nil, maskAnyf(combinationLimitError, "no more strategies left")
	}

	return newStrategy, nil
}

func (n *network) Shutdown() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Shutdown")
}

func (n *network) StoreNeuron(neu Neuron) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call StoreNeuron")

	raw, err := json.Marshal(neu.(*neuron))
	if err != nil {
		return maskAny(err)
	}
	err = n.Storage.Set(n.CtxKey("neuron:%s", neu.GetID()), string(raw))
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (n *network) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Trigger")

	strategyNeuron, err := n.GetBestNeuron()
	// TODO handle error in a proper way.
	// if storage.IsKeyNotFound(err) {
	// 	strategyNeuron, err = n.NewNeuron()
	// 	if err != nil {
	// 		return nil, maskAny(err)
	// 	}
	// } else if err != nil {
	// 	return nil, maskAny(err)
	// }
	if err != nil {
		return nil, maskAny(err)
	}

	// TODO create new neuron when score is not sufficient - evaluation network?

	neu := strategyNeuron.(spec.Neuron)
	for {
		imp, neu, err = neu.Trigger(imp)
		if err != nil {
			return nil, maskAny(err)
		}

		if neu == nil {
			// As soon as a neuron has decided to not forward an impulse to any other
			// neuron, the impulse went its way through the whole network. So we
			// break here to return the impulse.
			break
		}
	}

	return imp, nil
}
