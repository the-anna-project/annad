package core

import (
	"time"
)

type Neuron interface {
	SetState(state State)

	State() State

	Trigger(impuls Impuls) (Impuls, Connection, error)
}
