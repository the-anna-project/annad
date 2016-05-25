// Package profile provides implementations to generate CLG profiles.
package profile

import (
	"reflect"
	"sync"
	"time"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCLGProfile represents the object type of the CLG profile object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeCLGProfile spec.ObjectType = "clg-profile"
)

// Config represents the configuration used to create a new CLG profile object.
type Config struct {
	// Settings.

	// Body represents the CLG's implemented method body.
	Body string `json:"body,omitempty"`

	// HasChanged describes whether the CLG changed. A change might be a
	// renaming, a signature modification or even a reimplementation or bugfix.
	HasChanged bool `json:"hash_changed,omitempty"`

	// Hash represents the hashed value of the CLG's implemented method.
	Hash string `json:"hash,omitempty"`

	// InputsOutputs represents the CLG's implemented method input-output
	// argument pairs discovered during CLG profile creation. This list might
	// only holds some samples and no complete list.
	InputsOutputs spec.InputsOutputs `json:"inputs_outputs,omitempty"`

	// Name represents the CLG's implemented method name.
	Name string `json:"name,omitempty"`
}

// DefaultConfig provides a default configuration to create a new CLG profile
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Body:          "",
		HasChanged:    false,
		Hash:          "",
		InputsOutputs: spec.InputsOutputs{},
		Name:          "",
	}

	return newConfig
}

// New creates a new configured CLG profile object.
func New(config Config) (spec.CLGProfile, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newProfile := &profile{
		Config: config,

		CreatedAt: time.Now(),
		ID:        newID,
		Mutex:     sync.Mutex{},
		Type:      ObjectTypeCLGProfile,
	}

	if newProfile.Body == "" {
		return nil, maskAnyf(invalidConfigError, "method body of CLG profile must not be empty")
	}
	if newProfile.Hash == "" {
		return nil, maskAnyf(invalidConfigError, "hash of CLG profile must not be empty")
	}
	if newProfile.Name == "" {
		return nil, maskAnyf(invalidConfigError, "method name of CLG profile must not be empty")
	}

	return newProfile, nil
}

// NewEmptyProfile simply returns an empty, maybe invalid, profile object. This
// should only be used for things like unmarshaling.
func NewEmptyProfile() spec.CLGProfile {
	return &profile{}
}

type profile struct {
	Config

	// CreatedAt represents the creation time of the profile.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// ID represents the profile's identifier.
	ID spec.ObjectID `json:"id,omitempty"`

	Mutex sync.Mutex `json:"-"`

	// Type represents the profile's object type.
	Type spec.ObjectType `json:"type,omitempty"`
}

func (p *profile) Equals(other spec.CLGProfile) bool {
	if p.GetBody() != other.GetBody() {
		return false
	}
	if p.GetHash() != other.GetHash() {
		return false
	}
	if !reflect.DeepEqual(p.GetInputsOutputs(), other.GetInputsOutputs()) {
		return false
	}
	if p.GetName() != other.GetName() {
		return false
	}

	return true
}

func (p *profile) GetBody() string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Body
}

func (p *profile) GetHasChanged() bool {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.HasChanged
}

func (p *profile) GetHash() string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Hash
}

func (p *profile) GetInputsOutputs() spec.InputsOutputs {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.InputsOutputs
}

func (p *profile) GetName() string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Name
}

func (p *profile) SetHashChanged(hasChanged bool) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	p.HasChanged = hasChanged
}

func (p *profile) SetID(id spec.ObjectID) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	p.ID = id
}
