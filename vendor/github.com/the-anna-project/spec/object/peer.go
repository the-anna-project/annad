package spec

// Peer represents a unique member within the connection space of the neural
// network. It can either be of kind information or behaviour.
type Peer interface {
	Kind() string
	SetValue(value string)
	Value() string
}
