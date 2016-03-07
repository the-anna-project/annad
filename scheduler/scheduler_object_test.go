package scheduler

import (
	"testing"
)

func Test_Scheduler_GetType(t *testing.T) {
	newScheduler, err := NewScheduler(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newScheduler.GetType() != ObjectTypeScheduler {
		t.Fatal("expected", ObjectTypeScheduler, "got", newScheduler.GetType())
	}
}

func Test_Scheduler_GetID(t *testing.T) {
	firstScheduler, err := NewScheduler(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	secondScheduler, err := NewScheduler(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if firstScheduler.GetID() == secondScheduler.GetID() {
		t.Fatalf("IDs of schedulers are equal")
	}
}
