package islesser

import (
	"testing"

	"github.com/xh3b4sd/anna/object/context"
)

func Test_CLG_IsLesser(t *testing.T) {
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
			A:        12.5,
			B:        3.5,
			Expected: false,
		},
		{
			A:        14.5,
			B:        35.5,
			Expected: true,
		},
		{
			A:        7.5,
			B:        -3.5,
			Expected: false,
		},
		{
			A:        4.5,
			B:        12.5,
			Expected: true,
		},
		{
			A:        65,
			B:        17,
			Expected: false,
		},
		{
			A:        17,
			B:        65,
			Expected: true,
		},
	}

	newCLG := MustNew()

	for i, testCase := range testCases {
		f := newCLG.(*clg).calculate(context.MustNew(), testCase.A, testCase.B)
		if f != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", f)
		}
	}
}
