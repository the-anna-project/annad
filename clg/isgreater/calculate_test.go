package isgreater

import (
	"testing"

	"github.com/xh3b4sd/anna/object/context"
)

func Test_CLG_IsGreater(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected bool
	}{
		{
			A:        3.5,
			B:        3.5,
			Expected: false,
		},
		{
			A:        3.5,
			B:        12.5,
			Expected: false,
		},
		{
			A:        35.5,
			B:        14.5,
			Expected: true,
		},
		{
			A:        -3.5,
			B:        7.5,
			Expected: false,
		},
		{
			A:        12.5,
			B:        4.5,
			Expected: true,
		},
		{
			A:        17,
			B:        65,
			Expected: false,
		},
		{
			A:        65,
			B:        17,
			Expected: true,
		},
	}

	newCLG := MustNew()

	for i, testCase := range testCases {
		b := newCLG.(*clg).calculate(context.MustNew(), testCase.A, testCase.B)
		if b != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", b)
		}
	}
}
