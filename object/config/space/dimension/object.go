package dimension

// New creates a new dimension object. It provides configuration for space
// dimensions.
func New() *Object {
	return &Object{}
}

// Object represents the space dimension config object.
type Object struct {
	// Settings.

	// count is the default number of directional coordinates within the
	// connection space. E.g. a dice has 3 dimensions.
	count int
	// depth is the default size of each directional coordinate within the
	// connection space. E.g. using a depth of 3, the resulting volume being taken
	// by a 3 dimensional space would be 9.
	depth int
}

// Count returns the count of the dimension config.
func (o *Object) Count() int {
	return o.count
}

// Depth returns the depth of the dimension config.
func (o *Object) Depth() int {
	return o.depth
}

// SetCount sets the count for the dimension config.
func (o *Object) SetCount(count *int) {
	o.count = *count
}

// SetDepth sets the depth for the dimension config.
func (o *Object) SetDepth(depth *int) {
	o.depth = *depth
}
