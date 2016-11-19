package spec

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// OutputService provides a communication channel to send information sequences.
type OutputService interface {
	Boot()
	Channel() chan objectspec.TextOutput
	Metadata() map[string]string
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
}
