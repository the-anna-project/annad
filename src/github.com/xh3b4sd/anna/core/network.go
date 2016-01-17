package core

import (
	"time"
)

type Network interface {
	SetState(state State)

	State() State

	Trigger(impuls Impuls) (Impuls, Connection, error)
}
