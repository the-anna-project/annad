package spec

type Signal interface {
	GetError() error

	GetID() string

	GetInput() interface{}

	GetOutput() interface{}

	GetResponder() chan Signal

	SetError(err error)

	SetID(ID string)

	SetInput(input interface{})

	SetOutput(output interface{})
}
