// Package ctx implements spec.Ctx to provide a container for contextual
// information.
package ctx

import (
	"fmt"
	"sync"

	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCtx represents the object type of the context object. This is
	// used e.g. to register itself to the logger.
	ObjectTypeCtx = "ctx"
)

// Config represents the configuration used to create a new context object.
type Config struct {
	// ID represents the context's ID.
	ID spec.ObjectID

	// Object represents the object configuring and using this context.
	Object spec.Object
}

// DefaultConfig provides a default configuration to create a new context
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		ID:     spec.ObjectID("default"),
		Object: nil,
	}

	return newConfig
}

// NewCtx creates a new configured context object.
func NewCtx(config Config) spec.Ctx {
	newCtx := &ctx{
		Config: config,
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeCtx,
	}

	return newCtx
}

type ctx struct {
	Config
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (c *ctx) GetKey(f string, v ...interface{}) string {
	// keyValuePair is supposed to hold a structured key representation delimited
	// by colons, e.g. "strategy:successes".
	keyValuePair := fmt.Sprintf(f, v...)
	return fmt.Sprintf("o:%s:c:%s:%s", c.Object.GetType(), c.GetID(), keyValuePair)
}

func (c *ctx) SetID(ID spec.ObjectID) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.ID = ID
}
