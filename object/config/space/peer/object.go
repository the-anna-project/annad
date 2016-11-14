package peer

// New creates a new peer object. It provides configuration for peers within the
// connection space.
func New() *Object {
	return &Object{}
}

type Object struct {
	// Settings.

	// position describes the default position of new peers within the connection
	// space.
	position string
}

func (o *Object) Position() string {
	return o.position
}

func (o *Object) SetPosition(position string) {
	o.position = position
}
