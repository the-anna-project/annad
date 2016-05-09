// Package profile TODO
package profile

import (
	"reflect"
	"sync"
	"time"

	"github.com/xh3b4sd/anna/id"
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
	Body string

	// Hash represents the hashed value of the CLG's implemented method.
	Hash string

	// Inputs represents the CLG's implemented method input parameter types.
	Inputs []reflect.Kind

	// Name represents the CLG's implemented method name.
	Name string

	// Outputs represents the CLG's implemented method output parameter types.
	Outputs []reflect.Kind
}

// DefaultConfig provides a default configuration to create a new CLG index
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Body:    "",
		Hash:    "",
		Inputs:  nil,
		Name:    "",
		Outputs: nil,
	}

	return newConfig
}

// New creates a new configured CLG profile object.
func New(config Config) (spec.CLGProfile, error) {
	newProfile := &profile{
		Config: config,

		CreatedAt: time.Now(),
		ID:        id.NewObjectID(id.Hex128),
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
	if len(newProfile.Outputs) == 0 {
		return nil, maskAnyf(invalidConfigError, "output types of CLG profile must not be empty")
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
	if !reflect.DeepEqual(p.GetInputs(), other.GetInputs()) {
		return false
	}
	if p.GetName() != other.GetName() {
		return false
	}
	if !reflect.DeepEqual(p.GetOutputs(), other.GetOutputs()) {
		return false
	}

	return true
}

func (p *profile) GetBody() string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Body
}

func (p *profile) GetHash() string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Hash
}

func (p *profile) GetInputs() []reflect.Kind {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Inputs
}

func (p *profile) GetName() string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Name
}

func (p *profile) GetOutputs() []reflect.Kind {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Outputs
}
