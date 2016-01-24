package state

import (
	"sync"
	"time"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/spec"
)

type Config struct {
	Bytes      map[string][]byte              `json:"bytes,omitempty"`
	Cores      map[spec.ObjectID]spec.Core    `json:"cores,omitempty"`
	Impulses   map[spec.ObjectID]spec.Impulse `json:"impulses,omitempty"`
	ObjectID   spec.ObjectID                  `json:"object_id,omitempty"`
	ObjectType spec.ObjectType                `json:"object_type,omitempty"`
	Networks   map[spec.ObjectID]spec.Network `json:"networks,omitempty"`
	Neurons    map[spec.ObjectID]spec.Neuron  `json:"neurons,omitempty"`
	Order      []spec.Object                  `json:"order,omitempty"`
}

func DefaultConfig() Config {
	config := Config{
		Bytes: map[string][]byte{
			"request":  []byte{},
			"response": []byte{},
		},
		Cores:      map[spec.ObjectID]spec.Core{},
		Impulses:   map[spec.ObjectID]spec.Impulse{},
		ObjectID:   id.NewID(id.Hex512),
		ObjectType: spec.ObjectType(""),
		Networks:   map[spec.ObjectID]spec.Network{},
		Neurons:    map[spec.ObjectID]spec.Neuron{},
		Order:      []spec.Object{},
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
	Mutex     sync.Mutex `json:"mutex,omitempty"`
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

func (s *state) GetCore() (spec.Core, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if len(s.Cores) != 1 {
		return nil, maskAny(invalidCoreRelationError)
	}

	var core spec.Core
	for _, c := range s.Cores {
		core = c
		break
	}

	return core, nil
}

func (s *state) GetObjectID() spec.ObjectID {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.ObjectID
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

func (s *state) GetObjectType() spec.ObjectType {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.ObjectType
}

func (s *state) SetBytes(key string, bytes []byte) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Bytes[key] = bytes
}

func (s *state) SetCore(core spec.Core) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Order = append(s.Order, core)
	s.Cores[core.GetObjectID()] = core
}

func (s *state) SetImpulse(impulse spec.Impulse) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Order = append(s.Order, impulse)
	s.Impulses[impulse.GetObjectID()] = impulse
}

func (s *state) SetNetwork(network spec.Network) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Order = append(s.Order, network)
	s.Networks[network.GetObjectID()] = network
}

func (s *state) SetNeuron(neuron spec.Neuron) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Order = append(s.Order, neuron)
	s.Neurons[neuron.GetObjectID()] = neuron
}
