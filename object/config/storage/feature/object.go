package feature

// New creates a new feature object. It provides configuration for the feature
// storage.
func New() *Object {
	return &Object{}
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
