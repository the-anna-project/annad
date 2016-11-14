package connection

// New creates a new connection object. It provides configuration for the
// connection space.
func New() *Object {
	return &object{}
}

type Object struct {
	// Settings.

	// weight is the default score applied to a connection expressing its
	// importance.
	weight int
}

func (o *Object) Weight() int {
	return o.weight
}

func (o *Object) Configure() error {
	// Settings.

	return nil
}

func (o *Object) SetWeight(weight int) {
	o.weight = weight
}

func (o *Object) Validate() error {
	// Settings.

	return nil
}
