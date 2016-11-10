package spec

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

// Tracker represents a management object to track connection path patterns.
type Tracker interface {
	Configure() error

	// CLGIDs is a lookup function used by Track. It persists the single
	// connections between the destination and sources provided by networkPayload
	// in the underlying storage.
	CLGIDs(CLG CLG, networkPayload objectspec.NetworkPayload) error

	// CLGNames is a lookup function used by Track. It resolves the CLG names of
	// the destination and sources provided by networkPayload and persists the
	// single connections between them in the underlying storage.
	CLGNames(CLG CLG, networkPayload objectspec.NetworkPayload) error

	Service() Collection

	SetServiceCollection(sc Collection)

	SetStorageCollection(sc storagespec.Collection)

	Storage() storagespec.Collection

	// Track tracks connection path patterns.
	Track(CLG CLG, networkPayload objectspec.NetworkPayload) error

	Validate() error
}
