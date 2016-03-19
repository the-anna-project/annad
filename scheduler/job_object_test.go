package scheduler

import (
	"testing"
)

func Test_Job_NewJob_Error(t *testing.T) {
	_, err := NewJob(DefaultJobConfig())
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Job_GetType(t *testing.T) {
	newJobConfig := DefaultJobConfig()
	newJobConfig.ActionID = "test-action"
	newJob, err := NewJob(newJobConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newJob.GetType() != ObjectTypeJob {
		t.Fatal("expected", ObjectTypeJob, "got", newJob.GetType())
	}
}

func Test_Job_GetID(t *testing.T) {
	newJobConfig := DefaultJobConfig()
	newJobConfig.ActionID = "test-action"
	firstJob, err := NewJob(newJobConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newJobConfig = DefaultJobConfig()
	newJobConfig.ActionID = "test-action"
	secondJob, err := NewJob(newJobConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if firstJob.GetID() == secondJob.GetID() {
		t.Fatalf("IDs of jobs are equal")
	}
}
