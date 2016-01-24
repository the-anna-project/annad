package gateway

import (
	"sync"

	"github.com/xh3b4sd/anna/gateway/spec"
	"github.com/xh3b4sd/anna/id"
)

type SignalConfig struct {
	Bytes map[string][]byte
	ID    string
}

func DefaultSignalConfig() SignalConfig {
	newSignalConfig := SignalConfig{
		Bytes: map[string][]byte{
			"request":  []byte{},
			"response": []byte{},
		},
		ID: string(id.NewID(id.Hex512)),
	}

	return newSignalConfig
}

func NewSignal(config SignalConfig) spec.Signal {
	return &signal{
		Canceled:     false,
		SignalConfig: config,
		Mutex:        sync.Mutex{},
		Responder:    make(chan spec.Signal, 1000),
	}
}

type signal struct {
	Canceled bool

	SignalConfig

	Error     error
	Responder chan spec.Signal
	Mutex     sync.Mutex
}

func (s *signal) Cancel() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Canceled = true
	close(s.Responder)
}

func (s *signal) GetBytes(key string) ([]byte, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if b, ok := s.Bytes[key]; ok {
		return b, nil
	}

	return nil, maskAny(bytesNotFoundError)
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

func (s *signal) GetResponder() (chan spec.Signal, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if s.Canceled {
		return nil, maskAny(signalCanceledError)
	}

	return s.Responder, nil
}

func (s *signal) SetBytes(key string, bytes []byte) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Bytes[key] = bytes
}

func (s *signal) SetError(err error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Error = err
}

func (s *signal) SetResponder(responder chan spec.Signal) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Responder = responder
}
