package scheduler

import (
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/key"
)

func (s *scheduler) key(f string, v ...interface{}) string {
	return key.NewSysKey(s, f, v...)
}

func (s *scheduler) scheduleActiveJobs() {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call scheduleActiveJobs")

	err := s.Storage.WalkScoredElements(s.key("jobs:active"), nil, func(element string, score float64) error {
		newJob, err := s.FetchJob(spec.ObjectID(element))
		if err != nil {
			return maskAny(err)
		}
		err = s.Execute(newJob)
		if err != nil {
			return maskAny(err)
		}

		return nil
	})

	if err != nil {
		s.Log.WithTags(spec.Tags{L: "F", O: s, T: nil, V: 1}, "%#v", maskAny(err))
	}
}

func (s *scheduler) preventDuplicates(job spec.Job) error {
	values, err := s.Storage.GetHighestScoredElements(s.key("session:%s", job.GetSessionID()), -1)
	if err != nil {
		return maskAny(err)
	}
	if len(values) < 4 {
		// There is only one job running for this session. Nothing to do here.
		return nil
	}

	newestJobID := spec.ObjectID(values[0])
	if newestJobID != job.GetID() {
		// This is not the newest job. Nothing to do here.
		return nil
	}

	for i, _ := range values {
		if i%2 != 0 {
			continue
		}

		jobID := spec.ObjectID(values[i])
		if jobID == job.GetID() {
			// Ignore our own job ID.
			continue
		}

		newJob, err := s.FetchJob(jobID)
		if err != nil {
			return maskAny(err)
		}
		_, err = s.MarkAsReplaced(newJob)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}

func (s *scheduler) waitForStatus(job spec.Job, closer chan<- struct{}, replacer chan<- struct{}) error {
	job, err := s.WaitForFinalStatus(job.GetID(), nil)
	if err != nil {
		return maskAny(err)
	}
	if HasReplacedStatus(job) {
		// Signal the replacement of this job, because some other job took
		// over. Thus this job can tear down.
		replacer <- struct{}{}
	}

	// This job was either replaced or simply teared down by setting its
	// final status. Notifiy the job's action callback and retire.
	closer <- struct{}{}

	return nil
}

func (s *scheduler) waitForAction(job spec.Job, action spec.Action, closer <-chan struct{}, replacer <-chan struct{}) error {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call waitForAction")

	result, err := action(job.GetArgs(), closer)
	if err != nil {
		_, merr := s.MarkAsFailedWithError(job, err)
		if merr != nil {
			return maskAny(merr)
		}
		return maskAny(err)
	}
	// Refresh the job's state.
	job, err = s.FetchJob(job.GetID())
	if err != nil {
		return maskAny(err)
	}
	job.SetResult(result)

	select {
	case <-replacer:
		// The current job was replaced. Only persist the job containing it's
		// result.
		err := s.PersistJob(job)
		if err != nil {
			return maskAny(err)
		}
	default:
		_, err = s.MarkAsSucceeded(job)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}
