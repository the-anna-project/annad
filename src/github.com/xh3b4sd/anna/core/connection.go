package core

type Connection interface {
	SetState(state State)

	State() State

	Trigger(impuls Impuls) (Impuls, Neuron, error)
}
