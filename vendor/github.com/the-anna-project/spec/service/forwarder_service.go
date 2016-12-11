package service

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// ForwarderService represents an management layer to organize CLG forwarding
// rules. The forwarder obtains behaviour IDs for every single requested CLG of
// every possible CLG tree.
type ForwarderService interface {
	Boot()
	// Forward represents the public interface that bundles the following lookup
	// functions.
	//
	//     GetNetworkPayloads
	//     News
	//
	// The network payloads being found by any of the lookup functions listed
	// above are queued by Forward.
	Forward(clgService CLGService, networkPayload objectspec.NetworkPayload) error
	// MaxSignals returns the maximum number of signals being forwarded by one
	// CLG.
	MaxSignals() int
	// GetNetworkPayloads tries to lookup behaviour IDs that can be used to
	// forward a certain network payload from the requested CLG to other CLGs. If
	// there are behaviour IDs found, a network payload for each behaviour ID is
	// created and the list of new network payloads is returned. If there could
	// not any behaviour ID be found, an error is returned.
	GetNetworkPayloads(clgService CLGService, networkPayload objectspec.NetworkPayload) ([]objectspec.NetworkPayload, error)
	// NewNetworkpayloads creates new connections to other CLGs in a pseudo random
	// manner. For each connection one network payload is created. The resulting
	// list is returned.
	NewNetworkpayloads(clgService CLGService, networkPayload objectspec.NetworkPayload) ([]objectspec.NetworkPayload, error)
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
}
