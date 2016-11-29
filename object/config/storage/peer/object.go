package peer

// New creates a new peer object. It provides configuration for the peer
// storage.
func New() *Object {
	return &Object{}
}

// Object represents the peer storage config object.
type Object struct {
	// Settings.

	address string
	kind    string
	prefix  string
}

// Address returns the address the peer storage is listening on.
func (o *Object) Address() string {
	return o.address
}

// Kind returns the kind of the peer storage.
func (o *Object) Kind() string {
	return o.kind
}

// Prefix returns the prefix used to prefix keys of the peer storage.
func (o *Object) Prefix() string {
	return o.prefix
}

// SetAddress sets the address for the peer storage config.
func (o *Object) SetAddress(address *string) {
	o.address = *address
}

// SetKind sets the kind for the peer storage config.
func (o *Object) SetKind(kind *string) {
	o.kind = *kind
}

// SetPrefix sets the prefix for the peer storage config.
func (o *Object) SetPrefix(prefix *string) {
	o.prefix = *prefix
}
