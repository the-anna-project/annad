package spec

// Tracker represents a management object to track connection path patterns.
type Tracker interface {
	// Track tracks connection path patterns.
	Track(CLG CLG, networkPayload NetworkPayload) error

	// TODO document the algorithm
	//
	// Database is empty. We have one new connection path. After adding the new
	// connection path we have one connection path stored in the database. The
	// event identifying this case is called new-path.
	//
	//     before
	//
	//     new       1,2
	//
	//     after     1,2
	//
	// Database holds one connection path. We have one new connection path. The
	// new connection path equals the connection path already being stored in the
	// database. The database state does not change. The event identifying this
	// case is called match-path.
	//
	//     before    1,2
	//
	//     new       1,2
	//
	//     after     1,2
	//
	// Database holds one connection path. We have one new connection path. The
	// first peer of the new connection path is already stored as last peer in the
	// database: 2. This causes the connection path of which one peer of the new
	// connection path is matching to be extended. The event identifying this case
	// is called extend-tail.
	//
	//     before    1,2
	//
	//     new       2,3
	//
	//     after     1,2,3
	//
	// Database holds one connection path. We have one new connection path. Two
	// peers of the new connection path are already stored in the database: 1,2.
	// The new connection path matches the head of the connection path already
	// being stored in the database. The database state does not change. The event
	// identifying this case is called match-head.
	//
	//     before    1,2,3
	//
	//     new       1,2
	//
	//     after     1,2,3
	//
	// Database holds three connection paths. We have one new connection path. One
	// peer of the new connection path is already stored in the database: 1. This
	// causes the connection path of which one peer of the new connection path is
	// matching to be extended. The event identifying this case is called
	// extend-head.
	//
	//     before    1,2,3
	//
	//     new       0,1
	//
	//     after     0,1,2,3
	//
	// Database holds one connection path. We have one new connection path. The
	// new connection path matches the body of the connection path already being
	// stored in the database. The database state does not change. The event
	// identifying this case is called match-body.
	//
	//     before    0,1,2,3
	//
	//     new       1,2
	//
	//     after     0,1,2,3
	//
	// Database holds one connection path. We have one new connection path. One
	// peer of the new connection path shows the connection path stored in the
	// database must be split. Here the peer 2 is already present, while the
	// connection path 2,4 is not. This pattern shows peer 2 is a CLG that
	// forwards to more than one other CLG. Hence we have to split the stored
	// connection path to always only have linear connection paths stored. After
	// splitting the existing connection path we have three connection paths
	// stored in the database. The event identifying this case is called
	// split-path.
	//
	//     before    0,1,2,3
	//
	//     new       2,4
	//
	//     after     0,1,2
	//               2,3
	//               2,4
	//
}
