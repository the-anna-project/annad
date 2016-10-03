package spec

// TODO
type Forwarder interface {
	Forward(CLG CLG, networkPayload NetworkPayload) error
	GetBehaviourIDs(CLG spec.CLG, networkPayload spec.NetworkPayload) ([]string, error)
	NewBehaviourIDs(CLG spec.CLG, networkPayload spec.NetworkPayload) ([]string, error)
	ToInputCLG(CLG spec.CLG, networkPayload spec.NetworkPayload) error
}
