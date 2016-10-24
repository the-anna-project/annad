// TODO comment
package event

type Type string

// Event represents a container to track information about the event currently
// being tracked.
type Event interface {
	// Connection represents the new connection currently being tracked.
	GetConnection() string

	// ConnectionPath represents the stored connection path matching the new
	// connection according to the happening event.
	GetConnectionPath() string

	// Destination represents the destination of the network payload.
	GetDestination() string

	// Source represents one source of the network payload.
	GetSource() string

	// Type represents the event type currently happening.
	GetType() Type
}
