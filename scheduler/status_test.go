package scheduler

import (
	"testing"
)

func Test_Status_HasFinalStatus(t *testing.T) {
	testCases := []struct {
		Input    *job
		Expected bool
	}{
		// This status combination is invalid.
		{
			Input: &job{
				ActiveStatus: StatusStarted,
				Error:        "",
				FinalStatus:  StatusFailed,
				ID:           "",
			},
			Expected: false,
		},

		{
			Input: &job{
				ActiveStatus: StatusStopped,
				Error:        "",
				FinalStatus:  StatusFailed,
				ID:           "",
			},
			Expected: true,
		},

		// This status combination is invalid.
		{
			Input: &job{
				ActiveStatus: StatusStarted,
				Error:        "",
				FinalStatus:  StatusSucceeded,
				ID:           "",
			},
			Expected: false,
		},
		{
			Input: &job{
				ActiveStatus: StatusStopped,
				Error:        "",
				FinalStatus:  StatusSucceeded,
				ID:           "",
			},
			Expected: true,
		},
		{
			Input: &job{
				ActiveStatus: StatusStarted,
				Error:        "",
				FinalStatus:  "",
				ID:           "",
			},
			Expected: false,
		},
		{
			Input: &job{
				ActiveStatus: StatusStopped,
				Error:        "",
				FinalStatus:  "",
				ID:           "",
			},
			Expected: false,
		},
		{
			Input: &job{
				ActiveStatus: StatusStarted,
				Error:        "",
				FinalStatus:  "",
				ID:           "",
			},
			Expected: false,
		},
		{
			Input: &job{
				ActiveStatus: StatusStopped,
				Error:        "",
				FinalStatus:  "",
				ID:           "",
			},
			Expected: false,
		},
		{
			Input: &job{
				ActiveStatus: StatusStopped,
				Error:        "",
				FinalStatus:  StatusReplaced,
				ID:           "",
			},
			Expected: true,
		},
		{
			Input: &job{
				ActiveStatus: StatusStarted,
				Error:        "",
				FinalStatus:  StatusReplaced,
				ID:           "",
			},
			Expected: false,
		},
		{
			Input: &job{
				ActiveStatus: StatusStopped,
				Error:        "test error",
				FinalStatus:  "",
				ID:           "test-id",
			},
			Expected: false,
		},
	}

	for i, testCase := range testCases {
		output := HasFinalStatus(testCase.Input)

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}
