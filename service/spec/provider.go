package spec

// Provider should be implemented by every object which wants to use factories.
// This then creates an API between service implementations and service users.
type Provider interface {
	Service() Collection
}
