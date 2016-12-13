// Package context implements golang.org/x/net/context.Context and provides
// marshallable context primitives to distribute information across event
// queues.
package context

import (
	nativecontext "context"
	"time"
)

// Config represents the configuration used to create a new context.
type Config struct {
	// Settings.
	Context nativecontext.Context
}

// DefaultConfig provides a default configuration to create a new context by
// best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Context: nativecontext.Background(),
	}

	return newConfig
}

// New creates a new configured context object.
func New(config Config) (Context, error) {
	// Settings.
	if config.Context == nil {
		return nil, maskAnyf(invalidConfigError, "context must not be empty")
	}

	newContext := &context{
		// Internals.
		storage: map[interface{}]interface{}{},

		// Settings.
		context: config.Context,
	}

	return newContext, nil
}

type context struct {
	// Internals.
	storage map[interface{}]interface{}

	// Settings.
	context nativecontext.Context
}

func (c *context) Deadline() (time.Time, bool) {
	return c.context.Deadline()
}

func (c *context) Done() <-chan struct{} {
	return c.context.Done()
}

func (c *context) Err() error {
	return c.context.Err()
}

func (c *context) Value(key interface{}) interface{} {
	v, ok := c.storage[key]
	if ok {
		return v
	}

	return nil
}
