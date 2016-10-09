package multiply

import (
	"testing"

	"github.com/xh3b4sd/anna/context"
)

func Test_CLG_Multiply(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected float64
	}{
		{
			A:        3.5,
			B:        12.5,
			Expected: 43.75,
		},
		{
			A:        35.5,
			B:        14.5,
			Expected: 514.75,
		},
		{
			A:        -3.5,
			B:        7.5,
			Expected: -26.25,
		},
		{
			A:        12.5,
			B:        4.5,
			Expected: 56.25,
		},
		{
			A:        36.5,
			B:        6.5,
			Expected: 237.25,
		},
		{
			A:        17,
			B:        65,
			Expected: 1105,
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
