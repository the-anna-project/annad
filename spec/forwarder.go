package spec

// Forwarder represents an management layer to organize CLG forwarding rules.
// The forwarder obtains behaviour IDs for every single requested CLG of every
// possible CLG tree.
type Forwarder interface {
	FactoryProvider

	// Forward represents the public interface that bundles the following lookup
	// functions.
	//
	//     GetNetworkPayloads
	//     NewNetworkPayloads
	//
	// TODO
	Forward(CLG CLG, networkPayload NetworkPayload) error

	// GetMaxSignals returns the maximum number of signals being forwarded by one
	// CLG.
	GetMaxSignals() int

	GetNetworkPayloads(CLG spec.CLG, networkPayload spec.NetworkPayload) ([]NetworkPayload, error)

	NewNetworkPayloads(CLG spec.CLG, networkPayload spec.NetworkPayload) ([]NetworkPayload, error)

	ToInputCLG(CLG spec.CLG, networkPayload spec.NetworkPayload) error

	StorageProvider
}
