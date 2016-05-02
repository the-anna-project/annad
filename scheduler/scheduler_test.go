package scheduler

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/rafaeljusto/redigomock"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/redis"
)

func newTestScheduler(storage spec.Storage) spec.Scheduler {
	newConfig := DefaultConfig()
	if storage != nil {
		newConfig.Storage = storage
	}
	newConfig.WaitSleep = 10 * time.Millisecond
	newScheduler, err := NewScheduler(newConfig)
	if err != nil {
		panic(err)
	}

	return newScheduler
}

func newTestJob(args interface{}, sessionID string, actionID string) spec.Job {
	newJobConfig := DefaultJobConfig()
	newJobConfig.ActionID = actionID
	newJobConfig.Args = args
	newJobConfig.SessionID = sessionID
	newJob, err := NewJob(newJobConfig)
	if err != nil {
		panic(err)
	}

	return newJob
}

func Test_Scheduler_Boot_Reschedule(t *testing.T) {
	newScheduler := newTestScheduler(nil)
	var sleep int
	var m sync.Mutex
	newScheduler.Register("test-action", func(args interface{}, closer <-chan struct{}) (string, error) {
		// Note we increase the sleep time here so we have jobs simulating
		// different workloads. This ensures testing the expicit handling of
		// marking short living jobs as being replaced.
		m.Lock()
		sleep += 100
		m.Unlock()
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		return "test-result", nil
	})

	// Store some jobs for the same session ID and mark them as active.
	var newJobs []spec.Job
	for i := 0; i < 5; i++ {
		newJob := newTestJob(nil, "test-session", "test-action")
		newJob, err := newScheduler.MarkAsActive(newJob)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		err = newScheduler.StoreJob(newJob)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		newJobs = append(newJobs, newJob)
	}

	for _, newJob := range newJobs {
		// Check if the job is really active.
		newJob, err := newScheduler.GetJobByID(newJob.GetID())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		if HasReplacedStatus(newJob) {
			t.Fatal("expected", true, "got", false)
		}
		if newJob.GetActiveStatus() != StatusStarted {
			t.Fatal("expected", StatusStarted, "got", newJob.GetActiveStatus())
		}
	}

	var c int
	check := func() {
		c++

		for i, newJob := range newJobs {
			// Check if the job has successfully finished.
			newJob, err := newScheduler.GetJobByID(newJob.GetID())
			if err != nil {
				t.Fatal("call", c, "expected", nil, "got", err)
			}
			if i == len(newJobs)-1 {
				// The last job should have succeeded.
				if HasReplacedStatus(newJob) {
					t.Fatal("call", c, "expected", true, "got", false)
				}
			} else {
				// The other jobs should be replaced.
				if !HasReplacedStatus(newJob) {
					fmt.Printf("newJobs: %#v\n", newJobs)
					t.Fatal("call", c, "expected", true, "got", false)
				}
				if HasSucceededStatus(newJob) {
					t.Fatal("call", c, "expected", true, "got", false)
				}
			}
		}
	}

	newScheduler.Boot()
	time.Sleep(100 * time.Millisecond)
	check()

	newScheduler.Boot()
	check()
}

func Test_Scheduler_Execute_Success(t *testing.T) {
	newScheduler := newTestScheduler(nil)
	newScheduler.Register("test-action", func(args interface{}, closer <-chan struct{}) (string, error) {
		return "test-result", nil
	})

	newJob := newTestJob(nil, "test-session", "test-action")
	err := newScheduler.Execute(newJob)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newJob, err = newScheduler.WaitForFinalStatus(newJob.GetID(), nil)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !HasFinalStatus(newJob) {
		t.Fatal("expected", true, "got", false)
	}

	if newJob.GetResult() != "test-result" {
		t.Fatal("expected", "test-result", "got", newJob.GetResult())
	}

	if newJob.GetError() != nil {
		t.Fatal("expected", nil, "got", newJob.GetError())
	}
}

func Test_Scheduler_Execute_ActionRegisterError(t *testing.T) {
	newScheduler := newTestScheduler(nil)

	newJob := newTestJob(nil, "test-session", "test-action")
	err := newScheduler.Execute(newJob)
	if !IsActionNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Scheduler_Execute_StorageError(t *testing.T) {
	c := redigomock.NewConn()
	// Note PushToSet is called twice. Thus we need to mock the call.
	c.Command("ZADD").Expect(int64(1)).Expect(int64(1))
	// Note returning this specific error here makes no sense business wise. It
	// is only to verify the test.
	c.Command("SET").ExpectError(jobNotFoundError)
	newStorageConfig := redisstorage.DefaultConfigWithConn(c)
	newStorage := redisstorage.NewRedisStorage(newStorageConfig)

	newScheduler := newTestScheduler(newStorage)
	newScheduler.Register("test-action", func(args interface{}, closer <-chan struct{}) (string, error) {
		return "", nil
	})

	newJob := newTestJob(nil, "test-session", "test-action")
	err := newScheduler.Execute(newJob)
	if !IsJobNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Scheduler_Execute_Action_Error(t *testing.T) {
	newScheduler := newTestScheduler(nil)
	newScheduler.Register("test-action", func(args interface{}, closer <-chan struct{}) (string, error) {
		return "", fmt.Errorf("test error")
	})

	newJob := newTestJob(nil, "test-session", "test-action")
	err := newScheduler.Execute(newJob)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newJob, err = newScheduler.WaitForFinalStatus(newJob.GetID(), nil)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !HasFinalStatus(newJob) {
		t.Fatal("expected", true, "got", false)
	}

	if newJob.GetError().Error() != "test error" {
		t.Fatal("expected", "test error", "got", newJob.GetError())
	}
}

func Test_Scheduler_Execute_FetchState(t *testing.T) {
	newScheduler := newTestScheduler(nil)
	newScheduler.Register("test-action", func(args interface{}, closer <-chan struct{}) (string, error) {
		return "", nil
	})

	newJob := newTestJob(nil, "test-session", "test-action")
	err := newScheduler.Execute(newJob)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newJob, err = newScheduler.WaitForFinalStatus(newJob.GetID(), nil)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !HasFinalStatus(newJob) {
		t.Fatal("expected", true, "got", false)
	}

	// Fetching invalid state should not work.
	_, err = newScheduler.GetJobByID("invalid")
	if !IsJobNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}

	// Fetching valid state should work.
	newJob, err = newScheduler.GetJobByID(newJob.GetID())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !HasFinalStatus(newJob) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Scheduler_Execute_Wait(t *testing.T) {
	newScheduler := newTestScheduler(nil)
	newScheduler.Register("test-action", func(args interface{}, closer <-chan struct{}) (string, error) {
		// Just something to do, so the job blocks
		time.Sleep(300 * time.Millisecond)
		return "", nil
	})

	originalJob := newTestJob(nil, "test-session", "test-action")
	err := newScheduler.Execute(originalJob)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Directly close and end waiting.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	newJob, err := newScheduler.WaitForFinalStatus(originalJob.GetID(), closer)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if newJob != nil {
		t.Fatal("expected", nil, "got", newJob)
	}

	newJob, err = newScheduler.GetJobByID(originalJob.GetID())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// When we use the closer to end waiting before the job is finished, the
	// job object should not have a final state yet.
	if HasFinalStatus(newJob) {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_Scheduler_WaitForFinalStatus_Error(t *testing.T) {
	c := redigomock.NewConn()
	// Note returning this specific error here makes no sense business wise. It
	// is only to verify the test.
	c.Command("GET").ExpectError(jobNotFoundError)
	newStorageConfig := redisstorage.DefaultConfigWithConn(c)
	newStorage := redisstorage.NewRedisStorage(newStorageConfig)

	newScheduler := newTestScheduler(newStorage)

	closer := make(chan struct{}, 1)
	_, err := newScheduler.WaitForFinalStatus("some-id", closer)
	if !IsJobNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Scheduler_Execute_Scheduler_Ends_Wait(t *testing.T) {
	newScheduler := newTestScheduler(nil)
	newScheduler.Register("test-action", func(args interface{}, closer <-chan struct{}) (string, error) {
		select {
		case <-closer:
			// The test was successful.
		case <-time.After(1000 * time.Second):
			t.Fatal("expected", "no execution", "got", "execution")
		}
		return "", nil
	})

	originalJob := newTestJob(nil, "test-session", "test-action")
	err := newScheduler.Execute(originalJob)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Here we set the final state of the job. Thus the action's closer should be
	// triggered.
	originalJob, err = newScheduler.MarkAsSucceeded(originalJob)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	closer := make(chan struct{}, 1)
	newJob, err := newScheduler.WaitForFinalStatus(originalJob.GetID(), closer)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !HasFinalStatus(newJob) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Scheduler_MarkAsReplaced(t *testing.T) {
	c := redigomock.NewConn()
	// Note returning this specific error here makes no sense business wise. It
	// is only to verify the test.
	c.Command("ZREM").ExpectError(jobNotFoundError)
	newStorageConfig := redisstorage.DefaultConfigWithConn(c)
	newStorage := redisstorage.NewRedisStorage(newStorageConfig)

	newScheduler := newTestScheduler(newStorage)

	newJob := newTestJob(nil, "test-session", "test-action")
	_, err := newScheduler.(*scheduler).MarkAsReplaced(newJob)
	if !IsJobNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Scheduler_MarkAsSucceeded(t *testing.T) {
	c := redigomock.NewConn()
	// Note returning this specific error here makes no sense business wise. It
	// is only to verify the test.
	c.Command("ZREM").ExpectError(jobNotFoundError)
	newStorageConfig := redisstorage.DefaultConfigWithConn(c)
	newStorage := redisstorage.NewRedisStorage(newStorageConfig)

	newScheduler := newTestScheduler(newStorage)

	newJob := newTestJob(nil, "test-session", "test-action")
	_, err := newScheduler.(*scheduler).MarkAsSucceeded(newJob)
	if !IsJobNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Scheduler_MarkAsFailedWithError(t *testing.T) {
	c := redigomock.NewConn()
	// Note returning this specific error here makes no sense business wise. It
	// is only to verify the test.
	c.Command("ZREM").ExpectError(jobNotFoundError)
	newStorageConfig := redisstorage.DefaultConfigWithConn(c)
	newStorage := redisstorage.NewRedisStorage(newStorageConfig)

	newScheduler := newTestScheduler(newStorage)

	newJob := newTestJob(nil, "test-session", "test-action")
	_, err := newScheduler.(*scheduler).MarkAsFailedWithError(newJob, fmt.Errorf("test error"))
	if !IsJobNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Scheduler_Execute_MarkAsActiveError(t *testing.T) {
	c := redigomock.NewConn()
	// Note returning this specific error here makes no sense business wise. It
	// is only to verify the test.
	c.Command("ZADD").ExpectError(jobNotFoundError)
	newStorageConfig := redisstorage.DefaultConfigWithConn(c)
	newStorage := redisstorage.NewRedisStorage(newStorageConfig)

	newScheduler := newTestScheduler(newStorage)
	newScheduler.Register("test-action", func(args interface{}, closer <-chan struct{}) (string, error) {
		return "", nil
	})

	newJob := newTestJob(nil, "test-session", "test-action")
	err := newScheduler.(*scheduler).Execute(newJob)
	if !IsJobNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}
