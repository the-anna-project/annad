package scheduler

import (
	"fmt"
	"testing"
	"time"

	"github.com/rafaeljusto/redigomock"
	"github.com/xh3b4sd/anna/storage/redis"
)

func Test_Scheduler_Create_Success(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.WaitSleep = 10 * time.Millisecond
	newScheduler, err := NewScheduler(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	action := func(closer <-chan struct{}) (string, error) {
		return "test-result", nil
	}

	job, err := newScheduler.Create(action)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	job, err = newScheduler.WaitForFinalStatus(job.GetID(), nil)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !HasFinalStatus(job) {
		t.Fatal("expected", true, "got", false)
	}

	if job.GetResult() != "test-result" {
		t.Fatal("expected", "test-result", "got", job.GetResult())
	}

	if job.GetError() != nil {
		t.Fatal("expected", nil, "got", job.GetError())
	}
}

func Test_Scheduler_Create_Error(t *testing.T) {
	c := redigomock.NewConn()
	// Note returning this specific error here makes no sense business wise. It
	// is only to verify the test.
	c.Command("SET").ExpectError(jobNotFoundError)
	newStorageConfig := redisstorage.DefaultConfigWithConn(c)
	newStorage := redisstorage.NewRedisStorage(newStorageConfig)

	newConfig := DefaultConfig()
	newConfig.Storage = newStorage
	newConfig.WaitSleep = 10 * time.Millisecond
	newScheduler, err := NewScheduler(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	action := func(closer <-chan struct{}) (string, error) {
		return "", nil
	}
	_, err = newScheduler.Create(action)
	if !IsJobNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Scheduler_Create_Action_Error(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.WaitSleep = 10 * time.Millisecond
	newScheduler, err := NewScheduler(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	action := func(closer <-chan struct{}) (string, error) {
		return "", fmt.Errorf("test error")
	}
	job, err := newScheduler.Create(action)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	job, err = newScheduler.WaitForFinalStatus(job.GetID(), nil)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !HasFinalStatus(job) {
		t.Fatal("expected", true, "got", false)
	}

	if job.GetError().Error() != "test error" {
		t.Fatal("expected", "test error", "got", job.GetError())
	}
}

func Test_Scheduler_Create_FetchState(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.WaitSleep = 10 * time.Millisecond
	newScheduler, err := NewScheduler(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	action := func(closer <-chan struct{}) (string, error) {
		return "", nil
	}
	job, err := newScheduler.Create(action)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	job, err = newScheduler.WaitForFinalStatus(job.GetID(), nil)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !HasFinalStatus(job) {
		t.Fatal("expected", true, "got", false)
	}

	// Fetching invalid state should not work.
	_, err = newScheduler.FetchState("invalid")
	if !IsJobNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}

	// Fetching valid state should work.
	job, err = newScheduler.FetchState(job.GetID())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !HasFinalStatus(job) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Scheduler_Create_Wait(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.WaitSleep = 10 * time.Millisecond
	newScheduler, err := NewScheduler(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	action := func(closer <-chan struct{}) (string, error) {
		// Just something to do, so the job blocks
		time.Sleep(300 * time.Millisecond)
		return "", nil
	}
	originalJob, err := newScheduler.Create(action)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Directly close and end waiting.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	job, err := newScheduler.WaitForFinalStatus(originalJob.GetID(), closer)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if job != nil {
		t.Fatal("expected", nil, "got", job)
	}

	job, err = newScheduler.FetchState(originalJob.GetID())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// When we use the closer to end waiting before the job is finished, the
	// job object should not have a final state yet.
	if HasFinalStatus(job) {
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

	newConfig := DefaultConfig()
	newConfig.Storage = newStorage
	newConfig.WaitSleep = 10 * time.Millisecond
	newScheduler, err := NewScheduler(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	closer := make(chan struct{}, 1)
	_, err = newScheduler.WaitForFinalStatus("some-id", closer)
	if !IsJobNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Scheduler_Create_Scheduler_Ends_Wait(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.WaitSleep = 10 * time.Millisecond
	newScheduler, err := NewScheduler(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	action := func(closer <-chan struct{}) (string, error) {
		select {
		case <-closer:
			// The test was successful.
		case <-time.After(1000 * time.Second):
			t.Fatal("expected", "no execution", "got", "execution")
		}
		return "", nil
	}
	originalJob, err := newScheduler.Create(action)
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
	job, err := newScheduler.WaitForFinalStatus(originalJob.GetID(), closer)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !HasFinalStatus(job) {
		t.Fatal("expected", true, "got", false)
	}
}
