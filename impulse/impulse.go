// Package impulse implementes spec.Impulse. An impulse can walk through any
// spec.Core, spec.Network and spec.Neuron. Concrete implementations and their
// dynamic state decide about the way an impulse is going, resulting in
// behaviour.
package impulse

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeImpulse represents the object type of the impulse object. This
	// is used e.g. to register itself to the logger.
	ObjectTypeImpulse spec.ObjectType = "impulse"
)

// Config represents the configuration used to create a new impulse object.
type Config struct {
	// Dependencies.
	Log spec.Log

	// Settings.
	Expectation spec.Expectation
	Input       string
	Output      string
	SessionID   string
}

// DefaultConfig provides a default configuration to create a new impulse
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log: log.NewLog(log.DefaultConfig()),

		// Settings.
		Expectation: nil,
		Input:       "",
		Output:      "",
		SessionID:   string(id.MustNew()),
	}

	return newConfig
}

// New creates a new configured impulse object.
func New(config Config) (spec.Impulse, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		return nil, maskAny(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		return nil, maskAny(err)
	}

	newImpulse := &impulse{
		Config: config,
		ID:     newID,
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeImpulse,
	}

	if newImpulse.SessionID == "" {
		return nil, maskAnyf(invalidConfigError, "session ID must not be empty")
	}

	newImpulse.Log.Register(newImpulse.GetType())

	return newImpulse, nil
}

type impulse struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (i *impulse) GetExpectation() spec.Expectation {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.Expectation
}

func (i *impulse) GetInput() string {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.Input
}

func (i *impulse) GetOutput() string {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.Output
}

func (i *impulse) GetSessionID() string {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.SessionID
}

func (i *impulse) SetOutput(output string) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Output = output
}
