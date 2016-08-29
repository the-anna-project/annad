package spec

import (
	"golang.org/x/net/context"
)

// Context represents a container holding scope specific information.
type Context interface {
	context.Context

	// GetCLGTreeID returns the clg tree ID of the current Context.
	GetCLGTreeID() string

	// GetID returns the context's ID representing the very unique scope of its
	// own lifetime. This can be useful for e.g. gathering logs bound to one
	// request going through multiple independent sub-systems.
	GetID() string

	// GetSessionID returns the session ID of the current Context.
	GetSessionID() string

	// SetCLGTreeID sets the given clg tree ID to the current Context.
	SetCLGTreeID(clgTreeID string)

	// SetSessionID sets the given session ID to the current Context.
	SetSessionID(sessionID string)
}
