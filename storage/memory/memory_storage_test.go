package memorystorage

import (
	"testing"
)

func Test_MemoryStorage_GetID(t *testing.T) {
	firstStorage := NewMemoryStorage(DefaultConfig())
	secondStorage := NewMemoryStorage(DefaultConfig())

	if firstStorage.GetID() == secondStorage.GetID() {
		t.Fatal("expected", "different IDs", "got", "equal IDs")
	}
}

func Test_MemoryStorage_KeyValue(t *testing.T) {
	newStorage := NewMemoryStorage(DefaultConfig())

	value, err := newStorage.Get("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if value != "" {
		t.Fatal("expected", "empty string", "got", value)
	}

	err = newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	value, err = newStorage.Get("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if value != "bar" {
		t.Fatal("expected", "bar", "got", value)
	}
}

func Test_MemoryStorage_GetHighestScoredElements_Success(t *testing.T) {
	testCases := []struct {
		Key         string
		MaxElements int
		Elements    map[string]float64
		Expected    []string
	}{
		{
			Key:         "mykey",
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five": 0.50000,
			},
			Expected: []string{
				"zero.five",
				"0.50000",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five": 0.50000,
			},
			Expected: []string{
				"zero.five",
				"0.50000",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five":        0.50000,
				"zero.eight.one":   0.80000,
				"zero.eight.two":   0.80000,
				"zero.eight.three": 0.80000,
				"zero.one":         0.10000,
			},
			// Note the order because of the lexicographical order.
			Expected: []string{
				"zero.eight.one",
				"0.80000",
				"zero.eight.three",
				"0.80000",
				"zero.eight.two",
				"0.80000",
				"zero.five",
				"0.50000",
				"zero.one",
				"0.10000",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 0,
			Elements: map[string]float64{
				"zero.five":        0.50000,
				"zero.eight.one":   0.80000,
				"zero.eight.two":   0.80000,
				"zero.eight.three": 0.80000,
				"zero.one":         0.10000,
			},
			Expected: []string{},
		},
		{
			Key:         "mykey",
			MaxElements: 2,
			Elements: map[string]float64{
				"zero.five":        0.50000,
				"zero.eight.one":   0.80000,
				"zero.eight.two":   0.80000,
				"zero.eight.three": 0.80000,
				"zero.one":         0.10000,
			},
			// Note the order because of the lexicographical order.
			Expected: []string{
				"zero.eight.one",
				"0.80000",
				"zero.eight.three",
				"0.80000",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five":        0.50000,
				"zero.eight.one":   0.80000,
				"zero.eight.two":   0.80000,
				"zero.eight.three": 0.80000,
				"zero.one":         0.10000,
			},
			Expected: []string{
				"zero.eight.one",
				"0.80000",
			},
		},
	}

	for i, testCase := range testCases {
		// Setup
		newStorage := NewMemoryStorage(DefaultConfig())

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

		for i, e := range testCase.Expected {
			if output[i] != e {
				t.Fatal("case", i+1, "expected", e, "got", output[i])
			}
		}
	}
}

func Test_MemoryStorage_GetHighestScoredElements_NotFound(t *testing.T) {
	newStorage := NewMemoryStorage(DefaultConfig())

	output, err := newStorage.GetHighestScoredElements("not-found", 3)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if len(output) != 0 {
		t.Fatal("expected", 0, "got", len(output))
	}
}

func Test_MemoryStorage_GetElementsByScore_Success(t *testing.T) {
	testCases := []struct {
		Key         string
		Score       float64
		MaxElements int
		Elements    map[string]float64
		Expected    []string
	}{
		{
			Key:         "mykey",
			Score:       0.50000,
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five": 0.50000,
			},
			Expected: []string{
				"zero.five",
			},
		},
		{
			Key:         "mykey",
			Score:       0.50000,
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five": 0.50000,
			},
			Expected: []string{
				"zero.five",
			},
		},
		{
			Key:         "mykey",
			Score:       3.40000,
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five": 0.50000,
			},
			Expected: []string{},
		},
		{
			Key:         "mykey",
			Score:       0.80000,
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five":        0.50000,
				"zero.eight.one":   0.80000,
				"zero.eight.two":   0.80000,
				"zero.eight.three": 0.80000,
				"zero.one":         0.10000,
			},
			// Note the order because of the lexicographical order.
			Expected: []string{
				"zero.eight.one",
				"zero.eight.three",
				"zero.eight.two",
			},
		},
		{
			Key:         "mykey",
			Score:       3.40000,
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five":        0.50000,
				"zero.eight.one":   0.80000,
				"zero.eight.two":   0.80000,
				"zero.eight.three": 0.80000,
				"zero.one":         0.10000,
			},
			Expected: []string{},
		},
		{
			Key:         "mykey",
			Score:       0.80000,
			MaxElements: 2,
			Elements: map[string]float64{
				"zero.five":        0.50000,
				"zero.eight.one":   0.80000,
				"zero.eight.two":   0.80000,
				"zero.eight.three": 0.80000,
				"zero.one":         0.10000,
			},
			// Note the order because of the lexicographical order.
			Expected: []string{
				"zero.eight.one",
				"zero.eight.three",
			},
		},
		{
			Key:         "mykey",
			Score:       0.80000,
			MaxElements: 0,
			Elements: map[string]float64{
				"zero.five":        0.50000,
				"zero.eight.one":   0.80000,
				"zero.eight.two":   0.80000,
				"zero.eight.three": 0.80000,
				"zero.one":         0.10000,
			},
			Expected: []string{},
		},
		{
			Key:         "mykey",
			Score:       0.50000,
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five":        0.50000,
				"zero.eight.one":   0.80000,
				"zero.eight.two":   0.80000,
				"zero.eight.three": 0.80000,
				"zero.one":         0.10000,
			},
			Expected: []string{
				"zero.five",
			},
		},
	}

	for i, testCase := range testCases {
		// Setup
		newStorage := NewMemoryStorage(DefaultConfig())

		for e, s := range testCase.Elements {
			err := newStorage.SetElementByScore(testCase.Key, e, s)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
		}

		// Test
		output, err := newStorage.GetElementsByScore(testCase.Key, testCase.Score, testCase.MaxElements)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}

		// Assert
		if len(output) != len(testCase.Expected) {
			t.Fatal("case", i+1, "expected", len(testCase.Expected), "got", len(output))
		}

		for i, e := range testCase.Expected {
			if output[i] != e {
				t.Fatal("case", i+1, "expected", e, "got", output[i])
			}
		}
	}
}

func Test_MemoryStorage_GetElementsByScore_NotFound(t *testing.T) {
	newStorage := NewMemoryStorage(DefaultConfig())

	output, err := newStorage.GetElementsByScore("not-found", 0.8, 3)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if len(output) != 0 {
		t.Fatal("expected", 0, "got", len(output))
	}
}
