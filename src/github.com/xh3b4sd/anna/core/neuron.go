package core

type Neuron interface {
	SetState(state State)

	GetState() State

	Trigger(impulse Impulse) (Impulse, Connection, error)
}
