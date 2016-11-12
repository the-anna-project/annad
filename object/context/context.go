// Package context implements spec.Context and provides a wrapper of
// golang.org/x/net/context with additional business logic related to the Anna
// project.
package context

import (
	"time"

	netcontext "golang.org/x/net/context"

	"github.com/xh3b4sd/anna/object/spec"
)

// Config represents the configuration used to create a new context object.
type Config struct {
	// Settings.

	BehaviourID   string
	Context       netcontext.Context
	CLGName       string
	CLGTreeID     string
	Expectation   spec.Expectation
	InformationID string
	SessionID     string
}

// DefaultConfig provides a default configuration to create a new context
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		BehaviourID:   "",
		Context:       netcontext.Background(),
		CLGName:       "",
		CLGTreeID:     "",
		Expectation:   nil,
		InformationID: "",
		SessionID:     "",
	}

	return newConfig
}

// New creates a new configured context object.
func New(config Config) (spec.Context, error) {
	newContext := &context{
		Config: config,

		ID: "",
	}

	if newContext.Context == nil {
		return nil, maskAnyf(invalidConfigError, "context must not be empty")
	}

	return newContext, nil
}

// MustNew creates either a new default configured context object, or panics.
func MustNew() spec.Context {
	newContext, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newContext
}

type context struct {
	Config

	ID string
}

func (c *context) Clone() spec.Context {
	// At first we create a new context with its own very unique ID, which will
	// not be cloned. All properties but the context ID must be cloned below.
	newContext := MustNew()

	// We prepare a new underlying context to have a fresh storage.
	newContext.(*context).Context = netcontext.Background()

	// Add the other information to the new underlying context.
	behaviourID, _ := c.GetBehaviourID()
	newContext.SetBehaviourID(behaviourID)
	clgName, _ := c.GetCLGName()
	newContext.SetCLGName(clgName)
	clgTreeID, _ := c.GetCLGTreeID()
	newContext.SetCLGTreeID(clgTreeID)
	expectation, _ := c.GetExpectation()
	newContext.SetExpectation(expectation)
	informationID, _ := c.GetInformationID()
	newContext.SetInformationID(informationID)
	sessionID, _ := c.GetSessionID()
	newContext.SetSessionID(sessionID)

	return newContext
}

func (c *context) Deadline() (time.Time, bool) {
	return c.Context.Deadline()
}

func (c *context) Done() <-chan struct{} {
	return c.Context.Done()
}

func (c *context) Err() error {
	return c.Context.Err()
}

func (c *context) GetBehaviourID() (string, bool) {
	if c.BehaviourID == "" {
		return "", false
	}

	return c.BehaviourID, true
}

func (c *context) GetCLGName() (string, bool) {
	if c.CLGName == "" {
		return "", false
	}

	return c.CLGName, true
}

func (c *context) GetCLGTreeID() (string, bool) {
	if c.CLGTreeID == "" {
		return "", false
	}

	return c.CLGTreeID, true
}

func (c *context) GetExpectation() (spec.Expectation, bool) {
	if c.Expectation == nil {
		return nil, false
	}

	return c.Expectation, true
}

func (c *context) GetID() string {
	return c.ID
}

func (c *context) GetInformationID() (string, bool) {
	if c.InformationID == "" {
		return "", false
	}

	return c.InformationID, true
}

func (c *context) GetSessionID() (string, bool) {
	if c.SessionID == "" {
		return "", false
	}

	return c.SessionID, true
}

func (c *context) SetBehaviourID(behaviourID string) {
	c.BehaviourID = behaviourID
}

func (c *context) SetCLGName(clgName string) {
	c.CLGName = clgName
}

func (c *context) SetCLGTreeID(clgTreeID string) {
	c.CLGTreeID = clgTreeID
}

func (c *context) SetExpectation(expectation spec.Expectation) {
	c.Expectation = expectation
}

func (c *context) SetInformationID(informationID string) {
	c.InformationID = informationID
}

func (c *context) SetSessionID(sessionID string) {
	c.SessionID = sessionID
}

func (c *context) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}
