package tracker

import (
	"github.com/xh3b4sd/anna/spec"
)

type EventType string

var (
	EventExtendHead EventType = "extend-head"
	EventExtendTail EventType = "extend-tail"
	EventMatchGlob  EventType = "match-glob"
	EventMatchHead  EventType = "match-head"
	EventMatchPath  EventType = "match-path"
	EventMatchTail  EventType = "match-tail"
	EventNewPath    EventType = "new-path"
	EventSplitPath  EventType = "split-path"
)

// Event represents a container to track information about the tracking event
// currently happening.
type Event struct {
	// Connection represents the new connection currently being tracked.
	Connection string

	// ConnectionPath represents the stored connection path matching the new
	// connection according to the happening event.
	ConnectionPath string

	// Destination represents the destination of the network payload.
	Destination string

	// Source represents one source of the network payload.
	Source string

	// Type represents the event type currently happening.
	Type EventType
}

func initEvents(networkPayload spec.NetworkPayload) (map[string]Event, error) {
	events := map[string]Event{}

	d := string(networkPayload.GetDestination())
	for _, s := range networkPayload.GetSources() {
		events[cp] = Event{
			Connection:     string(s) + "," + d,
			ConnectionPath: "",
			Destination:    d,
			Source:         s,
			Type:           "",
		}
	}

	return events, nil
}

			var t EventType
			if cp == key {
				counter++
				t = EventMatchPath
			} else if strings.HasPrefix(key, cp) {
				t = EventMatchHead
			} else if strings.HasSuffix(key, cp) {
				t = EventMatchTail
			} else if strings.Contains(key, cp) {
				t = EventMatchGlob
			} else if strings.HasPrefix(key, meta.Source) {
				t = EventExtendHead
			} else if strings.HasSuffix(key, meta.Destination) {
				t = EventExtendTail
			} else if strings.HasSuffix(key, meta.Destination) {
				t = EventSplitPath
			} else {
				// When we get to here no other rule matched. That means
				counter++
				t = EventNewPath
			}

			events[cp].ConnectionPath = key
			events[cp].Type = t
