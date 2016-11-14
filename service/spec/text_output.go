package spec

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
)

// TextOutput provides a communication channel to send information sequences
// back to the client.
type TextOutput interface {
	Boot()
	// Channel returns a channel which is used to send text responses back to the
	// client.
	Channel() chan objectspec.TextOutput
	Metadata() map[string]string
	Service() Collection
	SetServiceCollection(serviceCollection Collection)
}
