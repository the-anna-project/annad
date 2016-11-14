package text

// New creates a new text object. It provides configuration for the text
// endpoint.
func New() *Object {
	return &object{}
}

type Object struct {
	// Settings.

	Address string
}

func (o *Object) Address() string {
	return o.Address
}

func (o *Object) Configure() error {
	// Settings.

	return nil
}

func (o *Object) SetAddress(address string) {
	o.Address = address
}

func (o *Object) Validate() error {
	// Settings.

	if len(o.Address) == "" {
		return maskAnyf(invalidConfigError, "address must not be empty")
	}

	return nil
}
