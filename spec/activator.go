package spec

// Activator represents an management layer to organize CLG activation rules.
// The activator obtains network payloads for every single requested CLG of
// every possible CLG tree. For more information on the activation mechanism see
// also Network.Activate.
type Activator interface {
	// Activate represents the public interface that bundles GetNetworkPayload and
	// NewNetworkPayload.
	Activate(CLG CLG, networkPayload NetworkPayload) (NetworkPayload, error)

	// GetNetworkPayload fetches the list of all queued network payloads of the
	// requested CLG from the underlying storage. The stored list is merged with
	// the given network payload. The resulting list is compared against the
	// single combination of requesting behavior IDs known to be already
	// successful in the past. In case list of network payloads contains network
	// payloads sent by the CLGs listed in the single combination of behavior IDs
	// stored for the requested CLG, the interface of the requested CLG is
	// fulfilled. Then a new network payload is created by merging the matching
	// network payloads of the stored queue. The matching network payloads are
	// then removed from the queue and the created network payload consisting of
	// the queued network payloads is returned. The modifications of the updated
	// queue are also persisted. In case no successful combination of behavior IDs
	// is stored, or no match using the stored configuration associated with the
	// requested CLG can be found, an error is returned.
	GetNetworkPayload(networkPayload spec.NetworkPayload, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error)

	// NewNetworkPayload fetches the list of all queued network payloads of the
	// requested CLG from the underlying storage. The stored list is merged with
	// the given network payload. The resulting list is used to find a combination
	// of network payloads that fulfill the interface of the requested CLG. This
	// creation process may be random or biased in some way.
	NewNetworkPayload(networkPayload spec.NetworkPayload, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error)
}
