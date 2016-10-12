package spec

// Tracker represents a management object to track connection path patterns.
type Tracker interface {
	// Track tracks connection path patterns.
	Track(CLG CLG, networkPayload NetworkPayload) error

	// TODO document the algorithm
	//
	// Database is empty. We have one new connection path. After adding the new
	// connection path we have one connection path stored in the database.
	//
	//     new       1,2
	//
	//                |
	//                v
	//
	//     stored
	//
	// Database holds one connection path. We have one new connection path. The
	// new connection path equals the connection path already being stored in the
	// database. The database state does not change.
	//
	//     new       1,2
	//
	//                |
	//                v
	//
	//     stored    1,2
	//
	// Database holds one connection path. We have one new connection path. One
	// peer of the new connection path is already stored in the database: 2. This
	// causes the connection path of which one peer of the new connection path is
	// matching to be extended.
	//
	//     new        2,3
	//
	//                 |
	//                 v
	//
	//     stored    1,2,3
	//
	// Database holds one connection path. We have one new connection path. Two
	// peers of the new connection path are already stored in the database: 2,3.
	// This causes the connection path of which two peers of the new connection
	// path are matching to be extended.
	//
	//     new        2,3,4
	//
	//                  |
	//                  v
	//
	//     stored    1,2,3,4
	//
	// Database holds one connection path. We have one new connection path. The
	// new connection path matches completely the connection path already being
	// stored in the database. The database state does not change.
	//
	//     new         2,3
	//
	//                  |
	//                  v
	//
	//     stored    1,2,3,4
	//
	// Database holds one connection path. We have one new connection path. One
	// peer of the new connection path shows the connection path stored in the
	// database must be split. Here the peer 3 is already present, while the
	// connection path 3,5 is not. This pattern shows peer 3 is a CLG that
	// forwards to more than one other CLG. Hence we have to split the stored
	// connection path to always only have linear connection paths stored. After
	// splitting the existing connection path we have three connection paths
	// stored in the database.
	//
	//     new        3,5
	//
	//                 |
	//                 v
	//
	//     stored    1,2,3
	//               3,4
	//               3,5
	//
	// Database holds one connection path. We have one new connection path. One
	// peer of the new connection path is already stored in the database: 1. This
	// causes the connection path of which one peer of the new connection path is
	// matching to be extended.
	//
	//     new        0,1
	//
	//                 |
	//                 v
	//
	//     stored    0,1,2
	//               2,3
	//               2,4
	//
	//
}
