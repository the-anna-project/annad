package information

import (
	"sync"

	objectspec "github.com/xh3b4sd/anna/object/spec"
)

// New creates a new peer object having type information. An information peer
// represents a unique information sequence within the connection space of the
// neural network.
func New() objectspec.Peer {
	return &object{
		kind:  "information",
		mutex: sync.Mutex{},
		value: "",
	}
}

type object struct {
	// Settings.

	kind  string
	mutex sync.Mutex
	value string
}

func (o *object) Kind() string {
	return o.kind
}

func (o *object) SetValue(value string) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	o.value = value
}

func (o *object) Value() string {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.value
}
