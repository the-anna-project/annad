// Package scheduler implements spec.Scheduler to provide job management.
package scheduler

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/xh3b4sd/anna/ctx"
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
	// Log represents the logger to use to log messages.
	Log spec.Log

	// Storage represents the storage to use for persisting data.
	Storage spec.Storage

	// WaitSleep represents the time to sleep between state-check cycles.
	WaitSleep time.Duration
}

// DefaultConfig provides a default configuration to create a new scheduler
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Log:       log.NewLog(log.DefaultConfig()),
		Storage:   memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),
		WaitSleep: 1 * time.Second,
	}

	return newConfig
}

// NewScheduler creates a new configured scheduler object.
func NewScheduler(config Config) (spec.Scheduler, error) {
	newScheduler := &scheduler{
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeScheduler,
	}

	newCtxConfig := ctx.DefaultConfig()
	newCtxConfig.Object = newScheduler
	newScheduler.Ctx = ctx.NewCtx(newCtxConfig)

	newScheduler.Log.Register(newScheduler.GetType())

	return newScheduler, nil
}

type scheduler struct {
	Config

	// Ctx is used for storage key management.
	Ctx spec.Ctx

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (s *scheduler) Create(action spec.Action) (spec.Job, error) {
	job, err := NewJob(DefaultJobConfig())
	if err != nil {
		return nil, maskAny(err)
	}

	go func() {
		closer := make(chan struct{}, 1)

		go func() {
			// This check sends a signal to the action's closer as soon as the
			// current job's final status was set.
			_, err := s.WaitForFinalStatus(job.GetID(), nil)
			if err != nil {
				s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
			}
			closer <- struct{}{}
		}()

		result, err := action(closer)
		if err != nil {
			_, err = s.MarkAsFailedWithError(job, err)
			if err != nil {
				s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
				return
			}
			return
		}

		job.SetResult(result)
		_, err = s.MarkAsSucceeded(job)
		if err != nil {
			s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
			return
		}
	}()

	err = s.PersistState(job)
	if err != nil {
		return nil, maskAny(err)
	}

	return job, nil
}

func (s *scheduler) FetchState(jobID spec.ObjectID) (spec.Job, error) {
	value, err := s.Storage.Get(s.Ctx.GetKey("job:%s", string(jobID)))
	if err != nil {
		return nil, maskAny(err)
	}

	if value == "" {
		return nil, maskAny(jobNotFoundError)
	}

	var newJob job
	err = json.Unmarshal([]byte(value), &newJob)
	if err != nil {
		return nil, maskAny(err)
	}

	return &newJob, nil
}

func (s *scheduler) MarkAsFailedWithError(job spec.Job, err error) (spec.Job, error) {
	job.SetActiveStatus(StatusStopped)
	job.SetError(err)
	job.SetFinalStatus(StatusFailed)

	err = s.PersistState(job)
	if err != nil {
		return nil, maskAny(err)
	}

	return job, nil
}

func (s *scheduler) MarkAsSucceeded(job spec.Job) (spec.Job, error) {
	job.SetActiveStatus(StatusStopped)
	job.SetFinalStatus(StatusSucceeded)

	err := s.PersistState(job)
	if err != nil {
		return nil, maskAny(err)
	}

	return job, nil
}

func (s *scheduler) PersistState(job spec.Job) error {
	raw, err := json.Marshal(job)
	if err != nil {
		return maskAny(err)
	}

	err = s.Storage.Set(s.Ctx.GetKey("job:%s", job.GetID()), string(raw))
	if err != nil {
		return maskAny(err)
	}

	return nil
}

// WaitForFinalStatus acts as described in the interface comments. Note that
// both, scheduler object and error will be nil in case the closer ends waiting for
// the scheduler to reach a final state.
func (s *scheduler) WaitForFinalStatus(jobID spec.ObjectID, closer <-chan struct{}) (spec.Job, error) {
	for {
		select {
		case <-closer:
			return nil, nil
		case <-time.After(s.WaitSleep):
			job, err := s.FetchState(jobID)
			if err != nil {
				return nil, maskAny(err)
			}

			if HasFinalStatus(job) {
				return job, nil
			}
		}
	}
}
