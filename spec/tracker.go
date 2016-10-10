package spec

// Tracker represents a management object to track connection path patterns and
// annotate network payloads accordingly.
type Tracker interface {
	// Track tracks connection path patterns and annotate network payloads
	// accordingly.
	Track(CLG CLG, networkPayload NetworkPayload) (NetworkPayload, error)
}
