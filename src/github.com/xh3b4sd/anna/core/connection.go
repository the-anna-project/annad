package core

type Connection interface {
	SetState(state State)

	GetState() State

	Trigger(impulse Impulse) (Impulse, Neuron, error)
}
