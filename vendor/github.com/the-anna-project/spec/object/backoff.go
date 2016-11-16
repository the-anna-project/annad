package spec

import (
	"time"
)

// Backoff represents the object managing backoff algorithms to retry actions.
type Backoff interface {
	// NextBackOff provides the duration expected to wait before retrying an
	// action. time.Duration = -1 indicates that no more retry should be
	// attempted.
	NextBackOff() time.Duration
	// Reset sets the backoff back to its initial state.
	Reset()
}
