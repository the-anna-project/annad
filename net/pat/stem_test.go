package patnet

import (
	"testing"
)

func Test_Stem(t *testing.T) {
	testCases := []struct {
		Input    []string
		Expected string
	}{
		{
			Input:    []string{},
			Expected: "",
		},
		{
			Input: []string{
				"abc",
				"abcd",
				"abcde",
				"abcdef",
			},
			Expected: "abc",
		},
		// Same test as before, but items are in different order.
		{
			Input: []string{
				"abcd",
				"abc",
				"abcdef",
				"abcde",
			},
			Expected: "abc",
		},
		{
			Input: []string{
				"xxx",
				"abcd",
				"abcde",
				"abcdef",
			},
			Expected: "",
		},
		{
			Input: []string{
				"abcd",
				"xxx",
				"abcdef",
				"abcde",
			},
			Expected: "",
		},
		{
			Input: []string{
				"ddggddgg",
				"ddxxddgg",
				"ddggddgg",
				"ddggddgg",
			},
			Expected: "dd",
		},
		{
			Input: []string{
				"ddgg",
				"ddxxddgg",
				"ddxx",
				"ddggddgg",
				"ddggddgg",
			},
			Expected: "dd",
		},
	}

	for i, testCase := range testCases {
		output := Stem(testCase.Input)

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}
