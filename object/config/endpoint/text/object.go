package text

// New creates a new text object. It provides configuration for the text
// endpoint.
func New() *Object {
	return &Object{}
}

type Object struct {
	// Settings.

	address string
}

func (o *Object) Address() string {
	return o.address
}

func (o *Object) SetAddress(address string) {
	o.address = address
}
