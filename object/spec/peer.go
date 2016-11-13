package spec

// Peer represents a unique member within the connection space of the neural
// network. It can either be of kind information or behaviour.
type Peer interface {
	Configure() error
	Kind() string
	SetValue(value string)
	Validate() error
	Value() string
}
