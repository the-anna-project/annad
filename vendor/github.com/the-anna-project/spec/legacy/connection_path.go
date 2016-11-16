package spec

// ConnectionPath represents a locatable object within the multi dimensional
// connection space. It provides functionality to qualify and calculate this
// location and the distance to other connection paths.
type ConnectionPath interface {
	// DistanceTo calculates the distance between the current and the provided
	// connection path by comparing their coordinates against each other. Paths
	// are normalized for comparison. The smaller path is aligned to the size of
	// the greater path by extending its outer vectors equally to both sides,
	// until its length is equally to the longer connection path. Once
	// normalized, the distance between each vector's dimension is measured and
	// summed up to the corresponding vector's dimension of the other path.  The
	// final distance represents the sum of all spin threads between the
	// connection path's vectors.
	//
	// Note that for this calculation dimensions within the connection space must
	// be defined by positive floating point numbers.
	DistanceTo(connectionPath ConnectionPath) float64

	// GetCoordinates returns the configured coordinates of the current
	// connection path.
	GetCoordinates() [][]float64

	// IsCloser calculates the distances of the two given connection paths
	// compared to the current connection path and compares the distance of a and
	// the distance of b against each other. The connection path being having the
	// smaller distance is qualified as being closer to the current connection
	// path and thus returned. In case the distance of a and b to the current
	// connection path is equal, one of a and b is chosen randomly.
	IsCloser(a, b ConnectionPath) (ConnectionPath, error)

	// String returns the string representation of the current connection path's
	// configured coordinates.
	String() (string, error)
}

// CLGTree TODO
type CLGTree interface {
	//
	GetValues() []CLGTree
}

// InputTree TODO
type InputTree interface {
	//
	GetValues() [][]string
}
