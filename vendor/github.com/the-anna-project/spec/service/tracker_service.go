package service

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// TrackerService represents a management object to track connection path
// patterns.
type TrackerService interface {
	Boot()
	// CLGIDs is a lookup function used by Track. It persists the single
	// connections between the destination and sources provided by networkPayload
	// in the underlying storage.
	CLGIDs(clgService CLGService, networkPayload objectspec.NetworkPayload) error
	// CLGNames is a lookup function used by Track. It resolves the CLG names of
	// the destination and sources provided by networkPayload and persists the
	// single connections between them in the underlying storage.
	CLGNames(clgService CLGService, networkPayload objectspec.NetworkPayload) error
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
	// Track tracks connection path patterns.
	Track(clgService CLGService, networkPayload objectspec.NetworkPayload) error
}
