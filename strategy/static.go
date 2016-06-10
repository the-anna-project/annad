package strategy

import (
	"reflect"
	"sync"

	"github.com/xh3b4sd/anna/clg"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

// StaticConfig represents the configuration used to create a new static
// strategy object.
type Config struct {
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
func New(config StaticConfig) (spec.Strategy, error) {
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
		Type:  ObjectTypeStrategy,
	}

	if newStrategy.Argument == nil {
		return nil, maskAnyf(invalidConfigError, "argument must not be empty")
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
	outputs := []reflect.Value(reflect.ValueOf(s.GetArgument()))

	return outputs, nil
}

func (s *static) GetArgument() interface{} {
	return s.Argument
}

func (s *static) GetRoot() spec.CLG {
	return ""
}

func (s *static) GetNodes() []spec.Strategy {
	return nil
}

func (s *static) RemoveNode(indizes []int, node spec.Strategy) error {
	return maskAnyf(notRemovableError, "static strategy")
}

func (s *static) SetNode(indizes []int, node spec.Strategy) error {
	return maskAnyf(notSetableError, "static strategy")
}

func (s *static) Validate() error {
	// Static strategies are always valid.
	return nil
}
