package spec

// Impulse is basically a container to carry around information. It walks
// through neural networks and is modified on its way to finally creating some
// output based on some input.
type Impulse interface {
	// GetInput returns the impulse's input.
	GetInput() (string, error)

	// GetOutput returns the impulse's output.
	GetOutput() (string, error)

	// GetCtx always returns a context associated with the given object. In case
	// there is no context associated with the given object, a new context is
	// created, stored and returned. In case there is already a context known for
	// the given object, this one is simply returned.
	GetCtx(object Object) Ctx

	Object

	// SetInput sets the impulse's input.
	SetInput(input string) error

	// SetOutput sets the impulse's output.
	SetOutput(output string) error
}
