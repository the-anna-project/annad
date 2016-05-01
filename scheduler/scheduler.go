// Package scheduler implements spec.Scheduler to provide job management.
package scheduler

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeScheduler represents the object type of the scheduler object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeScheduler spec.ObjectType = "scheduler"
)

// Config represents the configuration used to create a new scheduler object.
type Config struct {
	// Dependencies.

	// Log represents the logger to use to log messages.
	Log spec.Log

	// Storage represents the storage to use for persisting data.
	Storage spec.Storage

	// Settings.

	// Actions represents the configured actions a scheduler is able to execute.
	Actions map[string]spec.Action

	// WaitSleep represents the time to sleep between state-check cycles.
	WaitSleep time.Duration
}

// DefaultConfig provides a default configuration to create a new scheduler
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

		// Settings.
		Actions:   map[string]spec.Action{},
		WaitSleep: 1 * time.Second,
	}

	return newConfig
}

// NewScheduler creates a new configured scheduler object.
//
// The scheduler makes use of different key namespaces, that can be anything of
// the following.
//
//     job:<ID>
//
//         Holds a job's data as JSON string using the job's ID.
//
//         {...}
//
//     jobs:active
//
//         Holds the weighted list of job IDs ordered from oldest to newest
//         timestamp.
//
//         ID1,timestamp1,ID2,timestamp2,...
//
//     session:<ID>
//
//         Holds the weighted list scoped for a session ID, containing job IDs
//         ordered from oldest to newest timestamp.
//
//         ID1,timestamp1,ID2,timestamp2,...
//
func NewScheduler(config Config) (spec.Scheduler, error) {
	newScheduler := &scheduler{
		Config: config,
		Booted: false,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeScheduler,
	}

	newScheduler.Log.Register(newScheduler.GetType())

	return newScheduler, nil
}

type scheduler struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (s *scheduler) Boot() {
	// TODO fix that everywhere
	s.Mutex.Lock()
	if s.Booted {
		s.Mutex.Unlock()
		return
	}
	s.Booted = true
	s.Mutex.Unlock()

	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call Boot")

	s.scheduleActiveJobs()
}

func (s *scheduler) Execute(job spec.Job) error {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call Execute")

	action, ok := s.Actions[job.GetActionID()]
	if !ok {
		return maskAny(actionNotFoundError)
	}
	job, err := s.MarkAsActive(job)
	if err != nil {
		return maskAny(err)
	}
	err = s.PersistJob(job)
	if err != nil {
		return maskAny(err)
	}

	closer := make(chan struct{}, 1)
	replacer := make(chan struct{}, 1)

	// Check if there is already a job running for the given session ID.
	go func() {
		err := s.preventDuplicates(job)
		if err != nil {
			s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
		}
	}()

	// Check the current job's final status was set.
	go func() {
		err := s.waitForStatus(job, closer, replacer)
		if err != nil {
			s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
		}
	}()

	go func() {
		err := s.waitForAction(job, action, closer, replacer)
		if err != nil {
			s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
		}
	}()

	return nil
}

func (s *scheduler) FetchJob(jobID spec.ObjectID) (spec.Job, error) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call FetchJob")

	value, err := s.Storage.Get(s.key("job:%s", string(jobID)))
	if err != nil {
		return nil, maskAny(err)
	}

	if value == "" {
		return nil, maskAny(jobNotFoundError)
	}

	newJob := NewEmptyJob()
	err = json.Unmarshal([]byte(value), newJob)
	if err != nil {
		return nil, maskAny(err)
	}

	return newJob, nil
}

func (s *scheduler) MarkAsActive(job spec.Job) (spec.Job, error) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call MarkAsActive")

	job.SetActiveStatus(StatusStarted)

	// Note that the job IDs are ordered by their timestamps. That means when
	// iterating over the set of elements, the onces added first are accessed
	// first during iteration.
	err := s.Storage.SetElementByScore(s.key("jobs:active"), string(job.GetID()), float64(job.GetCreatedAt().UnixNano()))
	if err != nil {
		return nil, maskAny(err)
	}

	// Note that the job IDs are ordered by their timestamps. That means when
	// iterating over the set of elements, the onces added first are accessed
	// first during iteration.
	err = s.Storage.SetElementByScore(s.key("session:%s", job.GetSessionID()), string(job.GetID()), float64(job.GetCreatedAt().UnixNano()))
	if err != nil {
		return nil, maskAny(err)
	}

	return job, nil
}

func (s *scheduler) MarkAsFailedWithError(job spec.Job, err error) (spec.Job, error) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call MarkAsFailedWithError")

	job.SetError(err)
	job.SetFinalStatus(StatusFailed)

	job, err = s.MarkAsInactive(job)
	if err != nil {
		return nil, maskAny(err)
	}
	err = s.PersistJob(job)
	if err != nil {
		return nil, maskAny(err)
	}

	return job, nil
}

func (s *scheduler) MarkAsInactive(job spec.Job) (spec.Job, error) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call MarkAsInactive")

	job.SetActiveStatus(StatusStopped)

	err := s.Storage.RemoveScoredElement(s.key("jobs:active"), string(job.GetID()))
	if err != nil {
		return nil, maskAny(err)
	}

	err = s.Storage.RemoveScoredElement(s.key("session:%s", job.GetSessionID()), string(job.GetID()))
	if err != nil {
		return nil, maskAny(err)
	}

	return job, nil
}

func (s *scheduler) MarkAsReplaced(job spec.Job) (spec.Job, error) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call MarkAsSucceeded")

	job.SetFinalStatus(StatusReplaced)

	job, err := s.MarkAsInactive(job)
	if err != nil {
		return nil, maskAny(err)
	}
	err = s.PersistJob(job)
	if err != nil {
		return nil, maskAny(err)
	}

	return job, nil
}

func (s *scheduler) MarkAsSucceeded(job spec.Job) (spec.Job, error) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call MarkAsSucceeded")

	job.SetFinalStatus(StatusSucceeded)

	job, err := s.MarkAsInactive(job)
	if err != nil {
		return nil, maskAny(err)
	}
	err = s.PersistJob(job)
	if err != nil {
		return nil, maskAny(err)
	}

	return job, nil
}

func (s *scheduler) PersistJob(job spec.Job) error {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call PersistJob")

	raw, err := json.Marshal(job)
	if err != nil {
		return maskAny(err)
	}

	err = s.Storage.Set(s.key("job:%s", job.GetID()), string(raw))
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *scheduler) Register(actionID string, action spec.Action) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call Register")

	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Actions[actionID] = action
}

// WaitForFinalStatus acts as described in the interface comments. Note that
// both, job object and error will be nil in case the closer ends waiting for
// the job to reach a final state.
func (s *scheduler) WaitForFinalStatus(jobID spec.ObjectID, closer <-chan struct{}) (spec.Job, error) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call WaitForFinalStatus")

	for {
		select {
		case <-closer:
			return nil, nil
		case <-time.After(s.WaitSleep):
			job, err := s.FetchJob(jobID)
			if err != nil {
				return nil, maskAny(err)
			}

			if HasFinalStatus(job) {
				return job, nil
			}
		}
	}
}
