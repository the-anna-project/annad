package spec

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
)

// TextInput provides a communication channel to send information sequences
// back to the client.
type TextInput interface {
	Configure() error

	// Channel returns a channel which is used to send text responses back to the
	// client.
	Channel() chan objectspec.TextInput

	Metadata() map[string]string

	Service() Collection

	SetServiceCollection(serviceCollection Collection)


}
