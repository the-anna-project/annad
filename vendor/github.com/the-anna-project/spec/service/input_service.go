package spec

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// InputService provides a communication channel for information sequences.
type InputService interface {
	Boot()
	Channel() chan objectspec.TextInput
	Metadata() map[string]string
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
}
