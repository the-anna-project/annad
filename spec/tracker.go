package spec

import "github.com/xh3b4sd/anna/object/spec"

// Tracker represents a management object to track connection path patterns.
type Tracker interface {
	// CLGIDs is a lookup function used by Track. It persists the single
	// connections between the destination and sources provided by networkPayload
	// in the underlying storage.
	CLGIDs(CLG CLG, networkPayload spec.NetworkPayload) error

	// CLGNames is a lookup function used by Track. It resolves the CLG names of
	// the destination and sources provided by networkPayload and persists the
	// single connections between them in the underlying storage.
	CLGNames(CLG CLG, networkPayload spec.NetworkPayload) error

	// Track tracks connection path patterns.
	Track(CLG CLG, networkPayload spec.NetworkPayload) error
}
