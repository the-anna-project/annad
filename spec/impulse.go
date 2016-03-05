package spec

// Impulse is basically a container to carry around information. It walks
// through neural networks and is modified on its way to finally creating some
// output based on some input.
type Impulse interface {
	// AddObjectType pushes the given object type at the end of the impulse's
	// list of object types. Usually these object types come from Strategies.
	AddObjectType(objectType ObjectType) error

	// GetInput returns the impulse's input.
	GetInput() (string, error)

	// GetOutput returns the impulse's output.
	GetOutput() (string, error)

	// GetObjectType pops the next object type out of the impulse's list of
	// object types. Usually these object types come from Strategies.
	GetObjectType() (ObjectType, error)

	// GetCtx always returns a context associated with the given object. In case
	// there is no context associated with the given object, a new context is
	// created, stored and returned. In case there is already a context known for
	// the given object, this one is simply returned.
	GetCtx(object Object) Ctx

	Object

	// SetID sets the object ID of the impulse. This is the only exception across
	// all other objects. Only the impulse is allowed to modify its ID to make
	// the job network work.
	//
	// TODO having that said, it feels wrong. There should just be some job ID if
	// necessary.
	SetID(ID ObjectID) error

	// SetInput sets the impulse's input.
	SetInput(input string) error

	// SetOutput sets the impulse's output.
	SetOutput(output string) error
}
