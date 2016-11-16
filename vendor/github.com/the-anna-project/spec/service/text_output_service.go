package spec

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// TextOutputService provides a communication channel to send information
// sequences back to the client.
type TextOutputService interface {
	Boot()
	// Channel returns a channel which is used to send text responses back to the
	// client.
	Channel() chan objectspec.TextOutput
	Metadata() map[string]string
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
}
