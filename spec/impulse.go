package spec

type Impulse interface {
	AddObjectType(objectType ObjectType) error

	GetInput() (string, error)

	GetOutput() (string, error)

	GetObjectType() (ObjectType, error)

	Object

	SetID(ID ObjectID) error

	SetInput(input string) error

	SetOutput(output string) error
}
