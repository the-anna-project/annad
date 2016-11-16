package spec

import (
	objectspec "github.com/the-anna-project/spec/object"
)

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
// persisted connections.
//
//     Storage key of the peer "sum" and its corresponding hash map value.
//
//         peer:sum
//
//             created     1478992355
//             kind        behaviour
//             position    432.8,4342,54.334
//             updated     1478992355
//
//     Storage key of the peer "number" and its corresponding hash map value.
//
//         peer:number
//
//             created     1478992355
//             kind        information
//             position    432.8,4342,54.334
//             updated     1478992355
//
//     Storage key of the connection between the peer "sum" and "number" and its
//     corresponding hash map value. The two given peers being used to create
//     the connection key are ordered alpha numerically beforehand, regardless
//     of their kind.
//
//         peer:number:peer:sum
//
//             created    1478992355
//             updated    1478992355
//             weight     278.0082
//
//     Backreference from positions to peers and its corresponding list value.
//
//         position:432.8,4342,54.334    peer,peer
//
type ConnectionService interface {
	Boot()
	// Create manages the creation of a connection.
	//
	//     peer
	//     peer
	//     connection
	//     if not exist
	//     ensure transaction
	//
	Create(a, b objectspec.Peer) error
	Metadata() map[string]string
	Service() ServiceCollection
	SetDimensionCount(dimensionCount int)
	SetDimensionDepth(dimensionDepth int)
	SetServiceCollection(serviceCollection ServiceCollection)
	SetWeight(weight int)
}
