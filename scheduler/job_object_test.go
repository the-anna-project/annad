package scheduler

import (
	"testing"
)

func Test_Job_GetType(t *testing.T) {
	newJob, err := NewJob(DefaultJobConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newJob.GetType() != ObjectTypeJob {
		t.Fatal("expected", ObjectTypeJob, "got", newJob.GetType())
	}
}

func Test_Job_GetID(t *testing.T) {
	firstJob, err := NewJob(DefaultJobConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	secondJob, err := NewJob(DefaultJobConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if firstJob.GetID() == secondJob.GetID() {
		t.Fatalf("IDs of jobs are equal")
	}
}
