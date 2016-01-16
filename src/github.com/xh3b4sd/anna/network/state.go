package network

import (
	"time"
)

type State interface {
	// Capture provides a way to receive information that are supposed to be
	// included in the current state.
	Capture(v interface{}) State

	MarshalJSON() ([]byte, error)

	UnmarshalJSON([]byte) error
}

func NewState() State {
	return state{}
}

type state struct {
	// Blob is something received by Capture. This can either be everything we
	// want to gather state information for.
	Blob interface{} `json:"blob"`
}

func (n neuron) Age() time.Duration {
	return time.Since(n.CreatedAt)
}

func (s state) Capture(v interface{}) State {
	s.Blob = v
	return s
}

func (s state) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (s state) UnmarshalJSON([]byte) error {
	return nil
}
