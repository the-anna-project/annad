package spec

// Ctx represents a container for contextual information supposed to be carried
// around.
type Ctx interface {
	// NetKey returns a well configured key used to store and fetch data. Keys
	// generated with NetKey should only be used by objects related to the neural
	// network scope. This can be e.g. the CharNet, or the StratNet. These
	// objects generate and structure dynamic information used to learn. The
	// returned key has the following scheme. "s" stands for the scope, that is,
	// the network scope. "o" stands for the object requesting the key. "<key>"
	// stands for the key-value pair identifying the most specific part of the
	// key.
	//
	//     s:net:o:<object>:<key>
	//
	NetKey(f string, v ...interface{}) string

	Object

	// SysKey returns a well configured key used to store and fetch data. Keys
	// generated with SysKey should only be used by objects related to the system
	// scope. This can be e.g. the Scheduler. These objects generate and
	// structure fundamental information used to manage the system. The returned
	// key has the following scheme. "s" stands for the scope, that is, the
	// system scope. "o" stands for the object requesting the key. "<key>" stands
	// for the key-value pair identifying the most specific part of the key.
	//
	//     s:sys:o:<object>:<key>
	//
	SysKey(f string, v ...interface{}) string
}
