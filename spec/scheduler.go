package spec

// Action represents any work to be done when executing a job.
//
// The given closer signals the jobs end when the scheduler decides to do so.
// This happens e.g.  when the job's final status was manually set. That way an
// operator can finish a task and cancel the corresponding goroutine by setting
// the job's final state in the used storage.
//
// In case the work done within the action callback was successful, there might
// be some result returned. Then this is accessible using Job.GetResult.
//
// In case there was an error returned, it will be accessible using
// Job.GetError. Note that this only obtains the error's message, not the
// underlying structure of the original error object.
type Action func(closer <-chan struct{}) (string, error)

// Scheduler represents a scheduler object that manages jobs and their
// executions.
type Scheduler interface {
	// Create creates a new job object configured with the given action. The job
	// object is immediately returned and its corresponding action is executed
	// asynchronously.
	Create(action Action) (Job, error)

	// FetchState fetches the state for the given job ID.
	FetchState(jobID ObjectID) (Job, error)

	// MarkAsSucceeded marks the job object as succeeded and persists its state.
	// The returned job object is actually the refreshed version of the provided
	// one.
	MarkAsSucceeded(job Job) (Job, error)

	// MarkAsFailedWithError marks the given job as failed, adds information of
	// the given error to it and persists the job's state. The returned job is
	// actually the refreshed version of the provided one.
	MarkAsFailedWithError(job Job, err error) (Job, error)

	Object

	// PersistState writes the given job to the configured Storage.
	PersistState(job Job) error

	// WaitForFinalStatus blocks and waits for the given job to reach a final
	// status. The given closer can end the waiting and thus stop blocking the
	// call to WaitForFinalStatus.
	WaitForFinalStatus(jobID ObjectID, closer <-chan struct{}) (Job, error)
}
