package gateway

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

// SignalConfig represents the configuration used to create a new gateway
// signal.
type SignalConfig struct {
	ID     string
	Input  interface{}
	Output interface{}
}

// DefaultSignalConfig provides a default configuration to create a new gateway
// signal by best effort.
func DefaultSignalConfig() SignalConfig {
	newConfig := SignalConfig{
		ID:     string(id.MustNew()),
		Input:  nil,
		Output: nil,
	}

	return newConfig
}

// NewSignal creates a new configured gateway signal.
func NewSignal(config SignalConfig) spec.Signal {
	return &signal{
		SignalConfig: config,
		Mutex:        sync.Mutex{},
		Responder:    make(chan spec.Signal, 1000),
	}
}

type signal struct {
	Error error

	Mutex sync.Mutex

	SignalConfig

	Responder chan spec.Signal
}

func (s *signal) GetError() error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Error
}

func (s *signal) GetID() string {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.ID
}

func (s *signal) GetInput() interface{} {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	return s.Input
}

func (s *signal) GetOutput() interface{} {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	return s.Output
}

func (s *signal) GetResponder() chan spec.Signal {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	return s.Responder
}

func (s *signal) SetError(err error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Error = err
}

func (s *signal) SetOutput(output interface{}) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Output = output
}
