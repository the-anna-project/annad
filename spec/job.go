package spec

import (
	"encoding/json"
	"time"
)

// ActiveStatus represents a job's status indicating activation or
// deactivation.
type ActiveStatus string

// FinalStatus represents a job's status that is final. A job having this
// status will not change its status anymore.
type FinalStatus string

// Job represents a job that is executable by a Scheduler.
type Job interface {
	// GetActionID returns the job's action ID.
	GetActionID() string

	// GetActiveStatus returns the job's active status.
	GetActiveStatus() ActiveStatus

	// GetArgs returns the job's arguments.
	GetArgs() interface{}

	// GetCreatedAt returns the job's creation time.
	GetCreatedAt() time.Time

	// GetError returns the error returned during job execution, if any.
	GetError() error

	// GetFinalStatus returns the job's final status.
	GetFinalStatus() FinalStatus

	// GetResult returns the result returned during job execution, if any.
	GetResult() string

	// GetSessionID returns the job's session ID.
	GetSessionID() string

	json.Marshaler

	Object

	// SetActiveStatus sets the job's active status.
	SetActiveStatus(activeStatus ActiveStatus)

	// SetError sets the given error to the current job.
	SetError(err error)

	// SetFinalStatus sets the job's final status.
	SetFinalStatus(finalStatus FinalStatus)

	// SetResult sets the given result to the current job.
	SetResult(result string)

	json.Unmarshaler
}
