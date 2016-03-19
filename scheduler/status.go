package scheduler

import (
	"github.com/xh3b4sd/anna/spec"
)

const (
	// StatusStarted represents the status of a running job.
	StatusStarted spec.ActiveStatus = "started"

	// StatusStopped represents the status of a stopped job, that has not been
	// started yet.
	StatusStopped spec.ActiveStatus = "stopped"
)

const (
	// StatusFailed represents the status of a job where the action returned an
	// error.
	StatusFailed spec.FinalStatus = "failed"

	// StatusReplaced represents the status of a job where another job was
	// started for the same session ID.
	StatusReplaced spec.FinalStatus = "replaced"

	// StatusSucceeded represents the status of a job where the action returned
	// without an error.
	StatusSucceeded spec.FinalStatus = "succeeded"
)

// HasFailedStatus determines whether a job has failed or not. Note that this
// is about a final status.
func HasFailedStatus(job spec.Job) bool {
	if job.GetActiveStatus() == StatusStopped && job.GetFinalStatus() == StatusFailed {
		return true
	}

	return false
}

// HasFinalStatus determines whether a job has a final status or not.
func HasFinalStatus(job spec.Job) bool {
	if HasFailedStatus(job) || HasReplacedStatus(job) || HasSucceededStatus(job) {
		return true
	}

	return false
}

// HasReplacedStatus determines whether a job was replaced or not. Note that
// this is about a final status.
func HasReplacedStatus(job spec.Job) bool {
	if job.GetActiveStatus() == StatusStopped && job.GetFinalStatus() == StatusReplaced {
		return true
	}

	return false
}

// HasSucceededStatus determines whether a job has succeeded or not. Note that
// this is about a final status.
func HasSucceededStatus(job spec.Job) bool {
	if job.GetActiveStatus() == StatusStopped && job.GetFinalStatus() == StatusSucceeded {
		return true
	}

	return false
}
