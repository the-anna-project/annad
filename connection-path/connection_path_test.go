package connectionpath

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewConnectionPath(t *testing.T, coordinates [][]float64) spec.ConnectionPath {
	newConfig := DefaultConfig()
	newConfig.Coordinates = coordinates

	newConnectionPath, err := New(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newConnectionPath
}

func Test_ConnectionPath_NewFromString(t *testing.T) {
	s := "[[0.332,4,-8.5]]"
	expected := [][]float64{
		{0.332, 4, -8.5},
	}

	newConnectionPath, err := NewFromString(s)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	e := newConnectionPath.GetCoordinates()
	if !reflect.DeepEqual(e, expected) {
		t.Fatal("expected", expected, "got", e)
	}
}

func Test_ConnectionPath_NewFromString_Error(t *testing.T) {
	s := "0.332,4,-8.5]]"

	_, err := NewFromString(s)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ConnectionPath_DistanceTo(t *testing.T) {
	testCases := []struct {
		A        [][]float64
		B        [][]float64
		Distance float64
	}{
		{
			A:        [][]float64{{1, 1, 1}},
			B:        [][]float64{{1, 1, 1}},
			Distance: 0,
		},
		{
			A:        [][]float64{{1, 1, 1}},
			B:        [][]float64{{2, 2, 2}},
			Distance: 3,
		},
		{
			A:        [][]float64{{1, 1, 1}, {1, 1, 1}},
			B:        [][]float64{{2, 2, 2}, {2, 2, 2}},
			Distance: 6,
		},
		{
			A:        [][]float64{{2, 2, 2}, {2, 2, 2}},
			B:        [][]float64{{1, 1, 1}, {1, 1, 1}},
			Distance: 6,
		},
		{
			A:        [][]float64{{1, 1, 1}},
			B:        [][]float64{{2, 2, 2}, {2, 2, 2}},
			Distance: 6,
		},
		{
			A:        [][]float64{{2, 2, 2}, {2, 2, 2}},
			B:        [][]float64{{1, 1, 1}},
			Distance: 6,
		},
		{
			A:        [][]float64{{1, 1, 1}},
			B:        [][]float64{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
			Distance: 9,
		},
		{
			A:        [][]float64{{1, 1, 1}, {1, 1, 1}},
			B:        [][]float64{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
			Distance: 9,
		},
		{
			A:        [][]float64{{1, 1, 1}},
			B:        [][]float64{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
			Distance: 12,
		},
		{
			A:        [][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			B:        [][]float64{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
			Distance: 12,
		},
		{
			A:        [][]float64{{1.3, 2, 4.66}, {1, 1, 1}, {3, 17.8, 9}},
			B:        [][]float64{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
			Distance: 53.96,
		},
	}

	for i, testCase := range testCases {
		a := testMaybeNewConnectionPath(t, testCase.A)
		b := testMaybeNewConnectionPath(t, testCase.B)

		distance := a.DistanceTo(b)

		if distance != testCase.Distance {
			t.Fatal("case", i+1, "expected", testCase.Distance, "got", distance)
		}
	}
}

func Test_ConnectionPath_IsCloser(t *testing.T) {
	testCases := []struct {
		CP           [][]float64
		M            map[string][][]float64
		ID           string
		ErrorMatcher func(err error) bool
	}{
		{
			CP: [][]float64{{1, 1, 1}},
			M: map[string][][]float64{
				"A": {{1, 1, 1}},
				"B": {{2, 2, 2}},
			},
			ID:           "A",
			ErrorMatcher: nil,
		},
		{
			CP: [][]float64{{2, 2, 2}},
			M: map[string][][]float64{
				"A": {{1, 1, 1}},
				"B": {{2, 2, 2}},
			},
			ID:           "B",
			ErrorMatcher: nil,
		},
		{
			CP: [][]float64{{3, 3, 3}},
			M: map[string][][]float64{
				"A": {{1, 1, 1}},
				"B": {{2, 2, 2}},
			},
			ID:           "B",
			ErrorMatcher: nil,
		},
		{
			CP: [][]float64{{1, 2, 1}},
			M: map[string][][]float64{
				"A": {{1, 1, 1}},
				"B": {{2, 2, 2}},
			},
			ID:           "A",
			ErrorMatcher: nil,
		},
	}

	for i, testCase := range testCases {
		cp := testMaybeNewConnectionPath(t, testCase.CP)
		a := testMaybeNewConnectionPath(t, testCase.M["A"])
		b := testMaybeNewConnectionPath(t, testCase.M["B"])

		x, err := cp.IsCloser(a, b)

		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if testCase.ID == "A" {
				if !reflect.DeepEqual(a, x) {
					t.Fatal("case", i+1, "expected", a, "got", x)
				}
			}
			if testCase.ID == "B" {
				if !reflect.DeepEqual(b, x) {
					t.Fatal("case", i+1, "expected", b, "got", x)
				}
			}
		}
	}
}

func Test_ConnectionPath_IsCloser_Equal(t *testing.T) {
	testCases := []struct {
		CP           [][]float64
		M            map[string][][]float64
		ID           string
		ErrorMatcher func(err error) bool
	}{
		{
			CP: [][]float64{{1, 1, 1}},
			M: map[string][][]float64{
				"A": {{2, 2, 2}},
				"B": {{2, 2, 2}},
			},
			ID:           "A",
			ErrorMatcher: nil,
		},
		{
			CP: [][]float64{{1, 1, 1}},
			M: map[string][][]float64{
				"A": {{2, 2, 2}},
				"B": {{2, 2, 2}},
			},
			ID:           "B",
			ErrorMatcher: nil,
		},
	}

	for i, testCase := range testCases {
		cp := testMaybeNewConnectionPath(t, testCase.CP)
		a := testMaybeNewConnectionPath(t, testCase.M["A"])
		b := testMaybeNewConnectionPath(t, testCase.M["B"])

		var xs []spec.ConnectionPath
		for i := 0; i < 100; i++ {
			x, err := cp.IsCloser(a, b)
			if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
				t.Fatal("case", i+1, "expected", true, "got", false)
			}
			xs = append(xs, x)
		}
		if testCase.ErrorMatcher == nil {
			var found bool
			for _, x := range xs {
				if testCase.ID == "A" {
					if reflect.DeepEqual(a, x) {
						found = true
						break
					}
				}
				if testCase.ID == "B" {
					if reflect.DeepEqual(b, x) {
						found = true
						break
					}
				}
			}
			if !found {
				t.Fatal("case", i+1, "expected", true, "got", false)
			}
		}
	}
}

func Test_ConnectionPath_String(t *testing.T) {
	newCoordinates := [][]float64{
		{0.332, 4, -8.5},
	}
	expected := "[[0.332,4,-8.5]]"

	newConnectionPath := testMaybeNewConnectionPath(t, newCoordinates)

	s, err := newConnectionPath.String()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if s != expected {
		t.Fatal("expected", expected, "got", s)
	}
}

func Test_ConnectionPath_Validate(t *testing.T) {
	testCases := []struct {
		Coordinates  [][]float64
		ErrorMatcher func(err error) bool
	}{
		{
			Coordinates:  [][]float64{{4}},
			ErrorMatcher: nil,
		},
		{
			Coordinates:  [][]float64{{0.332, 4, -8.5}},
			ErrorMatcher: nil,
		},
		{
			Coordinates:  [][]float64{{4}, {0.332}},
			ErrorMatcher: nil,
		},
		{
			Coordinates:  [][]float64{{0.332, 4, -8.5}, {0.332, 4, -8.5}},
			ErrorMatcher: nil,
		},
		{
			Coordinates:  [][]float64{},
			ErrorMatcher: IsInvalidConnectionPath,
		},
		{
			Coordinates:  [][]float64{{}},
			ErrorMatcher: IsInvalidConnectionPath,
		},
		{
			Coordinates:  [][]float64{{}, {}},
			ErrorMatcher: IsInvalidConnectionPath,
		},
		{
			Coordinates:  [][]float64{{4}, {}},
			ErrorMatcher: IsInvalidConnectionPath,
		},
		{
			Coordinates:  [][]float64{{0.332, 4}, {-8}},
			ErrorMatcher: IsInvalidConnectionPath,
		},
	}

	for i, testCase := range testCases {
		newConnectionPath := testMaybeNewConnectionPath(t, testCase.Coordinates)
		err := newConnectionPath.Validate()

		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
	}
}
