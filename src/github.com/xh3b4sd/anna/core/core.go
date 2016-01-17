package core

import (
	"time"
)

type Core interface {
	Gateway() Gateway

	SetState(state State)

	State() State
}
