// Package state implements spec.State. All information an object holds is
// stored within its state. Business logic and state is fully decoupled. That
// way Anna is able to completely backup and restore her whole state.
package state

import (
	"sync"
	"time"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/file-system/fake"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

type Config struct {
	Bytes         map[string][]byte              `json:"bytes,omitempty"`
	Cores         map[spec.ObjectID]spec.Core    `json:"cores,omitempty"`
	FactoryClient spec.Factory                   `json:"-"`
	FileSystem    spec.FileSystem                `json:"-"`
	Impulses      map[spec.ObjectID]spec.Impulse `json:"impulses,omitempty"`
	Log           spec.Log                       `json:"-"`
	Networks      map[spec.ObjectID]spec.Network `json:"networks,omitempty"`
	Neurons       map[spec.ObjectID]spec.Neuron  `json:"neurons,omitempty"`
	ObjectID      spec.ObjectID                  `json:"object_id,omitempty"`
	ObjectType    spec.ObjectType                `json:"object_type,omitempty"`
	StateReader   spec.StateType                 `json:"state_reader,omitempty"`
	StateWriter   spec.StateType                 `json:"state_writer,omitempty"`
}

func DefaultConfig() Config {
	config := Config{
		Bytes: map[string][]byte{
			"request":  []byte{},
			"response": []byte{},
		},
		Cores:         map[spec.ObjectID]spec.Core{},
		FactoryClient: factoryclient.NewFactory(factoryclient.DefaultConfig()),
		FileSystem:    filesystemfake.NewFileSystem(),
		Impulses:      map[spec.ObjectID]spec.Impulse{},
		Log:           log.NewLog(log.DefaultConfig()),
		Networks:      map[spec.ObjectID]spec.Network{},
		Neurons:       map[spec.ObjectID]spec.Neuron{},
		ObjectID:      id.NewObjectID(id.Hex128),
		ObjectType:    spec.ObjectType(common.ObjectType.None),
		StateReader:   common.StateType.FSReader,
		StateWriter:   common.StateType.FSWriter,
	}

	return config
}

func NewState(config Config) spec.State {
	newState := &state{
		Config:    config,
		CreatedAt: time.Now(),
		Mutex:     sync.Mutex{},
	}

	return newState
}

type state struct {
	Config

	CreatedAt time.Time  `json:"created_at,omitempty"`
	Mutex     sync.Mutex `json:"-"`
}

func (s *state) GetAge() time.Duration {
	return time.Since(s.CreatedAt)
}

func (s *state) GetBytes(key string) ([]byte, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if b, ok := s.Bytes[key]; ok {
		return b, nil
	}

	return nil, maskAny(bytesNotFoundError)
}

func (s *state) GetCoreByID(objectID spec.ObjectID) (spec.Core, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if c, ok := s.Cores[objectID]; ok {
		return c, nil
	}

	return nil, maskAny(coreNotFoundError)
}

func (s *state) GetCores() map[spec.ObjectID]spec.Core {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Cores
}

func (s *state) GetImpulseByID(objectID spec.ObjectID) (spec.Impulse, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if i, ok := s.Impulses[objectID]; ok {
		return i, nil
	}

	return nil, maskAny(impulseNotFoundError)
}

func (s *state) GetImpulses() map[spec.ObjectID]spec.Impulse {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Impulses
}

func (s *state) GetNetworkByID(objectID spec.ObjectID) (spec.Network, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if n, ok := s.Networks[objectID]; ok {
		return n, nil
	}

	return nil, maskAny(networkNotFoundError)
}

func (s *state) GetNetworks() map[spec.ObjectID]spec.Network {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Networks
}

func (s *state) GetNeuronByID(objectID spec.ObjectID) (spec.Neuron, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if n, ok := s.Neurons[objectID]; ok {
		return n, nil
	}

	return nil, maskAny(neuronNotFoundError)
}

func (s *state) GetNeurons() map[spec.ObjectID]spec.Neuron {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Neurons
}

func (s *state) SetBytes(key string, bytes []byte) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Bytes[key] = bytes
}

func (s *state) SetCore(core spec.Core) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Cores[core.GetObjectID()] = core
}

func (s *state) SetImpulse(impulse spec.Impulse) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Impulses[impulse.GetObjectID()] = impulse
}

func (s *state) SetNetwork(network spec.Network) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Networks[network.GetObjectID()] = network
}

func (s *state) SetNeuron(neuron spec.Neuron) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Neurons[neuron.GetObjectID()] = neuron
}
