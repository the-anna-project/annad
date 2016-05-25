package scheduler

import (
	"testing"
)

func Test_Job_UnmarshalJSON_Error(t *testing.T) {
	newJob := NewEmptyJob()

	err := newJob.UnmarshalJSON([]byte("invalid"))
	if err == nil {
		t.Fatal("expected", "error", "got", nil)
	}
}
