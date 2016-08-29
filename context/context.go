// Package context implements spec.Context and provides a wrapper of
// golang.org/x/net/context with additional business logic related to the Anna
// project.
package context

import (
	"time"

	contextpkg "golang.org/x/net/context"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

type key string

const (
	clgTreeIDKey key = "clg-tree-id"
	sessionIDKey key = "session-id"
)

// Config represents the configuration used to create a new context object.
type Config struct {
	// Settings.

	Context   contextpkg.Context
	SessionID string
}

// DefaultConfig provides a default configuration to create a new context
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Context:   contextpkg.Background(),
		SessionID: "",
	}

	return newConfig
}

// New creates a new configured context object.
func New(config Config) (spec.Context, error) {
	newContext := &context{
		Config: config,

		ID: string(id.MustNew()),
	}

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

func (c *context) Deadline() (time.Time, bool) {
	return c.Context.Deadline()
}

func (c *context) Done() <-chan struct{} {
	return c.Context.Done()
}

func (c *context) Err() error {
	return c.Context.Err()
}

func (c *context) GetCLGTreeID() string {
	clgTreeID, ok := c.Context.Value(clgTreeIDKey).(string)
	if ok {
		return clgTreeID
	}

	return ""
}

func (c *context) GetID() string {
	return c.ID
}

func (c *context) GetSessionID() string {
	sessionID, ok := c.Context.Value(sessionIDKey).(string)
	if ok {
		return sessionID
	}

	return ""
}

func (c *context) SetCLGTreeID(clgTreeID string) {
	c.Context = contextpkg.WithValue(c.Context, clgTreeIDKey, clgTreeID)
}

func (c *context) SetSessionID(sessionID string) {
	c.Context = contextpkg.WithValue(c.Context, sessionIDKey, sessionID)
}

func (c *context) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}
