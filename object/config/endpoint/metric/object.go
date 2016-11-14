package metric

// New creates a new metric object. It provides configuration for the metric
// endpoint.
func New() *Object {
	return &Object{}
}

// Object represents the metric endpoint config object.
type Object struct {
	// Settings.

	address string
}

// Address returns the address of the endpoint config.
func (o *Object) Address() string {
	return o.address
}

// SetAddress sets the address for the endpoint config.
func (o *Object) SetAddress(address *string) {
	o.address = *address
}
