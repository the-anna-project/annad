package spec

// Ctx represents a container for contextual information supposed to be carried
// around.
type Ctx interface {
	// GetKey returns a well configured key used to store and fetch data. The
	// returned key has the following scheme. "s" stands for the scope, that is,
	// an object type.  "c" stands for the context identifier, e.g. "default" or
	// "math". "<key>" stands for the key-value pair identifying the most
	// specific part of the key.
	//
	//     s:<object-type>:c:<context>:<key>
	//
	GetKey(f string, v ...interface{}) string

	Object

	// SetID sets the contect's ID. This is configurable because the context
	// object is a container for contextual information. So even the ID needs to
	// be configured when e.g. storing and fetching contextual information from
	// a database.
	//
	// Note that this needs to be well known. The configured context makes sure
	// that storage keys are consistently created. Customize this carefully and
	// make sure you know what you are doing.
	SetID(ID ObjectID)
}
