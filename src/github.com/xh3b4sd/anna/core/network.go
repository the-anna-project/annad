package core

type Network interface {
	SetState(state State)

	GetState() State

	Trigger(impulse Impulse) (Impulse, Connection, error)
}
