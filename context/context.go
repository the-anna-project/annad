// Package context implements spec.Context and provides a wrapper of
// golang.org/x/net/context with additional business logic related to the Anna
// project.
package context

import (
	"time"

	netcontext "golang.org/x/net/context"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

type key string

const (
	behaviorIDKey    key = "behavior-id"
	clgTreeIDKey     key = "clg-tree-id"
	expectationIDKey key = "expectation"
	informationIDKey key = "information-id"
	sessionIDKey     key = "session-id"
)

// Config represents the configuration used to create a new context object.
type Config struct {
	// Settings.

	Context     netcontext.Context
	Expectation spec.Expectation
	SessionID   string

	// TODO we want to track the original input that was provided from the
	// outside. Further it would probably be interesting to also track the last 3
	// arguments of the current connection path.
}

// DefaultConfig provides a default configuration to create a new context
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Context:     netcontext.Background(),
		Expectation: nil,
		SessionID:   "",
	}

	return newConfig
}

// New creates a new configured context object.
func New(config Config) (spec.Context, error) {
	newContext := &context{
		Config: config,

		ID: string(id.MustNew()),
	}

	// If there is a session ID configured, we set it to the underlying context.
	// That way our standard configuration interface is obtained and the data
	// structures of the underlying implementation consistent.
	if config.SessionID != "" {
		newContext.SetSessionID(config.SessionID)
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

	// We set the session ID to our own context object and also make the
	// underlying context aware of it.
	newContext.(*context).SessionID = c.GetSessionID()
	newContext.SetSessionID(c.GetSessionID())

	// Add other information to the underlying context.
	newContext.SetCLGTreeID(c.GetCLGTreeID())
	newContext.SetExpectation(c.GetExpectation())

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

func (c *context) GetBehaviorID() string {
	behaviorID, ok := c.Context.Value(behaviorIDKey).(string)
	if ok {
		return behaviorID
	}

	return ""
}

func (c *context) GetCLGTreeID() string {
	clgTreeID, ok := c.Context.Value(clgTreeIDKey).(string)
	if ok {
		return clgTreeID
	}

	return ""
}

func (c *context) GetExpectation() spec.Expectation {
	expectation, ok := c.Context.Value(expectationIDKey).(spec.Expectation)
	if ok {
		return expectation
	}

	return nil
}

func (c *context) GetID() string {
	return c.ID
}

func (c *context) GetInformationID() string {
	informationID, ok := c.Context.Value(informationIDKey).(string)
	if ok {
		return informationID
	}

	return ""
}

func (c *context) GetSessionID() string {
	sessionID, ok := c.Context.Value(sessionIDKey).(string)
	if ok {
		return sessionID
	}

	return ""
}

func (c *context) SetBehaviorID(behaviorID string) {
	c.Context = netcontext.WithValue(c.Context, behaviorIDKey, behaviorID)
}

func (c *context) SetCLGTreeID(clgTreeID string) {
	c.Context = netcontext.WithValue(c.Context, clgTreeIDKey, clgTreeID)
}

func (c *context) SetExpectation(expectation spec.Expectation) {
	c.Context = netcontext.WithValue(c.Context, expectationIDKey, expectation)
}

func (c *context) SetInformationID(informationID string) {
	c.Context = netcontext.WithValue(c.Context, informationIDKey, informationID)
}

func (c *context) SetSessionID(sessionID string) {
	c.Context = netcontext.WithValue(c.Context, sessionIDKey, sessionID)
}

func (c *context) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}
