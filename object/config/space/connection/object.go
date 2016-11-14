package connection

// New creates a new connection object. It provides configuration for the
// connection space.
func New() *Object {
	return &Object{}
}

// Object represents the connection space config object.
type Object struct {
	// Settings.

	// weight is the default score applied to a connection expressing its
	// importance.
	weight int
}

// Weight returns the weight of the connection config.
func (o *Object) Weight() int {
	return o.weight
}

// SetWeight sets the weight for the connection config.
func (o *Object) SetWeight(weight *int) {
	o.weight = *weight
}
