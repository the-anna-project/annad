package clg

import (
	"testing"
)

func Test_CLG_IsBetween(t *testing.T) {
	testCases := []struct {
		N        float64
		Min      float64
		Max      float64
		Expected bool
	}{
		{
			N:        1,
			Min:      2,
			Max:      4,
			Expected: false,
		},
		{
			N:        2,
			Min:      2,
			Max:      4,
			Expected: true,
		},
		{
			N:        3,
			Min:      2,
			Max:      4,
			Expected: true,
		},
		{
			N:        4,
			Min:      2,
			Max:      4,
			Expected: true,
		},
		{
			N:        5,
			Min:      2,
			Max:      4,
			Expected: false,
		},
		{
			N:        35,
			Min:      -13,
			Max:      518,
			Expected: true,
		},
		{
			N:        -87,
			Min:      -413,
			Max:      -18,
			Expected: true,
		},
		{
			N:        -7,
			Min:      -413,
			Max:      -18,
			Expected: false,
		},
		{
			N:        -987,
			Min:      -413,
			Max:      -18,
			Expected: false,
		},
		{
			N:        1.8,
			Min:      2.34,
			Max:      4.944,
			Expected: false,
		},
		{
			N:        2.334,
			Min:      2.2,
			Max:      4.1,
			Expected: true,
		},
		{
			N:        3.9,
			Min:      2.003,
			Max:      4,
			Expected: true,
		},
		{
			N:        4,
			Min:      2.22,
			Max:      4.83,
			Expected: true,
		},
	}

	for i, testCase := range testCases {
		output := testMaybeNewCollection(t).IsBetween(testCase.N, testCase.Min, testCase.Max)

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}

func Test_CLG_Difference(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected float64
	}{
		{
			A:        3.5,
			B:        12.5,
			Expected: -9,
		},
		{
			A:        35.5,
			B:        14.5,
			Expected: 21,
		},
		{
			A:        -3.5,
			B:        -7.5,
			Expected: 4,
		},
		{
			A:        12.5,
			B:        4.5,
			Expected: 8,
		},
		{
			A:        36.5,
			B:        6.5,
			Expected: 30,
		},
		{
			A:        11.11,
			B:        10.10,
			Expected: 1.0099999999999998,
		},
	}

	for i, testCase := range testCases {
		output := testMaybeNewCollection(t).Difference(testCase.A, testCase.B)

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}

func Test_CLG_Greater(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected float64
	}{
		{
			A:        3.5,
			B:        3.5,
			Expected: 3.5,
		},
		{
			A:        3.5,
			B:        12.5,
			Expected: 12.5,
		},
		{
			A:        35.5,
			B:        14.5,
			Expected: 35.5,
		},
		{
			A:        -3.5,
			B:        7.5,
			Expected: 7.5,
		},
		{
			A:        12.5,
			B:        4.5,
			Expected: 12.5,
		},
		{
			A:        17,
			B:        65,
			Expected: 65,
		},
		{
			A:        65,
			B:        17,
			Expected: 65,
		},
	}

	for i, testCase := range testCases {
		output := testMaybeNewCollection(t).Greater(testCase.A, testCase.B)

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}

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

	for i, testCase := range testCases {
		output := testMaybeNewCollection(t).IsGreater(testCase.A, testCase.B)

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}

func Test_CLG_Product(t *testing.T) {
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

	for i, testCase := range testCases {
		output := testMaybeNewCollection(t).Product(testCase.A, testCase.B)

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}

func Test_CLG_Sum(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected float64
	}{
		{
			A:        3.5,
			B:        12.5,
			Expected: 16,
		},
		{
			A:        35.5,
			B:        14.5,
			Expected: 50,
		},
		{
			A:        -3.5,
			B:        7.5,
			Expected: 4,
		},
		{
			A:        12.5,
			B:        4.5,
			Expected: 17,
		},
		{
			A:        36.5,
			B:        6.5,
			Expected: 43,
		},
		{
			A:        36,
			B:        6.5,
			Expected: 42.5,
		},
		{
			A:        99.99,
			B:        12.15,
			Expected: 112.14,
		},
		{
			A:        17,
			B:        65,
			Expected: 82,
		},
	}

	for i, testCase := range testCases {
		output := testMaybeNewCollection(t).Sum(testCase.A, testCase.B)

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}

func Test_CLG_Quotient(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected float64
	}{
		{
			A:        3.5,
			B:        12.5,
			Expected: 0.28,
		},
		{
			A:        35.5,
			B:        14.5,
			Expected: 2.4482758620689653,
		},
		{
			A:        -3.5,
			B:        7.5,
			Expected: -0.4666666666666667,
		},
		{
			A:        12.5,
			B:        4.5,
			Expected: 2.7777777777777777,
		},
		{
			A:        36.5,
			B:        6.5,
			Expected: 5.615384615384615,
		},
	}

	for i, testCase := range testCases {
		output := testMaybeNewCollection(t).Quotient(testCase.A, testCase.B)

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}
