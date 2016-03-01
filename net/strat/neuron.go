package stratnet

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/strategy"
)

const (
	ObjectTypeStrategyNeuron spec.ObjectType = "strategy-neuron"
)

type NeuronConfig struct {
	// Dependencies.
	Log spec.Log

	// Settings.

	// Strategy is the neuron's associated strategy.
	Strategy spec.Strategy

	// Score represents the value by witch a strategy is evaluated.
	Score int
}

func DefaultNeuronConfig() NeuronConfig {
	newConfig := NeuronConfig{
		// Dependencies.
		Log: log.NewLog(log.DefaultConfig()),

		// Settings.
		Strategy: strategy.NewStrategy(strategy.DefaultConfig()),
		Score:    0,
	}

	return newConfig
}

type Neuron interface {
	// AddAction adds an existing action item to the strategy. i represents the
	// index of the action item within the strategy's action sequence.
	AddAction(objectType spec.ObjectType, i int) error

	// DecrScore decrements the score of the strategy by delta.
	DecrScore(objectType spec.ObjectType, delta int) error

	// DelAction deletes an existing action item from the strategy. i represents
	// the index of the action item within the strategy's action sequence.
	DelAction(objectType spec.ObjectType, i int) error

	// GetStrategy returns the neuron's action sequence.
	GetStrategy() (spec.Strategy, error)

	// IncrScore increments the score of the strategy by delta.
	IncrScore(objectType spec.ObjectType, delta int) error

	spec.Neuron
}

func NewNeuron(config NeuronConfig) Neuron {
	newNeuron := &neuron{
		NeuronConfig: config,
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		Type:         ObjectTypeStrategyNeuron,
	}

	newNeuron.Log.Register(newNeuron.GetType())

	return newNeuron
}

type neuron struct {
	ID spec.ObjectID

	Mutex sync.Mutex

	NeuronConfig

	Type spec.ObjectType
}

func (n *neuron) AddAction(objectType spec.ObjectType, i int) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call AddAction")
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	return nil
}

func (n *neuron) DecrScore(objectType spec.ObjectType, delta int) error {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call DecrScore")

	n.Score -= delta

	return nil
}

func (n *neuron) DelAction(objectType spec.ObjectType, i int) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call DelAction")
	return nil
}

func (n *neuron) GetStrategy() (spec.Strategy, error) {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call GetStrategy")

	return n.Strategy, nil
}

func (n *neuron) IncrScore(objectType spec.ObjectType, delta int) error {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call IncrScore")

	n.Score += delta

	return nil
}

func (n *neuron) Trigger(imp spec.Impulse) (spec.Impulse, spec.Neuron, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Trigger")

	strategy, err := n.GetStrategy()
	if err != nil {
		return nil, nil, maskAny(err)
	}

	for _, action := range strategy.GetActions() {
		imp.AddObjectType(action)
	}

	return imp, nil, nil
}
