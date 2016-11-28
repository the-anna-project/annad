package service

// LayerService provides business logic to manage connections inside network
// layers. There are two different kinds of layer service implementations in
// use. Each implementation takes care about connections between peers within
// the layer of its own responsibility. Connections across layers can be managed
// using the plain usage of the connection service (see ConnectionService).
//
//     behaviour
//
//         The behaviour layer service manages connections of the behaviour
//         layer of the connection space. Subject of this layer are connections
//         between behaviours. Note that behaviour is implemented in form of CLG
//         services (see CLGService).
//
//     information
//
//         The information layer service manages connections of the information
//         layer of the connection space. Subject of this layer are connections
//         between information. Note that information are provided in form of
//         input, which is received via input services (see InputService).
//
type LayerService interface {
	Boot()
	CreateConnection(peerA, peerB string) error
	DeleteConnection(peerA, peerB string) error
	Metadata() map[string]string
	Service() ServiceCollection
	SetKind(kind string)
	SetServiceCollection(serviceCollection ServiceCollection)
}
