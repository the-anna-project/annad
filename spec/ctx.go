package spec

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

	SetID(ID ObjectID)
}
