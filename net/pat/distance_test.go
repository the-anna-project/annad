package patnet

import (
	"testing"
)

func Test_Distance(t *testing.T) {
	testCases := []struct {
		Input    []string
		Expected int
	}{
		{
			Input:    []string{"", ""},
			Expected: 0,
		},
		{
			Input:    []string{"a", "a"},
			Expected: 0,
		},
		{
			Input:    []string{"abc", "abc"},
			Expected: 0,
		},

		{
			Input:    []string{"abc", "abcd"},
			Expected: 1,
		},
		{
			Input:    []string{"abc", "abx"},
			Expected: 1,
		},
		{
			Input:    []string{"abc", "axc"},
			Expected: 1,
		},
		{
			Input:    []string{"abc", "xbc"},
			Expected: 1,
		},
		{
			Input:    []string{"abx", "abc"},
			Expected: 1,
		},
		{
			Input:    []string{"axc", "abc"},
			Expected: 1,
		},
		{
			Input:    []string{"xbc", "abc"},
			Expected: 1,
		},

		{
			Input:    []string{"car", "egg"},
			Expected: 3,
		},

		{
			Input:    []string{"abcdef", "xcxe"},
			Expected: 4,
		},
		{
			Input:    []string{"hock", "shocking"},
			Expected: 4,
		},
	}

	for i, testCase := range testCases {
		output := Distance(testCase.Input[0], testCase.Input[1])

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}
