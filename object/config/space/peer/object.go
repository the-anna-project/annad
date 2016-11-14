package peer

// New creates a new peer object. It provides configuration for peers within the
// connection space.
func New() *Object {
	return &object{}
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

func (o *Object) Configure() error {
	// Settings.

	return nil
}

func (o *Object) SetPosition(position string) {
	o.position = position
}

func (o *Object) Validate() error {
	// Settings.

	if len(o.position) == "" {
		return maskAnyf(invalidConfigError, "position must not be empty")
	}

	return nil
}
