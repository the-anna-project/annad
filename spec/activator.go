package spec

type Activater interface {
	Activate(CLG CLG, networkPayload NetworkPayload) (NetworkPayload, error)
}
