package peer

// New creates a new peer object. It provides configuration for peers within the
// connection space.
func New() *Object {
	return &Object{}
}

// Object represents the space peer config object.
type Object struct {
	// Settings.

	// position describes the default position of new peers within the connection
	// space.
	position string
}

// Position returns the position of the peer config.
func (o *Object) Position() string {
	return o.position
}

// SetPosition sets the position for the peer config.
func (o *Object) SetPosition(position *string) {
	o.position = *position
}
