// Package impulse implementes spec.Impulse. An impulse can walk through any
// spec.Core, spec.Network and spec.Neuron. Concrete implementations and their
// dynamic state decide about the way an impulse is going, resulting in
// behaviour.
package impulse

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeImpulse represents the object type of the impulse object. This
	// is used e.g. to register itself to the logger.
	ObjectTypeImpulse spec.ObjectType = "impulse"
)

// Config represents the configuration used to create a new impulse object.
type Config struct {
	// Dependencies.
	Log spec.Log

	// Settings.
	Actions    []spec.ObjectType
	Inputs     map[spec.ObjectID]string
	Output     string
	Requestor  spec.ObjectType
	SessionID  string
	Strategies map[spec.ObjectType]spec.Strategy
}

// DefaultConfig provides a default configuration to create a new impulse
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log: log.NewLog(log.DefaultConfig()),

		// Settings.
		Actions:    []spec.ObjectType{},
		Inputs:     map[spec.ObjectID]string{},
		Output:     "",
		Requestor:  spec.ObjectType(""),
		SessionID:  string(id.NewObjectID(id.Hex128)),
		Strategies: map[spec.ObjectType]spec.Strategy{},
	}

	return newConfig
}

// NewImpulse creates a new configured impulse object.
func NewImpulse(config Config) (spec.Impulse, error) {
	newImpulse := &impulse{
		Config:            config,
		ID:                id.NewObjectID(id.Hex128),
		Mutex:             sync.Mutex{},
		OrderedStrategies: []spec.Strategy{},
		Type:              ObjectTypeImpulse,
	}

	if newImpulse.SessionID == "" {
		return nil, maskAnyf(invalidConfigError, "session ID must not be empty")
	}

	newImpulse.Log.Register(newImpulse.GetType())

	return newImpulse, nil
}

type impulse struct {
	Config

	ID                spec.ObjectID
	Mutex             sync.Mutex
	OrderedStrategies []spec.Strategy
	Type              spec.ObjectType
}

func (i *impulse) GetActions() []spec.ObjectType {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.Actions
}

func (i *impulse) GetAllInputs() map[spec.ObjectID]string {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.Inputs
}

func (i *impulse) GetAllStrategies() map[spec.ObjectType]spec.Strategy {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.Strategies
}

func (i *impulse) GetInputByImpulseID(impulseID spec.ObjectID) (string, error) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	if i, ok := i.Inputs[impulseID]; ok {
		return i, nil
	}

	return "", maskAny(inputNotFoundError)
}

func (i *impulse) GetOutput() string {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.Output
}

func (i *impulse) GetRequestor() spec.ObjectType {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.Requestor
}

func (i *impulse) GetSessionID() string {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.SessionID
}

func (i *impulse) GetStrategy() spec.Strategy {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	if len(i.OrderedStrategies) == 0 {
		return nil
	}

	return i.OrderedStrategies[len(i.OrderedStrategies)-1]
}

func (i *impulse) SetActions(actions []spec.ObjectType) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Actions = actions
}

func (i *impulse) SetInputByImpulseID(impulseID spec.ObjectID, input string) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Inputs[impulseID] = input
}

func (i *impulse) SetRequestor(requestor spec.ObjectType) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Requestor = requestor
}

func (i *impulse) SetOutput(output string) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Output = output
}

func (i *impulse) SetStrategyByRequestor(requestor spec.ObjectType, strategy spec.Strategy) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.OrderedStrategies = append(i.OrderedStrategies, strategy)
	i.Strategies[requestor] = strategy
}
