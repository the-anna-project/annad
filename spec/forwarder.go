package spec

// TODO
type Forwarder interface {
	Forward(CLG CLG, networkPayload NetworkPayload) error
}
