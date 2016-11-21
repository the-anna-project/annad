package text

// New creates a new text object. It provides configuration for the text
// endpoint.
func New() *Object {
	return &Object{}
}

// Object represents the text endpoint config object.
type Object struct {
	// Settings.

	address string
	echo    bool
}

// Address returns the address of the endpoint config.
func (o *Object) Address() string {
	return o.address
}

// Echo returns the echo config of the endpoint config.
func (o *Object) Echo() bool {
	return o.echo
}

// SetAddress sets the address for the endpoint config.
func (o *Object) SetAddress(address *string) {
	o.address = *address
}

// SetEcho sets the echo config for the endpoint config.
func (o *Object) SetEcho(echo *bool) {
	o.echo = *echo
}
