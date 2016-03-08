package scheduler

import (
	"fmt"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeJob represents the object type of the job object.
	ObjectTypeJob spec.ObjectType = "job"
)

// JobConfig represents the configuration used to create a new job object.
type JobConfig struct {
}

// DefaultJobConfig provides a default configuration to create a new job object
// by best effort.
func DefaultJobConfig() JobConfig {
	newConfig := JobConfig{}

	return newConfig
}

// NewJob creates a new configured job object.
func NewJob(config JobConfig) (spec.Job, error) {
	newJob := &job{
		JobConfig: config,

		ActiveStatus: StatusStarted,
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		FinalStatus:  "",
		Type:         ObjectTypeJob,
	}

	return newJob, nil
}

// Job represents a job that is executable.
type job struct {
	JobConfig

	// ActiveStatus represents a status indicating activation or deactivation.
	ActiveStatus spec.ActiveStatus `json:"active_status,omitempty"`

	// Error represents the message of an error occurred during job execution, if
	// any.
	Error string `json:"error,omitempty"`

	// FinalStatus represents any status that is final. A job having this status
	// will not change its status anymore.
	FinalStatus spec.FinalStatus `json:"final_status,omitempty"`

	// ID represents the job identifier.
	ID spec.ObjectID `json:"id,omitempty"`

	Mutex sync.Mutex `json:"-"`

	// Result represents the job's result returned by the corresponding action,
	// if any.
	Result string `json:"result,omitempty"`

	// Type represents the job's object type.
	Type spec.ObjectType `json:"type,omitempty"`
}

func (j *job) GetActiveStatus() spec.ActiveStatus {
	j.Mutex.Lock()
	defer j.Mutex.Unlock()

	return j.ActiveStatus
}

func (j *job) GetError() error {
	j.Mutex.Lock()
	defer j.Mutex.Unlock()

	if j.Error == "" {
		return nil
	}

	return fmt.Errorf(j.Error)
}

func (j *job) GetFinalStatus() spec.FinalStatus {
	j.Mutex.Lock()
	defer j.Mutex.Unlock()

	return j.FinalStatus
}

func (j *job) GetResult() string {
	j.Mutex.Lock()
	defer j.Mutex.Unlock()

	return j.Result
}

func (j *job) SetActiveStatus(activeStatus spec.ActiveStatus) {
	j.Mutex.Lock()
	defer j.Mutex.Unlock()

	j.ActiveStatus = activeStatus
}

func (j *job) SetError(err error) {
	j.Mutex.Lock()
	defer j.Mutex.Unlock()

	j.Error = err.Error()
}

func (j *job) SetFinalStatus(finalStatus spec.FinalStatus) {
	j.Mutex.Lock()
	defer j.Mutex.Unlock()

	j.FinalStatus = finalStatus
}

func (j *job) SetResult(result string) {
	j.Mutex.Lock()
	defer j.Mutex.Unlock()

	j.Result = result
}
