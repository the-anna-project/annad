// Package ctx implements spec.Ctx to provide a container for contextual
// information.
package ctx

import (
	"fmt"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCtx represents the object type of the context object. This is
	// used e.g. to register itself to the logger.
	ObjectTypeCtx = "ctx"
)

// Config represents the configuration used to create a new context object.
type Config struct {
	// Object represents the object configuring and using this context.
	Object spec.Object
}

// DefaultConfig provides a default configuration to create a new context
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Object: nil,
	}

	return newConfig
}

// NewCtx creates a new configured context object.
func NewCtx(config Config) spec.Ctx {
	newCtx := &ctx{
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeCtx,
	}

	return newCtx
}

type ctx struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

const (
	// NetKeyFormat represents the format used to create storage keys for the
	// network scope.
	NetKeyFormat = "s:net:%s:%s"
)

func (c *ctx) NetKey(f string, v ...interface{}) string {
	return fmt.Sprintf(NetKeyFormat, c.Object.GetType(), fmt.Sprintf(f, v...))
}

const (
	// SysKeyFormat represents the format used to create storage keys for the
	// system scope.
	SysKeyFormat = "s:sys:%s:%s"
)

func (c *ctx) SysKey(f string, v ...interface{}) string {
	return fmt.Sprintf(SysKeyFormat, c.Object.GetType(), fmt.Sprintf(f, v...))
}
