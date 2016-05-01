// Package strategy implements spec.Strategy to provide manageable action
// sequences.
package strategy

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeStrategy represents the object type of the strategy object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeStrategy = "strategy"
)

// Config represents the configuration used to create a new strategy object.
type Config struct {
	// CLGNames represents a list of ordered action items, that are CLG names.
	CLGNames []string

	// Requestor represents the object requesting a strategy. E.g. when the
	// character network requests a strategy to act on the given input, it will
	// instruct an impulse to go through the strategy network while being
	// configured with information about the character network. Here the
	// requestor would hold the object tyope of the character network.
	Requestor spec.ObjectType
}

// DefaultConfig provides a default configuration to create a new strategy
// object by best effort. Note that the list of CLG names is empty and needs to
// be properly set before the strategy creation.
func DefaultConfig() Config {
	newConfig := Config{
		CLGNames:  []string{},
		Requestor: "",
	}

	return newConfig
}

// NewStrategy creates a new configured strategy object.
func NewStrategy(config Config) (spec.Strategy, error) {
	newStrategy := &strategy{
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Type:   ObjectTypeStrategy,
	}

	newStrategy.CLGNames = randomizeCLGNames(newStrategy.CLGNames)

	if len(newStrategy.CLGNames) == 0 {
		return nil, maskAnyf(invalidConfigError, "CLG names must not be empty")
	}
	if newStrategy.ID == "" {
		return nil, maskAnyf(invalidConfigError, "ID must not be empty")
	}
	if newStrategy.Requestor == "" {
		return nil, maskAnyf(invalidConfigError, "requestor not be empty")
	}

	return newStrategy, nil
}

// NewEmptyStrategy simply returns an empty, maybe invalid, strategy object.
// This should only be used for things like unmarshaling.
func NewEmptyStrategy() spec.Strategy {
	return &strategy{}
}

type strategy struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (s *strategy) GetCLGNames() []string {
	return s.CLGNames
}

func (s *strategy) GetRequestor() spec.ObjectType {
	return s.Requestor
}
