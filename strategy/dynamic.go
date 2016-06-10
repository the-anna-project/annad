package strategy

import (
	"reflect"
	"sync"

	"github.com/xh3b4sd/anna/clg"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeDynamicStrategy represents the object type of the dynamic
	// strategy object. This is used e.g. to register itself to the logger.
	ObjectTypeDynamicStrategy = "dynamic-strategy"
)

// DynamicConfig represents the configuration used to create a new dynamic
// strategy object.
type DynamicConfig struct {
	// Settings.

	// Nodes are strategies that represent the arguments of the strategy's Root.
	// To turn Nodes into arguments, each strategy is executed. The outputs of
	// Nodes is then used to provide them as input to the Root of this strategy.
	Nodes []spec.Strategy

	// Root represents the base CLG of the strategy. Its arguments are the
	// strategies Nodes, if any. The output of the executed Root CLG will be
	// returned when calling Execute.
	Root spec.CLG
}

// DefaultDynamicConfig provides a default configuration to create a new
// dynamic strategy object by best effort.
//
// Note that Root is empty and must be properly configured for the strategy
// creation. Nodes is empty as well and needs to be configured, if desired.
func DefaultDynamicConfig() DynamicConfig {
	newConfig := DynamicConfig{
		// Settings.
		Nodes: nil,
		Root:  "",
	}

	return newConfig
}

// NewDynamic creates a new configured dynamic strategy object.
func NewDynamic(config DynamicConfig) (spec.Strategy, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newStrategy := &dynamic{
		DynamicConfig: config,

		ID:    newID,
		Mutex: sync.Mutex{},
		Type:  ObjectTypeDynamicStrategy,
	}

	if newStrategy.Root == "" {
		return nil, maskAnyf(invalidConfigError, "root must not be empty")
	}

	return newStrategy, nil
}

// NewEmptyDynamic simply returns an empty, maybe invalid, dynamic strategy
// object. This should only be used for things like unmarshaling.
func NewEmptyDynamic() spec.Strategy {
	return &dynamic{}
}

type dynamic struct {
	DynamicConfig

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (d *dynamic) Execute() ([]reflect.Value, error) {
	err := d.Validate()
	if err != nil {
		return nil, maskAny(err)
	}

	var inputs []reflect.Value

	for _, n := range d.Nodes {
		outputs, err := n.Execute()
		if err != nil {
			return nil, maskAny(err)
		}
		inputs = append(inputs, outputs...)
	}

	outputs, err := clg.Execute(d.Root, inputs)
	if err != nil {
		return nil, maskAny(err)
	}
	filtered, err := filterError(outputs)
	if err != nil {
		return nil, maskAny(err)
	}

	return filtered, nil
}

func (d *dynamic) GetOutputs() ([]reflect.Type, error) {
	outputs, err := clg.Outputs(d.Root)
	if err != nil {
		return nil, maskAny(err)
	}

	return outputs, nil
}

func (d *dynamic) IsStatic() bool {
	return false
}

func (d *dynamic) RemoveNode(indizes []int) error {
	if len(indizes) == 0 {
		return maskAny(indexOutOfRangeError)
	}

	if len(indizes) == 1 {
		// Here we want to remove the given node from the current strategy. We do
		// not need to redirect the removal to other nodes.
		index := indizes[0]

		if index < 0 || index > len(d.Nodes) {
			return maskAny(indexOutOfRangeError)
		}

		// See https://github.com/golang/go/wiki/SliceTricks.
		copy(d.Nodes[index:], d.Nodes[index+1:])
		d.Nodes[len(d.Nodes)-1] = nil
		d.Nodes = d.Nodes[:len(d.Nodes)-1]

		// We are done here.
		return nil
	}

	// There is more than one index given. Thus we remove our index and redirect
	// the removal to other nodes.
	index := indizes[0]
	newIndizes := indizes[1:]

	err := d.Nodes[index].RemoveNode(newIndizes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (d *dynamic) SetNode(indizes []int, node spec.Strategy) error {
	if len(indizes) == 0 {
		return maskAny(indexOutOfRangeError)
	}

	if len(indizes) == 1 {
		// Here we want to set the given node to the current strategy. We do not
		// need to redirect the setting to other nodes.
		index := indizes[0]

		if index < 0 || index > len(d.Nodes) {
			return maskAny(indexOutOfRangeError)
		}

		if len(d.Nodes) == 0 || index == len(d.Nodes) {
			// Append the desired node.
			d.Nodes = append(d.Nodes, node)
		} else {
			// Set the desired node.
			d.Nodes[index] = node
		}

		// We are done here.
		return nil
	}

	// There is more than one index given. Thus we remove our index and redirect
	// the setting to other nodes.
	index := indizes[0]
	newIndizes := indizes[1:]

	err := d.Nodes[index].SetNode(newIndizes, node)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (d *dynamic) Validate() error {
	// Check if interfaces of the current strategy'd Root and Nodes match.
	validInterfaces, err := isValidInterface(d.Root, d.Nodes)
	if err != nil {
		return maskAny(err)
	}
	if !validInterfaces {
		return maskAnyf(invalidStrategyError, "invalid interface")
	}

	// Check if current strategy is circular.
	if isCircular(d.GetID(), d.Nodes) {
		return maskAnyf(invalidStrategyError, "circular strategy")
	}

	// Validate Nodes recursively.
	for _, n := range d.Nodes {
		err := n.Validate()
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}
