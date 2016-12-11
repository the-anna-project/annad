package service

// LayerService provides business logic to manage peers inside network layers.
// There are three different kinds of layer service implementations in use. Each
// implementation takes care about peers within the layer of its own
// responsibility. Connections across layers can be managed using the plain
// usage of the connection service (see ConnectionService).
//
//     behaviour
//
//         The behaviour layer service manages peers of the behaviour layer of
//         the connection space. Subject of this layer are peers associated with
//         behaviours. Note that behaviour is implemented in form of CLG
//         services (see CLGService).
//
//     information
//
//         The information layer service manages peers of the information layer
//         of the connection space. Subject of this layer are peers associated
//         with information. Note that information are provided in form of
//         input, which is received via input services (see InputService).
//
//     position
//
//         The position layer service manages peers of the position layer of the
//         connection space. Subject of this layer are peers representing the
//         position of behaviour and information peers.
//
type LayerService interface {
	Boot()
	// CreatePeer creates the peer of the current layer using the given peer
	// argument. CreatePeer also creates the position peer for the actual peer
	// being created and connects both.
	CreatePeer(peer string) (string, error)
	// DeletePeer deletes the peer of the current layer using the given peer
	// argument. DeletePeer also deletes the position peer for the actual peer
	// being deleted and disconnects both.
	DeletePeer(peer string) (string, error)
	Kind() string
	Metadata() map[string]string
	PeerPosition(peer string) (string, error)
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
}
