// TODO comment
package event

type Type string

// TODO create events in separate files
var ExtendTail Type = "extend-tail"
var MatchBody Type = "match-body"
var MatchHead Type = "match-head"
var MatchPath Type = "match-path"
var MatchTail Type = "match-tail"
var NewPath Type = "new-path"
var SplitPath Type = "split-path"

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
