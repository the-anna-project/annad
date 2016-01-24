package spec

type Signal interface {
	Cancel()

	GetBytes(key string) ([]byte, error)

	GetError() error

	GetID() string

	GetResponder() (chan Signal, error)

	SetBytes(key string, bytes []byte)

	SetError(err error)

	SetResponder(responder chan Signal)
}
