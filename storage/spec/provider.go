package spec

// Provider should be implemented by every object which wants to use
// storages. This then creates an API between storage implementations and
// storage users.
type Provider interface {
	Storage() Collection
}
