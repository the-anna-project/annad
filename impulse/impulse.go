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
	ObjectTypeImpulse         spec.ObjectType = "impulse"
	ObjectTypeStrategyNetwork spec.ObjectType = "strategy-network"
)

type Config struct {
	Input string `json:"input"`

	Log spec.Log `json:"-"`

	ObjectTypes []spec.ObjectType `json:"object_types"`

	Output string `json:"output"`
}

func DefaultConfig() Config {
	newConfig := Config{
		Input: "",
		Log:   log.NewLog(log.DefaultConfig()),
		ObjectTypes: []spec.ObjectType{
			ObjectTypeStrategyNetwork,
		},
		Output: "",
	}

	return newConfig
}

func NewImpulse(config Config) (spec.Impulse, error) {
	newImpulse := &impulse{
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeImpulse,
	}

	if len(newImpulse.ObjectTypes) == 0 || newImpulse.ObjectTypes[0] != ObjectTypeStrategyNetwork {
		// The first network type an impulse needs to use is the strategy network
		// type. This ensures to provide strategies which guide the impulse thorugh
		// networks by best effort. In case a caller configures an impulse without
		// that knowledge, we just prevent accidents and return an error.
		return nil, maskAnyf(invalidNetworkTypeError, "")
	}

	newImpulse.Log.Register(newImpulse.GetType())

	return newImpulse, nil
}

type impulse struct {
	Config

	ID spec.ObjectID `json:"id"`

	Mutex sync.Mutex `json:"-"`

	Type spec.ObjectType `json:"type"`
}

func (i *impulse) AddObjectType(objectType spec.ObjectType) error {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.ObjectTypes = append(i.ObjectTypes, objectType)

	return nil
}

func (i *impulse) GetInput() (string, error) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.Input, nil
}

func (i *impulse) GetOutput() (string, error) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.Output, nil
}

func (i *impulse) GetObjectType() (spec.ObjectType, error) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	objectType := i.ObjectTypes[0]

	i.ObjectTypes = i.ObjectTypes[1:]

	return objectType, nil
}

func (i *impulse) SetID(ID spec.ObjectID) error {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	i.ID = ID
	return nil
}

func (i *impulse) SetInput(input string) error {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	i.Input = input
	return nil
}

func (i *impulse) SetOutput(output string) error {
	i.Mutex.Lock()
	i.Output = output
	return nil
}
