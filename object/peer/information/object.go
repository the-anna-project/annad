package information

import (
	"sync"

	objectspec "github.com/xh3b4sd/anna/object/spec"
)

// New creates a new peer object having type information. An information peer
// represents a unique information sequence within the connection space of the
// neural network.
func New() objectspec.Peer {
	return &object{}
}

type object struct {
	// Settings.

	mutex sync.Mutex
	value string
}

func (o *object) Configure() error {
	// Settings.

	o.metadata = map[string]string{
		"kind": "information",
		"name": "peer",
		"type": "object",
	}
	o.mutex = sync.Mutex{}

	return nil
}

func (o *object) Kind() string {
	return o.metadata["kind"]
}

func (o *object) SetValue(value string) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	o.value = value
}

func (o *object) Validate() error {
	// Settings.

	if len(o.value) == "" {
		return maskAnyf(invalidConfigError, "value must not be empty")
	}

	return nil
}

func (o *object) Value() string {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.value
}
