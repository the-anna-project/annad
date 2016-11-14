package connection

// New creates a new connection object. It provides configuration for the
// connection storage.
func New() *Object {
	return &Object{}
}

// Object represents the connection storage config object.
type Object struct {
	// Settings.

	address string
	kind    string
	prefix  string
}

// Address returns the address the connection storage is listening on.
func (o *Object) Address() string {
	return o.address
}

// Kind returns the kind of the connection storage.
func (o *Object) Kind() string {
	return o.kind
}

// Prefix returns the prefix used to prefix keys of the connection storage.
func (o *Object) Prefix() string {
	return o.prefix
}

// SetAddress sets the address for the connection storage config.
func (o *Object) SetAddress(address *string) {
	o.address = *address
}

// SetKind sets the kind for the connection storage config.
func (o *Object) SetKind(kind *string) {
	o.kind = *kind
}

// SetPrefix sets the prefix for the connection storage config.
func (o *Object) SetPrefix(prefix *string) {
	o.prefix = *prefix
}
