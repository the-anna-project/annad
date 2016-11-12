package spec

// Connection represents a service being able to manage connections within the
// connection space.
type Connection interface {
	Configure() error
	Create(a, b string) error
	Metadata() map[string]string
	Service() Collection
	SetServiceCollection(sc Collection)
	Validate() error
}
