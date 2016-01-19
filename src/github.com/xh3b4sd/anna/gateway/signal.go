package gateway

type Signal interface {
	GetBytes() []byte
	GetResponder() chan []byte
	SetBytes(bytes []byte)
	SetResponder(responder chan []byte)
}

func NewSignal(bytes []byte) Signal {
	return signal{
		Bytes:     bytes,
		Responder: make(chan []byte),
	}
}

type signal struct {
	Bytes     []byte
	Responder chan []byte
}

func (s signal) GetBytes() []byte {
	return s.Bytes
}

func (s signal) GetResponder() chan []byte {
	return s.Responder
}

func (s signal) SetBytes(bytes []byte) {
	s.Bytes = bytes
}

func (s signal) SetResponder(responder chan []byte) {
	s.Responder = responder
}
