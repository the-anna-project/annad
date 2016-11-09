package spec

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
)

// TextOutput provides a communication channel to send information sequences
// back to the client.
type TextOutput interface {
	// GetChannel returns a channel which is used to send text responses back to
	// the client.
	GetChannel() chan objectspec.TextOutput

	// GetMetadata returns the service's metadata.
	GetMetadata() map[string]string
}
