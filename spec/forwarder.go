package spec

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

// Forwarder represents an management layer to organize CLG forwarding rules.
// The forwarder obtains behaviour IDs for every single requested CLG of every
// possible CLG tree.
type Forwarder interface {
	// Forward represents the public interface that bundles the following lookup
	// functions.
	//
	//     GetNetworkPayloads
	//     News
	//
	// The network payloads being found by any of the lookup functions listed
	// above are queued by Forward.
	Forward(CLG CLG, networkPayload objectspec.NetworkPayload) error

	// GetMaxSignals returns the maximum number of signals being forwarded by one
	// CLG.
	GetMaxSignals() int

	// GetNetworkPayloads tries to lookup behaviour IDs that can be used to
	// forward a certain network payload from the requested CLG to other CLGs. If
	// there are behaviour IDs found, a network payload for each behaviour ID is
	// created and the list of new network payloads is returned. If there could
	// not any behaviour ID be found, an error is returned.
	GetNetworkPayloads(CLG CLG, networkPayload objectspec.NetworkPayload) ([]objectspec.NetworkPayload, error)

	// News creates new connections to other CLGs in a pseudo random
	// manner. For each connection one network payload is created. The resulting
	// list is returned.
	News(CLG CLG, networkPayload objectspec.NetworkPayload) ([]objectspec.NetworkPayload, error)

	servicespec.Provider

	storagespec.Provider
}
