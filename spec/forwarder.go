package spec

type Forwarder interface {
	Forward(CLG CLG, networkPayload NetworkPayload) error
}
