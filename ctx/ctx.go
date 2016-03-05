// Package ctx implements spec.Ctx to provide a container for contextual
// information.
package ctx

import (
	"fmt"
	"sync"

	"github.com/xh3b4sd/anna/spec"
)

const (
	ObjectTypeCtx = "ctx"
)

// Config represents the configuration used to create new context.
type Config struct {
	// ID represents the context's ID. This is configurable because the context
	// object is a container for contextual information. So even the ID needs to
	// be configured when e.g. storing and fetching contextual information from a
	// database.
	//
	// Note that this needs to be well known. The configured context makes sure
	// that storage keys are consistently created. Customize this carefully and
	// make sure you know what you are doing.
	ID spec.ObjectID

	// Object represents the object configuring and using this context.
	Object spec.Object
}

// DefaultConfig provides a default configuration to create new contexts by
// best effort.
func DefaultConfig() Config {
	newConfig := Config{
		ID:     spec.ObjectID("default"),
		Object: nil,
	}

	return newConfig
}

// NewCtx creates a new configured context.
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
