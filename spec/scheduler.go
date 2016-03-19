package spec

// Action represents any work to be done when executing a job.
//
// The given args represents some arbitrary structure being configured at job
// creation and passed through during job execution. The given closer signals
// the jobs end when the scheduler decides to do so. This happens e.g. when the
// job's final status was manually set. That way an operator can finish a task
// and cancel the corresponding goroutine by setting the job's final state in
// the used storage.
//
// In case the work done within the action callback was successful, there might
// be some result returned. Then this is accessible using Job.GetResult.
//
// In case there was an error returned, it will be accessible using
// Job.GetError. Note that this only obtains the error's message, not the
// underlying structure of the original error object.
type Action func(args interface{}, closer <-chan struct{}) (string, error)

// Scheduler represents a scheduler object that manages jobs and their
// executions.
type Scheduler interface {
	// Boot initializes and starts the whole scheduler like booting a machine.
	// The call to Boot blocks until the scheduler is completely initialized, so
	// you might want to call it in a separate goroutine.
	Boot()

	// Execute registers and executes a new job object configured as given. An
	// error is immediately returned in case the configured action is not
	// registered to the scheduler.
	Execute(job Job) error

	// FetchJob fetches the job for the given job ID.
	FetchJob(jobID ObjectID) (Job, error)

	// MarkAsActive marks the given job as active and persists the job's state.
	// The returned job is actually the refreshed version of the provided one.
	MarkAsActive(job Job) (Job, error)

	// MarkAsFailedWithError marks the given job as failed, adds information of
	// the given error to it and persists the job's state. The returned job is
	// actually the refreshed version of the provided one.
	MarkAsFailedWithError(job Job, err error) (Job, error)

	// MarkAsInactive marks the given job as inactive and persists the job's
	// state. The returned job is actually the refreshed version of the provided
	// one.
	MarkAsInactive(job Job) (Job, error)

	// MarkAsSucceeded marks the job object as succeeded and persists its state.
	// The returned job object is actually the refreshed version of the provided
	// one.
	MarkAsSucceeded(job Job) (Job, error)

	Object

	// PersistJob writes the given job to the configured Storage.
	PersistJob(job Job) error

	// Register registers the given action using the given actionID.
	Register(actionID string, action Action)

	// WaitForFinalStatus blocks and waits for the given job to reach a final
	// status. The given closer can end the waiting and thus stop blocking the
	// call to WaitForFinalStatus.
	WaitForFinalStatus(jobID ObjectID, closer <-chan struct{}) (Job, error)
}
