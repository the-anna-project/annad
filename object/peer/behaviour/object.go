package behaviour

import (
	"sync"

	objectspec "github.com/xh3b4sd/anna/object/spec"
)

// New creates a new peer object having type behaviour. A behaviour peer
// represents a unique CLG tree within the connection space of the neural
// network.
func New() objectspec.Peer {
	return &object{}
}

type object struct {
	// Settings.

	kind  string
	mutex sync.Mutex
	value string
}

func (o *object) Configure() error {
	// Settings.

	o.kind = "behaviour"
	o.mutex = sync.Mutex{}

	return nil
}

func (o *object) Kind() string {
	return o.kind
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
