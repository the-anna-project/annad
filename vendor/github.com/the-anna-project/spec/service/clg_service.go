package spec

// CLGService represents the CLGs interacting with each other within the neural
// network. Each CLG  is registered in the neural network. From there signal are
// dispatched in a dynamic fashion until some useful calculation took place.
type CLGService interface {
	Boot()
	// GetCalculate returns the CLG's calculate function which implements its
	// actual business logic.
	GetCalculate() interface{}
	Metadata() map[string]string
	Service() ServiceCollection
	// SetServiceCollection configures the CLG's factory collection. This is done
	// for all CLGs, regardless if a CLG is making use of the factory collection
	// or not.
	SetServiceCollection(serviceCollection ServiceCollection)
}
