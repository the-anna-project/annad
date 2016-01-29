package spec

import (
	"encoding/json"
)

type Core interface {
	// Boot initializes and starts the whole core like booting a machine. The
	// call to Boot blocks until the core is completely initialized, so you might
	// want to call it in a separate goroutine.
	Boot()

	json.Unmarshaler

	Object

	// Shutdown ends all processes of the core like shutting down a machine. The
	// call to Boot blocks until the core is completely shut down, so you might
	// want to call it in a separate goroutine.
	Shutdown()

	Trigger(imp Impulse) (Impulse, error)
}
