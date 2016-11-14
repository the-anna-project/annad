package connection

// New creates a new connection object. It provides configuration for the
// connection storage.
func New() *Object {
	return &object{}
}

type Object struct {
	// Settings.

	address string
	kind    string
	prefix  string
}

func (o *Object) Address() string {
	return o.address
}

func (o *Object) Configure() error {
	// Settings.

	return nil
}

func (o *Object) Kind() string {
	return o.kind
}

func (o *Object) Prefix() string {
	return o.prefix
}

func (o *Object) SetAddress(address string) {
	o.address = address
}

func (o *Object) SetKind(kind string) {
	o.kind = kind
}

func (o *Object) SetPrefix(prefix string) {
	o.prefix = prefix
}

func (o *Object) Validate() error {
	// Settings.

	// TODO more precise validation for address
	if len(o.address) == "" {
		return maskAnyf(invalidConfigError, "address must not be empty")
	}
	// TODO more precise validation for kind
	if len(o.kind) == "" {
		return maskAnyf(invalidConfigError, "kind must not be empty")
	}
	if len(o.prefix) == "" {
		return maskAnyf(invalidConfigError, "prefix must not be empty")
	}

	return nil
}
