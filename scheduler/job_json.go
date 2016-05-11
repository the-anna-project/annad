package scheduler

import (
	"encoding/json"
)

// jobClone is for making use of the stdlib json implementation. The job object
// implements its own marshaler and unmarshaler but only to provide json
// implementations for spec.Job. Note, not redirecting the type will cause
// infinite recursion.
type jobClone job

func (j *job) MarshalJSON() ([]byte, error) {
	newJob := jobClone(*j)

	raw, err := json.Marshal(newJob)
	if err != nil {
		return nil, maskAny(err)
	}

	return raw, nil
}

func (j *job) UnmarshalJSON(b []byte) error {
	newJob := jobClone{}

	err := json.Unmarshal(b, &newJob)
	if err != nil {
		return maskAny(err)
	}

	j.ActionID = newJob.ActionID
	j.ActiveStatus = newJob.ActiveStatus
	j.Args = newJob.Args
	j.CreatedAt = newJob.CreatedAt
	j.Error = newJob.Error
	j.FinalStatus = newJob.FinalStatus
	j.ID = newJob.ID
	j.Result = newJob.Result
	j.SessionID = newJob.SessionID
	j.Type = newJob.Type

	return nil
}
