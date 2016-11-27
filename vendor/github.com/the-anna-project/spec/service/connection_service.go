package service

// ConnectionService represents a service being able to manage connections
// within the connection space.
//
// Following is some explanation about the wording being used in documentation
// and code.
//
//     Coordinate
//
//         Coordinate is a point on a single dimension.
//
//     Depth
//
//         Depth is the size of each directional coordinate within the
//         connection space. E.g. using a Depth of 3, the resulting volume being
//         taken by a 3 dimensional space would be 9.
//
//     Dimension
//
//         Dimensions is the number of directional coordinates within the
//         connection space. E.g. a dice has 3 dimensions.
//
//     Position
//
//         Position is a point within the connection space. It is described by
//         one coordinate for each dimension of the connection space.
//
//     Weight
//
//         Weight is the score applied to a connection expressing its
//         importance.
//
// Following is an example of a possible storage key structure, to illustrate
// persisted connections. The given peers being used to create the peer and
// connection keys are ordered as such, that they form directed and reproducible
// storage keys. This means a connection can only be resolved into a certain,
// desired direction.
//
// This is a peer pointing to other peers. The key on the left is a string
// being formed by a prefix and a peer value. The value on the right is a
// list of strings being formed by peer values.
//
//     peer:sum     $id1
//     peer:$id1    sum,$id2,$id4
//     peer:$id2    $id3
//     peer:$id4    $id5
//
// This is a connection holding metadata. The key on the left is a string
// being formed by a prefix and the values of the two peers forming the
// connection. The order of the peers within the key expresses the
// connection direction. The value on the right is a map of strings.
//
//     connection:sum:$id1     weight 23.775
//     connection:$id1:sum     weight 23.775
//     connection:$id1:$id2    weight 23.775
//     connection:$id1:$id4    weight 23.775
//     connection:$id2:$id3    weight 23.775
//     connection:$id4:$id5    weight 23.775
//
// Following is a list of properties each peer has applied in form of
// connections to itself.
//
//     created
//     kind
//     position
//     updated
//
type ConnectionService interface {
	Boot()
	Create(peerA, peerB string) error
	Find(peerA, peerB string) (map[string]string, error)
	FindPeers(peer string) ([]string, error)
	Metadata() map[string]string
	Service() ServiceCollection
	SetDimensionCount(dimensionCount int)
	SetDimensionDepth(dimensionDepth int)
	SetServiceCollection(serviceCollection ServiceCollection)
	SetWeight(weight int)
}
