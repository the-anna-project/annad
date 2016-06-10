package strategy

import (
	"reflect"
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeStaticStrategy represents the object type of the dynamic
	// strategy object. This is used e.g. to register itself to the logger.
	ObjectTypeStaticStrategy = "static-strategy"
)

// StaticConfig represents the configuration used to create a new static
// strategy object.
type StaticConfig struct {
	// Settings.

	// Argument represents an arbitrary argument returned on strategy execution.
	Argument interface{}
}

// DefaultStaticConfig provides a default configuration to create a new static
// strategy object by best effort.
func DefaultStaticConfig() StaticConfig {
	newConfig := StaticConfig{
		// Settings.
		Argument: nil,
	}

	return newConfig
}

// NewStatic creates a new configured static strategy object.
func NewStatic(config StaticConfig) (spec.Strategy, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newStrategy := &static{
		StaticConfig: config,

		ID:    newID,
		Mutex: sync.Mutex{},
		Type:  ObjectTypeStaticStrategy,
	}

	return newStrategy, nil
}

// NewEmptyStatic simply returns an empty, maybe invalid, static strategy
// object. This should only be used for things like unmarshaling.
func NewEmptyStatic() spec.Strategy {
	return &static{}
}

type static struct {
	StaticConfig

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (s *static) Execute() ([]reflect.Value, error) {
	outputs := []reflect.Value{reflect.ValueOf(s.Argument)}

	return outputs, nil
}

func (s *static) GetOutputs() ([]reflect.Type, error) {
	outputs := []reflect.Type{reflect.TypeOf(s.Argument)}
	return outputs, nil
}

func (s *static) IsStatic() bool {
	return true
}

func (s *static) RemoveNode(indizes []int) error {
	return maskAnyf(notRemovableError, "static strategy")
}

func (s *static) SetNode(indizes []int, node spec.Strategy) error {
	return maskAnyf(notSettableError, "static strategy")
}

func (s *static) Validate() error {
	// Static strategies are always valid.
	return nil
}
