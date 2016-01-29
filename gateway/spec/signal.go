package gatewayspec

type Signal interface {
	Cancel()

	GetBytes(key string) ([]byte, error)

	GetError() error

	GetID() string

	GetObject(key string) (interface{}, error)

	GetResponder() (chan Signal, error)

	SetBytes(key string, bytes []byte)

	SetError(err error)

	SetObject(key string, object interface{})

	SetResponder(responder chan Signal)
}
