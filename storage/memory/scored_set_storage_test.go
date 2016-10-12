package memory

import (
	"reflect"
	"testing"
)

func Test_ScoredSetStorage_GetElementsByScore(t *testing.T) {
	testCases := []struct {
		Key          string
		Score        float64
		MaxElements  int
		Elements     map[string]float64
		Expected     []string
		ErrorMatcher func(err error) bool
	}{
		{
			Key:         "mykey",
			Score:       0.5,
			MaxElements: -1,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
			},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       0.5,
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
			},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       0.5,
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
			},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       3.4,
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected:     []string{},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       0.8,
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			// Note the order because of the descending lexicographical order.
			Expected: []string{
				"zero.eight.two",
				"zero.eight.three",
				"zero.eight.one",
			},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       3.4,
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			Expected:     []string{},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       0.8,
			MaxElements: 2,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			// Note the order because of the descending lexicographical order.
			Expected: []string{
				"zero.eight.two",
				"zero.eight.three",
			},
			ErrorMatcher: nil,
		},
		{
			Key:   "mykey",
			Score: 0.8,
			// Note we set MaxElements to zero, so nothing should be returned.
			MaxElements: 0,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			Expected:     []string{},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       0.5,
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			Expected: []string{
				"zero.five",
			},
			ErrorMatcher: nil,
		},
	}

	for i, testCase := range testCases {
		// Setup
		newStorage := MustNew()

		for e, s := range testCase.Elements {
			err := newStorage.SetElementByScore(testCase.Key, e, s)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
		}

		// Test
		output, err := newStorage.GetElementsByScore(testCase.Key, testCase.Score, testCase.MaxElements)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		// Assert
		if testCase.ErrorMatcher == nil {
			if len(output) != len(testCase.Expected) {
				t.Fatal("case", i+1, "expected", len(testCase.Expected), "got", len(output))
			}

			for j, e := range testCase.Expected {
				if output[j] != e {
					t.Fatal("case", i+1, "expected", e, "got", output[j])
				}
			}
		}

		newStorage.Shutdown()
	}
}

func Test_ScoredSetStorage_GetHighestScoredElements(t *testing.T) {
	testCases := []struct {
		Key         string
		MaxElements int
		Elements    map[string]float64
		Expected    []string
	}{
		{
			Key:         "mykey",
			MaxElements: -1,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
				"0.5",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
				"0.5",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
				"0.5",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			// Note the order because of the descending lexicographical order.
			Expected: []string{
				"zero.eight.two",
				"0.8",
				"zero.eight.three",
				"0.8",
				"zero.eight.one",
				"0.8",
				"zero.five",
				"0.5",
				"zero.one",
				"0.1",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 0,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			Expected: []string{
				"zero.eight.two",
				"0.8",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 2,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			// Note the order because of the descending lexicographical order.
			Expected: []string{
				"zero.eight.two",
				"0.8",
				"zero.eight.three",
				"0.8",
				"zero.eight.one",
				"0.8",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			Expected: []string{
				"zero.eight.two",
				"0.8",
				"zero.eight.three",
				"0.8",
			},
		},
	}

	for i, testCase := range testCases {
		// Setup
		newStorage := MustNew()

		for e, s := range testCase.Elements {
			err := newStorage.SetElementByScore(testCase.Key, e, s)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
		}

		// Test
		output, err := newStorage.GetHighestScoredElements(testCase.Key, testCase.MaxElements)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}

		// Assert
		if len(output) != len(testCase.Expected) {
			t.Fatal("case", i+1, "expected", len(testCase.Expected), "got", len(output))
		}

		for j, e := range testCase.Expected {
			if output[j] != e {
				t.Fatal("case", i+1, "expected", e, "got", output[j])
			}
		}

		newStorage.Shutdown()
	}
}

func Test_ScoredSetStorage_RemoveScoredElement(t *testing.T) {
	testCases := []struct {
		Elements      map[string]float64
		RemoveElement string
		Expected      []string
		ErrorMatcher  func(err error) bool
	}{
		{
			RemoveElement: "zero.five",
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected:     []string{},
			ErrorMatcher: nil,
		},
		{
			RemoveElement: "zero.five",
			Elements: map[string]float64{
				"zero.five":  0.5,
				"three.four": 3.4,
			},
			Expected: []string{
				"three.four",
				"3.4",
			},
			ErrorMatcher: nil,
		},
		{
			RemoveElement: "zero.five",
			Elements:      map[string]float64{},
			Expected:      []string{},
			ErrorMatcher:  nil,
		},
		{
			RemoveElement: "invalid",
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
				"0.5",
			},
			ErrorMatcher: nil,
		},
		{
			RemoveElement: "zero.eight.one",
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			// Note the order because of the descending lexicographical order.
			Expected: []string{
				"zero.eight.two",
				"0.8",
				"zero.eight.three",
				"0.8",
				"zero.five",
				"0.5",
				"zero.one",
				"0.1",
			},
			ErrorMatcher: nil,
		},
	}

	for i, testCase := range testCases {
		// Setup
		newStorage := MustNew()

		for e, s := range testCase.Elements {
			err := newStorage.SetElementByScore("test-key", e, s)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
		}

		// Test
		err := newStorage.RemoveScoredElement("test-key", testCase.RemoveElement)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		// Assert
		if testCase.ErrorMatcher == nil {
			values, err := newStorage.GetHighestScoredElements("test-key", -1)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}

			if len(values) != len(testCase.Expected) {
				t.Fatal("case", i+1, "expected", len(testCase.Expected), "got", len(values))
			}

			for j, e := range testCase.Expected {
				if values[j] != e {
					t.Fatal("case", i+1, "expected", e, "got", values[j])
				}
			}
		}

		newStorage.Shutdown()
	}
}

func Test_ScoredSetStorage_SetWalkRemove(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	err := newStorage.SetElementByScore("test-key", "test-value-1", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	var element2 []string
	var score2 []float64
	err = newStorage.WalkScoredSet("test-key", nil, func(element string, score float64) error {
		element2 = append(element2, element)
		score2 = append(score2, score)
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if !reflect.DeepEqual(element2, []string{"test-value-1"}) {
		t.Fatal("expected", []string{"test-value-1"}, "got", element2)
	}
	if !reflect.DeepEqual(score2, []float64{0.8}) {
		t.Fatal("expected", []float64{0.8}, "got", score2)
	}
	// Add second element.
	err = newStorage.SetElementByScore("test-key", "test-value-2", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	element2 = []string{}
	score2 = []float64{}
	err = newStorage.WalkScoredSet("test-key", nil, func(element string, score float64) error {
		element2 = append(element2, element)
		score2 = append(score2, score)
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if !reflect.DeepEqual(element2, []string{"test-value-1", "test-value-2"}) {
		t.Fatal("expected", []string{"test-value-1", "test-value-2"}, "got", element2)
	}
	if !reflect.DeepEqual(score2, []float64{0.8, 0.8}) {
		t.Fatal("expected", []float64{0.8, 0.8}, "got", score2)
	}

	// Check an element can be removed from a set.
	err = newStorage.RemoveScoredElement("test-key", "test-value-1")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.RemoveScoredElement("test-key", "test-value-2")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	var element3 string
	var score3 float64
	err = newStorage.WalkScoredSet("test-key", nil, func(element string, score float64) error {
		element3 = element
		score3 = score
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element3 != "" {
		t.Fatal("expected", "", "got", element3)
	}
	if score3 != 0 {
		t.Fatal("expected", "", "got", score3)
	}
}

func Test_ScoredSetStorage_WalkScoredSet(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	err := newStorage.SetElementByScore("test-key", "test-value", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Immediately close the walk.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	// Check that the walk does not happen, because we already ended it.
	var element1 string
	var score1 float64
	err = newStorage.WalkScoredSet("test-key", closer, func(element string, score float64) error {
		element1 = element
		score1 = score
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element1 != "" {
		t.Fatal("expected", "", "got", element1)
	}
	if score1 != 0 {
		t.Fatal("expected", "", "got", score1)
	}
}
