// Package strategy implements spec.Strategy to provide executable and
// permutable CLG tree structures.
package strategy

import (
	"sync"

	"github.com/xh3b4sd/anna/clg"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeStrategy represents the object type of the strategy object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeStrategy = "strategy"
)

// Config represents the configuration used to create a new strategy object.
type Config struct {
	// Nodes are strategies that represent the arguments of the strategy's Root.
	// To turn Nodes into arguments, each strategy is executed. The outputs of
	// Nodes is then used to provide them as input to the Root of this strategy.
	Nodes []spec.Strategy

	// Root represents the base CLG of the strategy. Its arguments are the
	// strategies Nodes, if any.
	Root spec.CLG
}

// DefaultConfig provides a default configuration to create a new strategy
// object by best effort.
//
// Note that Root is empty and must be properly configured for the strategy
// creation. Nodes is empty as well and needs to be configured, if desired.
func DefaultConfig() Config {
	newConfig := Config{
		Nodes: nil,
		Root:  nil,
	}

	return newConfig
}

// New creates a new configured strategy object.
func New(config Config) (spec.Strategy, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newStrategy := &strategy{
		Config: config,

		ID:    newID,
		Mutex: sync.Mutex{},
		Type:  ObjectTypeStrategy,
	}

	if newStrategy.Root == "" {
		return nil, maskAnyf(invalidConfigError, "root must not be empty")
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

func (s *strategy) Execute() ([]reflect.Value, error) {
	var inputs []reflect.Value

	for _, n := range s.GetNodes() {
		outputs, err := n.Execute()
		if err != nil {
			return nil, maskAny(err)
		}
		inputs = append(inputs, outputs...)
	}

	outputs, err := clg.Execute(s.GetRoot(), inputs)
	if err != nil {
		return nil, maskAny(err)
	}
	filtered, err := filterError(outputs)
	if err != nil {
		return nil, maskAny(err)
	}

	return filtered, nil
}

func (s *strategy) GetRoot() spec.CLG {
	return s.CLG
}

func (s *strategy) GetNodes() []spec.Strategy {
	return s.Nodes
}

func (s *strategy) SetNode(indizes []int, node spec.Strategy) error {
	if len(indizes) == 0 {
		return maskAny(indexOutOfRangeError)
	}

	if len(indizes) == 1 {
		// Here we want to set the given node to the current strategy. We do not
		// need to redirect the setting to other nodes.
		index := indizes[0]

		if index < 0 || index >= len(s.Nodes) {
			return maskAny(indexOutOfRangeError)
		}

		// Backup the current node in case something goes wrong.
		current := s.Nodes[index]

		s.Nodes[index] = node

		err := s.Validate()
		if err != nil {
			// In case the new node does not fit and causes the strategy interface to
			// be invalid, we revert the change.
			s.Nodes[index] = current
			return maskAny(err)
		}
	}

	// There is more than one index given. Thus we remove our index and redirect
	// the setting to other nodes.
	index := indizes[0]
	newIndizes := indizes[1:]

	err := s.Nodes[index].SetNode(newIndizes, node)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *strategy) Validate() error {
	validInterfaces, err := isValidInterface(s.GetRoot(), s.GetNodes())
	if err != nil {
		return maskAny(err)
	}
	if !validInterfaces {
		return maskAnyf(invalidStrategyError, "invalid interface")
	}

	if isCircular(s.GetID(), s.GetNodes()) {
		return maskAnyf(invalidStrategyError, "circular strategy")
	}

	return nil
}
